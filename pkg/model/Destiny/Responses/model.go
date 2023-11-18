package Responses

import (
	"github.com/d2checkpoint-com/bungienet/pkg/model/Destiny/Components/Profiles"
	"github.com/d2checkpoint-com/bungienet/pkg/model/Destiny/Entities/Characters"
	"time"
)

type SingleComponentResponseOfDestinyVendorReceiptsComponent any                   //TODO
type SingleComponentResponseOfDestinyInventoryComponent any                        //TODO
type SingleComponentResponseOfDestinyProfileComponent any                          //TODO
type SingleComponentResponseOfDestinyPlatformSilverComponent any                   //TODO
type SingleComponentResponseOfDestinyKiosksComponent any                           //TODO
type SingleComponentResponseOfDestinyPlugSetsComponent any                         //TODO
type SingleComponentResponseOfDestinyProfileProgressionComponent any               //TODO
type SingleComponentResponseOfDestinyPresentationNodesComponent any                //TODO
type SingleComponentResponseOfDestinyProfileRecordsComponent any                   //TODO
type SingleComponentResponseOfDestinyProfileCollectiblesComponent any              //TODO
type SingleComponentResponseOfDestinyMetricsComponent any                          //TODO
type SingleComponentResponseOfDestinyStringVariablesComponent any                  //TODO
type SingleComponentResponseOfDestinySocialCommendationsComponent any              //TODO
type DictionaryComponentResponseOfint64AndDestinyCharacterComponent any            //TODO
type DictionaryComponentResponseOfint64AndDestinyInventoryComponent any            //TODO
type DictionaryComponentResponseOfint64AndDestinyLoadoutsComponent any             //TODO
type DictionaryComponentResponseOfint64AndDestinyCharacterProgressionComponent any //TODO
type DictionaryComponentResponseOfint64AndDestinyCharacterRenderComponent any      //TODO
type DictionaryComponentResponseOfint64AndDestinyKiosksComponent any               //TODO
type DictionaryComponentResponseOfint64AndDestinyPlugSetsComponent any             //TODO
type DestinyBaseItemComponentSetOfuint32 any                                       //TODO
type DictionaryComponentResponseOfint64AndDestinyPresentationNodesComponent any    //TODO
type DictionaryComponentResponseOfint64AndDestinyCharacterRecordsComponent any     //TODO
type DictionaryComponentResponseOfint64AndDestinyCollectiblesComponent any         //TODO
type DictionaryComponentResponseOfint64AndDestinyStringVariablesComponent any      //TODO
type DictionaryComponentResponseOfint64AndDestinyCraftablesComponent any           //TODO
type DestinyItemComponentSetOfint64 any                                            //TODO
type DictionaryComponentResponseOfint64AndDestinyCurrenciesComponent any           //TODO
type SingleComponentResponseOfDestinyCharacterComponent any                        //TODO
type SingleComponentResponseOfDestinyCharacterProgressionComponent any             //TODO
type SingleComponentResponseOfDestinyCharacterRenderComponent any                  //TODO
type SingleComponentResponseOfDestinyCharacterActivitiesComponent any              //TODO
type SingleComponentResponseOfDestinyLoadoutsComponent any                         //TODO
type SingleComponentResponseOfDestinyCharacterRecordsComponent any                 //TODO
type SingleComponentResponseOfDestinyCollectiblesComponent any                     //TODO
type SingleComponentResponseOfDestinyCurrenciesComponent any                       //TODO

