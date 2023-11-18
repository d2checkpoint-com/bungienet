package Definitions

import (
	"github.com/d2checkpoint-com/bungienet/pkg/model/Destiny"
	"github.com/d2checkpoint-com/bungienet/pkg/model/Destiny/Constants"
	"github.com/d2checkpoint-com/bungienet/pkg/model/Destiny/Definitions/Common"
)

type DestinationHash interface {
	string
	int32
}

type DestinyActivityDefinition struct {
	DisplayProperties                *Common.DestinyDisplayPropertiesDefinition     `json:"displayProperties"`
	OriginalDisplayProperties        *Common.DestinyDisplayPropertiesDefinition     `json:"originalDisplayProperties"`
	SelectionScreenDisplayProperties *Common.DestinyDisplayPropertiesDefinition     `json:"selectionScreenDisplayProperties"`
	ReleaseIcon                      string                                         `json:"releaseIcon"`
	ReleaseTime                      int32                                          `json:"releaseTime"`
	ActivityLightLevel               int32                                          `json:"activityLightLevel"`
	DestinationHash                  any                                            `json:"destinationHash"`  // Mapped to Destiny.Definitions.DestinyDestinationDefinition // Errata: This is an int32 however one entry is a string
	PlaceHase                        uint32                                         `json:"placeHash"`        // Mapped to Destiny.Definitions.DestinyPlaceDefinition
	ActivityTypeHash                 uint32                                         `json:"activityTypeHash"` // Mapped to Destiny.Definitions.DestinyActivityTypeDefinition
	Tier                             int32                                          `json:"tier"`
	PgcrImage                        string                                         `json:"pgcrImage"`
	Rewards                          []*DestinyActivityRewardDefinition             `json:"rewards"`
	Modifiers                        []*DestinyActivityModifierReferenceDefinition  `json:"modifiers"`
	IsPlaylist                       bool                                           `json:"isPlaylist"`
	Challenges                       []*DestinyActivityChallengeDefinition          `json:"challenges"`
	OptionalUnlockStrings            []*DestinyActivityUnlockStringDefinition       `json:"optionalUnlockStrings"`
	PlaylistItems                    []*DestinyActivityPlaylistItemDefinition       `json:"playlistItems"`
	ActivityGraphList                []*DestinyActivityGraphListEntryDefinition     `json:"activityGraphList"`
	Matchmaking                      *DestinyActivityMatchmakingBlockDefinition     `json:"matchmaking"`
	GuidedGame                       *DestinyActivityGuidedBlockDefinition          `json:"guidedGame"`
	DirectActivityModeHash           uint32                                         `json:"directActivityModeHash"` // Mapped to Destiny.Definitions.DestinyActivityModeDefinition
	DirectActivityModeType           int32                                          `json:"directActivityModeType"`
	Loadouts                         []*DestinyActivityLoadoutRequirementSet        `json:"loadouts"`
	ActivityModeHashes               []uint32                                       `json:"activityModeHashes"` // Mapped to Destiny.Definitions.DestinyActivityModeDefinition
	ActivityModeTypes                []int32                                        `json:"activityModeTypes"`
	IsPvP                            bool                                           `json:"isPvP"`
	InsertionPoints                  []*DestinyActivityInsertionPointDefinition     `json:"insertionPoints"`
	ActivityLocationMappings         []*Constants.DestinyEnvironmentLocationMapping `json:"activityLocationMappings"`
	Hash                             uint32                                         `json:"hash"`
	Index                            int32                                          `json:"index"`
	Redacted                         bool                                           `json:"redacted"`
}

type DestinyActivityRewardDefinition struct {
	RewardText  string                        `json:"rewardText"`
	RewardItems []Destiny.DestinyItemQuantity `json:"rewardItems"`
}

type DestinyActivityModifierReferenceDefinition struct {
	ActivityModifierHash uint32 `json:"activityModifierHash"` // Mapped to Destiny.Definitions.DestinyActivityModifierDefinition
}

type DestinyActivityChallengeDefinition struct {
	ObjectiveHash uint32                         `json:"objectiveHash"` // Mapped to Destiny.Definitions.DestinyObjectiveDefinition
	DummyRewards  []*Destiny.DestinyItemQuantity `json:"dummyRewards"`
}

type DestinyActivityUnlockStringDefinition struct {
	DisplayString string `json:"displayString"`
}

type DestinyActivityPlaylistItemDefinition struct {
	ActivityHash           uint32   `json:"activityHash"`           // Mapped to Destiny.Definitions.DestinyActivityDefinition
	DirectActivityModeHash uint32   `json:"directActivityModeHash"` // Mapped to Destiny.Definitions.DestinyActivityModeDefinition
	DirectActivityModeType int32    `json:"directActivityModeType"`
	ActivityModeHashes     []uint32 `json:"activityModeHashes"` // Mapped to Destiny.Definitions.DestinyActivityModeDefinition
	ActivityModeTypes      []int32  `json:"activityModeTypes"`
}

type DestinyActivityGraphListEntryDefinition struct {
	ActivityGraphHash uint32 `json:"activityGraphHash"` // Mapped to Destiny.Definitions.DestinyActivityGraphDefinition
}

type DestinyActivityMatchmakingBlockDefinition struct {
	IsMatchmade          bool  `json:"isMatchmade"`
	MinParty             int32 `json:"minParty"`
	MaxParty             int32 `json:"maxParty"`
	MaxPlayers           int32 `json:"maxPlayers"`
	RequiresGuardianOath bool  `json:"requiresGuardianOath"`
}

type DestinyActivityGuidedBlockDefinition struct {
	GuidedMaxLobbySize int32 `json:"guidedMaxLobbySize"`
	GuidedMinLobbySize int32 `json:"guidedMinLobbySize"`
	GuidedDisbandCount int32 `json:"guidedDisbandCount"`
}

type DestinyActivityLoadoutRequirementSet struct {
	Requirements []*DestinyActivityLoadoutRequirement `json:"requirements"`
}

type DestinyActivityLoadoutRequirement struct {
	EquipmentSlotHash         uint32   `json:"equipmentSlotHash"`         // Mapped to Destiny.Definitions.DestinyEquipmentSlotDefinition
	AllowedEquippedItemHashes []uint32 `json:"allowedEquippedItemHashes"` // Mapped to Destiny.Definitions.DestinyInventoryItemDefinition
	AllowedWeaponSubTypes     []int32  `json:"allowedWeaponSubTypes"`
}

type DestinyActivityInsertionPointDefinition struct {
	PhaseHash uint32 `json:"phaseHash"`
}
