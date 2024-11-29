package rules

import (
	"github.com/stretchr/testify/assert"
	"go-receipt-processor/internal/application"
	"go-receipt-processor/internal/domain"
	"testing"
)

func TestAddPointsForMultipleOfQuarter(t *testing.T) {
	helper := application.NewPointsCalculatorHelper()

	tests := []struct {
		name           string
		receipt        domain.Receipt
		expectedPoints int
		expectedError  bool
	}{
		{
			name: "Total is a multiple of 0.25",
			receipt: domain.Receipt{
				Total: "25.00",
			},
			expectedPoints: 25, // 25.00 is a multiple of 0.25
			expectedError:  false,
		},
		{
			name: "Total is not a multiple of 0.25",
			receipt: domain.Receipt{
				Total: "25.30",
			},
			expectedPoints: 0, // 25.30 is not divisible by 0.25
			expectedError:  false,
		},
		{
			name: "Total is a multiple of 0.25 (with cents)",
			receipt: domain.Receipt{
				Total: "25.75",
			},
			expectedPoints: 25, // 25.75 is a multiple of 0.25
			expectedError:  false,
		},
		{
			name: "Invalid total",
			receipt: domain.Receipt{
				Total: "invalid",
			},
			expectedPoints: 0,
			expectedError:  true, // Error due to invalid total input
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			points, err := helper.AddPointsForMultipleOfQuarter(tt.receipt)
			if tt.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tt.expectedPoints, points)
		})
	}
}
