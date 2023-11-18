package Characters

import (
	"github.com/d2checkpoint-com/bungienet/pkg/model/Destiny"
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
