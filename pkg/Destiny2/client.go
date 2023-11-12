package Destiny2

import (
	"fmt"
	"github.com/d2checkpoint-com/BungieNet/internal/rest"
	"github.com/d2checkpoint-com/BungieNet/internal/shared"
	"github.com/d2checkpoint-com/BungieNet/pkg/Models"
)

const (
	bungieUrl = "https://www.bungie.net/Platform"
)

type Client struct {
	ApiKey       string
	UserAgent    string
	CacheBreaker bool
}

func NewClient(options ...any) *Client {
	return &Client{
		ApiKey:       options[0].(string),
		UserAgent:    shared.SetUserAgent(options),
		CacheBreaker: shared.SetCacheBreakOpt(options),
	}
}

func (c *Client) GetDestinyManifest() (*Models.Response, error) {
	url := bungieUrl + "/Destiny2/Manifest/"

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

func (c *Client) GetProfile(membershipId int64, membershipType int32, components ...any) (*Models.Response, error) {
	// Ensure at least one component is specified
	if len(components) == 0 {
		return nil, fmt.Errorf("no components specified")
	}

	// Build URL
	url := fmt.Sprintf("%s/Destiny2/%d/Profile/%d/?components=", bungieUrl, membershipType, membershipId)
	for _, component := range components {
		url += fmt.Sprintf("%d,", component)
	}
	url = url[:len(url)-1]

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

func (c *Client) GetCharacter(membershipType int32, membershipId, characterId int64, components ...any) (*Models.Response, error) {
	// Ensure at least one component is specified
	if len(components) == 0 {
		return nil, fmt.Errorf("no components specified")
	}

	// Build URL
	url := fmt.Sprintf("%s/Destiny2/%d/Profile/%d/Character/%d/?components=", bungieUrl, membershipType, membershipId, characterId)
	for _, component := range components {
		url += fmt.Sprintf("%d,", component)
	}
	url = url[:len(url)-1]

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
