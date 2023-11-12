package User

import (
	"fmt"
	"github.com/d2checkpoint-com/bungienet/internal/rest"
	"github.com/d2checkpoint-com/bungienet/internal/shared"
	"github.com/d2checkpoint-com/bungienet/pkg/Models"
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

func (c *Client) GetMembershipDataById(membershipId int64, membershipType int32) (*Models.Response, error) {
	// Build URL
	url := fmt.Sprintf("%s/User/GetMembershipsById/%d/%d/", bungieUrl, membershipId, membershipType)

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
