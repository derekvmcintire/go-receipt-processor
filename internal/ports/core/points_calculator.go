package http

import (
	"go-receipt-processor/internal/domain"
)

// PointsCalculator defines the method required for calculating points based on a receipt.
// Implementations of this interface encapsulate the business logic for determining
// points from receipt data.
type PointsCalculator interface {
	// CalculatePoints computes the points for the given receipt based on predefined business rules.
	// It returns:
	//   - The calculated points as an integer.
	//   - An error if the calculation fails, such as due to invalid input data.
	CalculatePoints(receipt domain.Receipt) (int, error)
}