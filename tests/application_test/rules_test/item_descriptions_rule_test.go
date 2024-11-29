package rules

import (
	"github.com/stretchr/testify/assert"
	"go-receipt-processor/internal/application"
	"go-receipt-processor/internal/domain"
	"testing"
)

func TestAddPointsForItemDescriptions(t *testing.T) {
	helper := application.NewPointsCalculatorHelper()

	tests := []struct {
		name           string
		receipt        domain.Receipt
		expectedPoints int
		expectedError  bool
	}{
		{
			name: "Valid description with multiple of 3 length",
			receipt: domain.Receipt{
				Items: []domain.Item{
					{ShortDescription: "abc", Price: "10.0"},
				},
			},
			expectedPoints: 2, // 10 * 0.2 = 2
			expectedError:  false,
		},
		{
			name: "Description not multiple of 3",
			receipt: domain.Receipt{
				Items: []domain.Item{
					{ShortDescription: "ab", Price: "10.0"},
				},
			},
			expectedPoints: 0,
			expectedError:  false,
		},
		{
			name: "Invalid price",
			receipt: domain.Receipt{
				Items: []domain.Item{
					{ShortDescription: "abc", Price: "invalid"},
				},
			},
			expectedPoints: 0,
			expectedError:  true, // Error due to invalid price
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			points, err := helper.AddPointsForItemDescriptions(tt.receipt)
			if tt.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tt.expectedPoints, points)
		})
	}
}
