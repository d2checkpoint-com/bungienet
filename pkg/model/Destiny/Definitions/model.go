package Definitions

import (
	"github.com/d2checkpoint-com/bungienet/pkg/model/Destiny"
	"github.com/d2checkpoint-com/bungienet/pkg/model/Destiny/Constants"
	"github.com/d2checkpoint-com/bungienet/pkg/model/Destiny/Definitions/Animations"
	"github.com/d2checkpoint-com/bungienet/pkg/model/Destiny/Definitions/Common"
	"github.com/d2checkpoint-com/bungienet/pkg/model/Destiny/Definitions/Items"
	"github.com/d2checkpoint-com/bungienet/pkg/model/Destiny/Misc"
	"github.com/d2checkpoint-com/bungienet/pkg/model/Links"
)

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

type DestinyItemTooltipNotification struct {
	DisplayString string `json:"displayString"`
	DisplayStyle  string `json:"displayStyle"`
}

type DestinyItemActionRequiredItemDefinition struct {
	Count          int32  `json:"count"`
	ItemHash       uint32 `json:"itemHash"` // Mapped to Destiny.Definitions.DestinyInventoryItemDefinition
	DeleteOnAction bool   `json:"deleteOnAction"`
}

type DestinyProgressionRewardDefinition struct {
	ProgressionMappingHash uint32 `json:"progressionMappingHash"` // Mapped to Destiny.Definitions.DestinyProgressionMappingDefinition
	Amount                 int32  `json:"amount"`
	ApplyThrottles         bool   `json:"applyThrottles"`
}

type DestinyItemActionBlockDefinition struct {
	VerbName                string                                     `json:"verbName"`
	VerbDescription         string                                     `json:"verbDescription"`
	IsPositive              bool                                       `json:"isPositive"`
	OverlayScreenName       string                                     `json:"overlayScreenName"`
	OverlayIcon             string                                     `json:"overlayIcon"`
	RequiredCooldownSeconds int32                                      `json:"requiredCooldownSeconds"`
	RequiredItems           []*DestinyItemActionRequiredItemDefinition `json:"requiredItems"`
	ProgressionRewards      []*DestinyProgressionRewardDefinition      `json:"progressionRewards"`
	ActionTypeLabel         string                                     `json:"actionTypeLabel"`
	RequiredLocation        string                                     `json:"requiredLocation"`
	RequiredCooldownHash    uint32                                     `json:"requiredCooldownHash"` // Mapped to Destiny.Definitions.DestinyInventoryItemDefinition
	DeleteOnAction          bool                                       `json:"deleteOnAction"`
	ConsumeEntireStack      bool                                       `json:"consumeEntireStack"`
	UseOnAcquire            bool                                       `json:"useOnAcquire"`
}

type DestinyItemCraftingBlockBonusPlugDefinition struct {
	SocketTypeHash uint32 `json:"socketTypeHash"` // Mapped to Destiny.Definitions.Sockets.DestinySocketTypeDefinition
	PlugItemHash   uint32 `json:"plugItemHash"`   // Mapped to Destiny.Definitions.DestinyInventoryItemDefinition
}

type DestinyItemCraftingBlockDefinition struct {
	OutputItemHash           uint32                                         `json:"outputItemHash"`           // Mapped to Destiny.Definitions.DestinyInventoryItemDefinition
	RequiredSocketTypeHashes []uint32                                       `json:"requiredSocketTypeHashes"` // Mapped to Destiny.Definitions.Sockets.DestinySocketTypeDefinition
	FailedRequirementStrings []string                                       `json:"failedRequirementStrings"`
	BaseMaterialRequirements uint32                                         `json:"baseMaterialRequirements"` // Mapped to Destiny.Definitions.DestinyMaterialRequirementSetDefinition
	BonusPlugs               []*DestinyItemCraftingBlockBonusPlugDefinition `json:"bonusPlugs"`
}

