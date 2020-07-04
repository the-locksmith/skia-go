package skia

/*
#include "skia.h"
*/
import "C"
import (
	"unsafe"

	"github.com/zhuah/skia-go/internal"
)

type skFontTableTag = C.sk_font_table_tag_t
type SkFontTableTag = skFontTableTag

type skFontId = C.sk_font_id_t
type SkFontId = skFontId

type skGlyphId = C.sk_glyph_id_t
type SkGlyphId = skGlyphId

type skTypeface = C.sk_typeface_t
type SkTypeface = skTypeface

func (v *SkTypeface) Unref() {
	C.sk_typeface_unref(v)
}

func (v *SkTypeface) GetUniqueId() SkFontId {
	return C.sk_typeface_get_unique_id(v)
}

func (v *SkTypeface) GetFontstyle() *SkFontstyle {
	return C.sk_typeface_get_fontstyle(v)
}

func (v *SkTypeface) GetFontWeight() int32 {
	return int32(C.sk_typeface_get_font_weight(v))
}

func (v *SkTypeface) GetFontWidth() int32 {
	return int32(C.sk_typeface_get_font_width(v))
}

func (v *SkTypeface) GetFontSlant() SkFontStyleSlant {
	return C.sk_typeface_get_font_slant(v)
}

func (v *SkTypeface) IsFixedPitch() bool {
	return bool(C.sk_typeface_is_fixed_pitch(v))
}

func NewTypefaceDefault() *SkTypeface {
	return C.sk_typeface_create_default()
}

func RefTypefaceDefault() *SkTypeface {
	return C.sk_typeface_ref_default()
}

func NewTypefaceFromName(familyName string, style *SkFontstyle) *SkTypeface {
	cstr, free := cstring(familyName)
	t := C.sk_typeface_create_from_name(cstr, style)
	free()
	return t
}

func NewTypefaceFromFile(path string, index int32) *SkTypeface {
	cstr, free := cstring(path)
	t := C.sk_typeface_create_from_file(cstr, cint32_t(index))
	free()
	return t
}

func NewTypefaceFromData(data *SkData, index int32) *SkTypeface {
	return C.sk_typeface_create_from_data(data, cint32_t(index))
}

func (v *SkTypeface) UnicharsToGlyphs(unichars []int32) []uint16 {
	glyphs := make([]uint16, len(unichars))
	C.sk_typeface_unichars_to_glyphs(v, (*cint32_t)(internal.SliceDataPtr(unichars)), cint_t(len(unichars)), (*cushort_t)(internal.SliceDataPtr(glyphs)))
	return glyphs
}

func (v *SkTypeface) UnicharToGlyph(unichar int32) uint16 {
	return uint16(C.sk_typeface_unichar_to_glyph(v, cint32_t(unichar)))
}

func (v *SkTypeface) CountGlyphs() int32 {
	return int32(C.sk_typeface_count_glyphs(v))
}

func (v *SkTypeface) CountTables() int32 {
	return int32(C.sk_typeface_count_tables(v))
}

func (v *SkTypeface) GetTableTags() []SkFontTableTag {
	n := C.sk_typeface_get_table_tags(v, nil)
	if n <= 0 {
		return nil
	}
	tags := make([]SkFontTableTag, n)
	C.sk_typeface_get_table_tags(v, (*SkFontTableTag)(internal.SliceDataPtr(tags)))
	return tags
}

func (v *SkTypeface) GetTableSize(tag SkFontTableTag) uint {
	return uint(C.sk_typeface_get_table_size(v, tag))
}

func (v *SkTypeface) GetTableData(tag SkFontTableTag, offset uint, length uint) []byte {
	data := make([]byte, length)
	n := uint(C.sk_typeface_get_table_data(v, tag, csize_t(offset), csize_t(length), internal.SliceDataPtr(data)))
	if n <= 0 {
		return nil
	}
	return data[:n]
}

func (v *SkTypeface) CopyTableData(tag SkFontTableTag) *SkData {
	return C.sk_typeface_copy_table_data(v, tag)
}

func (v *SkTypeface) GetUnitsPerEm() int32 {
	return int32(C.sk_typeface_get_units_per_em(v))
}

func (v *SkTypeface) GetKerningPairAdjustments(glyphs []uint16) ([]int32, bool) {
	adjustments := make([]int32, len(glyphs))
	ok := bool(C.sk_typeface_get_kerning_pair_adjustments(v, (*cuint16_t)(internal.SliceDataPtr(glyphs)), cint_t(len(glyphs)), (*cint32_t)(internal.SliceDataPtr(adjustments))))
	return adjustments, ok
}

func (v *SkTypeface) GetFamilyName() string {
	skStr := (*SkString)(C.sk_typeface_get_family_name(v))
	name := string(skStr.GetData())
	skStr.Destroy()
	return name
}

type skFontVariationPosition = C.sk_font_variation_position_t
type FontVariationPosition struct {
	Tag            uint32
	Value          float32
	refaef342f2    *skFontVariationPosition
	allocsaef342f2 interface{}
}

func fromNativeFontVariationPosition(v *skFontVariationPosition) *FontVariationPosition {
	return (*FontVariationPosition)(unsafe.Pointer(v))
}
func (v *FontVariationPosition) native() *skFontVariationPosition {
	return (*skFontVariationPosition)(unsafe.Pointer(v))
}

func (v *SkTypeface) GetVariationDesignPosition() []FontVariationPosition {
	n := int32(C.sk_typeface_get_variation_design_position(v, nil, 0))
	if n <= 0 {
		return nil
	}
	poses := make([]FontVariationPosition, n)
	C.sk_typeface_get_variation_design_position(v, (*skFontVariationPosition)(internal.SliceDataPtr(poses)), cint_t(n))
	return poses
}
