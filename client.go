package bungienet

import (
	"errors"
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/d2checkpoint-com/bungienet/pkg/client"
	"github.com/d2checkpoint-com/bungienet/pkg/client/Destiny2"
	"github.com/d2checkpoint-com/bungienet/pkg/client/User"
	"github.com/d2checkpoint-com/bungienet/pkg/model"
	"github.com/d2checkpoint-com/bungienet/pkg/model/Common/Models"
	"net/url"
)

const (
	bungieUrl  = "https://www.bungie.net/Platform"
	BreakCache = client.BreakCache
)

type ApiKey = client.ApiKey
type UserAgent = client.UserAgent

type Client struct {
	*Destiny2.Destiny2
	*User.User
	client *client.Client
}

// NewClient creates a new client with the specified options.
func NewClient(apiKey string, options ...any) *Client {
	for _, option := range options {
		switch option.(type) {
		case client.ApiKey:
			apiKey = option.(client.ApiKey).String()
		}
	}

	if apiKey == "" {
		panic("missing api key")
	}

	client.NewClient(client.ApiKey(apiKey), options)

	return &Client{Destiny2.NewClient(), User.NewClient(), &client.BungieClient}
}

type StreamingAlerts bool

const IncludeStreaming StreamingAlerts = true

func (c *Client) GetGlobalAlerts(includeStreaming ...any) (*GetGlobalAlertsResponse, error) {
	var q url.Values
	var cacheBreak client.CacheBreak
	for _, option := range includeStreaming {
		switch option.(type) {
		case client.CacheBreak:
			cacheBreak = option.(client.CacheBreak)
		case StreamingAlerts:
			if q == nil {
				q = url.Values{}
			}
			q.Add("includestreaming", fmt.Sprintf("%t,", includeStreaming))
		}
	}

	res, err := c.client.Send(client.Request{
		Method:      client.Get,
		BaseURL:     fmt.Sprintf("%s/GlobalAlerts/", bungieUrl),
		QueryParams: q,
		CacheBreak:  cacheBreak,
	})

	var r *GetGlobalAlertsResponse
	return r, errors.Join(err, sonic.Unmarshal(res, &r))
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
	var cacheBreak client.CacheBreak
	for _, option := range options {
		switch option.(type) {
		case client.CacheBreak:
			cacheBreak = option.(client.CacheBreak)
		}
	}

	res, err := c.client.Send(client.Request{
		Method:     client.Get,
		BaseURL:    fmt.Sprintf("%s/Settings/", bungieUrl),
		CacheBreak: cacheBreak,
	})

	var r *GetCommonSettingsResponse
	return r, errors.Join(err, sonic.Unmarshal(res, &r))
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
