package skia

/*
#include "skia.h"
*/
import "C"
import (
	"unsafe"

	"github.com/zhuah/skia-go/internal"
)

type skPath = C.sk_path_t
type SkPath = skPath

func NewPath() *SkPath {
	return C.sk_path_new()
}

func (v *SkPath) Destroy() {
	C.sk_path_delete(v)
}

func (v *SkPath) MoveTo(x float32, y float32) {
	C.sk_path_move_to(v, cfloat_t(x), cfloat_t(y))
}

func (v *SkPath) LineTo(x float32, y float32) {
	C.sk_path_line_to(v, cfloat_t(x), cfloat_t(y))
}

func (v *SkPath) QuadTo(x0 float32, y0 float32, x1 float32, y1 float32) {
	C.sk_path_quad_to(v, cfloat_t(x0), cfloat_t(y0), cfloat_t(x1), cfloat_t(y1))
}

func (v *SkPath) ConicTo(x0 float32, y0 float32, x1 float32, y1 float32, w float32) {
	C.sk_path_conic_to(v, cfloat_t(x0), cfloat_t(y0), cfloat_t(x1), cfloat_t(y1), cfloat_t(w))
}

func (v *SkPath) CubicTo(x0 float32, y0 float32, x1 float32, y1 float32, x2 float32, y2 float32) {
	C.sk_path_cubic_to(v, cfloat_t(x0), cfloat_t(y0), cfloat_t(x1), cfloat_t(y1), cfloat_t(x2), cfloat_t(y2))
}

type skPathDirection = C.sk_path_direction_t
type SkPathDirection = skPathDirection

const (
	SK_PATH_DIRECTION_CW  = C.CW_SK_PATH_DIRECTION
	SK_PATH_DIRECTION_CCW = C.CCW_SK_PATH_DIRECTION
)

type skPathArcSize = C.sk_path_arc_size_t
type SkPathArcSize = skPathArcSize

const (
	SK_PATH_ARC_SIZE_SMALL = C.SMALL_SK_PATH_ARC_SIZE
	SK_PATH_ARC_SIZE_LARGE = C.LARGE_SK_PATH_ARC_SIZE
)

func (v *SkPath) ArcTo(rx float32, ry float32, xAxisRotate float32, largeArc SkPathArcSize, sweep SkPathDirection, x float32, y float32) {
	C.sk_path_arc_to(v, cfloat_t(rx), cfloat_t(ry), cfloat_t(xAxisRotate), largeArc, sweep, cfloat_t(x), cfloat_t(y))
}

func (v *SkPath) RarcTo(rx float32, ry float32, xAxisRotate float32, largeArc SkPathArcSize, sweep SkPathDirection, x float32, y float32) {
	C.sk_path_rarc_to(v, cfloat_t(rx), cfloat_t(ry), cfloat_t(xAxisRotate), largeArc, sweep, cfloat_t(x), cfloat_t(y))
}

func (v *SkPath) ArcToWithOval(oval *Rect, startAngle float32, sweepAngle float32, forceMoveTo bool) {
	C.sk_path_arc_to_with_oval(v, oval.native(), cfloat_t(startAngle), cfloat_t(sweepAngle), cbool_t(forceMoveTo))
}

func (v *SkPath) ArcToWithPoints(x1 float32, y1 float32, x2 float32, y2 float32, radius float32) {
	C.sk_path_arc_to_with_points(v, cfloat_t(x1), cfloat_t(y1), cfloat_t(x2), cfloat_t(y2), cfloat_t(radius))
}

func (v *SkPath) Close() {
	C.sk_path_close(v)
}

func (v *SkPath) AddRect(rect *Rect, dir SkPathDirection) {
	C.sk_path_add_rect(v, rect.native(), dir)
}

func (v *SkPath) AddRrect(rrect *SkRrect, dir SkPathDirection) {
	C.sk_path_add_rrect(v, rrect, dir)
}

func (v *SkPath) AddRrectStart(rrect *SkRrect, dir SkPathDirection, start uint32) {
	C.sk_path_add_rrect_start(v, rrect, dir, cuint32_t(start))
}

func (v *SkPath) AddRoundedRect(rect *Rect, rx float32, ry float32, dir SkPathDirection) {
	C.sk_path_add_rounded_rect(v, rect.native(), cfloat_t(rx), cfloat_t(ry), dir)
}

func (v *SkPath) AddOval(oval *Rect, dir SkPathDirection) {
	C.sk_path_add_oval(v, oval.native(), dir)
}

func (v *SkPath) AddCircle(x float32, y float32, radius float32, dir SkPathDirection) {
	C.sk_path_add_circle(v, cfloat_t(x), cfloat_t(y), cfloat_t(radius), dir)
}

func (v *SkPath) GetBounds() Rect {
	var bounds Rect
	C.sk_path_get_bounds(v, bounds.native())
	return bounds
}

