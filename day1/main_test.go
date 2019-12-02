package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalculateFuel(t *testing.T) {
	cases := []struct {
		name     string
		mass     int
		expected int
	}{
		{
			name:     "12 divided by 2",
			mass:     12,
			expected: 2,
		},
		{
			name:     "14 should be 2",
			mass:     14,
			expected: 2,
		},
		{
			name:     "1969 should be 654",
			mass:     1969,
			expected: 654,
		},
		{
			name:     "100756 should be 33583",
			mass:     100756,
			expected: 33583,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := calculateFuelForModule(c.mass)

			assert.Equal(t, c.expected, got)
		})

	}
}

func TestCalculateFuelWithFuelMass(t *testing.T) {
	cases := []struct {
		name     string
		mass     int
		expected int
	}{
		{
			name:     "12 divided by 2",
			mass:     12,
			expected: 2,
		},
		{
			name:     "1969 should be 654",
			mass:     1969,
			expected: 966,
		},
		{
			name:     "100756 should be 50346",
			mass:     100756,
			expected: 50346,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := calculateFuelForModuleWithFuelMass(c.mass)

			assert.Equal(t, c.expected, got)
		})

	}
}
