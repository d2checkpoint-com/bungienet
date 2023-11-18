package Quests

type DestinyObjectiveProgress struct {
	ObjectiveHash   uint32 `json:"objectiveHash"`   // mapped to Destiny.Definitions.DestinyObjectiveDefinition
	DestinationHash uint32 `json:"destinationHash"` // mapped to Destiny.Definitions.DestinyDestinationDefinition
	ActivityHash    uint32 `json:"activityHash"`    // mapped to Destiny.Definitions.DestinyActivityDefinition
	Progress        int32  `json:"progress"`
	CompletionValue int32  `json:"completionValue"`
	Complete        bool   `json:"complete"`
	Visible         bool   `json:"visible"`
}
