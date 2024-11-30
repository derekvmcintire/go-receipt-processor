package domain

type Item struct {
	ShortDescription string `json:"shortDescription" binding:"required"`
	Price            string `json:"price" binding:"required"`
}

type Receipt struct {
	ID           string `json:"id"`
	Retailer     string `json:"retailer" binding:"required"`
	PurchaseDate string `json:"purchaseDate" binding:"required"`
	PurchaseTime string `json:"purchaseTime" binding:"required"`
	Items        []Item `json:"items" binding:"required,dive,required"` // Ensure `items` is not empty and each item is validated
	Total        string `json:"total" binding:"required"`
	Points       int    `json:"points"`
}
