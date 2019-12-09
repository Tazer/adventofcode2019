package main

import (
	"log"
	"strings"
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

func TestCalculateOrbitConnections(t *testing.T) {
	cases := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name: "12 divided by 2",
			input: `COM)B
			B)C
			C)D
			D)E
			E)F
			B)G
			G)H
			D)I
			E)J
			J)K
			K)L`,
			expected: 42,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {

			input := strings.Split(c.input, "\n")
			log.Print(len(input))
			got := calculateOrbitConnections(input)

			assert.Equal(t, c.expected, got)
		})

	}
}
