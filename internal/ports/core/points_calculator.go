package http

import (
	"go-receipt-processor/internal/domain"
)

// PointsCalculator defines the methods required for calculating points based on a receipt.
type PointsCalculator interface {
	CalculatePoints(receipt domain.Receipt) (int, error)
}
