// Package domain defines the core entities used in the receipt processing application.
// These domain models represent the essential data structures shared across the application.
package domain

// Item represents an individual item on a receipt.
// It includes a short description and its price.
type Item struct {
	ShortDescription string // A brief description of the item.
	Price            string // The price of the item in currency units.
}

// Receipt represents a customer's receipt and its associated data.
// It includes details about the retailer, purchase information, items, total amount, and points.
type Receipt struct {
	ID           string // A unique identifier for the receipt.
	Retailer     string // The name of the retailer where the receipt was issued.
	PurchaseDate string // The date of purchase in YYYY-MM-DD format.
	PurchaseTime string // The time of purchase in HH:MM format (24-hour clock).
	Items        []Item // A list of items purchased, represented by the Item struct.
	Total        string // The total amount for the receipt in currency units.
	Points       int    // The points calculated based on the receipt's data.
}