type DestinyProfileResponse struct {
	ResponseMintedTimestamp            time.Time                                                                  `json:"responseMintedTimestamp"`
	SecondaryComponentsMintedTimestamp time.Time                                                                  `json:"secondaryComponentsMintedTimestamp"`
	VendorReceipts                     *SingleComponentResponseOfDestinyVendorReceiptsComponent                   `json:"vendorReceipts"`
	ProfileInventory                   *SingleComponentResponseOfDestinyInventoryComponent                        `json:"profileInventory"`
	ProfileCurrencies                  *SingleComponentResponseOfDestinyInventoryComponent                        `json:"profileCurrencies"`
	Profile                            *SingleComponentResponseOfDestinyProfileComponent                          `json:"profile"`
	PlatformSilver                     *SingleComponentResponseOfDestinyPlatformSilverComponent                   `json:"platformSilver"`
	ProfileKiosks                      *SingleComponentResponseOfDestinyKiosksComponent                           `json:"profileKiosks"`
	ProfilePlugSets                    *SingleComponentResponseOfDestinyPlugSetsComponent                         `json:"profilePlugSets"`
	ProfileProgression                 *SingleComponentResponseOfDestinyProfileProgressionComponent               `json:"profileProgression"`
	ProfilePresentationNodes           *SingleComponentResponseOfDestinyPresentationNodesComponent                `json:"profilePresentationNodes"`
	ProfileRecords                     *SingleComponentResponseOfDestinyProfileRecordsComponent                   `json:"profileRecords"`
	ProfileCollectibles                *SingleComponentResponseOfDestinyProfileCollectiblesComponent              `json:"profileCollectibles"`
	ProfileTransitoryData              *SingleComponentResponseOfDestinyProfileTransitoryComponent                `json:"profileTransitoryData"`
	Metrics                            *SingleComponentResponseOfDestinyMetricsComponent                          `json:"metrics"`
	ProfileStringVariables             *SingleComponentResponseOfDestinyStringVariablesComponent                  `json:"profileStringVariables"`
	ProfileCommendations               *SingleComponentResponseOfDestinySocialCommendationsComponent              `json:"profileCommendations"`
	Characters                         *DictionaryComponentResponseOfint64AndDestinyCharacterComponent            `json:"characters"`
	CharacterInventories               *DictionaryComponentResponseOfint64AndDestinyInventoryComponent            `json:"characterInventories"`
	CharacterLoadouts                  *DictionaryComponentResponseOfint64AndDestinyLoadoutsComponent             `json:"characterLoadouts"`
	CharacterProgressions              *DictionaryComponentResponseOfint64AndDestinyCharacterProgressionComponent `json:"characterProgressions"`
	CharacterRenderData                *DictionaryComponentResponseOfint64AndDestinyCharacterRenderComponent      `json:"characterRenderData"`
	CharacterActivities                *DictionaryComponentResponseOfint64AndDestinyCharacterActivitiesComponent  `json:"characterActivities"`
	CharacterEquipment                 *DictionaryComponentResponseOfint64AndDestinyInventoryComponent            `json:"characterEquipment"`
	CharacterKiosks                    *DictionaryComponentResponseOfint64AndDestinyKiosksComponent               `json:"characterKiosks"`
	CharacterPlugSets                  *DictionaryComponentResponseOfint64AndDestinyPlugSetsComponent             `json:"characterPlugSets"`
	CharacterUninstancedItemComponents map[int64]*DestinyBaseItemComponentSetOfuint32                             `json:"characterUninstancedItemComponents"`
	CharacterPresentationNodes         *DictionaryComponentResponseOfint64AndDestinyPresentationNodesComponent    `json:"characterPresentationNodes"`
	CharacterRecords                   *DictionaryComponentResponseOfint64AndDestinyCharacterRecordsComponent     `json:"characterRecords"`
	CharacterCollectibles              *DictionaryComponentResponseOfint64AndDestinyCollectiblesComponent         `json:"characterCollectibles"`
	CharacterStringVariables           *DictionaryComponentResponseOfint64AndDestinyStringVariablesComponent      `json:"characterStringVariables"`
	CharacterCraftables                *DictionaryComponentResponseOfint64AndDestinyCurrenciesComponent           `json:"characterCraftables"`
	ItemComponents                     *DestinyItemComponentSetOfint64                                            `json:"itemComponents"`
	CharacterCurrencyLookups           *DictionaryComponentResponseOfint64AndDestinyCurrenciesComponent           `json:"characterCurrencyLookups"`
}

type DestinyCharacterResponse struct {
	Inventory                 *SingleComponentResponseOfDestinyInventoryComponent            `json:"inventory"`
	Character                 *SingleComponentResponseOfDestinyCharacterComponent            `json:"character"`
	Progressions              *SingleComponentResponseOfDestinyCharacterProgressionComponent `json:"progressions"`
	RenderData                *SingleComponentResponseOfDestinyCharacterRenderComponent      `json:"renderData"`
	Activities                *SingleComponentResponseOfDestinyCharacterActivitiesComponent  `json:"activities"`
	Equipment                 *SingleComponentResponseOfDestinyInventoryComponent            `json:"equipment"`
	Loadouts                  *SingleComponentResponseOfDestinyLoadoutsComponent             `json:"loadouts"`
	Kiosks                    *SingleComponentResponseOfDestinyKiosksComponent               `json:"kiosks"`
	PlugSets                  *SingleComponentResponseOfDestinyPlugSetsComponent             `json:"plugSets"`
	PresentationNodes         *SingleComponentResponseOfDestinyPresentationNodesComponent    `json:"presentationNodes"`
	Records                   *SingleComponentResponseOfDestinyCharacterRecordsComponent     `json:"records"`
	Collectibles              *SingleComponentResponseOfDestinyCollectiblesComponent         `json:"collectibles"`
	ItemComponents            *DestinyItemComponentSetOfint64                                `json:"itemComponents"`
	UninstancedItemComponents *DestinyBaseItemComponentSetOfuint32                           `json:"uninstancedItemComponents"`
	CurrencyLookups           *SingleComponentResponseOfDestinyCurrenciesComponent           `json:"currencyLookups"`
}

type SingleComponentResponseOfDestinyProfileTransitoryComponent struct {
	Data     *Profiles.DestinyProfileTransitoryComponent `json:"data"`
	Privacy  int32                                       `json:"privacy"`
	Disabled bool                                        `json:"disabled"`
}

type DictionaryComponentResponseOfint64AndDestinyCharacterActivitiesComponent struct {
	Data     map[int64]*Characters.DestinyCharacterActivitiesComponent `json:"data"`
	Privacy  int32                                                     `json:"privacy"`
	Disabled bool                                                      `json:"disabled"`
}
