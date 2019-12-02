package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntCode(t *testing.T) {
	cases := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "example 1",
			input:    "1,0,0,0,99",
			expected: "2,0,0,0,99",
		},
		{
			name:     "example 2",
			input:    "2,3,0,3,99",
			expected: "2,3,0,6,99",
		},
		{
			name:     "example 3",
			input:    "2,4,4,5,99,0",
			expected: "2,4,4,5,99,9801",
		},
		{
			name:     "example 4",
			input:    "1,1,1,4,99,5,6,0,99",
			expected: "30,1,1,4,2,5,6,0,99",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			intCodes := strings.Split(c.input, ",")
			got := processIntCode(intCodes)

			assert.Equal(t, c.expected, got)
		})

	}
}
