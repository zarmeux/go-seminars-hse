package tasks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDivide(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		a           float64
		b           float64
		expected    float64
		expectError bool
	}{
		{"normal division", 10.0, 2.0, 5.0, false},
		{"division by zero", 10.0, 0.0, 0.0, true},
		{"fraction result", 1.0, 2.0, 0.5, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result, err := Divide(tt.a, tt.b)

			if tt.expectError {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), "division by zero")
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, result)
			}
		})
	}
}
