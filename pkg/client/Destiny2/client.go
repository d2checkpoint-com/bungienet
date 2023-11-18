package Destiny2

import (
	"encoding/json"
	"fmt"
	"github.com/d2checkpoint-com/bungienet/internal/shared"
	"github.com/d2checkpoint-com/bungienet/pkg/model"
	"github.com/d2checkpoint-com/bungienet/pkg/model/Destiny/Config"
	"github.com/d2checkpoint-com/bungienet/pkg/model/Destiny/Responses"
	"github.com/sendgrid/rest"
	"time"
)

const (
	bungieUrl = "https://www.bungie.net/Platform"
)

type Client struct {
	ApiKey    string
	UserAgent string
}

func NewClient(options ...any) *Client {
	return &Client{
		ApiKey:    options[0].(string),
		UserAgent: shared.SetUserAgent(options),
	}
}

func (c *Client) GetDestinyManifest(options ...any) (*GetDestinyManifestResponse, error) {
	q := make(map[string]string)
	for _, option := range options {
		switch option.(type) {
		case model.CacheBreak:
			q["t"] = fmt.Sprintf("%d", time.Now().Unix())
		}
	}

	// Send and return request
	h, err := rest.Send(rest.Request{
		Method:      rest.Get,
		BaseURL:     bungieUrl + "/Destiny2/Manifest/",
		QueryParams: q,
		Headers: map[string]string{
			"X-API-Key":  c.ApiKey,
			"User-Agent": c.UserAgent,
		},
	})
	if err != nil {
		return nil, err
	}

	var r *GetDestinyManifestResponse
	err = json.Unmarshal([]byte(h.Body), &r)
	if err != nil {
		return nil, err
	}

	return r, shared.CheckResponse(r.ErrorCode, r.ErrorStatus, r.Message)
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

func (c *Client) GetProfile(membershipId int64, membershipType int32, components ...any) (*GetProfileResponse, error) {
	// Ensure at least one component is specified
	if len(components) == 0 {
		return nil, fmt.Errorf("no components specified")
	}

	q := make(map[string]string)
	for _, component := range components {
		switch component.(type) {
		case DestinyComponentType:
			q["components"] += fmt.Sprintf("%d,", component)
		case model.CacheBreak:
			q["t"] = fmt.Sprintf("%d", time.Now().Unix())
		default:
			return nil, fmt.Errorf("invalid component type")
		}
	}
	q["components"] = q["components"][:len(q["components"])-1]

	// Send and return request
	h, err := rest.Send(rest.Request{
		Method:      rest.Get,
		BaseURL:     fmt.Sprintf("%s/Destiny2/%d/Profile/%d/", bungieUrl, membershipType, membershipId),
		QueryParams: q,
		Headers: map[string]string{
			"X-API-Key":  c.ApiKey,
			"User-Agent": c.UserAgent,
		},
	})
	if err != nil {
		return nil, err
	}

	var r *GetProfileResponse
	err = json.Unmarshal([]byte(h.Body), &r)
	if err != nil {
		return nil, err
	}

	return r, shared.CheckResponse(r.ErrorCode, r.ErrorStatus, r.Message)
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

func (c *Client) GetCharacter(membershipType int32, membershipId, characterId int64, components ...any) (*GetCharacterResponse, error) {
	// Ensure at least one component is specified
	if len(components) == 0 {
		return nil, fmt.Errorf("no components specified")
	}

	q := make(map[string]string)
	for _, component := range components {
		switch component.(type) {
		case DestinyComponentType:
			q["components"] += fmt.Sprintf("%d,", component)
		case model.CacheBreak:
			q["t"] = fmt.Sprintf("%d", time.Now().Unix())
		default:
			return nil, fmt.Errorf("invalid component type")
		}
	}
	q["components"] = q["components"][:len(q["components"])-1]

	// Send and return request
	h, err := rest.Send(rest.Request{
		Method:      rest.Get,
		BaseURL:     fmt.Sprintf("%s/Destiny2/%d/Profile/%d/Character/%d/", bungieUrl, membershipType, membershipId, characterId),
		QueryParams: q,
		Headers: map[string]string{
			"X-API-Key":  c.ApiKey,
			"User-Agent": c.UserAgent,
		},
	})
	if err != nil {
		return nil, err
	}

	var r *GetCharacterResponse
	err = json.Unmarshal([]byte(h.Body), &r)
	if err != nil {
		return nil, err
	}

	return r, shared.CheckResponse(r.ErrorCode, r.ErrorStatus, r.Message)
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
