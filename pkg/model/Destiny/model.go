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
	ItemInstanceId           int64  `json:"itemInstanceId,string"`
	Quantity                 int32  `json:"quantity"`
	HasConditionalVisibility bool   `json:"hasConditionalVisibility"`
}

type DyeReference struct {
	ChannelHash uint32 `json:"channelHash"`
	DyeHash     uint32 `json:"dyeHash"`
}

type DestinyProgressionResetEntry struct {
	Season int32 `json:"season"`
	Resets int32 `json:"resets"`
}

type DestinyProgression struct {
	ProgressionHash     uint32                          `json:"progressionHash"` // mapped to Destiny.Definitions.DestinyProgressionDefinition
	DailyProgress       int32                           `json:"dailyProgress"`
	DailyLimit          int32                           `json:"dailyLimit"`
	WeeklyProgress      int32                           `json:"weeklyProgress"`
	WeeklyLimit         int32                           `json:"weeklyLimit"`
	CurrentProgress     int32                           `json:"currentProgress"`
	Level               int32                           `json:"level"`
	LevelCap            int32                           `json:"levelCap"`
	StepIndex           int32                           `json:"stepIndex"`
	ProgressToNextLevel int32                           `json:"progressToNextLevel"`
	NextLevelAt         int32                           `json:"nextLevelAt"`
	CurrentResetCount   int32                           `json:"currentResetCount"`
	SeasonResets        []*DestinyProgressionResetEntry `json:"seasonResets"`
	RewardItemStates    []int32                         `json:"rewardItemStates"`
}
