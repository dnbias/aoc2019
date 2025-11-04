package wires

import (
	"image"
	"image/color"
	"image/png"
	"math"
	"os"

	"aoc2019/day3/wires/geometry"
)

type Renderer struct {
	Path1         []geometry.Point
	Path2         []geometry.Point
	Intersections []geometry.Point
	Bounds        geometry.Bounds
	Colors        Colors
	Scale         float64
	Offset        geometry.Point
	CanvasWidth   int
	CanvasHeight  int
	ShowWire1     bool
	ShowWire2     bool
}

func NewRenderer(path1, path2 []geometry.Point, intersections []geometry.Point, bounds geometry.Bounds, colors Colors, scale float64, offset geometry.Point, canvasWidth, canvasHeight int) *Renderer {
	return &Renderer{
		Path1:         path1,
		Path2:         path2,
		Intersections: intersections,
		Bounds:        bounds,
		Colors:        colors,
		Scale:         scale,
		Offset:        offset,
		CanvasWidth:   canvasWidth,
		CanvasHeight:  canvasHeight,
		ShowWire1:     true,
		ShowWire2:     true,
	}
}

func (r *Renderer) RenderImage() *image.RGBA {
	img := image.NewRGBA(image.Rect(0, 0, r.CanvasWidth, r.CanvasHeight))

	r.drawBackground(img)

	if r.ShowWire1 {
		r.drawWire(img, r.Path1, r.Colors.Wire1, 4)
	}

	if r.ShowWire2 {
		r.drawWire(img, r.Path2, r.Colors.Wire2, 4)
	}

	r.drawIntersections(img)
	r.drawCentralPort(img)

	return img
}

func (r *Renderer) drawBackground(img *image.RGBA) {
	for y := 0; y < img.Bounds().Dy(); y++ {
		for x := 0; x < img.Bounds().Dx(); x++ {
			img.Set(x, y, r.Colors.Background)
		}
	}
}

func (r *Renderer) drawLine(img *image.RGBA, x0, y0, x1, y1 int, c color.Color, thickness int) {
	dx := geometry.Abs(x1 - x0)
	dy := -geometry.Abs(y1 - y0)
	sx := -1
	if x0 < x1 {
		sx = 1
	}
	sy := -1
	if y0 < y1 {
		sy = 1
	}
	err := dx + dy

	for {
		r.drawCircle(img, x0, y0, thickness/2, c)
		if x0 == x1 && y0 == y1 {
			break
		}
		e2 := 2 * err
		if e2 >= dy {
			err += dy
			x0 += sx
		}
		if e2 <= dx {
			err += dx
			y0 += sy
		}
	}
}

func (r *Renderer) drawWire(img *image.RGBA, path []geometry.Point, wireColor color.Color, thickness int) {
	if len(path) == 0 {
		return
	}

	for i := 1; i < len(path); i++ {
		x0, y0 := geometry.WorldToScreen(path[i-1], r.Bounds, r.Scale, r.Offset, r.CanvasWidth, r.CanvasHeight)
		x1, y1 := geometry.WorldToScreen(path[i], r.Bounds, r.Scale, r.Offset, r.CanvasWidth, r.CanvasHeight)
		r.drawLine(img, int(x0), int(y0), int(x1), int(y1), wireColor, thickness)
	}
}



func (r *Renderer) drawIntersections(img *image.RGBA) {
	for _, point := range r.Intersections {
		screenX, screenY := geometry.WorldToScreen(point, r.Bounds, r.Scale, r.Offset, r.CanvasWidth, r.CanvasHeight)

		if screenX >= 0 && screenX < float32(img.Bounds().Dx()) &&
			screenY >= 0 && screenY < float32(img.Bounds().Dy()) {

			r.drawCircle(img, int(screenX), int(screenY), 3, r.Colors.Intersection)
		}
	}
}

func (r *Renderer) drawCircle(img *image.RGBA, centerX, centerY, radius int, c color.Color) {
	for y := centerY - radius; y <= centerY+radius; y++ {
		for x := centerX - radius; x <= centerX+radius; x++ {
			distance := math.Sqrt(float64((x-centerX)*(x-centerX) + (y-centerY)*(y-centerY)))
			if distance <= float64(radius) {
				if x >= 0 && x < img.Bounds().Dx() && y >= 0 && y < img.Bounds().Dy() {
					img.Set(x, y, c)
				}
			}
		}
	}
}

func (r *Renderer) drawCentralPort(img *image.RGBA) {
	screenX, screenY := geometry.WorldToScreen(geometry.Point{X: 0, Y: 0}, r.Bounds, r.Scale, r.Offset, r.CanvasWidth, r.CanvasHeight)

	if screenX >= 0 && screenX < float32(img.Bounds().Dx()) &&
		screenY >= 0 && screenY < float32(img.Bounds().Dy()) {

		r.drawCircle(img, int(screenX), int(screenY), 4, r.Colors.CentralPort)
	}
}



func (r *Renderer) ExportToPNG(filename string) error {
	img := r.RenderImage()

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	return png.Encode(file, img)
}
