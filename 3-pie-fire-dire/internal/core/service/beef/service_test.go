package beef_test

import (
	"context"
	"errors"
	"pie-fire-dire/generated/mocks"
	"pie-fire-dire/internal/core/service/beef"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBeefSummary(t *testing.T) {
	testCases := []struct {
		name          string
		text          string
		expected      map[string]int
		expectedError error
	}{
		{
			name:     "Normal case",
			text:     "t-bone-bone t-bone",
			expected: map[string]int{"t-bone-bone": 1, "t-bone": 1},
		},
		{
			name:     "No valid words",
			text:     ". . . ! ! !",
			expected: map[string]int{},
		},
		{
			name:     "Duplicate word",
			text:     "t-bone-bone-.. . ..-t-bone-bone",
			expected: map[string]int{"t-bone-bone": 2},
		},
		{
			name:          "Error case",
			text:          "",
			expectedError: errors.New("error message"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var (
				ctx     = context.Background()
				mockAPI = mocks.NewBaconIpsumAPI(t)
				service = beef.NewService(beef.Dependencies{
					BaconIpsumAPI: mockAPI,
				})
			)

			mockAPI.On("GetText", ctx).Return(tc.text, tc.expectedError)

			result, err := service.Summary(ctx)

			if tc.expectedError != nil {
				assert.Error(t, err)
				assert.EqualError(t, err, tc.expectedError.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, len(tc.expected), len(result))
				for word, count := range tc.expected {
					assert.Equal(t, count, result[word])
				}
			}
		})
	}
}
