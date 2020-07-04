package skia

/*
#include "skia.h"
*/
import "C"

type skPaint = C.sk_paint_t
type SkPaint = skPaint

func NewPaint() *SkPaint {
	return C.sk_paint_new()
}

func (v *SkPaint) Clone() *SkPaint {
	return C.sk_paint_clone(v)
}

func (v *SkPaint) Destroy() {
	C.sk_paint_delete(v)
}

func (v *SkPaint) Reset() {
	C.sk_paint_reset(v)
}

func (v *SkPaint) IsAntialias() bool {
	return bool(C.sk_paint_is_antialias(v))
}

func (v *SkPaint) SetAntialias(aa bool) {
	C.sk_paint_set_antialias(v, cbool_t(aa))
}

func (v *SkPaint) GetColor() SkColor {
	return C.sk_paint_get_color(v)
}

func (v *SkPaint) SetColor(color SkColor) {
	C.sk_paint_set_color(v, color)
}

type skPaintStyle = C.sk_paint_style_t
type SkPaintStyle = skPaintStyle

const (
	SK_PAINT_STYLE_FILL            = C.FILL_SK_PAINT_STYLE
	SK_PAINT_STYLE_STROKE          = C.STROKE_SK_PAINT_STYLE
	SK_PAINT_STYLE_STROKE_AND_FILL = C.STROKE_AND_FILL_SK_PAINT_STYLE
)

func (v *SkPaint) GetStyle() SkPaintStyle {
	return C.sk_paint_get_style(v)
}

func (v *SkPaint) SetStyle(style SkPaintStyle) {
	C.sk_paint_set_style(v, style)
}

func (v *SkPaint) GetStrokeWidth() float32 {
	return float32(C.sk_paint_get_stroke_width(v))
}

func (v *SkPaint) SetStrokeWidth(width float32) {
	C.sk_paint_set_stroke_width(v, cfloat_t(width))
}

func (v *SkPaint) GetStrokeMiter() float32 {
	return float32(C.sk_paint_get_stroke_miter(v))
}

func (v *SkPaint) SetStrokeMiter(miter float32) {
	C.sk_paint_set_stroke_miter(v, cfloat_t(miter))
}

type skStrokeCap = C.sk_stroke_cap_t
type SkStrokeCap = skStrokeCap

const (
	SK_STROKE_CAP_BUTT   = C.BUTT_SK_STROKE_CAP
	SK_STROKE_CAP_ROUND  = C.ROUND_SK_STROKE_CAP
	SK_STROKE_CAP_SQUARE = C.SQUARE_SK_STROKE_CAP
)

type skStrokeJoin = C.sk_stroke_join_t
type SkStrokeJoin = skStrokeJoin

const (
	SK_STROKE_JOIN_MITER = C.MITER_SK_STROKE_JOIN
	SK_STROKE_JOIN_ROUND = C.ROUND_SK_STROKE_JOIN
	SK_STROKE_JOIN_BEVEL = C.BEVEL_SK_STROKE_JOIN
)

func (v *SkPaint) GetStrokeCap() SkStrokeCap {
	return C.sk_paint_get_stroke_cap(v)
}

func (v *SkPaint) SetStrokeCap(cap SkStrokeCap) {
	C.sk_paint_set_stroke_cap(v, cap)
}

func (v *SkPaint) GetStrokeJoin() SkStrokeJoin {
	return C.sk_paint_get_stroke_join(v)
}

func (v *SkPaint) SetStrokeJoin(join SkStrokeJoin) {
	C.sk_paint_set_stroke_join(v, join)
}

func (v *SkPaint) SetShader(shader *SkShader) {
	C.sk_paint_set_shader(v, shader)
}

func (v *SkPaint) SetMaskfilter(filter *SkMaskfilter) {
	C.sk_paint_set_maskfilter(v, filter)
}

type skBlendmode = C.sk_blendmode_t
type SkBlendmode = skBlendmode

