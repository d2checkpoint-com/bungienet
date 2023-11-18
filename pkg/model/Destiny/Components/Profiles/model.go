package Profiles

import "time"

type DestinyProfileTransitoryComponent struct {
	PartyMembers               []*DestinyProfileTransitoryPartyMember   `json:"partyMembers"`
	CurrentActivity            *DestinyProfileTransitoryCurrentActivity `json:"currentActivity"`
	Joinability                *DestinyProfileTransitoryJoinability     `json:"joinability"`
	Tracking                   []*DestinyProfileTransitoryTrackingEntry `json:"tracking"`
	LastOrbitedDestinationHash uint32                                   `json:"lastOrbitedDestinationHash"` // mapped to Destiny.Definitions.DestinyDestinationDefinition
}

type DestinyProfileTransitoryPartyMember struct {
	MembershipId int64  `json:"membershipId,string"`
	EmblemHash   uint32 `json:"emblemHash"` // mapped to Destiny.Definitions.DestinyInventoryItemDefinition
	DisplayName  string `json:"displayName"`
	Status       int32  `json:"status"`
}

type DestinyProfileTransitoryCurrentActivity struct {
	StartTime                   time.Time `json:"startTime"`
	EndTime                     time.Time `json:"endTime"`
	Score                       float32   `json:"score"`
	HighestOpposingFactionScore float32   `json:"highestOpposingFactionScore"`
	NumberOfOpponents           int32     `json:"numberOfOpponents"`
	NumberOfPlayers             int32     `json:"numberOfPlayers"`
}

type DestinyProfileTransitoryJoinability struct {
	OpenSlots      int32 `json:"openSlots"`
	PrivacySetting int32 `json:"privacySetting"`
	ClosedReasons  int32 `json:"closedReasons"`
}

type DestinyProfileTransitoryTrackingEntry struct {
	LocationHash      uint32    `json:"locationHash"`      // mapped to Destiny.Definitions.DestinyLocationDefinition
	ItemHash          uint32    `json:"itemHash"`          // mapped to Destiny.Definitions.DestinyInventoryItemDefinition
	ObjectiveHash     uint32    `json:"objectiveHash"`     // mapped to Destiny.Definitions.DestinyObjectiveDefinition
	ActivityHash      uint32    `json:"activityHash"`      // mapped to Destiny.Definitions.DestinyActivityDefinition
	QuestlineItemHash uint32    `json:"questlineItemHash"` // mapped to Destiny.Definitions.DestinyInventoryItemDefinition
	TrackedDate       time.Time `json:"trackedDate"`
}
