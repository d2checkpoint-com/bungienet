package User

import (
	"github.com/d2checkpoint-com/bungienet/pkg/model/GroupsV2"
	"github.com/d2checkpoint-com/bungienet/pkg/model/Ignores"
	"time"
)

type EmailSettings struct {
	OptInDefinitions        map[string]*EmailOptInDefinition        `json:"optInDefinitions"`
	SubscriptionDefinitions map[string]*EmailSubscriptionDefinition `json:"subscriptionDefinitions"`
	Views                   map[string]*EmailViewDefinition         `json:"views"`
}

type EmailOptInDefinition struct {
	Name                   string                         `json:"name"`
	Value                  int64                          `json:"value"`
	SetByDefault           bool                           `json:"setByDefault"`
	DependentSubscriptions []*EmailSubscriptionDefinition `json:"dependentSubscriptions"`
}

type EmailSubscriptionDefinition struct {
	Name         string                                           `json:"name"`
	Localization map[string]*EMailSettingSubscriptionLocalization `json:"localization"`
	Value        int64                                            `json:"value"`
}

type EmailViewDefinition struct {
	Name         string                        `json:"name"`
	ViewSettings []*EmailViewDefinitionSetting `json:"viewSettings"`
}

type EmailViewDefinitionSetting struct {
	Name                string                               `json:"name"`
	Localization        map[string]*EMailSettingLocalization `json:"localization"`
	SetByDefault        bool                                 `json:"setByDefault"`
	OptInAggregateValue int64                                `json:"optInAggregateValue"`
	Subscription        []*EmailSubscriptionDefinition       `json:"subscription"`
}

type EMailSettingLocalization struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type EMailSettingSubscriptionLocalization struct {
	UnknownUserDescription      string `json:"unknownUserDescription"`
	RegisteredUserDescription   string `json:"registeredUserDescription"`
	UnregisteredUserDescription string `json:"unregisteredUserDescription"`
	UnknownUserActionText       string `json:"unknownUserActionText"`
	KnownUserActionText         string `json:"knownUserActionText"`
	Title                       string `json:"title"`
	Description                 string `json:"description"`
}

type GeneralUser struct {
	MembershipId                      int64              `json:"membershipId"`
	UniqueName                        string             `json:"uniqueName"`
	NormalizedName                    string             `json:"normalizedName"`
	DisplayName                       string             `json:"displayName"`
	ProfilePicture                    string             `json:"profilePicture"`
	ProfileTheme                      int32              `json:"profileTheme"`
	UserTitle                         int32              `json:"userTitle"`
	SuccessMessageFlags               int32              `json:"successMessageFlags"`
	IsDeleted                         bool               `json:"isDeleted"`
	About                             string             `json:"about"`
	FirstAccess                       int64              `json:"firstAccess"`
	LastUpdate                        int64              `json:"lastUpdate"`
	LegacyPortalUID                   int64              `json:"legacyPortalUID"`
	Context                           *UserToUserContext `json:"context"`
	PsnDisplayName                    string             `json:"psnDisplayName"`
	XboxDisplayName                   string             `json:"xboxDisplayName"`
	FbDisplayName                     string             `json:"fbDisplayName"`
	ShowActivity                      bool               `json:"showActivity"`
	Locale                            string             `json:"locale"`
	LocaleInheritDefault              bool               `json:"localeInheritDefault"`
	LastBanReportId                   int64              `json:"lastBanReportId"`
	ShowGroupMessaging                bool               `json:"showGroupMessaging"`
	ProfilePicturePath                string             `json:"profilePicturePath"`
	ProfilePictureWidePath            string             `json:"profilePictureWidePath"`
	ProfileThemeName                  string             `json:"profileThemeName"`
	UserTitleDisplay                  string             `json:"userTitleDisplay"`
	StatusText                        string             `json:"statusText"`
	StatusDate                        time.Time          `json:"statusDate"`
	ProfileBanExpire                  time.Time          `json:"profileBanExpire"`
	BlizzardDisplayName               string             `json:"blizzardDisplayName"`
	SteamDisplayName                  string             `json:"steamDisplayName"`
	StadiaDisplayName                 string             `json:"stadiaDisplayName"`
	TwitchDisplayName                 string             `json:"twitchDisplayName"`
	CachedBungieGlobalDisplayName     string             `json:"cachedBungieGlobalDisplayName"`
	CachedBungieGlobalDisplayNameCode int16              `json:"cachedBungieGlobalDisplayNameCode"`
	EgsDisplayName                    string             `json:"egsDisplayName"`
}

type UserToUserContext struct {
	IsFollowing         bool                    `json:"isFollowing"`
	IgnoreStatus        *Ignores.IgnoreResponse `json:"ignoreStatus"`
	GlobalIgnoreEndDate int                     `json:"globalIgnoreEndDate"`
}

type UserMembershipData struct {
	DestinyMemberships  []*GroupsV2.GroupUserInfoCard `json:"destinyMemberships"`
	PrimaryMembershipId int64                         `json:"primaryMembershipId"`
	BungieNetUser       *GeneralUser                  `json:"bungieNetUser"`
}
