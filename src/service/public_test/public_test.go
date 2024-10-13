package public_test

import (
	"encoding/json"
	"mechta-tz.github.com/src/config"
	"mechta-tz.github.com/src/domain"
	"mechta-tz.github.com/src/service"
	"testing"
)

func TestCalculateJson(t *testing.T) {
	config.MainConfig = &config.Config{
		WorkersCount: 4,
	}

	tests := []struct {
		name        string
		input       []domain.TwoNumber
		expectedSum int
		expectError bool
	}{
		{
			name:        "Empty array",
			input:       []domain.TwoNumber{},
			expectedSum: 0,
			expectError: false,
		},
		{
			name: "Negative values",
			input: []domain.TwoNumber{
				{A: -5, B: -3},
				{A: -2, B: -4},
			},
			expectedSum: -14,
			expectError: false,
		},
		{
			name: "Positive values",
			input: []domain.TwoNumber{
				{A: 3, B: 5},
				{A: 10, B: 20},
			},
			expectedSum: 38,
			expectError: false,
		},
		{
			name: "Mixed values",
			input: []domain.TwoNumber{
				{A: 1, B: -1},
				{A: 2, B: 2},
				{A: -3, B: 3},
			},
			expectedSum: 4,
			expectError: false,
		},
		{
			name:        "Large number of elements",
			input:       generateLargeTestData(1000000),
			expectedSum: 0,
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, _ := json.Marshal(tt.input)

			result, err := service.CalculateJson(data)

			if (err != nil) != tt.expectError {
				t.Fatalf("CalculateJson returned an error: %v", err)
			}

			if !tt.expectError && *result != tt.expectedSum {
				t.Errorf("Expected sum %d, but got %d", tt.expectedSum, *result)
			}
		})
	}
}

func generateLargeTestData(n int) []domain.TwoNumber {
	data := make([]domain.TwoNumber, n)
	for i := 0; i < n; i++ {
		data[i] = domain.TwoNumber{
			A: i % 10,
			B: (i * -1) % 10,
		}
	}
	return data
}

func BenchmarkCalculateJson(b *testing.B) {
	config.MainConfig = &config.Config{
		WorkersCount: 4,
	}

	numbers := make([]domain.TwoNumber, 1000000)
	for i := 0; i < len(numbers); i++ {
		numbers[i] = domain.TwoNumber{
			A: i % 10,
			B: (i * -1) % 10,
		}
	}

	data, _ := json.Marshal(numbers)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := service.CalculateJson(data)
		if err != nil {
			b.Fatalf("CalculateJson returned an error: %v", err)
		}
	}
}
