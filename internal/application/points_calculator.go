// Logic for calculating points based on receipt rules
package application

import (
	"fmt"
	"go-receipt-processor/internal/domain"
	"go-receipt-processor/internal/ports/http"
	"math"
	"strings"
	"time"
)

/*
 * PointsCalculatorImpl is the concrete implementation of the PointsCalculator interface.
 * It contains the logic for calculating points on a receipt
 */
type PointsCalculatorImpl struct{}

/*
 * NewPointsCalculator creates a new instance of PointsCalculatorImpl
 * It's a factory function that returns a POINTER to a new PointsCalculatorImpl
 */
func NewPointsCalculator() http.PointsCalculator {
	return &PointsCalculatorImpl{}
}

/*
 * CalculatePoints is a method on the PointsCalculator that handles tallying up
 * points for a receipt
 */
func (c *PointsCalculatorImpl) CalculatePoints(receipt domain.Receipt) (int, error) {
	points := 0 // Initialize points to 0. This will accumulate points based on various rules.

	// parse the date
	// the way time formatting works here is bonkers. it NEEDS to be "2006-01-02"
	// the GO reference time is Mon Jan 2 15:04:05 2006 MST
	// and you must use this specific time as a template for formatting
	parsedDate, err := time.Parse("2006-01-02", receipt.PurchaseDate)
	if err != nil {
		return 0, fmt.Errorf("invalid purchase date format: %v", err)
	}

	// parse the time
	// the way time formatting works here is bonkers. it NEEDS to be "15:04"
	// the GO reference time is Mon Jan 2 15:04:05 2006 MST
	// and you must use this specific time as a template for formatting
	parsedTime, err := time.Parse("15:04", receipt.PurchaseTime)
	if err != nil {
		return 0, fmt.Errorf("invalid purchase time format: %v", err)
	}

	// Rule 1: One point for each alphanumeric character in the retailer name.
	// First, we remove any spaces from the retailer name using `strings.ReplaceAll`.
	// Then, we count the length of the remaining string and add that to the points.
	points += len(strings.ReplaceAll(receipt.Retailer, " ", ""))

	// Rule 2: 50 points if the total amount is a round dollar value (no cents).
	// This is checked by comparing the total to its floor value (i.e., the value without decimals).
	if receipt.Total == math.Floor(receipt.Total) {
		points += 50
	}

	// Rule 3: 25 points if the total is a multiple of 0.25.
	// We multiply the total by 100 to check if it's a multiple of 25 (i.e., total in cents).
	if int(receipt.Total*100)%25 == 0 {
		points += 25
	}

	// Rule 4: 5 points for every two items on the receipt.
	// The number of items is divided by 2, and the result is multiplied by 5 to calculate the points.
	points += (len(receipt.Items) / 2) * 5

	// Rule 5: For items with descriptions that are a multiple of 3 in length, apply the price multiplier rule.
	// We check if the length of the description (after trimming spaces) is a multiple of 3.
	// If true, we add points based on the price of the item, applying a 20% multiplier.
	for _, item := range receipt.Items {
		if len(strings.TrimSpace(item.ShortDescription))%3 == 0 {
			points += int(math.Ceil(item.Price * 0.2)) // Use math.Ceil to round up the points.
		}
	}

	// Rule 6: 6 points if the purchase day is odd (i.e., 1st, 3rd, 5th, etc. day of the month).
	// We check if the day of the purchase (from `receipt.PurchaseDate`) is an odd number.
	if parsedDate.Day()%2 != 0 {
		points += 6
	}

	// Rule 7: 10 points if the purchase time is between 2:00pm and 4:00pm.
	// We check the hour of the `PurchaseTime` to see if it falls within the 2-4pm range.
	if parsedTime.Hour() >= 14 && parsedTime.Hour() < 16 {
		points += 10
	}

	return points, nil
}
