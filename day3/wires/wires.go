package wires

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"

	"aoc2019/day3/wires/geometry"
)

func ReadPathsFromFile(filename string) [][]string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	csvReader := csv.NewReader(file)
	csvReader.Comma = ','

	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	return records
}

func GetWirePath(moves []string) map[geometry.Point]bool {
	path := make(map[geometry.Point]bool)
	currentX, currentY := 0, 0

	for _, move := range moves {
		direction := rune(move[0])
		distance, err := strconv.Atoi(move[1:])
		if err != nil {
			log.Fatal(err)
		}

		for i := 0; i < distance; i++ {
			switch direction {
			case 'R':
				currentX++
			case 'L':
				currentX--
			case 'U':
				currentY++
			case 'D':
				currentY--
			}
			path[geometry.Point{X: currentX, Y: currentY}] = true
		}
	}

	return path
}

func GetOrderedWirePath(moves []string) []geometry.Point {
	path := []geometry.Point{}
	currentX, currentY := 0, 0

	for _, move := range moves {
		direction := rune(move[0])
		distance, err := strconv.Atoi(move[1:])
		if err != nil {
			log.Fatal(err)
		}

		for i := 0; i < distance; i++ {
			switch direction {
			case 'R':
				currentX++
			case 'L':
				currentX--
			case 'U':
				currentY++
			case 'D':
				currentY--
			}
			path = append(path, geometry.Point{X: currentX, Y: currentY})
		}
	}

	return path
}

func Intersections(path1 map[geometry.Point]bool, path2 map[geometry.Point]bool) []geometry.Point {
	intersections := []geometry.Point{}

	for p := range path1 {
		if _, exists := path2[p]; exists {
			intersections = append(intersections, p)
		}
	}

	return intersections
}

func IntersectionsFromSlices(path1 []geometry.Point, path2 []geometry.Point) []geometry.Point {
	path2Set := make(map[geometry.Point]bool)
	for _, p := range path2 {
		path2Set[p] = true
	}

	intersections := []geometry.Point{}
	for _, p := range path1 {
		if path2Set[p] {
			intersections = append(intersections, p)
		}
	}

	return intersections
}
