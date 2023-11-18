package Profiles

import "github.com/d2checkpoint-com/bungienet/pkg/model/Destiny/Vendors"

type DestinyVendorReceiptsComponent struct {
	Receipts []*Vendors.DestinyVendorReceipt `json:"receipts"`
}