func (v *SkPath) ComputeTightBounds() Rect {
	var bounds Rect
	C.sk_path_compute_tight_bounds(v, bounds.native())
	return bounds
}

func (v *SkPath) RmoveTo(dx float32, dy float32) {
	C.sk_path_rmove_to(v, cfloat_t(dx), cfloat_t(dy))
}

func (v *SkPath) RlineTo(dx float32, yd float32) {
	C.sk_path_rline_to(v, cfloat_t(dx), cfloat_t(yd))
}

func (v *SkPath) RquadTo(dx0 float32, dy0 float32, dx1 float32, dy1 float32) {
	C.sk_path_rquad_to(v, cfloat_t(dx0), cfloat_t(dy0), cfloat_t(dx1), cfloat_t(dy1))
}

func (v *SkPath) RconicTo(dx0 float32, dy0 float32, dx1 float32, dy1 float32, w float32) {
	C.sk_path_rconic_to(v, cfloat_t(dx0), cfloat_t(dy0), cfloat_t(dx1), cfloat_t(dy1), cfloat_t(w))
}

func (v *SkPath) RcubicTo(dx0 float32, dy0 float32, dx1 float32, dy1 float32, dx2 float32, dy2 float32) {
	C.sk_path_rcubic_to(v, cfloat_t(dx0), cfloat_t(dy0), cfloat_t(dx1), cfloat_t(dy1), cfloat_t(dx2), cfloat_t(dy2))
}

func (v *SkPath) AddRectStart(crect *Rect, dir SkPathDirection, startIndex uint32) {
	C.sk_path_add_rect_start(v, crect.native(), dir, cuint32_t(startIndex))
}

func (v *SkPath) AddArc(crect *Rect, startAngle float32, sweepAngle float32) {
	C.sk_path_add_arc(v, crect.native(), cfloat_t(startAngle), cfloat_t(sweepAngle))
}

type skPathFilltype = C.sk_path_filltype_t
type SkPathFilltype = skPathFilltype

const (
	SK_PATH_FILLTYPE_WINDING         = C.WINDING_SK_PATH_FILLTYPE
	SK_PATH_FILLTYPE_EVENODD         = C.EVENODD_SK_PATH_FILLTYPE
	SK_PATH_FILLTYPE_INVERSE_WINDING = C.INVERSE_WINDING_SK_PATH_FILLTYPE
	SK_PATH_FILLTYPE_INVERSE_EVENODD = C.INVERSE_EVENODD_SK_PATH_FILLTYPE
)

func (v *SkPath) GetFilltype() SkPathFilltype {
	return C.sk_path_get_filltype(v)
}

func (v *SkPath) SetFilltype(typ SkPathFilltype) {
	C.sk_path_set_filltype(v, typ)
}

func (v *SkPath) Transform(matrix *Matrix) {
	C.sk_path_transform(v, matrix.native())
}

func (v *SkPath) TransformToDest(matrix *Matrix, destination *SkPath) {
	C.sk_path_transform_to_dest(v, matrix.native(), destination)
}

func (v *SkPath) Clone() *SkPath {
	return C.sk_path_clone(v)
}

type skPathVerb = C.sk_path_verb_t
type SkPathVerb = skPathVerb

const (
	SK_PATH_VERB_MOVE  = C.MOVE_SK_PATH_VERB
	SK_PATH_VERB_LINE  = C.LINE_SK_PATH_VERB
	SK_PATH_VERB_QUAD  = C.QUAD_SK_PATH_VERB
	SK_PATH_VERB_CONIC = C.CONIC_SK_PATH_VERB
	SK_PATH_VERB_CUBIC = C.CUBIC_SK_PATH_VERB
	SK_PATH_VERB_CLOSE = C.CLOSE_SK_PATH_VERB
	SK_PATH_VERB_DONE  = C.DONE_SK_PATH_VERB
)

type skPathAddMode = C.sk_path_add_mode_t
type SkPathAddMode = skPathAddMode

const (
	SK_PATH_ADD_MODE_APPEND = C.APPEND_SK_PATH_ADD_MODE
	SK_PATH_ADD_MODE_EXTEND = C.EXTEND_SK_PATH_ADD_MODE
)

func (v *SkPath) AddPathOffset(other *SkPath, dx float32, dy float32, addMode SkPathAddMode) {
	C.sk_path_add_path_offset(v, other, cfloat_t(dx), cfloat_t(dy), addMode)
}

func (v *SkPath) AddPathMatrix(other *SkPath, matrix *Matrix, addMode SkPathAddMode) {
	C.sk_path_add_path_matrix(v, other, matrix.native(), addMode)
}

func (v *SkPath) AddPath(other *SkPath, addMode SkPathAddMode) {
	C.sk_path_add_path(v, other, addMode)
}