type DestinyItemInventoryBlockDefinition struct {
	StackUniqueLabel                         string `json:"stackUniqueLabel"`
	MaxStackSize                             int32  `json:"maxStackSize"`
	BucketTypeHash                           uint32 `json:"bucketTypeHash"`         // Mapped to Destiny.Definitions.DestinyInventoryBucketDefinition
	RecoveryBucketTypeHash                   uint32 `json:"recoveryBucketTypeHash"` // Mapped to Destiny.Definitions.DestinyInventoryBucketDefinition
	TierTypeHash                             uint32 `json:"tierTypeHash"`           // Mapped to Destiny.Definitions.DestinyItemTierTypeDefinition
	IsInstanceItem                           bool   `json:"isInstanceItem"`
	TierTypeName                             string `json:"tierTypeName"`
	TierType                                 int32  `json:"tierType"`
	ExpirationTooltip                        string `json:"expirationTooltip"`
	ExpiredInActivityMessage                 string `json:"expiredInActivityMessage"`
	ExpiredInOrbitMessage                    string `json:"expiredInOrbitMessage"`
	SuppressExpirationWhenObjectivesComplete bool   `json:"suppressExpirationWhenObjectivesComplete"`
	RecipeItemHash                           uint32 `json:"recipeItemHash"` // Mapped to Destiny.Definitions.DestinyInventoryItemDefinition
}

type DestinyItemSetBlockEntryDefinition struct {
	TrackingValue int32  `json:"trackingValue"`
	ItemHash      uint32 `json:"itemHash"` // Mapped to Destiny.Definitions.DestinyInventoryItemDefinition
}

type DestinyItemSetBlockDefinition struct {
	ItemList                 []*DestinyItemSetBlockEntryDefinition `json:"itemList"`
	RequireOrderedSetItemAdd bool                                  `json:"requireOrderedSetItemAdd"`
	SetIsFeatured            bool                                  `json:"setIsFeatured"`
	SetType                  string                                `json:"setType"`
	QuestLineName            string                                `json:"questLineName"`
	QuestLineDescription     string                                `json:"questLineDescription"`
	QuestStepSummary         string                                `json:"questStepSummary"`
}

type DestinyInventoryItemStatDefinition struct {
	StatHash       uint32 `json:"statHash"` // Mapped to Destiny.Definitions.DestinyStatDefinition
	Value          int32  `json:"value"`
	Minimum        int32  `json:"minimum"`
	Maximum        int32  `json:"maximum"`
	DisplayMaximum int32  `json:"displayMaximum"`
}

type DestinyItemStatBlockDefinition struct {
	DisablePrimaryStatDisplay bool                                           `json:"disablePrimaryStatDisplay"`
	StatGroupHash             uint32                                         `json:"statGroupHash"` // Mapped to Destiny.Definitions.DestinyStatGroupDefinition
	Stats                     map[uint32]*DestinyInventoryItemStatDefinition `json:"stats"`
	HasDisplayableStats       bool                                           `json:"hasDisplayableStats"`
	PrimaryBaseStatHash       uint32                                         `json:"primaryBaseStatHash"` // Mapped to Destiny.Definitions.DestinyStatDefinition
}

type DestinyEquippingBlockDefinition struct {
	GearsetItemHash       uint32   `json:"gearsetItemHash"` // Mapped to Destiny.Definitions.DestinyInventoryItemDefinition
	UniqueLabel           string   `json:"uniqueLabel"`
	UniqueLabelHash       uint32   `json:"uniqueLabelHash"`       // Mapped to Destiny.Definitions.DestinyInventoryItemDefinition
	EquipmentSlotTypeHash uint32   `json:"equipmentSlotTypeHash"` // Mapped to Destiny.Definitions.DestinyEquipmentSlotDefinition
	Attributes            int32    `json:"attributes"`
	AmmoType              int32    `json:"ammoType"`
	DisplayStrings        []string `json:"displayStrings"`
}

type DestinyGearArtArrangementReference struct {
	ClassHash          uint32 `json:"classHash"` // Mapped to Destiny.Definitions.DestinyClassDefinition
	ArtArrangementHash uint32 `json:"artArrangementHash"`
}

type DestinyItemTranslationBlockDefinition struct {
	WeaponPatternIdentifier string                                `json:"weaponPatternIdentifier"`
	WeaponPatternHash       uint32                                `json:"weaponPatternHash"` // Mapped to Destiny.Definitions.DestinyWeaponPatternDefinition
	DefaultDyes             []*Destiny.DyeReference               `json:"defaultDyes"`
	LockedDyes              []*Destiny.DyeReference               `json:"lockedDyes"`
	CustomDyes              []*Destiny.DyeReference               `json:"customDyes"`
	Arrangements            []*DestinyGearArtArrangementReference `json:"arrangements"`
	HasGeometry             bool                                  `json:"hasGeometry"`
}

