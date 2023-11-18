package Constants

type DestinyEnvironmentLocationMapping struct {
	LocationHash     uint32 `json:"locationHash"` // Mapped to Destiny.Definitions.DestinyLocationDefinition
	ActivationSource string `json:"activationSource"`
	ItemHash         uint32 `json:"itemHash"`      // Mapped to Destiny.Definitions.DestinyInventoryItemDefinition
	ObjectiveHash    uint32 `json:"objectiveHash"` // Mapped to Destiny.Definitions.DestinyObjectiveDefinition
	ActivityHash     uint32 `json:"activityHash"`  // Mapped to Destiny.Definitions.DestinyActivityDefinition
}
