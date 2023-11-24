package client

import (
	"context"
	"fmt"
	"net/http"
	"net/http/cookiejar"
)

var BungieClient Client

// Client allows modification of client headers, redirect policy and other settings
type Client struct {
	*http.Client
	*RateLimiter
	ApiKey
	UserAgent
}

// NewClient initializes the BungieClient variable.
func NewClient(options ...any) {
	var apiKey ApiKey
	var userAgent UserAgent
	var httpClient *http.Client

	for _, option := range options {
		switch option.(type) {
		case ApiKey:
			apiKey = option.(ApiKey)
		case UserAgent:
			userAgent = option.(UserAgent)
		case *http.Client:
			httpClient = option.(*http.Client)
		}
	}

	BungieClient = Client{
		ApiKey:      apiKey,
		UserAgent:   userAgent,
		Client:      newHttpClient(httpClient),
		RateLimiter: newRateLimiter(),
	}
	BungieClient.Client.Jar = newCookieJar()
}

func newHttpClient(c *http.Client) *http.Client {
	if c == nil {
		return &http.Client{}
	}
	return c
}

func newCookieJar() *cookiejar.Jar {
	jar, err := cookiejar.New(nil)
	if err != nil {
		panic(fmt.Errorf("unable to create cookie jar: %w", err))
	}
	return jar
}

// Send will build your request, make the request, and build your response.
func (c *Client) Send(request Request) ([]byte, error) {
	return c.SendWithContext(context.Background(), request)
}

// SendWithContext will build your request passing in the provided context, make the request, and build your response.
func (c *Client) SendWithContext(ctx context.Context, request Request) ([]byte, error) {
	// Throttle using the rate limiter.
	err := c.RateLimiter.Wait(request.IsStatRequest, ctx)
	if err != nil {
		return nil, err
	}

	// Build the HTTP request object.
	req, err := buildRequestObject(request)
	if err != nil {
		return nil, err
	}
	// Pass in the user provided context
	req = req.WithContext(ctx)

	// Build the HTTP client and make the request.
	res, err := c.makeRequest(req)
	if err != nil {
		return nil, err
	}

	// Build HttpResponse object.
	return buildResponse(res)
}

type ApiKey string

func (a ApiKey) String() string {
	return string(a)
}

type UserAgent string

func (u UserAgent) String() string {
	return string(u)
}
