package skia

/*
#include "skia.h"
*/
import "C"
import (
	"github.com/zhuah/skia-go/internal"
)

type skImagefilter = C.sk_imagefilter_t
type SkImagefilter = skImagefilter

type skImagefilterCroprect = C.sk_imagefilter_croprect_t
type SkImagefilterCroprect = skImagefilterCroprect

func NewSkImagefilterCroprect() *SkImagefilterCroprect {
	return C.sk_imagefilter_croprect_new()
}

func NewSkImagefilterCroprectWithRect(rect *Rect, flags uint32) *SkImagefilterCroprect {
	return C.sk_imagefilter_croprect_new_with_rect(rect.native(), cuint32_t(flags))
}

func (v *SkImagefilterCroprect) Destroy() {
	C.sk_imagefilter_croprect_destructor(v)
}

func (v *SkImagefilterCroprect) GetRect(rect *Rect) {
	C.sk_imagefilter_croprect_get_rect(v, rect.native())
}

func (v *SkImagefilterCroprect) GetFlags() uint32 {
	return uint32(C.sk_imagefilter_croprect_get_flags(v))
}

func (v *SkImagefilter) Unref() {
	C.sk_imagefilter_unref(v)
}

func NewImagefilterMatrix(matrix *Matrix, quality SkFilterQuality, input *SkImagefilter) *SkImagefilter {
	return C.sk_imagefilter_new_matrix(matrix.native(), quality, input)
}

func NewImagefilterAlphaThreshold(region *SkRegion, innerThreshold float32, outerThreshold float32, input *SkImagefilter) *SkImagefilter {
	return C.sk_imagefilter_new_alpha_threshold(region, cfloat_t(innerThreshold), cfloat_t(outerThreshold), input)
}

func NewImagefilterBlur(sigmaX float32, sigmaY float32, input *SkImagefilter, cropRect *SkImagefilterCroprect) *SkImagefilter {
	return C.sk_imagefilter_new_blur(cfloat_t(sigmaY), cfloat_t(sigmaY), input, cropRect)
}

func NewImagefilterColorFilter(cf *SkColorfilter, input *SkImagefilter, cropRect *SkImagefilterCroprect) *SkImagefilter {
	return C.sk_imagefilter_new_color_filter(cf, input, cropRect)
}

func NewImagefilterCompose(outer *SkImagefilter, inner *SkImagefilter) *SkImagefilter {
	return C.sk_imagefilter_new_compose(outer, inner)
}

type skDropShadowImageFilterShadowMode = C.sk_drop_shadow_image_filter_shadow_mode_t
type SkDropShadowImageFilterShadowMode = skDropShadowImageFilterShadowMode

const (
	SK_DROP_SHADOW_IMAGE_FILTER_SHADOW_MODE_DRAW_SHADOW_AND_FOREGROUND = C.DRAW_SHADOW_AND_FOREGROUND_SK_DROP_SHADOW_IMAGE_FILTER_SHADOW_MODE
	SK_DROP_SHADOW_IMAGE_FILTER_SHADOW_MODE_DRAW_SHADOW_ONLY           = C.DRAW_SHADOW_ONLY_SK_DROP_SHADOW_IMAGE_FILTER_SHADOW_MODE
)

type skDisplacementMapEffectChannelSelectorType = C.sk_displacement_map_effect_channel_selector_type_t
type SkDisplacementMapEffectChannelSelectorType = skDisplacementMapEffectChannelSelectorType

const (
	SK_DISPLACEMENT_MAP_EFFECT_CHANNEL_SELECTOR_TYPE_UNKNOWN = C.UNKNOWN_SK_DISPLACEMENT_MAP_EFFECT_CHANNEL_SELECTOR_TYPE
	SK_DISPLACEMENT_MAP_EFFECT_CHANNEL_SELECTOR_TYPE_R       = C.R_SK_DISPLACEMENT_MAP_EFFECT_CHANNEL_SELECTOR_TYPE
	SK_DISPLACEMENT_MAP_EFFECT_CHANNEL_SELECTOR_TYPE_G       = C.G_SK_DISPLACEMENT_MAP_EFFECT_CHANNEL_SELECTOR_TYPE
	SK_DISPLACEMENT_MAP_EFFECT_CHANNEL_SELECTOR_TYPE_B       = C.B_SK_DISPLACEMENT_MAP_EFFECT_CHANNEL_SELECTOR_TYPE
	SK_DISPLACEMENT_MAP_EFFECT_CHANNEL_SELECTOR_TYPE_A       = C.A_SK_DISPLACEMENT_MAP_EFFECT_CHANNEL_SELECTOR_TYPE
)