type DestinyItemPreviewBlockDefinition struct {
	ScreenStyle           string                                        `json:"screenStyle"`
	PreviewVendorHash     uint32                                        `json:"previewVendorHash"` // Mapped to Destiny.Definitions.DestinyVendorDefinition
	ArtifactHash          uint32                                        `json:"artifactHash"`      // Mapped to Destiny.Definitions.DestinyArtifactDefinition
	PreviewActionString   string                                        `json:"previewActionString"`
	DerivedItemCategories []*Items.DestinyDerivedItemCategoryDefinition `json:"derivedItemCategories"`
}

type DestinyItemVersionDefinition struct {
	PowerCapHash uint32 `json:"powerCapHash"` // Mapped to Destiny.Definitions.DestinyPowerCapDefinition
}

type DestinyItemQualityBlockDefinition struct {
	ItemLevels                      []int32                         `json:"itemLevels"`
	QualityLevel                    int32                           `json:"qualityLevel"`
	InfusionCategoryName            string                          `json:"infusionCategoryName"`
	InfusionCategoryHash            uint32                          `json:"infusionCategoryHash"`            // Mapped to Destiny.Definitions.DestinyInfusionCategoryDefinition
	InfusionCategoryHashes          []uint32                        `json:"infusionCategoryHashes"`          // Mapped to Destiny.Definitions.DestinyInfusionCategoryDefinition
	ProgressionLevelRequirementHash uint32                          `json:"progressionLevelRequirementHash"` // Mapped to Destiny.Definitions.DestinyProgressionLevelRequirementDefinition
	CurrentVersion                  int32                           `json:"currentVersion"`
	Versions                        []*DestinyItemVersionDefinition `json:"versions"`
	DisplayVersionWatermarkIcons    []string                        `json:"displayVersionWatermarkIcons"`
}

type DestinyItemValueBlockDefinition struct {
	ItemValue        []*Destiny.DestinyItemQuantity `json:"itemValue"`
	ValueDescription string                         `json:"valueDescription"`
}

type DestinyItemVendorSourceReference struct {
	VendorHash        uint32  `json:"vendorHash"` // Mapped to Destiny.Definitions.DestinyVendorDefinition
	VendorItemIndexes []int32 `json:"vendorItemIndexes"`
}

type DestinyItemSourceBlockDefinition struct {
	SourceHashes  []uint32                              `json:"sourceHashes"` // Mapped to Destiny.Definitions.DestinyInventoryItemDefinition
	Sources       []*SourcesDestinyItemSourceDefinition `json:"sources"`
	Exclusive     int32                                 `json:"exclusive"`
	VendorSources []*DestinyItemVendorSourceReference
}

type DestinyObjectiveDisplayProperties struct {
	ActivityHash               uint32 `json:"activityHash"` // Mapped to Destiny.Definitions.DestinyActivityDefinition
	DisplayOnItemPreviewScreen bool   `json:"displayOnItemPreviewScreen"`
}

type DestinyItemObjectiveBlockDefinition struct {
	ObjectiveHashes                []uint32                             `json:"objectiveHashes"`       // Mapped to Destiny.Definitions.DestinyObjectiveDefinition
	DisplayActivityHashes          []uint32                             `json:"displayActivityHashes"` // Mapped to Destiny.Definitions.DestinyActivityDefinition
	RequireFullObjectiveCompletion bool                                 `json:"requireFullObjectiveCompletion"`
	QuestlineItemHash              uint32                               `json:"questlineItemHash"` // Mapped to Destiny.Definitions.DestinyInventoryItemDefinition
	Narrative                      string                               `json:"narrative"`
	ObjectiveVerbName              string                               `json:"objectiveVerbName"`
	QuestTypeIdentifier            string                               `json:"questTypeIdentifier"`
	QuestTypeHash                  uint32                               `json:"questTypeHash"` // Mapped to Destiny.Definitions.DestinyProgressionDefinition
	PerObjectiveDisplayProperties  []*DestinyObjectiveDisplayProperties `json:"perObjectiveDisplayProperties"`
	DisplayAsStatTracker           bool                                 `json:"displayAsStatTracker"`
}

