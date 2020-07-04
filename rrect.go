package skia

/*
#include "skia.h"
*/
import "C"
import (
	"unsafe"
)

type skRrect = C.sk_rrect_t
type SkRrect = skRrect

func NewRrect() *SkRrect {
	return C.sk_rrect_new()
}

func NewRrectCopy(rrect *SkRrect) *SkRrect {
	return C.sk_rrect_new_copy(rrect)
}

func (v *SkRrect) Destroy() {
	C.sk_rrect_delete(v)
}

type skRrectType = C.sk_rrect_type_t
type SkRrectType = skRrectType

const (
	SK_RRECT_TYPE_EMPTY      = C.EMPTY_SK_RRECT_TYPE
	SK_RRECT_TYPE_RECT       = C.RECT_SK_RRECT_TYPE
	SK_RRECT_TYPE_OVAL       = C.OVAL_SK_RRECT_TYPE
	SK_RRECT_TYPE_SIMPLE     = C.SIMPLE_SK_RRECT_TYPE
	SK_RRECT_TYPE_NINE_PATCH = C.NINE_PATCH_SK_RRECT_TYPE
	SK_RRECT_TYPE_COMPLEX    = C.COMPLEX_SK_RRECT_TYPE
)

func (v *SkRrect) GetType() SkRrectType {
	return C.sk_rrect_get_type(v)
}

func (v *SkRrect) GetRect() Rect {
	var rect Rect
	C.sk_rrect_get_rect(v, rect.native())
	return rect
}

type skRrectCorner = C.sk_rrect_corner_t
type SkRrectCorner = skRrectCorner

const (
	SK_RRECT_CORNER_UPPER_LEFT  = C.UPPER_LEFT_SK_RRECT_CORNER
	SK_RRECT_CORNER_UPPER_RIGHT = C.UPPER_RIGHT_SK_RRECT_CORNER
	SK_RRECT_CORNER_LOWER_RIGHT = C.LOWER_RIGHT_SK_RRECT_CORNER
	SK_RRECT_CORNER_LOWER_LEFT  = C.LOWER_LEFT_SK_RRECT_CORNER
)

func (v *SkRrect) GetRadii(corner SkRrectCorner) Vector {
	var radii Vector
	C.sk_rrect_get_radii(v, corner, radii.native())
	return radii
}

func (v *SkRrect) GetWidth() float32 {
	return float32(C.sk_rrect_get_width(v))
}

func (v *SkRrect) GetHeight() float32 {
	return float32(C.sk_rrect_get_height(v))
}

func (v *SkRrect) SetEmpty() {
	C.sk_rrect_set_empty(v)
}

func (v *SkRrect) SetRect(rect *Rect) {
	C.sk_rrect_set_rect(v, rect.native())
}

func (v *SkRrect) SetOval(rect *Rect) {
	C.sk_rrect_set_oval(v, rect.native())
}

func (v *SkRrect) SetRectXy(rect *Rect, xRad float32, yRad float32) {
	C.sk_rrect_set_rect_xy(v, rect.native(), cfloat_t(xRad), cfloat_t(yRad))
}

func (v *SkRrect) SetNinePatch(rect *Rect, leftRad float32, topRad float32, rightRad float32, bottomRad float32) {
	C.sk_rrect_set_nine_patch(v, rect.native(), cfloat_t(leftRad), cfloat_t(topRad), cfloat_t(rightRad), cfloat_t(bottomRad))
}

func (v *SkRrect) SetRectRadii(rect *Rect, radii *[4]Vector) {
	C.sk_rrect_set_rect_radii(v, rect.native(), (*skVector)(unsafe.Pointer(radii)))
}

func (v *SkRrect) Inset(dx float32, dy float32) {
	C.sk_rrect_inset(v, cfloat_t(dx), cfloat_t(dy))
}

func (v *SkRrect) Outset(dx float32, dy float32) {
	C.sk_rrect_outset(v, cfloat_t(dx), cfloat_t(dy))
}

func (v *SkRrect) Offset(dx float32, dy float32) {
	C.sk_rrect_offset(v, cfloat_t(dx), cfloat_t(dy))
}

func (v *SkRrect) Contains(rect *Rect) bool {
	return bool(C.sk_rrect_contains(v, rect.native()))
}

func (v *SkRrect) IsValid() bool {
	return bool(C.sk_rrect_is_valid(v))
}

func (v *SkRrect) Transform(matrix *Matrix, dest *SkRrect) bool {
	return bool(C.sk_rrect_transform(v, matrix.native(), dest))
}