type skMatrixConvolutionTilemode = C.sk_matrix_convolution_tilemode_t
type SkMatrixConvolutionTilemode = skMatrixConvolutionTilemode

const (
	SK_MATRIX_CONVOLUTION_TILEMODE_CLAMP          = C.CLAMP_SK_MATRIX_CONVOLUTION_TILEMODE
	SK_MATRIX_CONVOLUTION_TILEMODE_REPEAT         = C.REPEAT_SK_MATRIX_CONVOLUTION_TILEMODE
	SK_MATRIX_CONVOLUTION_TILEMODE_CLAMP_TO_BLACK = C.CLAMP_TO_BLACK_SK_MATRIX_CONVOLUTION_TILEMODE
)

func NewImagefilterDisplacementMapEffect(xChannelSelector SkDisplacementMapEffectChannelSelectorType, yChannelSelector SkDisplacementMapEffectChannelSelectorType, scale float32, displacement *SkImagefilter, color *SkImagefilter, cropRect *SkImagefilterCroprect) *SkImagefilter {
	return C.sk_imagefilter_new_displacement_map_effect(xChannelSelector, yChannelSelector, cfloat_t(scale), displacement, color, cropRect)
}

func NewImagefilterDropShadow(dx float32, dy float32, sigmaX float32, sigmaY float32, color SkColor, shadowMode SkDropShadowImageFilterShadowMode, input *SkImagefilter, cropRect *SkImagefilterCroprect) *SkImagefilter {
	return C.sk_imagefilter_new_drop_shadow(cfloat_t(dx), cfloat_t(dy), cfloat_t(sigmaY), cfloat_t(sigmaY), color, shadowMode, input, cropRect)
}

func NewImagefilterDistantLitDiffuse(direction *Point3, lightColor SkColor, surfaceScale float32, kd float32, input *SkImagefilter, cropRect *SkImagefilterCroprect) *SkImagefilter {
	return C.sk_imagefilter_new_distant_lit_diffuse(direction.native(), lightColor, cfloat_t(surfaceScale), cfloat_t(kd), input, cropRect)
}

func NewImagefilterPointLitDiffuse(location *Point3, lightColor SkColor, surfaceScale float32, kd float32, input *SkImagefilter, cropRect *SkImagefilterCroprect) *SkImagefilter {
	return C.sk_imagefilter_new_point_lit_diffuse(location.native(), lightColor, cfloat_t(surfaceScale), cfloat_t(kd), input, cropRect)
}

func NewImagefilterSpotLitDiffuse(location *Point3, target *Point3, specularExponent float32, cutoffAngle float32, lightColor SkColor, surfaceScale float32, kd float32, input *SkImagefilter, cropRect *SkImagefilterCroprect) *SkImagefilter {
	return C.sk_imagefilter_new_spot_lit_diffuse(location.native(), target.native(), cfloat_t(specularExponent), cfloat_t(cutoffAngle), lightColor, cfloat_t(surfaceScale), cfloat_t(kd), input, cropRect)
}

func NewImagefilterDistantLitSpecular(direction *Point3, lightColor SkColor, surfaceScale float32, ks float32, shininess float32, input *SkImagefilter, cropRect *SkImagefilterCroprect) *SkImagefilter {
	return C.sk_imagefilter_new_distant_lit_specular(direction.native(), lightColor, cfloat_t(surfaceScale), cfloat_t(ks), cfloat_t(shininess), input, cropRect)
}

func NewImagefilterPointLitSpecular(location *Point3, lightColor SkColor, surfaceScale float32, ks float32, shininess float32, input *SkImagefilter, cropRect *SkImagefilterCroprect) *SkImagefilter {
	return C.sk_imagefilter_new_point_lit_specular(location.native(), lightColor, cfloat_t(surfaceScale), cfloat_t(ks), cfloat_t(shininess), input, cropRect)
}