type DestinyItemMetricBlockDefinition struct {
	AvailableMetricCategoryNodeHashes []uint32 `json:"availableMetricCategoryNodeHashes"` // Mapped to Destiny.Definitions.DestinyMetricCategoryDefinition
}

type DestinyItemPlugDefinition struct {
	InsertionRules                   []*Items.DestinyPlugRuleDefinition `json:"insertionRules"`
	PlugCategoryIdentifier           string                             `json:"plugCategoryIdentifier"`
	PlugCategoryHash                 uint32                             `json:"plugCategoryHash"` // Mapped to Destiny.Definitions.DestinyPlugCategoryDefinition
	OnActionRecreateSelf             bool                               `json:"onActionRecreateSelf"`
	InsertionMaterialRequirementHash uint32                             `json:"insertionMaterialRequirementHash"` // Mapped to Destiny.Definitions.DestinyInventoryItemDefinition
	PreviewItemOverrideHash          uint32                             `json:"previewItemOverrideHash"`          // Mapped to Destiny.Definitions.DestinyInventoryItemDefinition
	EnabledMaterialRequirementHash   uint32                             `json:"enabledMaterialRequirementHash"`   // Mapped to Destiny.Definitions.DestinyInventoryItemDefinition
	EnabledRules                     []*Items.DestinyPlugRuleDefinition `json:"enabledRules"`
	UiPlugLabel                      string                             `json:"uiPlugLabel"`
	PlugStyle                        int32                              `json:"plugStyle"`
	PlugAvailability                 int32                              `json:"plugAvailability"`
	AlternateUiPlugLabel             string                             `json:"alternateUiPlugLabel"`
	AlternatePlugStyle               int32                              `json:"alternatePlugStyle"`
	IsDummyPlug                      bool                               `json:"isDummyPlug"`
	ParentItemOverride               *Items.DestinyParentItemOverride   `json:"parentItemOverride"`
	EnergyCapacity                   *Items.DestinyEnergyCapacityEntry  `json:"energyCapacity"`
	EnergyCost                       *Items.DestinyEnergyCostEntry      `json:"energyCost"`
}

type DestinyItemGearsetBlockDefinition struct {
	TrackingValueMax int32    `json:"trackingValueMax"`
	ItemList         []uint32 `json:"itemList"` // Mapped to Destiny.Definitions.DestinyInventoryItemDefinition
}

type DestinyItemSackBlockDefinition struct {
	DetailAction    string `json:"detailAction"`
	OpenAction      string `json:"openAction"`
	SelectItemCount int32  `json:"selectItemCount"`
	VendorSackType  string `json:"vendorSackType"`
	OpenOnAcquire   bool   `json:"openOnAcquire"`
}

type DestinyItemSocketEntryPlugItemDefinition struct {
	PlugItemHash uint32 `json:"plugItemHash"` // Mapped to Destiny.Definitions.DestinyInventoryItemDefinition
}

type DestinyItemSocketEntryDefinition struct {
	SocketTypeHash                        uint32                                      `json:"socketTypeHash"`        // Mapped to Destiny.Definitions.DestinySocketTypeDefinition
	SingleInitialItemHash                 uint32                                      `json:"singleInitialItemHash"` // Mapped to Destiny.Definitions.DestinyInventoryItemDefinition
	ReusablePlugItems                     []*DestinyItemSocketEntryPlugItemDefinition `json:"reusablePlugItems"`
	PreventInitializationOnVendorPurchase bool                                        `json:"preventInitializationOnVendorPurchase"`
	HidePerksInItemTooltip                bool                                        `json:"hidePerksInItemTooltip"`
	PlugSources                           int32                                       `json:"plugSources"`
	ReusablePlugSetHash                   uint32                                      `json:"reusablePlugSetHash"`   // Mapped to Destiny.Definitions.DestinyPlugSetDefinition
	RandomizedPlugSetHash                 uint32                                      `json:"randomizedPlugSetHash"` // Mapped to Destiny.Definitions.DestinyPlugSetDefinition
	DefaultVisible                        bool                                        `json:"defaultVisible"`
}

