package application

import (
	"go-receipt-processor/internal/domain"
	"go-receipt-processor/internal/ports/http"
)

// PointsCalculatorImpl implements the PointsCalculator interface,
// responsible for calculating points based on receipt data.
type PointsCalculatorImpl struct {
	helpers http.PointsCalculatorRules // Dependency injection of helper methods
}

// NewPointsCalculator creates and returns a new instance of PointsCalculatorImpl.
func NewPointsCalculator(helpers http.PointsCalculatorRules) http.PointsCalculator {
	return &PointsCalculatorImpl{
		helpers: helpers,
	}
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
	points += c.helpers.AddPointsForRetailerName(receipt)

	// Rule 2.
	roundDollarPoints, err := c.helpers.AddPointsForRoundDollarTotal(receipt)
	if err != nil {
		return 0, err
	}
	points += roundDollarPoints

	// Rule 3.
	multipleOfQuarterPoints, err := c.helpers.AddPointsForMultipleOfQuarter(receipt)
	if err != nil {
		return 0, err
	}
	points += multipleOfQuarterPoints

	// Rule 4.
	points += c.helpers.AddPointsForItemCount(receipt)

	// Rule 5.
	itemPoints, err := c.helpers.AddPointsForItemDescriptions(receipt)
	if err != nil {
		return 0, err
	}
	points += itemPoints

	// Rule 6.
	points += c.helpers.AddPointsForOddDay(parsedDateAndTime)

	// Rule 7.
	points += c.helpers.AddPointsForAfternoonPurchaseTime(parsedDateAndTime)

	return points, nil
}
