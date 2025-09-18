package tasks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseString(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"empty string", "", ""},
		{"single character", "a", "a"},
		{"palindrome", "radar", "radar"},
		{"normal string", "hello", "olleh"},
		{"unicode string", "привет", "тевирп"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := ReverseString(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}
