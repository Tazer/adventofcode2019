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
			input:    "1002,4,3,4,33",
			expected: "1002,4,3,4,99",
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
