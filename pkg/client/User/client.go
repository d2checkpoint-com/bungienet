package User

import (
	"encoding/json"
	"fmt"
	"github.com/d2checkpoint-com/bungienet/internal/shared"
	"github.com/d2checkpoint-com/bungienet/pkg/model"
	"github.com/d2checkpoint-com/bungienet/pkg/model/User"
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

func (c *Client) GetMembershipDataById(membershipId int64, membershipType int32, options ...any) (*GetMembershipDataByIdResponse, error) {
	q := make(map[string]string)
	for _, option := range options {
		switch option.(type) {
		case model.CacheBreak:
			q["t"] = fmt.Sprintf("%d", time.Now().Unix())
		}
	}
	h, err := rest.Send(rest.Request{
		Method:      rest.Get,
		BaseURL:     fmt.Sprintf("%s/User/GetMembershipsById/%d/%d/", bungieUrl, membershipId, membershipType),
		QueryParams: q,
		Headers: map[string]string{
			"X-API-Key":  c.ApiKey,
			"User-Agent": c.UserAgent,
		},
	})
	if err != nil {
		return nil, err
	}

	var r *GetMembershipDataByIdResponse
	err = json.Unmarshal([]byte(h.Body), &r)
	if err != nil {
		return nil, err
	}

	return r, shared.CheckResponse(r.ErrorCode, r.ErrorStatus, r.Message)
}

type GetMembershipDataByIdResponse struct {
	Response           *User.UserMembershipData `json:"Response"`
	ErrorCode          int32                    `json:"ErrorCode"`
	ThrottleSeconds    int32                    `json:"ThrottleSeconds"`
	ErrorStatus        string                   `json:"ErrorStatus"`
	Message            string                   `json:"Message"`
	MessageData        map[string]string        `json:"MessageData"`
	DetailedErrorTrace string                   `json:"DetailedErrorTrace"`
}
