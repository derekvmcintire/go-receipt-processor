package rules

import (
	"go-receipt-processor/internal/application"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAddPointsForOddDay(t *testing.T) {
	helper := application.NewPointsCalculatorHelper()

	tests := []struct {
		name           string
		parsedDate     time.Time
		expectedPoints int
	}{
		{
			name:           "Odd day",
			parsedDate:     time.Date(2024, time.March, 3, 14, 0, 0, 0, time.UTC),
			expectedPoints: 6, // 3rd is an odd day
		},
		{
			name:           "Even day",
			parsedDate:     time.Date(2024, time.March, 4, 14, 0, 0, 0, time.UTC),
			expectedPoints: 0, // 4th is an even day
		},
		{
			name:           "Edge case: first day of the month",
			parsedDate:     time.Date(2024, time.March, 1, 14, 0, 0, 0, time.UTC),
			expectedPoints: 6, // 1st is an odd day
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			points := helper.AddPointsForOddDay(tt.parsedDate)
			assert.Equal(t, tt.expectedPoints, points)
		})
	}
}
