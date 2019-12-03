package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCrossinglinesBasic(t *testing.T) {
	cases := []struct {
		name     string
		line1    string
		line2    string
		expected int
	}{
		{
			name:     "example 1",
			line1:    "R8,U5,L5,D3",
			line2:    "U7,R6,D4,L4",
			expected: 6,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got, _, err := calculateManhattanDistance(c.line1, c.line2)
			assert.Nil(t, err)
			assert.Equal(t, c.expected, got)
		})

	}
}

func TestCrossinglines(t *testing.T) {
	cases := []struct {
		name     string
		line1    string
		line2    string
		expected int
	}{
		{
			name:     "example 1",
			line1:    "R8,U5,L5,D3",
			line2:    "U7,R6,D4,L4",
			expected: 6,
		},
		{
			name:     "example 2",
			line1:    "R75,D30,R83,U83,L12,D49,R71,U7,L72",
			line2:    "U62,R66,U55,R34,D71,R55,D58,R83",
			expected: 159,
		},
		{
			name:     "example 3",
			line1:    "R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51",
			line2:    "U98,R91,D20,R16,D67,R40,U7,R15,U6,R7",
			expected: 135,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got, _, err := calculateManhattanDistance(c.line1, c.line2)
			assert.Nil(t, err)
			assert.Equal(t, c.expected, got)
		})

	}
}

func TestCrossinglinesSimple(t *testing.T) {
	cases := []struct {
		name     string
		line1    string
		line2    string
		expected int
	}{
		{
			name:     "example 1",
			line1:    "R2,U2,D2",
			line2:    "U2,R8",
			expected: 6,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got, _, err := calculateManhattanDistance(c.line1, c.line2)
			assert.Nil(t, err)
			assert.Equal(t, c.expected, got)
		})

	}
}

func TestSteps(t *testing.T) {
	cases := []struct {
		name     string
		line1    string
		line2    string
		expected int
	}{
		{
			name:     "example 1",
			line1:    "R75,D30,R83,U83,L12,D49,R71,U7,L72",
			line2:    "U62,R66,U55,R34,D71,R55,D58,R83",
			expected: 610,
		},
		{
			name:     "example 1",
			line1:    "R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51",
			line2:    "U98,R91,D20,R16,D67,R40,U7,R15,U6,R7",
			expected: 410,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			_, got, err := calculateManhattanDistance(c.line1, c.line2)
			assert.Nil(t, err)
			assert.Equal(t, c.expected, got)
		})

	}
}
