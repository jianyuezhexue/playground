package qrCode

import (
	"image"
	"image/color"
)

type Circle struct {
	P image.Point
	R int
}

func (c *Circle) ColorModel() color.Model {
	return color.AlphaModel
}

func (c *Circle) Bounds() image.Rectangle {
	return image.Rect(c.P.X-c.R, c.P.Y-c.R, c.P.X+c.R, c.P.Y+c.R)
}

func (c *Circle) At(x, y int) color.Color {
	xx, yy, rr := float64(x-c.P.X), float64(y-c.P.Y), float64(c.R)
	if xx*xx+yy*yy < rr*rr {
		return color.Alpha{A: 255}
	}
	return color.Alpha{}
}
