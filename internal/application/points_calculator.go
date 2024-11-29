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
// It provides the concrete logic for calculating points based on receipt data.
type PointsCalculatorImpl struct{}

// NewPointsCalculator is a factory function that creates and returns a new instance
// of PointsCalculatorImpl, implementing the PointsCalculator interface.
func NewPointsCalculator() http.PointsCalculator {
	return &PointsCalculatorImpl{}
}

// CalculatePoints calculates the points for a receipt based on specific business rules.
// It returns:
//   - The total points as an integer.
//   - An error if the input receipt data is invalid (e.g., date/time parsing fails).
func (c *PointsCalculatorImpl) CalculatePoints(receipt domain.Receipt) (int, error) {
	points := 0 // Initialize the total points to 0.

	// Parse the purchase date and time
	parsedDate, parsedTime, err := c.parseDateTime(receipt)
	if err != nil {
		return 0, err
	}

	// Apply all rules to the receipt and accumulate points
	points += c.calculateAlphaNumericRetailerNamePoints(receipt)

	roundDollarPoints, err := c.addPointsForRoundDollarTotal(receipt)

	if err != nil {
		return 0, err // Return the error if item descriptions parsing fails
	}

	points += roundDollarPoints

	multipleOfQuarterPoints, err := c.addPointsForMultipleOfQuarter(receipt)

	if err != nil {
		return 0, err // Return the error if item descriptions parsing fails
	}

	points += multipleOfQuarterPoints

	points += c.addPointsForItemCount(receipt)

	itemPoints, err := c.addPointsForItemDescriptions(receipt)
	if err != nil {
		return 0, err // Return the error if item descriptions parsing fails
	}
	points += itemPoints

	points += c.addPointsForOddDay(parsedDate)
	points += c.addPointsForAfternoonPurchaseTime(parsedTime)

	return points, nil
}

// parseDateTime parses the date and time from the receipt.
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

// calculateAlphaNumericRetailerNamePoints adds points based on the retailer name length.
func (c *PointsCalculatorImpl) calculateAlphaNumericRetailerNamePoints(receipt domain.Receipt) int {
	cleanedName := ""
	for _, ch := range receipt.Retailer {
		if unicode.IsLetter(ch) || unicode.IsDigit(ch) {
			cleanedName += string(ch)
		}
	}
	fmt.Println("Cleaned retailer name:", cleanedName) // Debug log
	return len(cleanedName)
}

// addPointsForRoundDollarTotal adds points if the total is a round dollar amount (no cents).
func (c *PointsCalculatorImpl) addPointsForRoundDollarTotal(receipt domain.Receipt) (int, error) {
	total, err := strconv.ParseFloat(receipt.Total, 64)
	if err != nil {
		return 0, fmt.Errorf("unable to convert total '%s' to a float: %v", receipt.Total, err)
	}

	if total == math.Floor(total) {
		return 50, nil
	}
	return 0, nil
}

// addPointsForMultipleOfQuarter adds points if the total is a multiple of 0.25.
func (c *PointsCalculatorImpl) addPointsForMultipleOfQuarter(receipt domain.Receipt) (int, error) {
	total, err := strconv.ParseFloat(receipt.Total, 64)
	if err != nil {
		return 0, fmt.Errorf("unable to convert total '%s' to a float: %v", receipt.Total, err)
	}

	if int(total*100)%25 == 0 {
		return 25, nil
	}
	return 0, nil
}

// addPointsForItemCount adds points for every two items on the receipt.
func (c *PointsCalculatorImpl) addPointsForItemCount(receipt domain.Receipt) int {
	return (len(receipt.Items) / 2) * 5
}

// addPointsForItemDescriptions adds points for items with descriptions whose lengths are multiples of 3.
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
			points += int(math.Ceil(price * 0.2)) // Add points based on item price
		}
	}
	return points, nil
}

// addPointsForOddDay adds points if the purchase day is an odd number.
func (c *PointsCalculatorImpl) addPointsForOddDay(parsedDate time.Time) int {
	if parsedDate.Day()%2 != 0 {
		return 6
	}
	return 0
}

// addPointsForAfternoonPurchaseTime adds points if the purchase time is between 2:00 PM and 4:00 PM.
func (c *PointsCalculatorImpl) addPointsForAfternoonPurchaseTime(parsedTime time.Time) int {
	if parsedTime.Hour() >= 14 && parsedTime.Hour() < 16 {
		return 10
	}
	return 0
}
