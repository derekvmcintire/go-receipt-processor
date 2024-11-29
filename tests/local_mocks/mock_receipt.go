package local_mocks

import "go-receipt-processor/internal/domain"

var MockReceipt = domain.Receipt{
	Retailer:     "StoreABC",
	PurchaseDate: "2024-11-29",
	PurchaseTime: "15:30",
	Items: []domain.Item{
		{ShortDescription: "Item 1", Price: "5.00"},
		{ShortDescription: "Item 2", Price: "5.00"},
	},
	Total:  "10.00",
	Points: 0,
}
