package rules

import (
	"github.com/stretchr/testify/assert"
	"go-receipt-processor/internal/application"
	"go-receipt-processor/internal/domain"
	"testing"
)

func TestAddPointsForRoundDollarTotal(t *testing.T) {
	helper := application.NewPointsCalculatorHelper()

	tests := []struct {
		name           string
		receipt        domain.Receipt
		expectedPoints int
		expectedError  bool
	}{
		{
			name: "Round dollar total",
			receipt: domain.Receipt{
				Total: "100.00",
			},
			expectedPoints: 50,
			expectedError:  false,
		},
		{
			name: "Total with cents",
			receipt: domain.Receipt{
				Total: "99.99",
			},
			expectedPoints: 0,
			expectedError:  false,
		},
		{
			name: "Invalid total",
			receipt: domain.Receipt{
				Total: "invalid",
			},
			expectedPoints: 0,
			expectedError:  true, // Error case due to invalid input
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			points, err := helper.AddPointsForRoundDollarTotal(tt.receipt)
			if tt.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tt.expectedPoints, points)
		})
	}
}