const (
	SK_BLENDMODE_CLEAR      = C.CLEAR_SK_BLENDMODE
	SK_BLENDMODE_SRC        = C.SRC_SK_BLENDMODE
	SK_BLENDMODE_DST        = C.DST_SK_BLENDMODE
	SK_BLENDMODE_SRCOVER    = C.SRCOVER_SK_BLENDMODE
	SK_BLENDMODE_DSTOVER    = C.DSTOVER_SK_BLENDMODE
	SK_BLENDMODE_SRCIN      = C.SRCIN_SK_BLENDMODE
	SK_BLENDMODE_DSTIN      = C.DSTIN_SK_BLENDMODE
	SK_BLENDMODE_SRCOUT     = C.SRCOUT_SK_BLENDMODE
	SK_BLENDMODE_DSTOUT     = C.DSTOUT_SK_BLENDMODE
	SK_BLENDMODE_SRCATOP    = C.SRCATOP_SK_BLENDMODE
	SK_BLENDMODE_DSTATOP    = C.DSTATOP_SK_BLENDMODE
	SK_BLENDMODE_XOR        = C.XOR_SK_BLENDMODE
	SK_BLENDMODE_PLUS       = C.PLUS_SK_BLENDMODE
	SK_BLENDMODE_MODULATE   = C.MODULATE_SK_BLENDMODE
	SK_BLENDMODE_SCREEN     = C.SCREEN_SK_BLENDMODE
	SK_BLENDMODE_OVERLAY    = C.OVERLAY_SK_BLENDMODE
	SK_BLENDMODE_DARKEN     = C.DARKEN_SK_BLENDMODE
	SK_BLENDMODE_LIGHTEN    = C.LIGHTEN_SK_BLENDMODE
	SK_BLENDMODE_COLORDODGE = C.COLORDODGE_SK_BLENDMODE
	SK_BLENDMODE_COLORBURN  = C.COLORBURN_SK_BLENDMODE
	SK_BLENDMODE_HARDLIGHT  = C.HARDLIGHT_SK_BLENDMODE
	SK_BLENDMODE_SOFTLIGHT  = C.SOFTLIGHT_SK_BLENDMODE
	SK_BLENDMODE_DIFFERENCE = C.DIFFERENCE_SK_BLENDMODE
	SK_BLENDMODE_EXCLUSION  = C.EXCLUSION_SK_BLENDMODE
	SK_BLENDMODE_MULTIPLY   = C.MULTIPLY_SK_BLENDMODE
	SK_BLENDMODE_HUE        = C.HUE_SK_BLENDMODE
	SK_BLENDMODE_SATURATION = C.SATURATION_SK_BLENDMODE
	SK_BLENDMODE_COLOR      = C.COLOR_SK_BLENDMODE
	SK_BLENDMODE_LUMINOSITY = C.LUMINOSITY_SK_BLENDMODE
)

func (v *SkPaint) SetBlendmode(mode SkBlendmode) {
	C.sk_paint_set_blendmode(v, mode)
}

func (v *SkPaint) IsDither() bool {
	return bool(C.sk_paint_is_dither(v))
}

func (v *SkPaint) SetDither(dither bool) {
	C.sk_paint_set_dither(v, cbool_t(dither))
}

func (v *SkPaint) GetShader() *SkShader {
	return C.sk_paint_get_shader(v)
}

func (v *SkPaint) GetMaskfilter() *SkMaskfilter {
	return C.sk_paint_get_maskfilter(v)
}

func (v *SkPaint) SetColorfilter(filter *SkColorfilter) {
	C.sk_paint_set_colorfilter(v, filter)
}

func (v *SkPaint) GetColorfilter() *SkColorfilter {
	return C.sk_paint_get_colorfilter(v)
}

func (v *SkPaint) SetImagefilter(filter *SkImagefilter) {
	C.sk_paint_set_imagefilter(v, filter)
}

func (v *SkPaint) GetImagefilter() *SkImagefilter {
	return C.sk_paint_get_imagefilter(v)
}

func (v *SkPaint) GetBlendmode() SkBlendmode {
	return C.sk_paint_get_blendmode(v)
}

type skFilterQuality = C.sk_filter_quality_t
type SkFilterQuality = skFilterQuality

const (
	SK_FILTER_QUALITY_NONE   = C.NONE_SK_FILTER_QUALITY
	SK_FILTER_QUALITY_LOW    = C.LOW_SK_FILTER_QUALITY
	SK_FILTER_QUALITY_MEDIUM = C.MEDIUM_SK_FILTER_QUALITY
	SK_FILTER_QUALITY_HIGH   = C.HIGH_SK_FILTER_QUALITY
)

func (v *SkPaint) SetFilterQuality(arg1 SkFilterQuality) {
	C.sk_paint_set_filter_quality(v, arg1)
}

func (v *SkPaint) GetFilterQuality() SkFilterQuality {
	return C.sk_paint_get_filter_quality(v)
}

func (v *SkPaint) GetPathEffect() *SkPathEffect {
	return C.sk_paint_get_path_effect(v)
}

func (v *SkPaint) SetPathEffect(effect *SkPathEffect) {
	C.sk_paint_set_path_effect(v, effect)
}

func (v *SkPaint) GetFillPath(src *SkPath, dst *SkPath, cullRect *Rect, resScale float32) bool {
	return bool(C.sk_paint_get_fill_path(v, src, dst, cullRect.native(), cfloat_t(resScale)))
}
