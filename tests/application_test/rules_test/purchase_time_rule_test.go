package rules

import (
	"go-receipt-processor/internal/application"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAddPointsForAfternoonPurchaseTime(t *testing.T) {
	helper := application.NewPointsCalculatorHelper()

	tests := []struct {
		name           string
		parsedTime     time.Time
		expectedPoints int
	}{
		{
			name:           "Purchase time in afternoon (between 2:00 PM and 4:00 PM)",
			parsedTime:     time.Date(2024, time.March, 3, 14, 30, 0, 0, time.UTC), // 2:30 PM
			expectedPoints: 10,
		},
		{
			name:           "Purchase time exactly at 2:00 PM",
			parsedTime:     time.Date(2024, time.March, 3, 14, 0, 0, 0, time.UTC), // 2:00 PM
			expectedPoints: 10,
		},
		{
			name:           "Purchase time exactly at 4:00 PM",
			parsedTime:     time.Date(2024, time.March, 3, 16, 0, 0, 0, time.UTC), // 4:00 PM
			expectedPoints: 0,                                                     // 4:00 PM is not within the range (2:00 PM to 4:00 PM)
		},
		{
			name:           "Purchase time outside afternoon range",
			parsedTime:     time.Date(2024, time.March, 3, 10, 30, 0, 0, time.UTC), // 10:30 AM
			expectedPoints: 0,                                                      // 10:30 AM is outside the afternoon range
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			points := helper.AddPointsForAfternoonPurchaseTime(tt.parsedTime)
			assert.Equal(t, tt.expectedPoints, points)
		})
	}
}
