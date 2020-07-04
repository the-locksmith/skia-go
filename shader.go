package skia

/*
#include "skia.h"
*/
import "C"
import (
	"unsafe"

	"github.com/zhuah/skia-go/internal"
)

type skShader = C.sk_shader_t
type SkShader = skShader

func (v *SkShader) Ref() {
	C.sk_shader_ref(v)
}

func (v *SkShader) Unref() {
	C.sk_shader_unref(v)
}

func (v *SkShader) WithLocalMatrix(localMatrix *Matrix) *SkShader {
	return C.sk_shader_with_local_matrix(v, localMatrix.native())
}

func (v *SkShader) WithColorFilter(filter *SkColorfilter) *SkShader {
	return C.sk_shader_with_color_filter(v, filter)
}

func NewShaderEmpty() *SkShader {
	return C.sk_shader_new_empty()
}

func NewShaderColor(color SkColor) *SkShader {
	return C.sk_shader_new_color(color)
}

func NewShaderColor4f(color *Color4f, colorspace *SkColorspace) *SkShader {
	return C.sk_shader_new_color4f(color.native(), colorspace)
}

func NewShaderBlend(mode SkBlendmode, dst *SkShader, src *SkShader, localMatrix *Matrix) *SkShader {
	return C.sk_shader_new_blend(mode, dst, src, localMatrix.native())
}

func NewShaderLerp(t float32, dst *SkShader, src *SkShader, localMatrix *Matrix) *SkShader {
	return C.sk_shader_new_lerp(cfloat_t(t), dst, src, localMatrix.native())
}

type skShaderTilemode = C.sk_shader_tilemode_t
type SkShaderTilemode = skShaderTilemode

const (
	SK_SHADER_TILEMODE_CLAMP  = C.CLAMP_SK_SHADER_TILEMODE
	SK_SHADER_TILEMODE_REPEAT = C.REPEAT_SK_SHADER_TILEMODE
	SK_SHADER_TILEMODE_MIRROR = C.MIRROR_SK_SHADER_TILEMODE
	SK_SHADER_TILEMODE_DECAL  = C.DECAL_SK_SHADER_TILEMODE
)

func NewShaderLinearGradient(points [2]Point, colors []SkColor, colorPos []float32, tileMode SkShaderTilemode, localMatrix *Matrix) *SkShader {
	return C.sk_shader_new_linear_gradient((*skPoint)(unsafe.Pointer(&points)), (*SkColor)(internal.SliceDataPtr(colors)), (*cfloat_t)(internal.SliceDataPtr(colorPos)), cint_t(len(colors)), tileMode, localMatrix.native())
}

func NewShaderLinearGradientColor4f(points *[2]Point, colors []Color4f, colorspace *SkColorspace, colorPos []float32, tileMode SkShaderTilemode, localMatrix *Matrix) *SkShader {
	return C.sk_shader_new_linear_gradient_color4f((*skPoint)(unsafe.Pointer(&points)), (*skColor4f)(internal.SliceDataPtr(colors)), colorspace, (*cfloat_t)(internal.SliceDataPtr(colorPos)), cint_t(len(colors)), tileMode, localMatrix.native())
}

func NewShaderRadialGradient(center *Point, radius float32, colors []SkColor, colorPos []float32, tileMode SkShaderTilemode, localMatrix *Matrix) *SkShader {
	return C.sk_shader_new_radial_gradient(center.native(), cfloat_t(radius), (*SkColor)(internal.SliceDataPtr(colors)), (*cfloat_t)(internal.SliceDataPtr(colorPos)), cint_t(len(colors)), tileMode, localMatrix.native())
}

func NewShaderRadialGradientColor4f(center *Point, radius float32, colors []Color4f, colorspace *SkColorspace, colorPos []float32, tileMode SkShaderTilemode, localMatrix *Matrix) *SkShader {
	return C.sk_shader_new_radial_gradient_color4f(center.native(), cfloat_t(radius), (*skColor4f)(internal.SliceDataPtr(colors)), colorspace, (*cfloat_t)(internal.SliceDataPtr(colorPos)), cint_t(len(colors)), tileMode, localMatrix.native())
}

func NewShaderSweepGradient(center *Point, colors []SkColor, colorPos []float32, tileMode SkShaderTilemode, startAngle float32, endAngle float32, localMatrix *Matrix) *SkShader {
	return C.sk_shader_new_sweep_gradient(center.native(), (*SkColor)(internal.SliceDataPtr(colors)), (*cfloat_t)(internal.SliceDataPtr(colorPos)), cint_t(len(colors)), tileMode, cfloat_t(startAngle), cfloat_t(endAngle), localMatrix.native())
}

func NewShaderSweepGradientColor4f(center *Point, colors []Color4f, colorspace *SkColorspace, colorPos []float32, tileMode SkShaderTilemode, startAngle float32, endAngle float32, localMatrix *Matrix) *SkShader {
	return C.sk_shader_new_sweep_gradient_color4f(center.native(), (*skColor4f)(internal.SliceDataPtr(colors)), colorspace, (*cfloat_t)(internal.SliceDataPtr(colorPos)), cint_t(len(colors)), tileMode, cfloat_t(startAngle), cfloat_t(endAngle), localMatrix.native())
}

func NewShaderTwoPointConicalGradient(start *Point, startRadius float32, end *Point, endRadius float32, colors []SkColor, colorPos []float32, tileMode SkShaderTilemode, localMatrix *Matrix) *SkShader {
	return C.sk_shader_new_two_point_conical_gradient(start.native(), cfloat_t(startRadius), end.native(), cfloat_t(endRadius), (*SkColor)(internal.SliceDataPtr(colors)), (*cfloat_t)(internal.SliceDataPtr(colorPos)), cint_t(len(colors)), tileMode, localMatrix.native())
}

func NewShaderTwoPointConicalGradientColor4f(start *Point, startRadius float32, end *Point, endRadius float32, colors []Color4f, colorspace *SkColorspace, colorPos []float32, tileMode SkShaderTilemode, localMatrix *Matrix) *SkShader {
	return C.sk_shader_new_two_point_conical_gradient_color4f(start.native(), cfloat_t(startRadius), end.native(), cfloat_t(endRadius), (*skColor4f)(internal.SliceDataPtr(colors)), colorspace, (*cfloat_t)(internal.SliceDataPtr(colorPos)), cint_t(len(colors)), tileMode, localMatrix.native())
}

func NewShaderPerlinNoiseFractalNoise(baseFrequencyX float32, baseFrequencyY float32, numOctaves int32, seed float32, tileSize *Isize) *SkShader {
	return C.sk_shader_new_perlin_noise_fractal_noise(cfloat_t(baseFrequencyX), cfloat_t(baseFrequencyY), cint32_t(numOctaves), cfloat_t(seed), tileSize.native())
}

func NewShaderPerlinNoiseTurbulence(baseFrequencyX float32, baseFrequencyY float32, numOctaves int32, seed float32, tileSize *Isize) *SkShader {
	return C.sk_shader_new_perlin_noise_turbulence(cfloat_t(baseFrequencyX), cfloat_t(baseFrequencyY), cint32_t(numOctaves), cfloat_t(seed), tileSize.native())
}

func NewShaderPerlinNoiseImprovedNoise(baseFrequencyX float32, baseFrequencyY float32, numOctaves int32, z float32) *SkShader {
	return C.sk_shader_new_perlin_noise_improved_noise(cfloat_t(baseFrequencyX), cfloat_t(baseFrequencyY), cint32_t(numOctaves), cfloat_t(z))
}