type DestinyItemIntrinsicSocketEntryDefinition struct {
	PlugItemHash   uint32 `json:"plugItemHash"`   // Mapped to Destiny.Definitions.DestinyInventoryItemDefinition
	SocketTypeHash uint32 `json:"socketTypeHash"` // Mapped to Destiny.Definitions.DestinySocketTypeDefinition
	DefaultVisible bool   `json:"defaultVisible"`
}

type DestinyItemSocketCategoryDefinition struct {
	SocketCategoryHash uint32  `json:"socketCategoryHash"` // Mapped to Destiny.Definitions.DestinySocketCategoryDefinition
	SocketIndexes      []int32 `json:"socketIndexes"`
}

type DestinyItemSocketBlockDefinition struct {
	Detail           string                                       `json:"detail"`
	SocketEntries    []*DestinyItemSocketEntryDefinition          `json:"socketEntries"`
	IntrinsicSockets []*DestinyItemIntrinsicSocketEntryDefinition `json:"intrinsicSockets"`
	SocketCategories []*DestinyItemSocketCategoryDefinition       `json:"socketCategories"`
}

type DestinyItemSummaryBlockDefinition struct {
	SortPriority int32 `json:"sortPriority"`
}

type DestinyItemTalentGridBlockDefinition struct {
	TalentGridHash   uint32 `json:"talentGridHash"` // Mapped to Destiny.Definitions.DestinyTalentGridDefinition
	ItemDetailString string `json:"itemDetailString"`
	BuildName        string `json:"buildName"`
	HudDamageType    int32  `json:"hudDamageType"`
	HudIcon          string `json:"hudIcon"`
}

type DestinyItemInvestmentStatDefinition struct {
	StatTypeHash          uint32 `json:"statTypeHash"` // Mapped to Destiny.Definitions.DestinyStatDefinition
	Value                 int32  `json:"value"`
	IsConditionallyActive bool   `json:"isConditionallyActive"`
}

type DestinyItemPerkEntryDefinition struct {
	RequirementDisplayString string `json:"requirementDisplayString"`
	PerkHash                 uint32 `json:"perkHash"` // Mapped to Destiny.Definitions.DestinySandboxPerkDefinition
	PerkVisibility           int32  `json:"perkVisibility"`
}

