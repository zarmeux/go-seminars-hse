package tasks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountVowels(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{"no vowels", "bcdfg", 0},
		{"all vowels", "aeiou", 5},
		{"mixed case", "aEiOu", 5},
		{"with consonants", "hello world", 3},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := CountVowels(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}
