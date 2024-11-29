package http

import (
	"go-receipt-processor/internal/domain"
	"time"
)

// PointsCalculatorRules is the interface that defines the helper methods for calculating points.
type PointsCalculatorRules interface {
	AddPointsForRetailerName(receipt domain.Receipt) int
	AddPointsForRoundDollarTotal(receipt domain.Receipt) (int, error)
	AddPointsForMultipleOfQuarter(receipt domain.Receipt) (int, error)
	AddPointsForItemCount(receipt domain.Receipt) int
	AddPointsForItemDescriptions(receipt domain.Receipt) (int, error)
	AddPointsForOddDay(parsedDateAndTime time.Time) int
	AddPointsForAfternoonPurchaseTime(parsedDateAndTime time.Time) int
}
