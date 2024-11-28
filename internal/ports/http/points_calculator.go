package http

import "go-receipt-processor/internal/domain"

// PointsCalculator defines the interface for calculating points
type PointsCalculator interface {
	CalculatePoints(receipt domain.Receipt) (int, error)
}
