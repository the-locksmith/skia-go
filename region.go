package skia

/*
#include "skia.h"
*/
import "C"
import (
	"github.com/uiez/skia-go/internal"
)

type skRegion = C.sk_region_t
type SkRegion = skRegion

type skRegionIterator = C.sk_region_iterator_t
type SkRegionIterator = skRegionIterator

type skRegionCliperator = C.sk_region_cliperator_t
type SkRegionCliperator = skRegionCliperator

type skRegionSpanerator = C.sk_region_spanerator_t
type SkRegionSpanerator = skRegionSpanerator

func NewRegion() *SkRegion {
	return C.sk_region_new()
}

func (v *SkRegion) Destroy() {
	C.sk_region_delete(v)
}

func (v *SkRegion) IsEmpty() bool {
	return bool(C.sk_region_is_empty(v))
}

func (v *SkRegion) IsRect() bool {
	return bool(C.sk_region_is_rect(v))
}

func (v *SkRegion) IsComplex() bool {
	return bool(C.sk_region_is_complex(v))
}

func (v *SkRegion) GetBounds(rect *Irect) {
	C.sk_region_get_bounds(v, rect.native())
}

func (v *SkRegion) GetBoundaryPath(path *SkPath) bool {
	return bool(C.sk_region_get_boundary_path(v, path))
}

func (v *SkRegion) SetEmpty() bool {
	return bool(C.sk_region_set_empty(v))
}

func (v *SkRegion) SetRect(rect *Irect) bool {
	return bool(C.sk_region_set_rect(v, rect.native()))
}

func (v *SkRegion) SetRects(rects []Irect) bool {
	return bool(C.sk_region_set_rects(v, (*skIrect)(internal.SliceDataPtr(rects)), cint_t(len(rects))))
}

func (v *SkRegion) SetRegion(region *SkRegion) bool {
	return bool(C.sk_region_set_region(v, region))
}

func (v *SkRegion) SetPath(path *SkPath, clip *SkRegion) bool {
	return bool(C.sk_region_set_path(v, path, clip))
}

func (v *SkRegion) IntersectsRect(rect *Irect) bool {
	return bool(C.sk_region_intersects_rect(v, rect.native()))
}

func (v *SkRegion) Intersects(src *SkRegion) bool {
	return bool(C.sk_region_intersects(v, src))
}

func (v *SkRegion) ContainsPoint(x int32, y int32) bool {
	return bool(C.sk_region_contains_point(v, cint32_t(x), cint32_t(y)))
}

func (v *SkRegion) ContainsRect(rect *Irect) bool {
	return bool(C.sk_region_contains_rect(v, rect.native()))
}

func (v *SkRegion) Contains(region *SkRegion) bool {
	return bool(C.sk_region_contains(v, region))
}

func (v *SkRegion) QuickContains(rect *Irect) bool {
	return bool(C.sk_region_quick_contains(v, rect.native()))
}

func (v *SkRegion) QuickRejectRect(rect *Irect) bool {
	return bool(C.sk_region_quick_reject_rect(v, rect.native()))
}

func (v *SkRegion) QuickReject(region *SkRegion) bool {
	return bool(C.sk_region_quick_reject(v, region))
}

func (v *SkRegion) Translate(x int32, y int32) {
	C.sk_region_translate(v, cint32_t(x), cint32_t(y))
}

type skRegionOp = C.sk_region_op_t
type SkRegionOp = skRegionOp

const (
	SK_REGION_OP_DIFFERENCE         = C.DIFFERENCE_SK_REGION_OP
	SK_REGION_OP_INTERSECT          = C.INTERSECT_SK_REGION_OP
	SK_REGION_OP_UNION              = C.UNION_SK_REGION_OP
	SK_REGION_OP_XOR                = C.XOR_SK_REGION_OP
	SK_REGION_OP_REVERSE_DIFFERENCE = C.REVERSE_DIFFERENCE_SK_REGION_OP
	SK_REGION_OP_REPLACE            = C.REPLACE_SK_REGION_OP
)

func (v *SkRegion) OpCheckRect(rect *Irect, op SkRegionOp) bool {
	return bool(C.sk_region_op_rect(v, rect.native(), op))
}

func (v *SkRegion) OpCheck(region *SkRegion, op SkRegionOp) bool {
	return bool(C.sk_region_op(v, region, op))
}

func NewRegionIterator(v *SkRegion) *SkRegionIterator {
	return C.sk_region_iterator_new(v)
}

func (v *SkRegionIterator) Destroy() {
	C.sk_region_iterator_delete(v)
}

func (v *SkRegionIterator) Rewind() bool {
	return bool(C.sk_region_iterator_rewind(v))
}

func (v *SkRegionIterator) Done() bool {
	return bool(C.sk_region_iterator_done(v))
}

func (v *SkRegionIterator) Next() {
	C.sk_region_iterator_next(v)
}

func (v *SkRegionIterator) Rect(rect *Irect) {
	C.sk_region_iterator_rect(v, rect.native())
}

func NewRegionCliperator(region *SkRegion, clip *Irect) *SkRegionCliperator {
	return C.sk_region_cliperator_new(region, clip.native())
}

func (v *SkRegionCliperator) Destroy() {
	C.sk_region_cliperator_delete(v)
}

func (v *SkRegionCliperator) Done() bool {
	return bool(C.sk_region_cliperator_done(v))
}

func (v *SkRegionCliperator) Next() {
	C.sk_region_cliperator_next(v)
}

func (v *SkRegionCliperator) Rect(rect *Irect) {
	C.sk_region_cliperator_rect(v, rect.native())
}

func NewSkRegionSpanerator(region *SkRegion, y int32, left int32, right int32) *SkRegionSpanerator {
	return C.sk_region_spanerator_new(region, cint32_t(y), cint32_t(left), cint32_t(right))
}

func (v *SkRegionSpanerator) Destroy() {
	C.sk_region_spanerator_delete(v)
}

func (v *SkRegionSpanerator) Next() (left, right int32, ok bool) {
	var cleft, cright cint32_t
	ok = bool(C.sk_region_spanerator_next(v, &cleft, &cright))
	if ok {
		left = int32(cleft)
		right = int32(cright)
	}
	return
}
