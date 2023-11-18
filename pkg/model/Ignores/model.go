package Ignores

type IgnoreResponse struct {
	IsIgnored   bool  `json:"isIgnored"`
	IgnoreFlags int32 `json:"ignoreFlags"`
}
