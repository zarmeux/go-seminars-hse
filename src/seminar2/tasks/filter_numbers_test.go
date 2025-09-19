package tasks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilterNumbers(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		numbers   []int
		predicate func(int) bool
		expected  []int
	}{
		{
			name:      "even numbers",
			numbers:   []int{1, 2, 3, 4, 5, 6},
			predicate: func(n int) bool { return n%2 == 0 },
			expected:  []int{2, 4, 6},
		},
		{
			name:      "numbers greater than 3",
			numbers:   []int{1, 2, 3, 4, 5},
			predicate: func(n int) bool { return n > 3 },
			expected:  []int{4, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := FilterNumbers(tt.numbers, tt.predicate)
			assert.Equal(t, tt.expected, result)
		})
	}
}