func (v *SkPath) AddPathReverse(other *SkPath) {
	C.sk_path_add_path_reverse(v, other)
}

func (v *SkPath) Reset() {
	C.sk_path_reset(v)
}

func (v *SkPath) Rewind() {
	C.sk_path_rewind(v)
}

func (v *SkPath) CountPoints() int32 {
	return int32(C.sk_path_count_points(v))
}

func (v *SkPath) CountVerbs() int32 {
	return int32(C.sk_path_count_verbs(v))
}

func (v *SkPath) GetPoint(index int32) Point {
	var point Point
	C.sk_path_get_point(v, cint32_t(index), point.native())
	return point
}

func (v *SkPath) GetPoints(points []Point) int {
	return int(C.sk_path_get_points(v, (*skPoint)(internal.SliceDataPtr(points)), cint_t(len(points))))
}

func (v *SkPath) Contains(x float32, y float32) bool {
	return bool(C.sk_path_contains(v, cfloat_t(x), cfloat_t(y)))
}

func (v *SkPath) GetConvexity() SkPathConvexity {
	return C.sk_path_get_convexity(v)
}

func (v *SkPath) SetConvexity(convexity SkPathConvexity) {
	C.sk_path_set_convexity(v, convexity)
}

func (v *SkPath) GetLastPoint() (Point, bool) {
	var point Point
	if bool(C.sk_path_get_last_point(v, point.native())) {
		return point, true
	}
	return point, false
}

func (v *SkPath) AddPoly(points []Point, close bool) {
	C.sk_path_add_poly(v, (*skPoint)(internal.SliceDataPtr(points)), cint_t(len(points)), cbool_t(close))
}

func (v *SkPath) GetSegmentMasks() uint32 {
	return uint32(C.sk_path_get_segment_masks(v))
}

func (v *SkPath) IsOval(bounds *Rect) bool {
	return bool(C.sk_path_is_oval(v, bounds.native()))
}

func (v *SkPath) IsRrect(bounds *SkRrect) bool {
	return bool(C.sk_path_is_rrect(v, bounds))
}

func (v *SkPath) IsLine(line *[2]Point) bool {
	return bool(C.sk_path_is_line(v, (*skPoint)(unsafe.Pointer(line))))
}

func (v *SkPath) IsRect(rect *Rect) (dir SkPathDirection, closed bool, ok bool) {
	var cclosed cbool_t
	ok = bool(C.sk_path_is_rect(v, rect.native(), &cclosed, &dir))
	closed = bool(cclosed)
	return
}

func SkPathopOp(one *SkPath, two *SkPath, op SkPathop, result *SkPath) bool {
	return bool(C.sk_pathop_op(one, two, op, result))
}

func SkPathopSimplify(path *SkPath, result *SkPath) bool {
	return bool(C.sk_pathop_simplify(path, result))
}

func SkPathopTightBounds(path *SkPath) (Rect, bool) {
	var result Rect
	ok := bool(C.sk_pathop_tight_bounds(path, result.native()))
	return result, ok
}

type skOpbuilder = C.sk_opbuilder_t
type SkOpbuilder = skOpbuilder

func NewOpbuilder() *SkOpbuilder {
	return C.sk_opbuilder_new()
}

func (v *SkOpbuilder) Destroy() {
	C.sk_opbuilder_destroy(v)
}

func (v *SkOpbuilder) Add(path *SkPath, op SkPathop) {
	C.sk_opbuilder_add(v, path, op)
}

func (v *SkOpbuilder) Resolve(result *SkPath) bool {
	return bool(C.sk_opbuilder_resolve(v, result))
}

type skPathop = C.sk_pathop_t
type SkPathop = skPathop

func fromGrSurfaceorigin(v C.sk_pathop_t) SkPathop {
	return SkPathop(v)
}
func (v SkPathop) native() C.sk_pathop_t {
	return C.sk_pathop_t(v)
}

const (
	SK_PATHOP_DIFFERENCE         = C.DIFFERENCE_SK_PATHOP
	SK_PATHOP_INTERSECT          = C.INTERSECT_SK_PATHOP
	SK_PATHOP_UNION              = C.UNION_SK_PATHOP
	SK_PATHOP_XOR                = C.XOR_SK_PATHOP
	SK_PATHOP_REVERSE_DIFFERENCE = C.REVERSE_DIFFERENCE_SK_PATHOP
)

type skPathConvexity = C.sk_path_convexity_t
type SkPathConvexity = skPathConvexity

const (
	SK_PATH_CONVEXITY_UNKNOWN = C.UNKNOWN_SK_PATH_CONVEXITY
	SK_PATH_CONVEXITY_CONVEX  = C.CONVEX_SK_PATH_CONVEXITY
	SK_PATH_CONVEXITY_CONCAVE = C.CONCAVE_SK_PATH_CONVEXITY
)
