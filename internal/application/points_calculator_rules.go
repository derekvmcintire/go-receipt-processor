// points_calculator_helpers.go
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

// Helper functions for calculating points based on specific business rules:
//   1. One point for every alphanumeric character in the retailer name.
//   2. 50 points if the total is a round dollar amount (no cents).
//   3. 25 points if the total is a multiple of 0.25.
//   4. 5 points for every two items on the receipt.
//   5. If the length of the trimmed item description is a multiple of 3,
//     multiply the item price by 0.2 and round up to the nearest integer to determine the points for that item.
//   6. 6 points if the day in the purchase date is odd (e.g., 1st, 3rd, 5th, etc.).
//   7. 10 points if the purchase time is between 2:00 PM and 4:00 PM (inclusive).

// PointsCalculatorRulesImpl is the concrete implementation of the PointsCalculatorHelpers interface.
type PointsCalculatorRulesImpl struct{}

// NewPointsCalculatorHelper creates and returns a new instance of PointsCalculatorHelperImpl.
func NewPointsCalculatorHelper() http.PointsCalculatorRules {
	return &PointsCalculatorRulesImpl{}
}

// Rule 1: Calculate points based on alphanumeric characters in retailer name
func (h *PointsCalculatorRulesImpl) AddPointsForRetailerName(receipt domain.Receipt) int {
	cleanedName := ""
	for _, ch := range receipt.Retailer {
		if unicode.IsLetter(ch) || unicode.IsDigit(ch) {
			cleanedName += string(ch)
		}
	}
	return len(cleanedName)
}

// Rule 2: Points for round dollar amounts
func (h *PointsCalculatorRulesImpl) AddPointsForRoundDollarTotal(receipt domain.Receipt) (int, error) {
	total, err := strconv.ParseFloat(receipt.Total, 64)
	if err != nil {
		return 0, fmt.Errorf("unable to convert total '%s' to a float: %v", receipt.Total, err)
	}
	if total == math.Floor(total) {
		return 50, nil
	}
	return 0, nil
}

// Rule 3: Points for totals that are multiples of 0.25
func (h *PointsCalculatorRulesImpl) AddPointsForMultipleOfQuarter(receipt domain.Receipt) (int, error) {
	total, err := strconv.ParseFloat(receipt.Total, 64)
	if err != nil {
		return 0, fmt.Errorf("unable to convert total '%s' to a float: %v", receipt.Total, err)
	}
	if int(total*100)%25 == 0 {
		return 25, nil
	}
	return 0, nil
}

// Rule 4: Points based on item count
func (h *PointsCalculatorRulesImpl) AddPointsForItemCount(receipt domain.Receipt) int {
	return (len(receipt.Items) / 2) * 5
}

// Rule 5: Points based on item descriptions
func (h *PointsCalculatorRulesImpl) AddPointsForItemDescriptions(receipt domain.Receipt) (int, error) {
	points := 0
	for _, item := range receipt.Items {
		price, err := strconv.ParseFloat(item.Price, 64)
		if err != nil {
			return 0, fmt.Errorf("unable to convert item price '%s' to a float: %v", item.Price, err)
		}
		if len(strings.TrimSpace(item.ShortDescription))%3 == 0 {
			points += int(math.Ceil(price * 0.2))
		}
	}
	return points, nil
}

// Rule 6: Points for odd days
func (h *PointsCalculatorRulesImpl) AddPointsForOddDay(parsedDate time.Time) int {
	if parsedDate.Day()%2 != 0 {
		return 6
	}
	return 0
}

// Rule 7: Points for afternoon purchase time
func (h *PointsCalculatorRulesImpl) AddPointsForAfternoonPurchaseTime(parsedTime time.Time) int {
	if parsedTime.Hour() >= 14 && parsedTime.Hour() < 16 {
		return 10
	}
	return 0
}
