package Vendors

import (
	"github.com/d2checkpoint-com/bungienet/pkg/model/Destiny"
	"time"
)

type DestinyVendorReceipt struct {
	CurrencyPaid           []*Destiny.DestinyItemQuantity `json:"currencyPaid"`
	ItemReceived           *Destiny.DestinyItemQuantity   `json:"itemReceived"`
	LicenseUnlockHash      uint32                         `json:"licenseUnlockHash"`
	PurchasedByCharacterId int64                          `json:"purchasedByCharacterId,string"`
	RefundPolicy           int32                          `json:"refundPolicy"`
	SequenceNumber         int32                          `json:"sequenceNumber"`
	TimeToExpiration       int64                          `json:"timeToExpiration,string"`
	ExpiresOn              time.Time                      `json:"expiresOn"`
}
