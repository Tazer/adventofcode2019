package main

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCrossinglinesBasic(t *testing.T) {
	cases := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "example 1",
			input:    "112233",
			expected: true,
		},
		{
			name:     "example 2",
			input:    "123444",
			expected: false,
		},
		{
			name:     "example 3",
			input:    "111122",
			expected: true,
		},
		{
			name:     "example 4",
			input:    "111122",
			expected: true,
		},
		{
			name:     "example 5",
			input:    "221444",
			expected: true,
		},
		{
			name:     "example 6",
			input:    "42244",
			expected: true,
		},
		{
			name:     "example 7",
			input:    "24144",
			expected: true,
		},
		{
			name:     "example 8",
			input:    "44144",
			expected: true,
		},
		{
			name:     "example 9",
			input:    "111114",
			expected: false,
		},
		{
			name:     "example 10",
			input:    "411111",
			expected: false,
		},
		{
			name:     "example 11",
			input:    "411122",
			expected: true,
		},
		{
			name:     "example 12",
			input:    "111111",
			expected: false,
		},
		{
			name:     "example 13",
			input:    "255515",
			expected: false,
		},
		{
			name:     "example 14",
			input:    "331331",
			expected: true,
		},
		{
			name:     "example 15",
			input:    "331133",
			expected: true,
		},
		{
			name:     "example 16",
			input:    "688999",
			expected: true,
		},
		{
			name:     "example 17",
			input:    "441111",
			expected: true,
		},
		{
			name:     "example 18",
			input:    "421111",
			expected: false,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			res := false
			got := getIncorrectList(c.input)

			log.Printf("got: %+v , %s", got, c.name)

			isValid := false
			for k, v := range got {
				if v == 2 {
					log.Printf("correct %v: %d", k, v)
					isValid = true
				}
			}
			if isValid {
				res = true
			} else {
				res = false
			}

			assert.Equal(t, c.expected, res)
		})

	}
}
