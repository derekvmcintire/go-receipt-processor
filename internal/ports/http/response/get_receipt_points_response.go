package response

// ReceiptProcessResponse represents the response data for processing a receipt.
type GetReceiptPointsResponse struct {
	Points int `json:"points"` // JSON binding Points to lowercase points
}
