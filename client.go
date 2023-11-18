package bungienet

import (
	"encoding/json"
	"fmt"
	"github.com/d2checkpoint-com/bungienet/internal/shared"
	"github.com/d2checkpoint-com/bungienet/pkg/client/Destiny2"
	"github.com/d2checkpoint-com/bungienet/pkg/client/User"
	"github.com/d2checkpoint-com/bungienet/pkg/model"
	"github.com/d2checkpoint-com/bungienet/pkg/model/Common/Models"
	"github.com/sendgrid/rest"
	"time"
)

const (
	bungieUrl = "https://www.bungie.net/Platform"
)

type Client struct {
	ApiKey    string
	UserAgent string
	Destiny2  *Destiny2.Client
	User      *User.Client
}

// NewClient creates a new client with the specified options.
func NewClient(options ...any) *Client {
	return &Client{
		ApiKey:    options[0].(string),
		UserAgent: shared.SetUserAgent(options),
		Destiny2:  Destiny2.NewClient(options...),
		User:      User.NewClient(options...),
	}
}

type StreamingAlerts bool

const IncludeStreaming StreamingAlerts = true

func (c *Client) GetGlobalAlerts(includeStreaming ...any) (*GetGlobalAlertsResponse, error) {
	q := make(map[string]string)
	for _, option := range includeStreaming {
		switch option.(type) {
		case model.CacheBreak:
			q["t"] = time.Now().String()
		case StreamingAlerts:
			q["includestreaming"] = "true"
		}
	}

	// Send and return request
	h, err := rest.Send(rest.Request{
		Method:      rest.Get,
		BaseURL:     fmt.Sprintf("%s/GlobalAlerts/", bungieUrl),
		QueryParams: q,
		Headers: map[string]string{
			"X-API-Key":  c.ApiKey,
			"User-Agent": c.UserAgent,
		},
	})
	if err != nil {
		return nil, err
	}

	var r *GetGlobalAlertsResponse
	err = json.Unmarshal([]byte(h.Body), &r)
	if err != nil {
		return nil, err
	}

	return r, shared.CheckResponse(r.ErrorCode, r.ErrorStatus, r.Message)
}

type GetGlobalAlertsResponse struct {
	Response           []*model.GlobalAlert `json:"Response"`
	ErrorCode          int32                `json:"ErrorCode"`
	ThrottleSeconds    int32                `json:"ThrottleSeconds"`
	ErrorStatus        string               `json:"ErrorStatus"`
	Message            string               `json:"Message"`
	MessageData        map[string]string    `json:"MessageData"`
	DetailedErrorTrace string               `json:"DetailedErrorTrace"`
}

func (c *Client) GetCommonSettings(options ...any) (*GetCommonSettingsResponse, error) {
	q := make(map[string]string)
	for _, option := range options {
		switch option.(type) {
		case model.CacheBreak:
			q["t"] = time.Now().String()
		}
	}

	// Send and return request
	h, err := rest.Send(rest.Request{
		Method:      rest.Get,
		BaseURL:     fmt.Sprintf("%s/Settings/", bungieUrl),
		QueryParams: q,
		Headers: map[string]string{
			"X-API-Key":  c.ApiKey,
			"User-Agent": c.UserAgent,
		},
	})
	if err != nil {
		return nil, err
	}

	var r *GetCommonSettingsResponse
	err = json.Unmarshal([]byte(h.Body), &r)
	if err != nil {
		return nil, err
	}

	return r, shared.CheckResponse(r.ErrorCode, r.ErrorStatus, r.Message)
}

type GetCommonSettingsResponse struct {
	Response           []*Models.CoreSettingsConfiguration `json:"Response"`
	ErrorCode          int32                               `json:"ErrorCode"`
	ThrottleSeconds    int32                               `json:"ThrottleSeconds"`
	ErrorStatus        string                              `json:"ErrorStatus"`
	Message            string                              `json:"Message"`
	MessageData        map[string]string                   `json:"MessageData"`
	DetailedErrorTrace string                              `json:"DetailedErrorTrace"`
}
