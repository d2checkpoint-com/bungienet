package Characters

import (
	"github.com/d2checkpoint-com/bungienet/pkg/model/Destiny"
	"github.com/d2checkpoint-com/bungienet/pkg/model/Destiny/Misc"
	"time"
)

type DestinyCharacterActivitiesComponent struct {
	DateActivityStarted         time.Time                  `json:"dateActivityStarted"`
	AvailableActivities         []*Destiny.DestinyActivity `json:"availableActivities"`
	CurrentActivityHash         uint32                     `json:"currentActivityHash"` // mapped to Destiny.Definitions.DestinyActivityDefinition
	CurrentActivityMode         uint32                     `json:"currentActivityMode"` // mapped to  Destiny.Definitions.DestinyActivityModeDefinition
	CurrentActivityModeType     int32                      `json:"currentActivityModeType"`
	CurrentActivityModeHashes   []uint32                   `json:"currentActivityModeHashes"` // mapped to Destiny.Definitions.DestinyActivityModeDefinition
	CurrentActivityModeTypes    []int32                    `json:"currentActivityModeTypes"`
	CurrentPlaylistActivityHash uint32                     `json:"currentPlaylistActivityHash"` // mapped to Destiny.Definitions.DestinyActivityDefinition
	LastCompletedStoryHash      uint32                     `json:"lastCompletedStoryHash"`      // mapped to Destiny.Definitions.DestinyActivityDefinition
}

type DestinyCharacterComponent struct {
	MembershipId             int64                       `json:"membershipId,string"`
	MembershipType           int32                       `json:"membershipType"`
	CharacterId              int64                       `json:"characterId,string"`
	DateLastPlayed           time.Time                   `json:"dateLastPlayed"`
	MinutesPlayedThisSession int64                       `json:"minutesPlayedThisSession"`
	MinutesPlayedTotal       int64                       `json:"minutesPlayedTotal"`
	Light                    int32                       `json:"light"`
	Stats                    map[uint32]int32            `json:"stats"`
	RaceHash                 uint32                      `json:"raceHash"`
	GenderHash               uint32                      `json:"genderHash"`
	ClassHash                uint32                      `json:"classHash"`
	RaceType                 int32                       `json:"raceType"`
	ClassType                int32                       `json:"classType"`
	GenderType               int32                       `json:"genderType"`
	EmblemPath               string                      `json:"emblemPath"`
	EmblemBackgroundPath     string                      `json:"emblemBackgroundPath"`
	EmblemHash               uint32                      `json:"emblemHash"`
	EmblemColor              *Misc.DestinyColor          `json:"emblemColor"`
	LevelProgression         *Destiny.DestinyProgression `json:"levelProgression"`
	BaseCharacterLevel       int32                       `json:"baseCharacterLevel"`
	PercentToNextLevel       float64                     `json:"percentToNextLevel"`
	TitleRecordHash          int32                       `json:"titleRecordHash"`
}
