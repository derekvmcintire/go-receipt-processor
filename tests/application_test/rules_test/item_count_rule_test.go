package rules

import (
	"github.com/stretchr/testify/assert"
	"go-receipt-processor/internal/application"
	"go-receipt-processor/internal/domain"
	"testing"
)

func TestAddPointsForItemCount(t *testing.T) {
	helper := application.NewPointsCalculatorHelper()

	tests := []struct {
		name           string
		receipt        domain.Receipt
		expectedPoints int
	}{
		{
			name: "Even number of items",
			receipt: domain.Receipt{
				Items: []domain.Item{{}, {}, {}, {}},
			},
			expectedPoints: 10, // 4 items -> 2 pairs, each worth 5 points
		},
		{
			name: "Odd number of items",
			receipt: domain.Receipt{
				Items: []domain.Item{{}, {}, {}, {}, {}},
			},
			expectedPoints: 10, // 5 items -> 2 pairs, each worth 5 points
		},
		{
			name: "No items",
			receipt: domain.Receipt{
				Items: []domain.Item{},
			},
			expectedPoints: 0, // No items, no points
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			points := helper.AddPointsForItemCount(tt.receipt)
			assert.Equal(t, tt.expectedPoints, points)
		})
	}
}
