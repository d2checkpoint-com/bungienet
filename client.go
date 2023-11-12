package bungienet

import (
	"github.com/d2checkpoint-com/bungienet/internal/rest"
	"github.com/d2checkpoint-com/bungienet/internal/shared"
	"github.com/d2checkpoint-com/bungienet/pkg/Destiny2"
	"github.com/d2checkpoint-com/bungienet/pkg/Models"
	"github.com/d2checkpoint-com/bungienet/pkg/User"
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

// NewClient creates a new client with the specified options.
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
