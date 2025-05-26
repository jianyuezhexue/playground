package qrCode

import (
	"image"
	"image/color"
)

type Rectangle struct {
	P image.Point
	W int
	H int
}

func (r *Rectangle) ColorModel() color.Model {
	return color.AlphaModel
}

func (r *Rectangle) Bounds() image.Rectangle {
	return image.Rect(r.P.X-r.W, r.P.Y-r.H, r.P.X+r.W, r.P.Y+r.H)
}

func (r *Rectangle) At(x, y int) color.Color {
	xx, yy, ww, hh := float64(x-r.P.X), float64(y-r.P.Y), float64(r.W), float64(r.H)
	if xx*xx+yy*yy < ww*ww+hh*hh {
		return color.Alpha{A: 255}
	}
	return color.Alpha{}
}
