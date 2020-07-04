package skia

/*
#include "skia.h"
*/
import "C"
import (
	"github.com/zhuah/skia-go/internal"
)

type skPathEffect = C.sk_path_effect_t
type SkPathEffect = skPathEffect

func (v *SkPathEffect) Unref() {
	C.sk_path_effect_unref(v)
}

func NewPathEffectCompose(outer *SkPathEffect, inner *SkPathEffect) *SkPathEffect {
	return C.sk_path_effect_create_compose(outer, inner)
}

func NewPathEffectSum(first *SkPathEffect, second *SkPathEffect) *SkPathEffect {
	return C.sk_path_effect_create_sum(first, second)
}

func NewPathEffectDiscrete(segLength float32, deviation float32, seedAssist uint32) *SkPathEffect {
	return C.sk_path_effect_create_discrete(cfloat_t(segLength), cfloat_t(deviation), cuint_t(seedAssist))
}

func NewPathEffectCorner(radius float32) *SkPathEffect {
	return C.sk_path_effect_create_corner(cfloat_t(radius))
}

type skPathEffect1dStyle = C.sk_path_effect_1d_style_t
type SkPathEffect1dStyle = skPathEffect1dStyle

const (
	SK_PATH_EFFECT_1D_STYLE_TRANSLATE = C.TRANSLATE_SK_PATH_EFFECT_1D_STYLE
	SK_PATH_EFFECT_1D_STYLE_ROTATE    = C.ROTATE_SK_PATH_EFFECT_1D_STYLE
	SK_PATH_EFFECT_1D_STYLE_MORPH     = C.MORPH_SK_PATH_EFFECT_1D_STYLE
)

type skPathEffectTrimMode = C.sk_path_effect_trim_mode_t
type SkPathEffectTrimMode = skPathEffectTrimMode

const (
	SK_PATH_EFFECT_TRIM_MODE_NORMAL   = C.NORMAL_SK_PATH_EFFECT_TRIM_MODE
	SK_PATH_EFFECT_TRIM_MODE_INVERTED = C.INVERTED_SK_PATH_EFFECT_TRIM_MODE
)

func NewPathEffect1dPath(path *SkPath, advance float32, phase float32, style SkPathEffect1dStyle) *SkPathEffect {
	return C.sk_path_effect_create_1d_path(path, cfloat_t(advance), cfloat_t(phase), style)
}

func NewPathEffect2dLine(width float32, matrix *Matrix) *SkPathEffect {
	return C.sk_path_effect_create_2d_line(cfloat_t(width), matrix.native())
}

func NewPathEffect2dPath(matrix *Matrix, path *SkPath) *SkPathEffect {
	return C.sk_path_effect_create_2d_path(matrix.native(), path)
}

func NewPathEffectDash(intervals []float32, phase float32) *SkPathEffect {
	return C.sk_path_effect_create_dash((*cfloat_t)(internal.SliceDataPtr(intervals)), cint_t(len(intervals)), cfloat_t(phase))
}

func NewPathEffectTrim(start float32, stop float32, mode SkPathEffectTrimMode) *SkPathEffect {
	return C.sk_path_effect_create_trim(cfloat_t(start), cfloat_t(stop), mode)
}
