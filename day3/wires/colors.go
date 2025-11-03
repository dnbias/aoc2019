package wires

import (
	"image/color"
)

type Colors struct {
	Wire1        color.Color
	Wire2        color.Color
	Intersection color.Color
	Background   color.Color
	Grid         color.Color
	CentralPort  color.Color
}

var DefaultColors = Colors{
	Wire1:        color.NRGBA{R: 0, G: 102, B: 204, A: 255}, // Blue
	Wire2:        color.NRGBA{R: 204, G: 0, B: 102, A: 255}, // Red
	Intersection: color.NRGBA{R: 255, G: 204, B: 0, A: 255}, // Yellow
	Background:   color.NRGBA{R: 26, G: 26, B: 26, A: 255},  // Dark gray
	Grid:         color.NRGBA{R: 51, G: 51, B: 51, A: 255},  // Light gray
	CentralPort:  color.NRGBA{R: 0, G: 255, B: 0, A: 255},   // Green
}

var HighContrastColors = Colors{
	Wire1:        color.NRGBA{R: 0, G: 0, B: 255, A: 255},     // Bright Blue
	Wire2:        color.NRGBA{R: 255, G: 0, B: 0, A: 255},     // Bright Red
	Intersection: color.NRGBA{R: 255, G: 255, B: 0, A: 255},   // Bright Yellow
	Background:   color.NRGBA{R: 0, G: 0, B: 0, A: 255},       // Black
	Grid:         color.NRGBA{R: 128, G: 128, B: 128, A: 255}, // Gray
	CentralPort:  color.NRGBA{R: 0, G: 255, B: 0, A: 255},     // Bright Green
}
