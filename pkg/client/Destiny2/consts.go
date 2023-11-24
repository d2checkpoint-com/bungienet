package Destiny2

type DestinyComponentType int32

const (
	None                  DestinyComponentType = 0
	Profiles              DestinyComponentType = 100
	VendorReceipts        DestinyComponentType = 101
	ProfileInventories    DestinyComponentType = 102
	ProfileCurrencies     DestinyComponentType = 103
	ProfileProgression    DestinyComponentType = 104
	PlatformSilver        DestinyComponentType = 105
	Characters            DestinyComponentType = 200
	CharacterInventories  DestinyComponentType = 201
	CharacterProgressions DestinyComponentType = 202
	CharacterRenderData   DestinyComponentType = 203
	CharacterActivities   DestinyComponentType = 204
	CharacterEquipment    DestinyComponentType = 205
	ItemInstances         DestinyComponentType = 300
	ItemObjectives        DestinyComponentType = 301
	ItemPerks             DestinyComponentType = 302
	ItemRenderData        DestinyComponentType = 303
	ItemStats             DestinyComponentType = 304
	ItemSockets           DestinyComponentType = 305
	ItemTalentGrids       DestinyComponentType = 306
	ItemCommonData        DestinyComponentType = 307
	ItemPlugStates        DestinyComponentType = 308
	ItemPlugObjectives    DestinyComponentType = 309
	ItemReusablePlugs     DestinyComponentType = 310
	Vendors               DestinyComponentType = 400
	VendorCategories      DestinyComponentType = 401
	VendorSales           DestinyComponentType = 402
	Kiosks                DestinyComponentType = 500
	CurrencyLookups       DestinyComponentType = 600
	PresentationNodes     DestinyComponentType = 700
	Collectibles          DestinyComponentType = 800
	Records               DestinyComponentType = 900
	Transitory            DestinyComponentType = 1000
	Metrics               DestinyComponentType = 1100
	StringVariables       DestinyComponentType = 1200
	Craftables            DestinyComponentType = 1300
	SocialCommendations   DestinyComponentType = 1400
)
