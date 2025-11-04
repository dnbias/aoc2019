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

	wire1 := wires.GetWirePath(path1)
	wire2 := wires.GetWirePath(path2)

	intersections := wires.Intersections(wire1, wire2)

	fmt.Println("Intersections:", intersections)

	intersection, distance := geometry.FindClosestIntersection(intersections)

	fmt.Println("Closest intersection:", intersection, "with distance:", distance)
}
