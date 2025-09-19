package tasks

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInvertMap(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    map[string]int
		expected map[int]string
	}{
		{
			name:     "simple inversion",
			input:    map[string]int{"a": 1, "b": 2, "c": 3},
			expected: map[int]string{1: "a", 2: "b", 3: "c"},
		},
		{
			name:     "empty map",
			input:    map[string]int{},
			expected: map[int]string{},
		},
		{
			name:     "duplicate values",
			input:    map[string]int{"a": 1, "b": 1, "c": 2},
			expected: map[int]string{1: "b", 2: "c"}, // Последнее значение перезаписывается
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := invertMap(tt.input)

			if len(result) != len(tt.expected) {
				t.Errorf("Expected %d elements, got %d", len(tt.expected), len(result))
			}

			for k, v := range tt.expected {
				assert.Equal(t, v, result[k], fmt.Sprintf("Expected %s, got %s", tt.expected[k], result[k]))
			}
		})
	}
}
