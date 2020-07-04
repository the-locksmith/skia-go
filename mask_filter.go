package skia

/*
#include "skia.h"
*/
import "C"
import (
	"unsafe"
)

type skMaskfilter = C.sk_maskfilter_t
type SkMaskfilter = skMaskfilter

func (v *SkMaskfilter) Ref() {
	C.sk_maskfilter_ref(v)
}

func (v *SkMaskfilter) Unref() {
	C.sk_maskfilter_unref(v)
}

type skBlurstyle = C.sk_blurstyle_t
type SkBlurstyle = skBlurstyle

const (
	SK_BLUR_STYLE_NORMAL = C.NORMAL_SK_BLUR_STYLE
	SK_BLUR_STYLE_SOLID  = C.SOLID_SK_BLUR_STYLE
	SK_BLUR_STYLE_OUTER  = C.OUTER_SK_BLUR_STYLE
	SK_BLUR_STYLE_INNER  = C.INNER_SK_BLUR_STYLE
)

func NewMaskfilterBlur(style SkBlurstyle, sigma float32) *SkMaskfilter {
	return C.sk_maskfilter_new_blur(style, cfloat_t(sigma))
}

func NewMaskfilterBlurWithFlags(style SkBlurstyle, sigma float32, respectCTM bool) *SkMaskfilter {
	return C.sk_maskfilter_new_blur_with_flags(style, cfloat_t(sigma), cbool_t(respectCTM))
}

func NewMaskfilterTable(table *[256]byte) *SkMaskfilter {
	return C.sk_maskfilter_new_table((*cuchar_t)(unsafe.Pointer(table)))
}

func NewMaskfilterGamma(gamma float32) *SkMaskfilter {
	return C.sk_maskfilter_new_gamma(cfloat_t(gamma))
}

func NewMaskfilterClip(min byte, max byte) *SkMaskfilter {
	return C.sk_maskfilter_new_clip(cuchar_t(min), cuchar_t(max))
}

func NewMaskfilterShader(shader *SkShader) *SkMaskfilter {
	return C.sk_maskfilter_new_shader(shader)
}
