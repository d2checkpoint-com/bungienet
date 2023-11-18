package Models

import "github.com/d2checkpoint-com/bungienet/pkg/model/User"

type CoreSettingsConfiguration struct {
	Environment                    string                  `json:"environment"`
	Systems                        map[string]*CoreSystem  `json:"systems"`
	IgnoreReasons                  []*CoreSetting          `json:"ignoreReasons"`
	ForumCategories                []*CoreSetting          `json:"forumCategories"`
	GroupAvatars                   []*CoreSetting          `json:"groupAvatars"`
	DefaultGroupTheme              *CoreSetting            `json:"defaultGroupTheme"`
	DestinyMembershipTypes         []*CoreSetting          `json:"destinyMembershipTypes"`
	RecruitmentPlatformTags        []*CoreSetting          `json:"recruitmentPlatformTags"`
	RecruitmentMiscTags            []*CoreSetting          `json:"recruitmentMiscTags"`
	RecruitmentActivities          []*CoreSetting          `json:"recruitmentActivities"`
	UserContentLocales             []*CoreSetting          `json:"userContentLocales"`
	SystemContentLocales           []*CoreSetting          `json:"systemContentLocales"`
	ClanBannerDecals               []*CoreSetting          `json:"clanBannerDecals"`
	ClanBannerDecalColors          []*CoreSetting          `json:"clanBannerDecalColors"`
	ClanBannerGonfalons            []*CoreSetting          `json:"clanBannerGonfalons"`
	ClanBannerGonfalonColors       []*CoreSetting          `json:"clanBannerGonfalonColors"`
	ClanBannerGonfalonDetails      []*CoreSetting          `json:"clanBannerGonfalonDetails"`
	ClanBannerGonfalonDetailColors []*CoreSetting          `json:"clanBannerGonfalonDetailColors"`
	ClanBannerStandards            []*CoreSetting          `json:"clanBannerStandards"`
	Destiny2CoreSettings           []*Destiny2CoreSettings `json:"destiny2CoreSettings"`
	EmailSettings                  *User.EmailSettings     `json:"emailSettings"`
	FireteamActivities             []*CoreSetting          `json:"fireteamActivities"`
}

type CoreSystem struct {
	Enabled    bool              `json:"enabled"`
	Parameters map[string]string `json:"parameters"`
}

type CoreSetting struct {
	Identifier    string         `json:"identifier"`
	IsDefault     bool           `json:"isDefault"`
	DisplayName   string         `json:"displayName"`
	Summary       string         `json:"summary"`
	ImagePath     string         `json:"imagePath"`
	ChildSettings []*CoreSetting `json:"childSettings"`
}

type Destiny2CoreSettings struct {
	CollectionRootNode                     uint32   `json:"collectionRootNode"`                    // Destiny.Definitions.Presentation.DestinyPresentationNodeDefinition
	BadgesRootNode                         uint32   `json:"badgesRootNode"`                        // Destiny.Definitions.Presentation.DestinyPresentationNodeDefinition
	RecordsRootNode                        uint32   `json:"recordsRootNode"`                       // Destiny.Definitions.Presentation.DestinyPresentationNodeDefinition
	MedalsRootNode                         uint32   `json:"medalsRootNode"`                        // Destiny.Definitions.Presentation.DestinyPresentationNodeDefinition
	MetricsRootNode                        uint32   `json:"metricsRootNode"`                       // Destiny.Definitions.Presentation.DestinyPresentationNodeDefinition
	ActiveTriumphsRootNodeHash             uint32   `json:"activeTriumphsRootNodeHash"`            // Destiny.Definitions.Presentation.DestinyPresentationNodeDefinition
	ActiveSealsRootNodeHash                uint32   `json:"activeSealsRootNodeHash"`               // Destiny.Definitions.Presentation.DestinyPresentationNodeDefinition
	LegacyTriumphsRootNodeHash             uint32   `json:"legacyTriumphsRootNodeHash"`            // Destiny.Definitions.Presentation.DestinyPresentationNodeDefinition
	LegacySealsRootNodeHash                uint32   `json:"legacySealsRootNodeHash"`               // Destiny.Definitions.Presentation.DestinyPresentationNodeDefinition
	MedalsRootNodeHash                     uint32   `json:"medalsRootNodeHash"`                    // Destiny.Definitions.Presentation.DestinyPresentationNodeDefinition
	ExoticCatalystsRootNodeHash            uint32   `json:"exoticCatalystsRootNodeHash"`           // Destiny.Definitions.Presentation.DestinyPresentationNodeDefinition
	LoreRootNodeHash                       uint32   `json:"loreRootNodeHash"`                      // Destiny.Definitions.Presentation.DestinyPresentationNodeDefinition
	CraftingRootNodeHash                   uint32   `json:"craftingRootNodeHash"`                  // Destiny.Definitions.Presentation.DestinyPresentationNodeDefinition
	LoadoutConstantsHash                   uint32   `json:"loadoutConstantsHash"`                  //  Destiny.Definitions.Loadouts.DestinyLoadoutConstantsDefinition
	GuardianRankConstantsHash              uint32   `json:"guardianRankConstantsHash"`             //  Destiny.Definitions.GuardianRanks.DestinyGuardianRankConstantsDefinition
	GuardianRanksRootNodeHash              uint32   `json:"guardianRanksRootNodeHash"`             // Destiny.Definitions.Presentation.DestinyPresentationNodeDefinition
	CurrentRankProgressionHashes           []uint32 `json:"currentRankProgressionHashes"`          // Destiny.Definitions.Progression.DestinyProgressionDefinition
	InsertPlugFreeProtectedPlugItemHashes  []uint32 `json:"insertPlugFreeProtectedPlugItemHashes"` // Destiny.Definitions.Items.DestinyItemDefinition
	InsertPlugFreeBlockedSocketTypeHashes  []uint32 `json:"insertPlugFreeBlockedSocketTypeHashes"` // Destiny.Definitions.Sockets.DestinySocketTypeDefinition
	UndiscoveredCollectibleImage           string   `json:"undiscoveredCollectibleImage"`
	AmmoTypeHeavyIcon                      string   `json:"ammoTypeHeavyIcon"`
	AmmoTypeSpecialIcon                    string   `json:"ammoTypeSpecialIcon"`
	AmmoTypePrimaryIcon                    string   `json:"ammoTypePrimaryIcon"`
	CurrentSeasonalArtifactHash            uint32   `json:"currentSeasonalArtifactHash"`            // Destiny.Definitions.DestinyVendorDefinition
	CurrentSeasonHash                      uint32   `json:"currentSeasonHash"`                      // Destiny.Definitions.Seasons.DestinySeasonDefinition
	SeasonalChallengesPresentationNodeHash uint32   `json:"seasonalChallengesPresentationNodeHash"` // Destiny.Definitions.Presentation.DestinyPresentationNodeDefinition
	FutureSeasonHashes                     []uint32 `json:"futureSeasonHashes"`                     // Destiny.Definitions.Seasons.DestinySeasonDefinition
	PastSeasonHashes                       []uint32 `json:"pastSeasonHashes"`                       // Destiny.Definitions.Seasons.DestinySeasonDefinition
}
