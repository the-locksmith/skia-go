package skia

/*
#include "skia.h"
*/
import "C"
import (
	"unsafe"

	"github.com/zhuah/skia-go/internal"
)

type skFont = C.sk_font_t
type SkFont = skFont

func NewFont() *SkFont {
	return C.sk_font_new()
}

func NewFontWithValues(typeface *SkTypeface, size float32, scaleX float32, skewX float32) *SkFont {
	return C.sk_font_new_with_values(typeface, cfloat_t(size), cfloat_t(scaleX), cfloat_t(skewX))
}

func (v *SkFont) Destroy() {
	C.sk_font_delete(v)
}

func (v *SkFont) IsForceAutoHinting() bool {
	return bool(C.sk_font_is_force_auto_hinting(v))
}

func (v *SkFont) SetForceAutoHinting(value bool) {
	C.sk_font_set_force_auto_hinting(v, cbool_t(value))
}

func (v *SkFont) IsEmbeddedBitmaps() bool {
	return bool(C.sk_font_is_embedded_bitmaps(v))
}

func (v *SkFont) SetEmbeddedBitmaps(value bool) {
	C.sk_font_set_embedded_bitmaps(v, cbool_t(value))
}

func (v *SkFont) IsSubpixel() bool {
	return bool(C.sk_font_is_subpixel(v))
}

func (v *SkFont) SetSubpixel(value bool) {
	C.sk_font_set_subpixel(v, cbool_t(value))
}

func (v *SkFont) IsLinearMetrics() bool {
	return bool(C.sk_font_is_linear_metrics(v))
}

func (v *SkFont) SetLinearMetrics(value bool) {
	C.sk_font_set_linear_metrics(v, cbool_t(value))
}

func (v *SkFont) IsEmbolden() bool {
	return bool(C.sk_font_is_embolden(v))
}

func (v *SkFont) SetEmbolden(value bool) {
	C.sk_font_set_embolden(v, cbool_t(value))
}

func (v *SkFont) IsBaselineSnap() bool {
	return bool(C.sk_font_is_baseline_snap(v))
}

func (v *SkFont) SetBaselineSnap(value bool) {
	C.sk_font_set_baseline_snap(v, cbool_t(value))
}

type skFontHinting = C.sk_font_hinting_t
type SkFontHinting = skFontHinting

const (
	SK_FONT_HINTING_NONE   = C.NONE_SK_FONT_HINTING
	SK_FONT_HINTING_SLIGHT = C.SLIGHT_SK_FONT_HINTING
	SK_FONT_HINTING_NORMAL = C.NORMAL_SK_FONT_HINTING
	SK_FONT_HINTING_FULL   = C.FULL_SK_FONT_HINTING
)

type skFontEdging = C.sk_font_edging_t
type SkFontEdging = skFontEdging

const (
	SK_FONT_EDGING_ALIAS              = C.ALIAS_SK_FONT_EDGING
	SK_FONT_EDGING_ANTIALIAS          = C.ANTIALIAS_SK_FONT_EDGING
	SK_FONT_EDGING_SUBPIXEL_ANTIALIAS = C.SUBPIXEL_ANTIALIAS_SK_FONT_EDGING
)

func (v *SkFont) GetEdging() SkFontEdging {
	return C.sk_font_get_edging(v)
}

func (v *SkFont) SetEdging(value SkFontEdging) {
	C.sk_font_set_edging(v, value)
}

func (v *SkFont) GetHinting() SkFontHinting {
	return C.sk_font_get_hinting(v)
}

func (v *SkFont) SetHinting(value SkFontHinting) {
	C.sk_font_set_hinting(v, value)
}

func (v *SkFont) GetTypeface() *SkTypeface {
	return C.sk_font_get_typeface(v)
}

func (v *SkFont) SetTypeface(value *SkTypeface) {
	C.sk_font_set_typeface(v, value)
}

func (v *SkFont) GetSize() float32 {
	return float32(C.sk_font_get_size(v))
}

func (v *SkFont) SetSize(value float32) {
	C.sk_font_set_size(v, cfloat_t(value))
}

func (v *SkFont) GetScaleX() float32 {
	return float32(C.sk_font_get_scale_x(v))
}

func (v *SkFont) SetScaleX(value float32) {
	C.sk_font_set_scale_x(v, cfloat_t(value))
}

func (v *SkFont) GetSkewX() float32 {
	return float32(C.sk_font_get_skew_x(v))
}

func (v *SkFont) SetSkewX(value float32) {
	C.sk_font_set_skew_x(v, cfloat_t(value))
}

type skTextEncoding = C.sk_text_encoding_t
type SkTextEncoding = skTextEncoding

