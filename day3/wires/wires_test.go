package wires

import (
	"aoc2019/day3/wires/geometry"
	"strings"
	"testing"
)

func TestFindClosestIntersection(t *testing.T) {
	testCases := []struct {
		name             string
		path1            string
		path2            string
		expectedDistance int
	}{
		{
			name:             "Test Case 1",
			path1:            "R75,D30,R83,U83,L12,D49,R71,U7,L72",
			path2:            "U62,R66,U55,R34,D71,R55,D58,R83",
			expectedDistance: 159,
		},
		{
			name:             "Test Case 2",
			path1:            "R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51",
			path2:            "U98,R91,D20,R16,D67,R40,U7,R15,U6,R7",
			expectedDistance: 135,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			moves1 := strings.Split(tc.path1, ",")
			moves2 := strings.Split(tc.path2, ",")

			wire1Points := GetWirePath(moves1)
			wire2Points := GetWirePath(moves2)

			intersections := Intersections(wire1Points, wire2Points)
			_, distance := geometry.FindClosestIntersection(intersections)

			if distance != tc.expectedDistance {
				t.Errorf("Expected distance %d, but got %d", tc.expectedDistance, distance)
			}
		})
	}
}
