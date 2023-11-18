package Config

type DestinyManifest struct {
	Version                        string                         `json:"version"`
	MobileAssetContentPath         string                         `json:"mobileAssetContentPath"`
	MobileGearAssetDataBases       []*GearAssetDataBaseDefinition `json:"mobileGearAssetDataBases"`
	MobileWorldContentPaths        map[string]string              `json:"mobileWorldContentPaths"`
	JsonWorldContentPaths          map[string]string              `json:"jsonWorldContentPaths"`
	JsonWorldComponentContentPaths map[string]map[string]string   `json:"jsonWorldComponentContentPaths"`
	MobileClanBannerDatabasePath   string                         `json:"mobileClanBannerDatabasePath"`
	MobileGearCDN                  map[string]string              `json:"mobileGearCDN"`
	IconImagePyramidInfo           []*ImagePyramidEntry           `json:"iconImagePyramidInfo"`
}

type GearAssetDataBaseDefinition struct {
	Version int32  `json:"version"`
	Path    string `json:"path"`
}

type ImagePyramidEntry struct {
	Name   string  `json:"name"`
	Factor float32 `json:"factor"`
}
