package domain

type Item struct {
	ShortDescription string `json:"shortDescription" binding:"required"` // JSON key mapping and validation
	Price            string `json:"price" binding:"required"`            // Mark field as required
}

type Receipt struct {
	ID           string `json:"id"`                                     // Optional, not marked as required
	Retailer     string `json:"retailer" binding:"required"`            // Required field
	PurchaseDate string `json:"purchaseDate" binding:"required"`        // Required field
	PurchaseTime string `json:"purchaseTime" binding:"required"`        // Required field
	Items        []Item `json:"items" binding:"required,dive,required"` // Ensure `items` is not empty and each item is validated
	Total        string `json:"total" binding:"required"`               // Required field
	Points       int    `json:"points"`                                 // Optional, likely calculated later
}
