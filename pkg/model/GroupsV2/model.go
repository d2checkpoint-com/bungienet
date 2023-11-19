package GroupsV2

type GroupUserInfoCard struct {
	LastSeenDisplayName         string  `json:"LastSeenDisplayName"`
	LastSeenDisplayNameType     int32   `json:"LastSeenDisplayNameType"`
	SupplementalDisplayName     string  `json:"SupplementalDisplayName"`
	IconPath                    string  `json:"iconPath"`
	CrossSaveOverride           int32   `json:"crossSaveOverride"`
	ApplicableMembershipTypes   []int32 `json:"applicableMembershipTypes"`
	IsPublic                    bool    `json:"isPublic"`
	MembershipType              int32   `json:"membershipType"`
	MembershipId                int64   `json:"membershipId,string"`
	DisplayName                 string  `json:"displayName"`
	BungieGlobalDisplayName     string  `json:"bungieGlobalDisplayName"`
	BungieGlobalDisplayNameCode int16   `json:"bungieGlobalDisplayNameCode"`
}
