package rules_test

import (
	"github.com/stretchr/testify/assert"
	"go-receipt-processor/internal/application"
	"go-receipt-processor/internal/domain"
	"testing"
)

func TestAddPointsForRetailerName(t *testing.T) {
	helper := application.NewPointsCalculatorHelper()

	tests := []struct {
		name           string
		receipt        domain.Receipt
		expectedPoints int
	}{
		{
			name: "Valid alphanumeric retailer name",
			receipt: domain.Receipt{
				Retailer: "Store123",
			},
			expectedPoints: 8, // "Store123" has 8 alphanumeric characters
		},
		{
			name: "Retailer with special characters",
			receipt: domain.Receipt{
				Retailer: "Store@123!",
			},
			expectedPoints: 8, // "Store123" has 8 alphanumeric characters (ignoring '@' and '!')
		},
		{
			name: "Empty retailer name",
			receipt: domain.Receipt{
				Retailer: "",
			},
			expectedPoints: 0, // No points for empty name
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			points := helper.AddPointsForRetailerName(tt.receipt)
			assert.Equal(t, tt.expectedPoints, points)
		})
	}
}