type DestinyInventoryItemDefinition struct {
	DisplayProperties                 *Common.DestinyDisplayPropertiesDefinition `json:"displayProperties"`
	TooltipNotifications              []*DestinyItemTooltipNotification          `json:"tooltipNotifications"`
	CollectibleHash                   uint32                                     `json:"collectibleHash"` // Mapped to Destiny.Definitions.DestinyCollectibleDefinition
	IconWatermark                     string                                     `json:"iconWatermark"`
	IconWatermarkShelved              string                                     `json:"iconWatermarkShelved"`
	SecondaryIcon                     string                                     `json:"secondaryIcon"`
	SecondaryOverlay                  string                                     `json:"secondaryOverlay"`
	SecondarySpecial                  string                                     `json:"secondarySpecial"`
	BackgroundColor                   *Misc.DestinyColor                         `json:"backgroundColor"`
	Screenshot                        string                                     `json:"screenshot"`
	ItemTypeDisplayName               string                                     `json:"itemTypeDisplayName"`
	FlavorText                        string                                     `json:"flavorText"`
	UiItemDisplayStyle                string                                     `json:"uiItemDisplayStyle"`
	ItemTypeAndTierDisplayName        string                                     `json:"itemTypeAndTierDisplayName"`
	DisplaySource                     string                                     `json:"displaySource"`
	TooltipStyle                      string                                     `json:"tooltipStyle"`
	Action                            *DestinyItemActionBlockDefinition          `json:"action"`
	Crafting                          *DestinyItemCraftingBlockDefinition        `json:"crafting"`
	Inventory                         *DestinyItemInventoryBlockDefinition       `json:"inventory"`
	SetData                           *DestinyItemSetBlockDefinition             `json:"setData"`
	Stats                             *DestinyItemStatBlockDefinition            `json:"stats"`
	EmblemObjectiveHash               uint32                                     `json:"emblemObjectiveHash"` // Mapped to Destiny.Definitions.DestinyObjectiveDefinition
	EquippingBlock                    *DestinyEquippingBlockDefinition           `json:"equippingBlock"`
	TranslationBlock                  *DestinyItemTranslationBlockDefinition     `json:"translationBlock"`
	Preview                           *DestinyItemPreviewBlockDefinition         `json:"preview"`
	Quality                           *DestinyItemQualityBlockDefinition         `json:"quality"`
	Value                             *DestinyItemValueBlockDefinition           `json:"value"`
	SourceData                        *DestinyItemSourceBlockDefinition          `json:"sourceData"`
	ObjectiveData                     *DestinyItemObjectiveBlockDefinition       `json:"objectiveData"`
	Metrics                           *DestinyItemMetricBlockDefinition          `json:"metrics"`
	Plug                              *DestinyItemPlugDefinition                 `json:"plug"`
	Gearset                           *DestinyItemGearsetBlockDefinition         `json:"gearset"`
	Sack                              *DestinyItemSackBlockDefinition            `json:"sack"`
	Sockets                           *DestinyItemSocketBlockDefinition          `json:"sockets"`
	Summary                           *DestinyItemSummaryBlockDefinition         `json:"summary"`
	TalentGrid                        *DestinyItemTalentGridBlockDefinition      `json:"talentGrid"`
	InvestmentStats                   []*DestinyItemInvestmentStatDefinition     `json:"investmentStats"`
	Perks                             []*DestinyItemPerkEntryDefinition          `json:"perks"`
	LoreHash                          uint32                                     `json:"loreHash"`        // Mapped to Destiny.Definitions.DestinyLoreDefinition
	SummaryItemHash                   uint32                                     `json:"summaryItemHash"` // Mapped to Destiny.Definitions.DestinyInventoryItemDefinition
	Animations                        []*Animations.DestinyAnimationReference    `json:"animations"`
	AllowActions                      bool                                       `json:"allowActions"`
	Links                             []*Links.HyperlinkReference                `json:"links"`
	DoesPostmasterPullHaveSideEffects bool                                       `json:"doesPostmasterPullHaveSideEffects"`
	NonTransferrable                  bool                                       `json:"nonTransferrable"`
	ItemCategoryHashes                []uint32                                   `json:"itemCategoryHashes"` // Mapped to Destiny.Definitions.DestinyItemCategoryDefinition
	SpecialItemType                   int32                                      `json:"specialItemType"`
	ItemType                          int32                                      `json:"itemType"`
	ItemSubType                       int32                                      `json:"itemSubType"`
	ClassType                         int32                                      `json:"classType"`
	BreakerType                       int32                                      `json:"breakerType"`
	BreakerTypeHash                   uint32                                     `json:"breakerTypeHash"` // Mapped to Destiny.Definitions.DestinyBreakerTypeDefinition
	Equippable                        bool                                       `json:"equippable"`
	DamageTypeHashes                  []uint32                                   `json:"damageTypeHashes"` // Mapped to Destiny.Definitions.DestinyDamageTypeDefinition
	DamageTypes                       []int32                                    `json:"damageTypes"`
	DefaultDamageType                 uint32                                     `json:"defaultDamageType"`
	SeasonHash                        uint32                                     `json:"seasonHash"` // Mapped to Destiny.Definitions.DestinySeasonDefinition
	IsWrapper                         bool                                       `json:"isWrapper"`
	TraitHashes                       []uint32                                   `json:"traitHashes"` // Mapped to Destiny.Definitions.DestinyTraitDefinition
	Hash                              uint32                                     `json:"hash"`
	Index                             int32                                      `json:"index"`
	Redacted                          bool                                       `json:"redacted"`
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

type SourcesDestinyItemSourceDefinition struct {
	Level            int32                                          `json:"level"`
	MinQuality       int32                                          `json:"minQuality"`
	MaxQuality       int32                                          `json:"maxQuality"`
	MinLevelRequired int32                                          `json:"minLevelRequired"`
	MaxLevelRequired int32                                          `json:"maxLevelRequired"`
	ComputedStats    map[uint32]*DestinyInventoryItemStatDefinition `json:"computedStats"`
	SourceHashes     []uint32                                       `json:"sourceHashes"`
}
