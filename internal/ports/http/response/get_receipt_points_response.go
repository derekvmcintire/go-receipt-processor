package response

// ReceiptProcessResponse represents the response data for processing a receipt.
type GetReceiptPointsResponse struct {
	Points int `json:"points"` // Points from the requested receipt
}
