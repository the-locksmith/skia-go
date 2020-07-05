package skia

/*
#include "skia.h"
*/
import "C"
import (
	"unsafe"

	"github.com/uiez/skia-go/internal"
)

type skMatrix = C.sk_matrix_t
type Matrix struct {
	ScaleX float32
	SkewX  float32
	TransX float32
	SkewY  float32
	ScaleY float32
	TransY float32
	Persp0 float32
	Persp1 float32
	Persp2 float32
}

func fromNativeMatrix(v *skMatrix) *Matrix {
	return (*Matrix)(unsafe.Pointer(v))
}

func MatrixConcat(first, second *Matrix) *Matrix {
	var res Matrix
	C.sk_matrix_concat(res.native(), first.native(), second.native())
	return &res
}

func (v *Matrix) native() *skMatrix {
	return (*skMatrix)(unsafe.Pointer(v))
}

func (v *Matrix) CanInvert() bool {
	return bool(C.sk_matrix_try_invert(v.native(), nil))
}

func (v *Matrix) TryInvert() (*Matrix, bool) {
	var m Matrix
	if bool(C.sk_matrix_try_invert(v.native(), m.native())) {
		return &m, true
	}

	return nil, false
}

func (v *Matrix) PreConcat(matrix *Matrix) {
	C.sk_matrix_pre_concat(v.native(), matrix.native())
}

func (v *Matrix) PostConcat(matrix *Matrix) {
	C.sk_matrix_post_concat(v.native(), matrix.native())
}

func (v *Matrix) MapRect(dest *Rect, source *Rect) {
	C.sk_matrix_map_rect(v.native(), dest.native(), source.native())
}

func (v *Matrix) MapPoints(points []Point) {
	ptr := (*skPoint)(internal.SliceDataPtr(points))
	C.sk_matrix_map_points(v.native(), ptr, ptr, cint32_t(len(points)))
}

func (v *Matrix) MapVectors(vectors []Point) {
	ptr := (*skPoint)(internal.SliceDataPtr(vectors))
	C.sk_matrix_map_vectors(v.native(), ptr, ptr, cint32_t(len(vectors)))
}

func (v *Matrix) MapXy(x float32, y float32) Point {
	var p Point
	C.sk_matrix_map_xy(v.native(), cfloat_t(x), cfloat_t(y), p.native())
	return p
}

func (v *Matrix) MapVector(x float32, y float32) Point {
	var p Point
	C.sk_matrix_map_vector(v.native(), cfloat_t(x), cfloat_t(y), p.native())
	return p
}

func (v *Matrix) MapRadius(radius float32) float32 {
	return float32(C.sk_matrix_map_radius(v.native(), cfloat_t(radius)))
}
