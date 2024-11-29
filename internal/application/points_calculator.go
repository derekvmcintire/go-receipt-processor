package application

import (
	"fmt"
	"go-receipt-processor/internal/domain"
	"go-receipt-processor/internal/ports/http"
	"math"
	"strconv"
	"strings"
	"time"
	"unicode"
)

// PointsCalculatorImpl implements the PointsCalculator interface.
// It provides the logic for calculating points based on receipt data.
type PointsCalculatorImpl struct{}

// NewPointsCalculator creates and returns a new instance of PointsCalculatorImpl.
// It implements the PointsCalculator interface.
func NewPointsCalculator() http.PointsCalculator {
	return &PointsCalculatorImpl{}
}

// CalculatePoints calculates the total points for a receipt based on specific business rules.
// It returns:
//   - The total points as an integer.
//   - An error if any issue occurs (e.g., invalid data or date/time parsing failure).
func (c *PointsCalculatorImpl) CalculatePoints(receipt domain.Receipt) (int, error) {
	points := 0 // Initialize the total points to 0.

	// Parse the purchase date and time from the receipt.
	parsedDate, parsedTime, err := c.parseDateTime(receipt)
	if err != nil {
		return 0, err // Return error if date or time parsing fails
	}

	// Apply the business rules and accumulate points
	points += c.calculateAlphaNumericRetailerNamePoints(receipt)

	// Add points if the total is a round dollar amount (no cents)
	roundDollarPoints, err := c.addPointsForRoundDollarTotal(receipt)
	if err != nil {
		return 0, err // Return error if unable to convert the total to a float
	}
	points += roundDollarPoints

	// Add points if the total is a multiple of 0.25
	multipleOfQuarterPoints, err := c.addPointsForMultipleOfQuarter(receipt)
	if err != nil {
		return 0, err // Return error if unable to convert the total to a float
	}
	points += multipleOfQuarterPoints

	// Add points for every two items on the receipt
	points += c.addPointsForItemCount(receipt)

	// Add points based on item descriptions (length multiple of 3)
	itemPoints, err := c.addPointsForItemDescriptions(receipt)
	if err != nil {
		return 0, err // Return error if unable to parse any item price
	}
	points += itemPoints

	// Add points if the purchase day is odd
	points += c.addPointsForOddDay(parsedDate)

	// Add points if the purchase time is in the afternoon (2:00 PM - 4:00 PM)
	points += c.addPointsForAfternoonPurchaseTime(parsedTime)

	return points, nil
}

// parseDateTime parses the purchase date and time from the receipt.
// It returns two time.Time objects (parsedDate and parsedTime), and an error if parsing fails.
func (c *PointsCalculatorImpl) parseDateTime(receipt domain.Receipt) (time.Time, time.Time, error) {
	parsedDate, err := time.Parse("2006-01-02", receipt.PurchaseDate)
	if err != nil {
		return time.Time{}, time.Time{}, fmt.Errorf("invalid purchase date format: %v", err)
	}

	parsedTime, err := time.Parse("15:04", receipt.PurchaseTime)
	if err != nil {
		return time.Time{}, time.Time{}, fmt.Errorf("invalid purchase time format: %v", err)
	}

	return parsedDate, parsedTime, nil
}

// calculateAlphaNumericRetailerNamePoints calculates points based on the number of alphanumeric characters
// in the retailer name, ignoring spaces. It returns the number of points earned.
func (c *PointsCalculatorImpl) calculateAlphaNumericRetailerNamePoints(receipt domain.Receipt) int {
	cleanedName := ""
	// Extract only alphanumeric characters (letters and digits) from the retailer name
	for _, ch := range receipt.Retailer {
		if unicode.IsLetter(ch) || unicode.IsDigit(ch) {
			cleanedName += string(ch)
		}
	}
	// Debug log the cleaned retailer name (alphanumeric characters only)
	fmt.Println("Cleaned retailer name:", cleanedName)
	return len(cleanedName) // Return the count of alphanumeric characters
}

// addPointsForRoundDollarTotal adds points if the total is a round dollar amount (no cents).
// It returns the points earned (50) or 0 if the total is not a round dollar.
func (c *PointsCalculatorImpl) addPointsForRoundDollarTotal(receipt domain.Receipt) (int, error) {
	total, err := strconv.ParseFloat(receipt.Total, 64)
	if err != nil {
		return 0, fmt.Errorf("unable to convert total '%s' to a float: %v", receipt.Total, err)
	}

	// Check if the total is a round dollar amount (no cents)
	if total == math.Floor(total) {
		return 50, nil
	}
	return 0, nil
}

// addPointsForMultipleOfQuarter adds points if the total is a multiple of 0.25.
// It returns 25 points if the total is a multiple of 0.25, otherwise 0 points.
func (c *PointsCalculatorImpl) addPointsForMultipleOfQuarter(receipt domain.Receipt) (int, error) {
	total, err := strconv.ParseFloat(receipt.Total, 64)
	if err != nil {
		return 0, fmt.Errorf("unable to convert total '%s' to a float: %v", receipt.Total, err)
	}

	// Check if the total is a multiple of 0.25 (rounded to two decimal places)
	if int(total*100)%25 == 0 {
		return 25, nil
	}
	return 0, nil
}

// addPointsForItemCount adds points for every two items on the receipt.
// It returns the total points earned, based on the number of items.
func (c *PointsCalculatorImpl) addPointsForItemCount(receipt domain.Receipt) int {
	// Every two items contribute 5 points
	return (len(receipt.Items) / 2) * 5
}

// addPointsForItemDescriptions adds points for items with descriptions whose lengths are multiples of 3.
// It calculates points by multiplying the item price by 0.2 and rounding up to the nearest integer.
// It returns an error if there is an issue parsing the price of any item.
func (c *PointsCalculatorImpl) addPointsForItemDescriptions(receipt domain.Receipt) (int, error) {
	points := 0
	for _, item := range receipt.Items {
		// Attempt to parse the item price
		price, err := strconv.ParseFloat(item.Price, 64)
		if err != nil {
			return 0, fmt.Errorf("unable to convert item price '%s' to a float: %v", item.Price, err)
		}

		// Add points if the item description length is a multiple of 3
		if len(strings.TrimSpace(item.ShortDescription))%3 == 0 {
			// Round up to the nearest integer for points based on item price
			points += int(math.Ceil(price * 0.2))
		}
	}
	return points, nil
}

// addPointsForOddDay adds points if the purchase day is an odd number.
// It returns 6 points if the purchase day is odd, otherwise 0.
func (c *PointsCalculatorImpl) addPointsForOddDay(parsedDate time.Time) int {
	if parsedDate.Day()%2 != 0 {
		return 6
	}
	return 0
}

// addPointsForAfternoonPurchaseTime adds points if the purchase time is between 2:00 PM and 4:00 PM.
// It returns 10 points if the purchase time is within the specified range, otherwise 0.
func (c *PointsCalculatorImpl) addPointsForAfternoonPurchaseTime(parsedTime time.Time) int {
	if parsedTime.Hour() >= 14 && parsedTime.Hour() < 16 {
		return 10
	}
	return 0
}
