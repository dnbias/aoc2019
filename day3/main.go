package main

import (
	"aoc2019/day3/wires"
	"aoc2019/day3/wires/geometry"
	"fmt"
)

func main() {
	paths := wires.ReadPathsFromFile("input")
	path1 := paths[0]
	path2 := paths[1]

	// part1(path1, path2)
	part2(path1, path2)
}

func part1(path1, path2 []string) {
	wire1 := wires.GetWirePath(path1)
	wire2 := wires.GetWirePath(path2)

	intersections := wires.Intersections(wire1, wire2)

	fmt.Println("Intersections:", intersections)

	intersection, distance := geometry.FindClosestIntersection(intersections)

	fmt.Println("Closest intersection:", intersection, "with distance:", distance)
}

func part2(path1, path2 []string) {
	wire1 := wires.GetOrderedWirePath(path1)
	wire2 := wires.GetOrderedWirePath(path2)

	intersections := wires.IntersectionsFromSlices(wire1, wire2)

	fmt.Println("Intersections:", intersections)

	intersection, steps := geometry.FindLowestNumberOfStepsIntersection(wire1, wire2, intersections)
	fmt.Println("Lowest cumulative steps intersection:", intersection, "with steps:", steps)
}
