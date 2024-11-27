// The Receipt domain model (contains fields like retailer, items, total, etc.)
package domain

import "time"

type Item struct {
	ShortDescription string
	Price            float64
}

type Receipt struct {
    ID          string
    Retailer    string
    PurchaseDate time.Time
    PurchaseTime time.Time
    Items       []Item
    Total       float64
}