const (
	SK_TEXT_ENCODING_UTF8     = C.UTF8_SK_TEXT_ENCODING
	SK_TEXT_ENCODING_UTF16    = C.UTF16_SK_TEXT_ENCODING
	SK_TEXT_ENCODING_UTF32    = C.UTF32_SK_TEXT_ENCODING
	SK_TEXT_ENCODING_GLYPH_ID = C.GLYPH_ID_SK_TEXT_ENCODING
)

func (v *SkFont) TextToGlyphs(text unsafe.Pointer, byteLength uint, encoding SkTextEncoding, glyphs []uint16) int {
	return int(C.sk_font_text_to_glyphs(v, text, csize_t(byteLength), encoding, (*cuint16_t)(internal.SliceDataPtr(glyphs)), cint_t(len(glyphs))))
}

func (v *SkFont) UnicharToGlyph(c int32) uint16 {
	return uint16(C.sk_font_unichar_to_glyph(v, cint32_t(c)))
}

func (v *SkFont) UnicharsToGlyphs(uni []int32) []uint16 {
	glyphs := make([]uint16, len(uni))
	C.sk_font_unichars_to_glyphs(v, (*cint32_t)(internal.SliceDataPtr(uni)), cint_t(len(uni)), (*cushort_t)(internal.SliceDataPtr(glyphs)))
	return glyphs
}

func (v *SkFont) MeasureText(text unsafe.Pointer, byteLength uint, encoding SkTextEncoding, bounds *Rect, paint *SkPaint) float32 {
	return float32(C.sk_font_measure_text(v, text, csize_t(byteLength), encoding, bounds.native(), paint))
}

//
//func (v *SkFont) BreakText(text unsafe.Pointer, byteLength uint, encoding SkTextEncoding, maxWidth float32, measuredWidth *float32, paint *SkPaint) uint {
//	cText := Text
//	__ret := C.sk_font_break_text(v, cText, cByteLength, cEncoding, cMaxWidth, cMeasuredWidth, cPaint)
//}

func (v *SkFont) GetWidthsBounds(glyphs []uint16, paint *SkPaint, needsWidth, needsBounds bool) (widths []float32, bounds []Rect) {
	if needsWidth {
		widths = make([]float32, len(glyphs))
	}
	if needsBounds {
		bounds = make([]Rect, len(glyphs))
	}
	C.sk_font_get_widths_bounds(v, (*cuint16_t)(internal.SliceDataPtr(glyphs)), cint_t(len(glyphs)), (*cfloat_t)(internal.SliceDataPtr(widths)), (*skRect)(internal.SliceDataPtr(bounds)), paint)
	return
}

func (v *SkFont) GetPos(glyphs []uint16, origin Point) []Point {
	poses := make([]Point, len(glyphs))
	C.sk_font_get_pos(v, (*cuint16_t)(internal.SliceDataPtr(glyphs)), cint_t(len(glyphs)), (*skPoint)(internal.SliceDataPtr(poses)), origin.native())
	return poses
}

func (v *SkFont) GetXpos(glyphs []uint16, origin float32) []float32 {
	poses := make([]float32, len(glyphs))
	C.sk_font_get_xpos(v, (*cuint16_t)(internal.SliceDataPtr(glyphs)), cint_t(len(glyphs)), (*cfloat_t)(internal.SliceDataPtr(poses)), cfloat_t(origin))
	return poses
}

func (v *SkFont) GetPath(glyph uint16, path *SkPath) bool {
	return bool(C.sk_font_get_path(v, cuint16_t(glyph), path))
}

func (v *SkFont) GetPaths(glyphs []uint16, glyphPathProc SkGlyphPathProc, context interface{}) {
	cproc, ccontext, free := skiaGlyphPathProcConvert(glyphPathProc, context)
	C.sk_font_get_paths(v, (*cuint16_t)(internal.SliceDataPtr(glyphs)), cint_t(len(glyphs)), cproc, ccontext)
	free()
}

type skFontmetrics = C.sk_fontmetrics_t
type Fontmetrics struct {
	FFlags              uint32
	FTop                float32
	FAscent             float32
	FDescent            float32
	FBottom             float32
	FLeading            float32
	FAvgCharWidth       float32
	FMaxCharWidth       float32
	FXMin               float32
	FXMax               float32
	FXHeight            float32
	FCapHeight          float32
	FUnderlineThickness float32
	FUnderlinePosition  float32
	FStrikeoutThickness float32
	FStrikeoutPosition  float32
}

func fromNativeFontmetrics(v *skFontmetrics) *Fontmetrics {
	return (*Fontmetrics)(unsafe.Pointer(v))
}
func (v *Fontmetrics) native() *skFontmetrics {
	return (*skFontmetrics)(unsafe.Pointer(v))
}

func (v *SkFont) GetMetrics() (metrics Fontmetrics, lineSpacing float32) {
	lineSpacing = float32(C.sk_font_get_metrics(v, metrics.native()))
	return
}
