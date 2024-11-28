// The Receipt domain model (contains fields like retailer, items, total, etc.)
package domain

type Item struct {
	ShortDescription string
	Price            float64
}

type Receipt struct {
	ID           string
	Retailer     string
	PurchaseDate string
	PurchaseTime string
	Items        []Item
	Total        float64
}
