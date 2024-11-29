package application

import (
	"fmt"
	"go-receipt-processor/internal/domain"
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

// Rule 1.
// CalculateAlphaNumericRetailerNamePoints calculates points based on the number of alphanumeric characters
// (letters and digits) in the retailer's name, ignoring spaces or other special characters.
func CalculateAlphaNumericRetailerNamePoints(receipt domain.Receipt) int {
	cleanedName := ""
	for _, ch := range receipt.Retailer {
		if unicode.IsLetter(ch) || unicode.IsDigit(ch) {
			cleanedName += string(ch)
		}
	}
	return len(cleanedName)
}

// Rule 2.
// AddPointsForRoundDollarTotal checks if the receipt's total is a round dollar amount (no cents).
// If the total is a whole number (i.e., the cents portion is 0), it awards 50 points.
func AddPointsForRoundDollarTotal(receipt domain.Receipt) (int, error) {
	total, err := strconv.ParseFloat(receipt.Total, 64)
	if err != nil {
		return 0, fmt.Errorf("unable to convert total '%s' to a float: %v", receipt.Total, err)
	}
	if total == math.Floor(total) {
		return 50, nil
	}
	return 0, nil
}

// Rule 3.
// AddPointsForMultipleOfQuarter checks if the receipt's total is a multiple of 0.25.
// If the total is divisible by 0.25, it awards 25 points.
//
// Parameters:
//   - receipt: The domain.Receipt object containing receipt details.
//
// Returns: A time.TIme value representing the the exact time of the purchase
//   - points: The number of points awarded.
//   - err: An error if calculating points fails
func AddPointsForMultipleOfQuarter(receipt domain.Receipt) (int, error) {
	total, err := strconv.ParseFloat(receipt.Total, 64)
	if err != nil {
		return 0, fmt.Errorf("unable to convert total '%s' to a float: %v", receipt.Total, err)
	}
	if int(total*100)%25 == 0 {
		return 25, nil
	}
	return 0, nil
}

// Rule 4.
// AddPointsForItemCount awards 5 points for every two items on the receipt.
// For example, a receipt with 4 items will receive 10 points, while a receipt with 5 items will receive 10 points as well.
//
// Parameters:
//   - receipt: The domain.Receipt object containing receipt details.
//
// Returns: A time.TIme value representing the the exact time of the purchase
//   - points: The number of points awarded.
//   - err: An error if calculating points fails
func AddPointsForItemCount(receipt domain.Receipt) int {
	return (len(receipt.Items) / 2) * 5
}

// Rule 5.
// AddPointsForItemDescriptions checks each item description in the receipt. If the length of the
// trimmed description is a multiple of 3, it awards points based on the item price.
// The points for each item are calculated by multiplying the price by 0.2 and rounding up to the nearest integer.
//
// Parameters:
//   - receipt: The domain.Receipt object containing receipt details.
//
// Returns: A time.TIme value representing the the exact time of the purchase
//   - points: The number of points awarded.
//   - err: An error if calculating points fails
func AddPointsForItemDescriptions(receipt domain.Receipt) (int, error) {
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

// Rule 6.
// AddPointsForOddDay checks if the day of the month for the purchase date is odd (e.g., 1st, 3rd, 5th, etc.)
// If the day is odd, it awards 6 points.
//
// Parameters:
//   - parsedDate: The time.Time object that represents purchase time.
//
// Returns: A time.TIme value representing the the exact time of the purchase
//   - points: The number of points awarded.
//   - err: An error if calculating points fails
func AddPointsForOddDay(parsedDate time.Time) int {
	if parsedDate.Day()%2 != 0 {
		return 6
	}
	return 0
}

// Rule 7.
// AddPointsForAfternoonPurchaseTime checks if the purchase time is between 2:00 PM and 4:00 PM (inclusive).
// If the purchase time is in the afternoon, it awards 10 points.
//
// Parameters:
//   - parsedDate: The time.Time object that represents purchase time.
//
// Returns: A time.TIme value representing the the exact time of the purchase
//   - points: The number of points awarded.
//   - err: An error if calculating points fails
func AddPointsForAfternoonPurchaseTime(parsedTime time.Time) int {
	if parsedTime.Hour() >= 14 && parsedTime.Hour() < 16 {
		return 10
	}
	return 0
}
