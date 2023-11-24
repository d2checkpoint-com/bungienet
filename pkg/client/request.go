package client

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/bytedance/sonic"
	"io"
	"net/http"
	"net/url"
	"time"
)

// makeRequest makes the API call.
func (c *Client) makeRequest(req *http.Request) (*http.Response, error) {
	return c.Client.Do(req)
}

// Request holds the request to an API Call.
type Request struct {
	Method      Method
	BaseURL     string
	Headers     map[string]string
	QueryParams url.Values
	Body        []byte
	CacheBreak
	IsStatRequest
}

// Method contains the supported HTTP verbs.
type Method string

// Supported HTTP verbs.
const (
	Get    Method = "GET"
	Post   Method = "POST"
	Put    Method = "PUT"
	Patch  Method = "PATCH"
	Delete Method = "DELETE"
)

type CacheBreak bool

// BreakCache is a value that can be passed to through the Request to break the cache.
const BreakCache CacheBreak = true

type IsStatRequest bool

const StatRequest IsStatRequest = true

// buildRequestObject creates the HTTP request object.
func buildRequestObject(request Request) (*http.Request, error) {
	// Add a CacheBreak query parameter if the request is a stat request.
	if request.CacheBreak {
		t := fmt.Sprintf("%d", time.Now().UnixNano())
		if request.QueryParams == nil {
			request.QueryParams = url.Values{}
		}
		request.QueryParams.Add(t[len(t)/2:], t[:len(t)/2])
	}
	// Add any query parameters to the URL.
	if request.QueryParams != nil {
		request.BaseURL = addQueryParameters(request.BaseURL, request.QueryParams)
	}
	req, err := http.NewRequest(string(request.Method), request.BaseURL, bytes.NewBuffer(request.Body))
	if err != nil {
		return req, err
	}
	for key, value := range request.Headers {
		req.Header.Set(key, value)
	}
	if _, exists := req.Header["X-API-Key"]; !exists {
		req.Header.Set("X-API-Key", BungieClient.ApiKey.String())
	}
	if _, exists := req.Header["User-Agent"]; !exists && BungieClient.UserAgent.String() != "" {
		req.Header.Set("User-Agent", BungieClient.UserAgent.String())
	}
	if _, exists := req.Header["Content-Type"]; len(request.Body) > 0 && !exists {
		req.Header.Set("Content-Type", "application/json")
	}
	return req, err
}

// addQueryParameters adds query parameters to the URL.
func addQueryParameters(baseURL string, params url.Values) string {
	return baseURL + "?" + params.Encode()
}

// buildResponse builds the response struct.
func buildResponse(res *http.Response) ([]byte, error) {
	body, err := io.ReadAll(res.Body)
	err = errors.Join(err, res.Body.Close())
	if res.StatusCode != 200 {
		err = errors.Join(err, fmt.Errorf("HttpStatus: %d %s", res.StatusCode, res.Status))
	}
	return body, checkResponse(&body, err)
}

func checkResponse(b *[]byte, err error) error {
	type response struct {
		ErrorCode          int32             `json:"ErrorCode"`
		ThrottleSeconds    int32             `json:"ThrottleSeconds"`
		ErrorStatus        string            `json:"ErrorStatus"`
		Message            string            `json:"Message"`
		MessageData        map[string]string `json:"MessageData"`
		DetailedErrorTrace string            `json:"DetailedErrorTrace"`
	}
	var r response
	err = errors.Join(err, sonic.Unmarshal(*b, &r))

	var s string
	if r.ErrorCode != 1 && r.ErrorCode != 0 {
		s = fmt.Sprintf("ErrorCode: %d", r.ErrorCode)
	}

	if r.ErrorStatus != "" && r.ErrorStatus != "Success" {
		s = fmt.Sprintf("%s ErrorStatus: %s", s, r.ErrorStatus)
	}

	if r.Message != "" && r.Message != "Ok" {
		s = fmt.Sprintf("%s Message: %s", s, r.Message)
	}

	if len(r.MessageData) > 0 {
		s = fmt.Sprintf("%s MessageData: %v", s, r.MessageData)
	}

	if r.DetailedErrorTrace != "" {
		s = fmt.Sprintf("%s DetailedErrorTrace: %s", s, r.DetailedErrorTrace)
	}

	BungieClient.ThrottleSeconds.Set(r.ThrottleSeconds)
	if r.ThrottleSeconds != 0 {
		s = fmt.Sprintf("%s ThrottleSeconds: %d", s, r.ThrottleSeconds)
	}
	if s != "" {
		err = errors.Join(err, errors.New(s))
	}

	return err
}
