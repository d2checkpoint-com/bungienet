package Destiny2

import (
	"errors"
	"fmt"
	"github.com/bytedance/sonic"
	client2 "github.com/d2checkpoint-com/bungienet/internal/client"
	"github.com/d2checkpoint-com/bungienet/pkg/model/Destiny/Config"
	"github.com/d2checkpoint-com/bungienet/pkg/model/Destiny/Responses"
	"net/url"
)

const (
	bungieUrl = "https://www.bungie.net/Platform"
	statsUrl  = "https://stats.bungie.net/Platform"
)

type Destiny2 struct {
	client *client2.Client
}

func NewClient(c *client2.Client) *Destiny2 {
	return &Destiny2{c}
}

func (c *Destiny2) GetDestinyManifest(options ...any) (*GetDestinyManifestResponse, error) {
	var cacheBreak client2.CacheBreak
	for _, option := range options {
		switch option.(type) {
		case client2.CacheBreak:
			cacheBreak = option.(client2.CacheBreak)
		}
	}

	// Send and return request
	res, err := c.client.Send(client2.Request{
		Method:     client2.Get,
		BaseURL:    bungieUrl + "/Destiny2/Manifest/",
		CacheBreak: cacheBreak,
	})

	var r *GetDestinyManifestResponse
	return r, errors.Join(err, sonic.Unmarshal(res, &r))
}

type GetDestinyManifestResponse struct {
	Response           *Config.DestinyManifest `json:"Response"`
	ErrorCode          int32                   `json:"ErrorCode"`
	ThrottleSeconds    int32                   `json:"ThrottleSeconds"`
	ErrorStatus        string                  `json:"ErrorStatus"`
	Message            string                  `json:"Message"`
	MessageData        map[string]string       `json:"MessageData"`
	DetailedErrorTrace string                  `json:"DetailedErrorTrace"`
}

func (c *Destiny2) GetProfile(membershipId int64, membershipType int32, components ...any) (*GetProfileResponse, error) {
	if len(components) == 0 {
		return nil, fmt.Errorf("no components specified")
	}

	var q url.Values
	var cacheBreak client2.CacheBreak
	for _, component := range components {
		switch component.(type) {
		case DestinyComponentType:
			if q == nil {
				q = url.Values{}
			}
			q.Add("components", fmt.Sprintf("%d,", component))
		case int32:
			if q == nil {
				q = url.Values{}
			}
			q.Add("components", fmt.Sprintf("%d,", component))
		case client2.CacheBreak:
			cacheBreak = component.(client2.CacheBreak)
		}
	}

	// Send and return request
	res, err := c.client.Send(client2.Request{
		Method:      client2.Get,
		BaseURL:     fmt.Sprintf("%s/Destiny2/%d/Profile/%d/", bungieUrl, membershipType, membershipId),
		QueryParams: q,
		CacheBreak:  cacheBreak,
	})

	var r *GetProfileResponse
	return r, errors.Join(err, sonic.Unmarshal(res, &r))
}

type GetProfileResponse struct {
	Response           *Responses.DestinyProfileResponse `json:"Response"`
	ErrorCode          int32                             `json:"ErrorCode"`
	ThrottleSeconds    int32                             `json:"ThrottleSeconds"`
	ErrorStatus        string                            `json:"ErrorStatus"`
	Message            string                            `json:"Message"`
	MessageData        map[string]string                 `json:"MessageData"`
	DetailedErrorTrace string                            `json:"DetailedErrorTrace"`
}

func (c *Destiny2) GetCharacter(membershipType int32, membershipId, characterId int64, components ...any) (*GetCharacterResponse, error) {
	// Ensure at least one component is specified
	if len(components) == 0 {
		return nil, fmt.Errorf("no components specified")
	}

	var q url.Values
	var cacheBreak client2.CacheBreak
	for _, component := range components {
		switch component.(type) {
		case DestinyComponentType:
			if q == nil {
				q = url.Values{}
			}
			q.Add("components", fmt.Sprintf("%d,", component))
		case int32:
			if q == nil {
				q = url.Values{}
			}
			q.Add("components", fmt.Sprintf("%d,", component))
		case client2.CacheBreak:
			cacheBreak = component.(client2.CacheBreak)
		}
	}

	res, err := c.client.Send(client2.Request{
		Method:      client2.Get,
		BaseURL:     fmt.Sprintf("%s/Destiny2/%d/Profile/%d/Character/%d/", bungieUrl, membershipType, membershipId, characterId),
		QueryParams: q,
		CacheBreak:  cacheBreak,
	})

	var r *GetCharacterResponse
	return r, errors.Join(err, sonic.Unmarshal(res, &r))
}

type GetCharacterResponse struct {
	Response           *Responses.DestinyCharacterResponse `json:"Response"`
	ErrorCode          int32                               `json:"ErrorCode"`
	ThrottleSeconds    int32                               `json:"ThrottleSeconds"`
	ErrorStatus        string                              `json:"ErrorStatus"`
	Message            string                              `json:"Message"`
	MessageData        map[string]string                   `json:"MessageData"`
	DetailedErrorTrace string                              `json:"DetailedErrorTrace"`
}