func NewImagefilterSpotLitSpecular(location *Point3, target *Point3, specularExponent float32, cutoffAngle float32, lightColor SkColor, surfaceScale float32, ks float32, shininess float32, input *SkImagefilter, cropRect *SkImagefilterCroprect) *SkImagefilter {
	return C.sk_imagefilter_new_spot_lit_specular(location.native(), target.native(), cfloat_t(specularExponent), cfloat_t(cutoffAngle), lightColor, cfloat_t(surfaceScale), cfloat_t(ks), cfloat_t(shininess), input, cropRect)
}

func NewImagefilterMagnifier(src *Rect, inset float32, input *SkImagefilter, cropRect *SkImagefilterCroprect) *SkImagefilter {
	return C.sk_imagefilter_new_magnifier(src.native(), cfloat_t(inset), input, cropRect)
}

func NewImagefilterMatrixConvolution(kernelSize Isize, kernel []float32, gain float32, bias float32, kernelOffset *Ipoint, tileMode SkMatrixConvolutionTilemode, convolveAlpha bool, input *SkImagefilter, cropRect *SkImagefilterCroprect) *SkImagefilter {
	return C.sk_imagefilter_new_matrix_convolution(kernelSize.native(), (*cfloat_t)(internal.SliceDataPtr(kernel)), cfloat_t(gain), cfloat_t(bias), kernelOffset.native(), tileMode, cbool_t(convolveAlpha), input, cropRect)
}

func NewImagefilterMerge(filters []*SkImagefilter, cropRect *SkImagefilterCroprect) *SkImagefilter {
	return C.sk_imagefilter_new_merge((**SkImagefilter)(internal.SliceDataPtr(filters)), cint32_t(len(filters)), cropRect)
}

func NewImagefilterDilate(radiusX int32, radiusY int32, input *SkImagefilter, cropRect *SkImagefilterCroprect) *SkImagefilter {
	return C.sk_imagefilter_new_dilate(cint32_t(radiusX), cint32_t(radiusY), input, cropRect)
}

func NewImagefilterErode(radiusX int32, radiusY int32, input *SkImagefilter, cropRect *SkImagefilterCroprect) *SkImagefilter {
	return C.sk_imagefilter_new_erode(cint32_t(radiusX), cint32_t(radiusY), input, cropRect)
}

func NewImagefilterOffset(dx float32, dy float32, input *SkImagefilter, cropRect *SkImagefilterCroprect) *SkImagefilter {
	return C.sk_imagefilter_new_offset(cfloat_t(dx), cfloat_t(dy), input, cropRect)
}

func NewImagefilterPicture(picture *SkPicture) *SkImagefilter {
	return C.sk_imagefilter_new_picture(picture)
}

func NewImagefilterPictureWithCroprect(picture *SkPicture, cropRect *Rect) *SkImagefilter {
	return C.sk_imagefilter_new_picture_with_croprect(picture, cropRect.native())
}

func NewImagefilterTile(src *Rect, dst *Rect, input *SkImagefilter) *SkImagefilter {
	return C.sk_imagefilter_new_tile(src.native(), dst.native(), input)
}

func NewImagefilterXfermode(mode SkBlendmode, background *SkImagefilter, foreground *SkImagefilter, cropRect *SkImagefilterCroprect) *SkImagefilter {
	return C.sk_imagefilter_new_xfermode(mode, background, foreground, cropRect)
}

func NewImagefilterArithmetic(k1 float32, k2 float32, k3 float32, k4 float32, enforcePMColor bool, background *SkImagefilter, foreground *SkImagefilter, cropRect *SkImagefilterCroprect) *SkImagefilter {
	return C.sk_imagefilter_new_arithmetic(cfloat_t(k1), cfloat_t(k2), cfloat_t(k3), cfloat_t(k4), cbool_t(enforcePMColor), background, foreground, cropRect)
}

func NewImagefilterImageSource(image *SkImage, srcRect *Rect, dstRect *Rect, filterQuality SkFilterQuality) *SkImagefilter {
	return C.sk_imagefilter_new_image_source(image, srcRect.native(), dstRect.native(), filterQuality)
}

func NewImagefilterImageSourceDefault(image *SkImage) *SkImagefilter {
	return C.sk_imagefilter_new_image_source_default(image)
}

func NewImagefilterPaint(paint *SkPaint, cropRect *SkImagefilterCroprect) *SkImagefilter {
	return C.sk_imagefilter_new_paint(paint, cropRect)
}
