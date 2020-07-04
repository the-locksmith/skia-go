package skia

/*
#include "skia.h"
*/
import "C"
import (
	"image/color"
	"unsafe"
)

type skColor = C.sk_color_t
type SkColor = skColor

func GetBitShift() (a, r, g, b int) {
	var ca, cr, cg, cb C.int
	C.sk_color_get_bit_shift(&ca, &cr, &cg, &cb)
	return int(ca), int(cr), int(cg), int(cb)
}

func makeColor(r, g, b, a uint8) uint32 {
	return (uint32(a) << 24) | (uint32(r) << 16) | (uint32(g) << 8) | uint32(b)
}
func NewColor(c color.Color) SkColor {
	if c == nil {
		return 0
	}
	rgba := color.NRGBAModel.Convert(c).(color.NRGBA)
	return SkColor(makeColor(rgba.R, rgba.G, rgba.B, rgba.A))
}

type skPmcolor = C.sk_pmcolor_t
type SkPmcolor = skPmcolor

func NewPMColor(c color.Color) SkPmcolor {
	if c == nil {
		return 0
	}
	rgba := color.RGBAModel.Convert(c).(color.RGBA)
	return SkPmcolor(makeColor(rgba.R, rgba.G, rgba.B, rgba.A))
}

type skColor4f = C.sk_color4f_t
type Color4f struct {
	FR float32
	FG float32
	FB float32
	FA float32
}

func NewColor4f(c color.Color) Color4f {
	if c == nil {
		return Color4f{}
	}
	rgba := color.NRGBAModel.Convert(c).(color.NRGBA)
	return Color4f{
		FR: float32(rgba.R) / float32(256),
		FG: float32(rgba.G) / float32(256),
		FB: float32(rgba.B) / float32(256),
		FA: float32(rgba.A) / float32(256),
	}
}
func fromNativeColor4f(v *skColor4f) *Color4f {
	return (*Color4f)(unsafe.Pointer(v))
}
func (v *Color4f) native() *skColor4f {
	return (*skColor4f)(unsafe.Pointer(v))
}

func (c *Color4f) ToColor() SkColor {
	return C.sk_color4f_to_color(c.native())
}

func (c SkColor) ToColor4f() Color4f {
	var v Color4f
	C.sk_color4f_from_color(c, v.native())
	return v
}

func (c SkPmcolor) Unpremultiply() SkColor {
	return C.sk_color_unpremultiply(c)
}

func (c SkColor) Premultiply() SkPmcolor {
	return C.sk_color_premultiply(c)
}
