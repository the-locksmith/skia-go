package skia

/*
#include "skia.h"
*/
import "C"
import (
	"unsafe"
)

type skColorfilter = C.sk_colorfilter_t
type SkColorfilter = skColorfilter

func (c *SkColorfilter) Unref() {
	C.sk_colorfilter_unref(c)
}

func NewColorFilterMode(c SkColor, mode SkBlendmode) *SkColorfilter {
	return C.sk_colorfilter_new_mode(c, mode)
}

func NewColorFilterLighting(mul SkColor, add SkColor) *SkColorfilter {
	return C.sk_colorfilter_new_lighting(mul, add)
}

func NewColorFilterCompose(outer *SkColorfilter, inner *SkColorfilter) *SkColorfilter {
	return C.sk_colorfilter_new_compose(outer, inner)
}

func NewColorFilterColorMatrix(array *[20]float32) *SkColorfilter {
	return C.sk_colorfilter_new_color_matrix((*cfloat_t)(unsafe.Pointer(array)))
}

func NewColorColorFilterLuma() *SkColorfilter {
	return C.sk_colorfilter_new_luma_color()
}

type skHighcontrastconfigInvertstyle = C.sk_highcontrastconfig_invertstyle_t
type SkHighcontrastconfigInvertstyle = skHighcontrastconfigInvertstyle

const (
	SK_HIGH_CONTRAST_CONFIG_INVERT_STYLE_NO_INVERT         = C.NO_INVERT_SK_HIGH_CONTRAST_CONFIG_INVERT_STYLE
	SK_HIGH_CONTRAST_CONFIG_INVERT_STYLE_INVERT_BRIGHTNESS = C.INVERT_BRIGHTNESS_SK_HIGH_CONTRAST_CONFIG_INVERT_STYLE
	SK_HIGH_CONTRAST_CONFIG_INVERT_STYLE_INVERT_LIGHTNESS  = C.INVERT_LIGHTNESS_SK_HIGH_CONTRAST_CONFIG_INVERT_STYLE
)

type skHighcontrastconfig = C.sk_highcontrastconfig_t
type Highcontrastconfig struct {
	FGrayscale   bool
	FInvertStyle SkHighcontrastconfigInvertstyle
	FContrast    float32
}

func fromNativeHighcontrastconfig(v *skHighcontrastconfig) *Highcontrastconfig {
	return (*Highcontrastconfig)(unsafe.Pointer(v))
}
func (v *Highcontrastconfig) native() *skHighcontrastconfig {
	return (*skHighcontrastconfig)(unsafe.Pointer(v))
}

func NewColorFilterHighContrast(config *Highcontrastconfig) *SkColorfilter {
	return C.sk_colorfilter_new_high_contrast(config.native())
}

func NewColorFilterTable(table *[256]byte) *SkColorfilter {
	return C.sk_colorfilter_new_table((*cuchar_t)(unsafe.Pointer(table)))
}

func NewColorFilterTableArgb(tableA *[256]byte, tableR *[256]byte, tableG *[256]byte, tableB *[256]byte) *SkColorfilter {
	return C.sk_colorfilter_new_table_argb((*cuchar_t)(unsafe.Pointer(tableA)), (*cuchar_t)(unsafe.Pointer(tableR)), (*cuchar_t)(unsafe.Pointer(tableG)), (*cuchar_t)(unsafe.Pointer(tableB)))
}
