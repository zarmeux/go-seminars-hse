package tasks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCounter(t *testing.T) {
	tests := []struct {
		name     string
		initial  int
		actions  func(*counter)
		expected int
	}{
		{
			name:     "initial value",
			initial:  0,
			actions:  func(c *counter) {},
			expected: 0,
		},
		{
			name:    "increment once",
			initial: 0,
			actions: func(c *counter) {
				c.Increment()
			},
			expected: 1,
		},
		{
			name:    "decrement once",
			initial: 5,
			actions: func(c *counter) {
				c.Decrement()
			},
			expected: 4,
		},
		{
			name:    "reset counter",
			initial: 10,
			actions: func(c *counter) {
				c.Reset()
			},
			expected: 0,
		},
		{
			name:    "add number",
			initial: 5,
			actions: func(c *counter) {
				c.Add(10)
			},
			expected: 15,
		},
		{
			name:    "subtract number",
			initial: 20,
			actions: func(c *counter) {
				c.Subtract(7)
			},
			expected: 13,
		},
		{
			name:    "multiple operations",
			initial: 0,
			actions: func(c *counter) {
				c.Increment()
				c.Increment()
				c.Add(5)
				c.Decrement()
				c.Subtract(3)
			},
			expected: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cnt := newCounter(tt.initial)
			tt.actions(cnt)

			assert.Equal(t, tt.expected, cnt.GetValue(), "Counter value should match expected")
		})
	}
}
