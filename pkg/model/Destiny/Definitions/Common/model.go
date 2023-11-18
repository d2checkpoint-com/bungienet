package Common

type DestinyDisplayPropertiesDefinition struct {
	Description   string                           `json:"description"`
	Name          string                           `json:"name"`
	Icon          string                           `json:"icon"`
	IconSequences []*DestinyIconSequenceDefinition `json:"iconSequences"`
	HighResIcon   string                           `json:"highResIcon"`
	HasIcon       bool                             `json:"hasIcon"`
}

type DestinyIconSequenceDefinition struct {
	Frames []string `json:"frames"`
}
