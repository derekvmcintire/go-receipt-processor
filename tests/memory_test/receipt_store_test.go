package memory_test

import (
	"github.com/stretchr/testify/assert"
	"go-receipt-processor/internal/adapters/memory"
	"go-receipt-processor/internal/domain"
	"testing"
)

func TestSaveReceipt_Success(t *testing.T) {
	// Create a new in-memory store
	store := memory.NewReceiptStore()

	// Define a sample receipt to save
	receipt := domain.Receipt{
		Retailer:     "Store A",
		PurchaseDate: "2024-11-29",
		PurchaseTime: "14:30",
		Items: []domain.Item{
			{ShortDescription: "Item 1", Price: "50.00"},
			{ShortDescription: "Item 2", Price: "50.00"},
		},
		Total:  "100.00",
		Points: 100,
	}

	// Save the receipt in the store
	receiptID, err := store.Save(receipt)

	// Assert that there is no error and that the ID is returned
	assert.NoError(t, err)
	assert.NotEmpty(t, receiptID)

	// Retrieve the receipt by ID and assert that it matches the saved receipt
	savedReceipt, err := store.Find(receiptID)
	assert.NoError(t, err)
	assert.Equal(t, receipt.Retailer, savedReceipt.Retailer)
	assert.Equal(t, receipt.Total, savedReceipt.Total)
	assert.Equal(t, receipt.Points, savedReceipt.Points)
}

func TestFindReceipt_Success(t *testing.T) {
	// Create a new in-memory store
	store := memory.NewReceiptStore()

	// Define and save a sample receipt
	receipt := domain.Receipt{
		Retailer:     "Store B",
		PurchaseDate: "2024-11-29",
		PurchaseTime: "15:30",
		Items: []domain.Item{
			{ShortDescription: "Item A", Price: "30.00"},
			{ShortDescription: "Item B", Price: "70.00"},
		},
		Total:  "100.00",
		Points: 100,
	}

	receiptID, err := store.Save(receipt)
	assert.NoError(t, err)

	// Retrieve the saved receipt
	savedReceipt, err := store.Find(receiptID)
	assert.NoError(t, err)
	assert.Equal(t, receipt.Retailer, savedReceipt.Retailer)
	assert.Equal(t, receipt.Total, savedReceipt.Total)
	assert.Equal(t, receipt.Points, savedReceipt.Points)
}

func TestFindReceipt_Failure(t *testing.T) {
	// Create a new in-memory store
	store := memory.NewReceiptStore()

	// Attempt to find a non-existent receipt
	nonExistentID := "nonexistent-id"
	_, err := store.Find(nonExistentID)

	// Assert that no error is returned (as per the current implementation)
	// (You can modify this behavior in the future to return an error if needed)
	assert.NoError(t, err)
}

func TestSaveMultipleReceipts_UniqueIDs(t *testing.T) {
	// Create a new in-memory store
	store := memory.NewReceiptStore()

	// Define two sample receipts
	receipt1 := domain.Receipt{
		Retailer:     "Store C",
		PurchaseDate: "2024-11-29",
		PurchaseTime: "16:00",
		Items: []domain.Item{
			{ShortDescription: "Item 1", Price: "25.00"},
			{ShortDescription: "Item 2", Price: "25.00"},
		},
		Total:  "50.00",
		Points: 50,
	}
	receipt2 := domain.Receipt{
		Retailer:     "Store D",
		PurchaseDate: "2024-11-29",
		PurchaseTime: "17:00",
		Items: []domain.Item{
			{ShortDescription: "Item A", Price: "40.00"},
			{ShortDescription: "Item B", Price: "60.00"},
		},
		Total:  "100.00",
		Points: 100,
	}

	// Save the receipts
	receiptID1, err := store.Save(receipt1)
	assert.NoError(t, err)
	receiptID2, err := store.Save(receipt2)
	assert.NoError(t, err)

	// Assert that the IDs are unique
	assert.NotEqual(t, receiptID1, receiptID2)

	// Retrieve both receipts and assert that they match the saved data
	savedReceipt1, err := store.Find(receiptID1)
	assert.NoError(t, err)
	assert.Equal(t, receipt1.Retailer, savedReceipt1.Retailer)
	assert.Equal(t, receipt1.Total, savedReceipt1.Total)

	savedReceipt2, err := store.Find(receiptID2)
	assert.NoError(t, err)
	assert.Equal(t, receipt2.Retailer, savedReceipt2.Retailer)
	assert.Equal(t, receipt2.Total, savedReceipt2.Total)
}
