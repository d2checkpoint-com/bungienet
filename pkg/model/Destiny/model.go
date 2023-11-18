package Destiny

import "github.com/d2checkpoint-com/bungienet/pkg/model/Destiny/Challenges"

type DestinyActivity struct {
	ActivityHash            uint32                               `json:"activityHash"` // mapped to Destiny.Definitions.DestinyActivityDefinition
	IsNew                   bool                                 `json:"isNew"`
	CanLead                 bool                                 `json:"canLead"`
	CanJoin                 bool                                 `json:"canJoin"`
	IsCompleted             bool                                 `json:"isCompleted"`
	IsVisible               bool                                 `json:"isVisible"`
	DisplayLevel            int32                                `json:"displayLevel"`
	RecommendedLight        int32                                `json:"recommendedLight"`
	DifficultyTier          int32                                `json:"difficultyTier"`
	Challenges              []*Challenges.DestinyChallengeStatus `json:"challenges"`
	ModifierHashes          []uint32                             `json:"modifierHashes"` // mapped to Destiny.Definitions.ActivityModifiers.DestinyActivityModifierDefinition
	BooleanActivityOptions  map[uint32]bool                      `json:"booleanActivityOptions"`
	LoadoutRequirementIndex int32                                `json:"loadoutRequirementIndex"`
}

type DestinyItemQuantity struct {
	ItemHash                 uint32 `json:"itemHash"` // mapped to Destiny.Definitions.DestinyInventoryItemDefinition
	ItemInstanceId           int64  `json:"itemInstanceId"`
	Quantity                 int32  `json:"quantity"`
	HasConditionalVisibility bool   `json:"hasConditionalVisibility"`
}

type DyeReference struct {
	ChannelHash uint32 `json:"channelHash"`
	DyeHash     uint32 `json:"dyeHash"`
}
