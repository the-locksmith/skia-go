package skia

/*
#include "skia.h"
*/
import "C"
import (
	"github.com/uiez/skia-go/internal"
)

type skFontmgr = C.sk_fontmgr_t
type SkFontmgr = skFontmgr

type skFontstyleset = C.sk_fontstyleset_t
type SkFontstyleset = skFontstyleset

func NewDefaultFontmgr() *SkFontmgr {
	return C.sk_fontmgr_create_default()
}

func RefDefaultFontmgr() *SkFontmgr {
	return C.sk_fontmgr_ref_default()
}

func (v *SkFontmgr) Unref() {
	C.sk_fontmgr_unref(v)
}

func (v *SkFontmgr) CountFamilies() int {
	return int(C.sk_fontmgr_count_families(v))
}

func (v *SkFontmgr) GetFamilyName(index int) string {
	str := NewStringEmpty()
	C.sk_fontmgr_get_family_name(v, cint32_t(index), str)
	name := string(str.GetData())
	str.Destroy()
	return name
}

func (v *SkFontmgr) CreateStyleset(index int) *SkFontstyleset {
	return C.sk_fontmgr_create_styleset(v, cint32_t(index))
}

func (v *SkFontmgr) MatchFamily(familyName string) *SkFontstyleset {
	cstr, free := cstring(familyName)
	sset := C.sk_fontmgr_match_family(v, cstr)
	free()

	return sset
}

func (v *SkFontmgr) MatchFaceStyle(face *SkTypeface, style *SkFontstyle) *SkTypeface {
	return C.sk_fontmgr_match_face_style(v, face, style)
}

func (v *SkFontmgr) CreateFromData(data *SkData, index int32) *SkTypeface {
	return C.sk_fontmgr_create_from_data(v, data, cint32_t(index))
}

func (v *SkFontmgr) CreateFromFile(path string, index int32) *SkTypeface {
	cpath, free := cstring(path)
	t := C.sk_fontmgr_create_from_file(v, cpath, cint32_t(index))
	free()

	return t
}

func (v *SkFontmgr) MatchFamilyStyle(familyName string, style *SkFontstyle) *SkTypeface {
	cname, free := cstring(familyName)
	t := C.sk_fontmgr_match_family_style(v, cname, style)
	free()
	return t
}

func (v *SkFontmgr) MatchFamilyStyleCharacter(familyName string, style *SkFontstyle, bcp47 []string, character int32) *SkTypeface {
	cname, free := cstring(familyName)

	safeStringFrees := make([]func(), len(bcp47))
	cBcp47Arr := make([]*cchar_t, len(bcp47))
	for i := range bcp47 {
		s, f := cstring(bcp47[i])
		safeStringFrees[i] = f
		cBcp47Arr[i] = s
	}
	t := C.sk_fontmgr_match_family_style_character(v, cname, style, (**cchar_t)(internal.SliceDataPtr(cBcp47Arr)), cint_t(len(bcp47)), cint32_t(character))
	free()
	for i := range safeStringFrees {
		safeStringFrees[i]()
	}

	return t
}
