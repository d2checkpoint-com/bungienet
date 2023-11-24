package User

import (
	"errors"
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/d2checkpoint-com/bungienet/pkg/client"
	modelUser "github.com/d2checkpoint-com/bungienet/pkg/model/User"
)

const (
	bungieUrl = "https://www.bungie.net/Platform"
)

type User struct {
	client *client.Client
}

func NewClient() *User {
	return &User{&client.BungieClient}
}

func (c *User) GetMembershipDataById(membershipId int64, membershipType int32, options ...any) (*GetMembershipDataByIdResponse, error) {
	var cacheBreak client.CacheBreak
	for _, option := range options {
		switch option.(type) {
		case client.CacheBreak:
			cacheBreak = option.(client.CacheBreak)
		}
	}

	res, err := c.client.Send(client.Request{
		Method:     client.Get,
		BaseURL:    fmt.Sprintf("%s/User/GetMembershipsById/%d/%d/", bungieUrl, membershipId, membershipType),
		CacheBreak: cacheBreak,
	})

	var r *GetMembershipDataByIdResponse
	return r, errors.Join(err, sonic.Unmarshal(res, &r))
}

type GetMembershipDataByIdResponse struct {
	Response           *modelUser.UserMembershipData `json:"Response"`
	ErrorCode          int32                         `json:"ErrorCode"`
	ThrottleSeconds    int32                         `json:"ThrottleSeconds"`
	ErrorStatus        string                        `json:"ErrorStatus"`
	Message            string                        `json:"Message"`
	MessageData        map[string]string             `json:"MessageData"`
	DetailedErrorTrace string                        `json:"DetailedErrorTrace"`
}
