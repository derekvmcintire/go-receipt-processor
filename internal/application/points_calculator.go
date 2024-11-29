// points_calculator_impl.go
package application

import (
	"go-receipt-processor/internal/domain"
	"go-receipt-processor/internal/ports/http"
)

// PointsCalculatorImpl implements the PointsCalculator interface,
// responsible for calculating points based on receipt data.
type PointsCalculatorImpl struct{}

// NewPointsCalculator creates and returns a new instance of PointsCalculatorImpl.
func NewPointsCalculator() http.PointsCalculator {
	return &PointsCalculatorImpl{}
}

// CalculatePoints calculates the total points for a receipt based on specific business rules.
//
// Parameters:
//   - receipt: The domain.Receipt object containing receipt details.
//
// Returns: A time.TIme value representing the the exact time of the purchase
//   - points: The total points awarded for the receipt.
//   - err: An error if calculating points fails
func (c *PointsCalculatorImpl) CalculatePoints(receipt domain.Receipt) (int, error) {
	points := 0

	// Parse purchase date and time from the receipt.
	parsedDateAndTime, err := ParseReceiptDateTime(receipt)
	if err != nil {
		return 0, err // Return error if parsing fails
	}

	// Rule 1.
	points += CalculateAlphaNumericRetailerNamePoints(receipt)

	// Rule 2.
	roundDollarPoints, err := AddPointsForRoundDollarTotal(receipt)
	if err != nil {
		return 0, err
	}
	points += roundDollarPoints

	// Rule 3.
	multipleOfQuarterPoints, err := AddPointsForMultipleOfQuarter(receipt)
	if err != nil {
		return 0, err
	}
	points += multipleOfQuarterPoints

	// Rule 4.
	points += AddPointsForItemCount(receipt)

	// Rule 5.
	itemPoints, err := AddPointsForItemDescriptions(receipt)
	if err != nil {
		return 0, err
	}
	points += itemPoints

	// Rule 6.
	points += AddPointsForOddDay(parsedDateAndTime)

	// Rule 7.
	points += AddPointsForAfternoonPurchaseTime(parsedDateAndTime)

	return points, nil
}
