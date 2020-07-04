package skia

/*
#include "skia.h"
*/
import "C"

type skFontstyle = C.sk_fontstyle_t
type SkFontstyle = skFontstyle

type skFontStyleSlant = C.sk_font_style_slant_t
type SkFontStyleSlant = skFontStyleSlant

const (
	SK_FONT_STYLE_SLANT_UPRIGHT = C.UPRIGHT_SK_FONT_STYLE_SLANT
	SK_FONT_STYLE_SLANT_ITALIC  = C.ITALIC_SK_FONT_STYLE_SLANT
	SK_FONT_STYLE_SLANT_OBLIQUE = C.OBLIQUE_SK_FONT_STYLE_SLANT
)

func NewFontstyle(weight int32, width int32, slant SkFontStyleSlant) *SkFontstyle {
	return C.sk_fontstyle_new(cint32_t(weight), cint32_t(width), slant)
}

func (v *SkFontstyle) Destroy() {
	C.sk_fontstyle_delete(v)
}

func (v *SkFontstyle) GetWeight() int32 {
	return int32(C.sk_fontstyle_get_weight(v))
}

func (v *SkFontstyle) GetWidth() int32 {
	return int32(C.sk_fontstyle_get_width(v))
}

func (v *SkFontstyle) GetSlant() SkFontStyleSlant {
	return C.sk_fontstyle_get_slant(v)
}

func CreateEmptyFontstyleset() *SkFontstyleset {
	return C.sk_fontstyleset_create_empty()
}

func (v *SkFontstyleset) Unref() {
	C.sk_fontstyleset_unref(v)
}

func (v *SkFontstyleset) GetCount() int {
	return int(C.sk_fontstyleset_get_count(v))
}

func (v *SkFontstyleset) GetStyle(index int, style *SkFontstyle) {
	C.sk_fontstyleset_get_style(v, cint32_t(index), style, nil)
}

func (v *SkFontstyleset) CreateTypeface(index int) *SkTypeface {
	return C.sk_fontstyleset_create_typeface(v, cint32_t(index))
}

func (v *SkFontstyleset) MatchStyle(style *SkFontstyle) *SkTypeface {
	return C.sk_fontstyleset_match_style(v, style)
}
