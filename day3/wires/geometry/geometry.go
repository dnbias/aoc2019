package geometry

import (
	"math"
)

type Point struct {
	X int
	Y int
}

type Bounds struct {
	MinX, MaxX, MinY, MaxY int
}

var CentralPort Point = Point{X: 0, Y: 0}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func ManhattanDistance(pos1 Point, pos2 Point) int {
	return Abs((pos1.X - pos2.X)) + Abs((pos1.Y - pos2.Y))
}

func FindClosestIntersection(intersections []Point) (Point, int) {
	distance := math.MaxInt
	closest_intersection := Point{}
	for _, p := range intersections {
		r := ManhattanDistance(CentralPort, p)
		if r < distance {
			distance = r
			closest_intersection = p
		}
	}
	return closest_intersection, distance
}

// wire1 and wire2 should be ordered
func FindLowestNumberOfStepsIntersection(wire1, wire2 []Point, intersections []Point) (Point, int) {
	steps := math.MaxInt
	closest_intersection := Point{}
	for _, p := range intersections {
		current_steps := 0
		for i, p1 := range wire1 {
			if p1 == p {
				current_steps += i + 1
			}
		}
		for i, p2 := range wire2 {
			if p2 == p {
				current_steps += i + 1
			}
		}
		if current_steps < steps {
			steps = current_steps
			closest_intersection = p
		}
	}
	return closest_intersection, steps
}

func CalculateBounds(path1, path2 map[Point]bool) Bounds {
	bounds := Bounds{
		MinX: 0,
		MaxX: 0,
		MinY: 0,
		MaxY: 0,
	}

	updateBounds := func(p Point) {
		if p.X < bounds.MinX {
			bounds.MinX = p.X
		}
		if p.X > bounds.MaxX {
			bounds.MaxX = p.X
		}
		if p.Y < bounds.MinY {
			bounds.MinY = p.Y
		}
		if p.Y > bounds.MaxY {
			bounds.MaxY = p.Y
		}
	}

	for p := range path1 {
		updateBounds(p)
	}
	for p := range path2 {
		updateBounds(p)
	}

	return bounds
}

func CalculateBoundsFromSlices(path1, path2 []Point) Bounds {
	bounds := Bounds{
		MinX: 0,
		MaxX: 0,
		MinY: 0,
		MaxY: 0,
	}

	updateBounds := func(p Point) {
		if p.X < bounds.MinX {
			bounds.MinX = p.X
		}
		if p.X > bounds.MaxX {
			bounds.MaxX = p.X
		}
		if p.Y < bounds.MinY {
			bounds.MinY = p.Y
		}
		if p.Y > bounds.MaxY {
			bounds.MaxY = p.Y
		}
	}

	for _, p := range path1 {
		updateBounds(p)
	}
	for _, p := range path2 {
		updateBounds(p)
	}

	return bounds
}

func WorldToScreen(world Point, bounds Bounds, scale float64, offset Point, canvasWidth, canvasHeight int) (float32, float32) {
	centerX := float64(canvasWidth) / 2.0
	centerY := float64(canvasHeight) / 2.0

	screenX := centerX + float64(world.X-offset.X)*scale
	screenY := centerY - float64(world.Y-offset.Y)*scale

	return float32(screenX), float32(screenY)
}

func ScreenToWorld(screenX, screenY float32, bounds Bounds, scale float64, offset Point, canvasWidth, canvasHeight int) Point {
	centerX := float64(canvasWidth) / 2.0
	centerY := float64(canvasHeight) / 2.0

	worldX := int((float64(screenX)-centerX)/scale) + offset.X
	worldY := int((centerY-float64(screenY))/scale) + offset.Y

	return Point{X: worldX, Y: worldY}
}

func AutoScaleAndCenter(bounds Bounds, canvasWidth, canvasHeight int, padding float64) (float64, Point) {
	pathWidth := float64(bounds.MaxX - bounds.MinX)
	pathHeight := float64(bounds.MaxY - bounds.MinY)

	scaleX := float64(canvasWidth) * (1 - padding) / pathWidth
	scaleY := float64(canvasHeight) * (1 - padding) / pathHeight

	scale := math.Min(scaleX, scaleY)

	offsetX := bounds.MinX + int(pathWidth/2)
	offsetY := bounds.MinY + int(pathHeight/2)

	return scale, Point{X: offsetX, Y: offsetY}
}
