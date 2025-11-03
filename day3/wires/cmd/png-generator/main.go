package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"aoc2019/day3/wires"
	"aoc2019/day3/wires/geometry"
)

func main() {
	var (
		file1      = flag.String("file1", "", "Path to first wire path file")
		file2      = flag.String("file2", "", "Path to second wire path file")
		path1      = flag.String("path1", "", "First wire path (comma-separated)")
		path2      = flag.String("path2", "", "Second wire path (comma-separated)")
		testData   = flag.Bool("test", false, "Use test data")
		outputFile = flag.String("out", "wires.png", "Output PNG file name")
		width      = flag.Int("width", 1200, "Width of the output image")
		height     = flag.Int("height", 800, "Height of the output image")
	)
	flag.Parse()

	var path1Moves, path2Moves []string

	if *testData {
		path1Moves = strings.Split("R75,D30,R83,U83,L12,D49,R71,U7,L72", ",")
		path2Moves = strings.Split("U62,R66,U55,R34,D71,R55,D58,R83", ",")
	} else if *file1 != "" && *file2 != "" {
		paths := wires.ReadPathsFromFile(*file1)
		if len(paths) > 0 {
			path1Moves = paths[0]
		}
		paths = wires.ReadPathsFromFile(*file2)
		if len(paths) > 0 {
			path2Moves = paths[0]
		}
	} else if *path1 != "" && *path2 != "" {
		path1Moves = strings.Split(*path1, ",")
		path2Moves = strings.Split(*path2, ",")
	} else {
		fmt.Println("Usage: go run wires/cmd/png-generator/main.go [flags]")
		fmt.Println("Please provide an input source.")
		fmt.Println("Flags:")
		flag.PrintDefaults()
		return
	}

	if len(path1Moves) == 0 || len(path2Moves) == 0 {
		log.Fatal("Failed to load wire paths. Please provide input via -test, -file, or -path flags.")
	}

	// --- Core Logic ---
	wire1Path := wires.GetOrderedWirePath(path1Moves)
	wire2Path := wires.GetOrderedWirePath(path2Moves)
	intersections := wires.IntersectionsFromSlices(wire1Path, wire2Path)

	bounds := geometry.CalculateBoundsFromSlices(wire1Path, wire2Path)

	scale, offset := geometry.AutoScaleAndCenter(bounds, *width, *height, 0.1)

	renderer := wires.NewRenderer(
		wire1Path,
		wire2Path,
		intersections,
		bounds,
		wires.DefaultColors,
		scale,
		offset,
		*width,
		*height,
	)

	err := renderer.ExportToPNG(*outputFile)
	if err != nil {
		log.Fatalf("Failed to export PNG: %v", err)
	}

	fmt.Printf("Successfully generated wire visualization to %s\n", *outputFile)
}
