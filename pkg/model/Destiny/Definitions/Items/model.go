package Items

type DestinyDerivedItemDefinition struct {
	ItemHash        uint32 `json:"itemHash"`
	ItemName        string `json:"itemName"`
	ItemDetail      string `json:"itemDetail"`
	ItemDescription string `json:"itemDescription"`
	IconPath        string `json:"iconPath"`
	VendorItemIndex int32  `json:"vendorItemIndex"`
}

type DestinyDerivedItemCategoryDefinition struct {
	CategoryDescription string                          `json:"categoryDescription"`
	Items               []*DestinyDerivedItemDefinition `json:"items"`
}

type DestinyPlugRuleDefinition struct {
	FailureMessage string `json:"failureMessage"`
}

type DestinyParentItemOverride struct {
	AdditionalEquipRequirementsDisplayStrings []string `json:"additionalEquipRequirementsDisplayStrings"`
	PipIcon                                   string   `json:"pipIcon"`
}

type DestinyEnergyCapacityEntry struct {
	CapacityValue  int32  `json:"capacityValue"`
	EnergyTypeHash uint32 `json:"energyTypeHash"`
	EnergyType     int32  `json:"energyType"`
}

type DestinyEnergyCostEntry struct {
	EnergyCost     int32  `json:"energyCost"`
	EnergyTypeHash uint32 `json:"energyTypeHash"`
	EnergyType     int32  `json:"energyType"`
}
