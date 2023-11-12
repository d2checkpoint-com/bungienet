package bungienet

import (
	"github.com/d2checkpoint-com/BungieNet/internal/rest"
	"github.com/d2checkpoint-com/BungieNet/internal/shared"
	"github.com/d2checkpoint-com/BungieNet/pkg/Destiny2"
	"github.com/d2checkpoint-com/BungieNet/pkg/Models"
	"github.com/d2checkpoint-com/BungieNet/pkg/User"
)

const (
	bungieUrl = "https://www.bungie.net/Platform"
)

type Client struct {
	ApiKey       string
	UserAgent    string
	CacheBreaker bool
	Destiny2     *Destiny2.Client
	User         *User.Client
}

func NewClient(options ...any) *Client {
	return &Client{
		ApiKey:       options[0].(string),
		UserAgent:    shared.SetUserAgent(options),
		CacheBreaker: shared.SetCacheBreakOpt(options),
		Destiny2:     Destiny2.NewClient(options...),
		User:         User.NewClient(options...),
	}
}

func (c *Client) GetGlobalAlerts(includeStreaming ...bool) (*Models.Response, error) {
	// Build URL
	url := bungieUrl + "/GlobalAlerts/"
	if len(includeStreaming) > 0 && includeStreaming[0] {
		url += "?includeStreaming=true"
		if c.CacheBreaker {
			url += "&t=%d"
		} else if c.CacheBreaker {
			url += "?t=%d"
		}
	}

	// Build request
	request := rest.Request{
		Method:  rest.Get,
		BaseURL: url,
		Headers: map[string]string{
			"X-API-Key":  c.ApiKey,
			"User-Agent": c.UserAgent,
		},
	}

	// Send request
	response, err := rest.Send(request)
	if err != nil {
		return nil, err
	}

	return shared.UnmarshalResponse(response)
}

func (c *Client) GetCommonSettings() (*Models.Response, error) {
	// Build URL
	url := bungieUrl + "/Platform/Settings/"

	// Send and return request
	return shared.SendRequest(rest.Request{
		Method:  rest.Get,
		BaseURL: shared.SetCacheBreakUrl(url, c.CacheBreaker),
		Headers: map[string]string{
			"X-API-Key":  c.ApiKey,
			"User-Agent": c.UserAgent,
		},
	})
}