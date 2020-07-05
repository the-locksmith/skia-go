package skia

/*
#include "skia.h"
*/
import "C"
import (
	"github.com/uiez/skia-go/internal"
)

type skVertices = C.sk_vertices_t
type SkVertices = skVertices

func (v *SkVertices) Unref() {
	C.sk_vertices_unref(v)
}

func (v *SkVertices) Ref() {
	C.sk_vertices_ref(v)
}

type skVerticesVertexMode = C.sk_vertices_vertex_mode_t
type SkVerticesVertexMode = skVerticesVertexMode

const (
	SK_VERTICES_VERTEX_MODE_TRIANGLES      = C.TRIANGLES_SK_VERTICES_VERTEX_MODE
	SK_VERTICES_VERTEX_MODE_TRIANGLE_STRIP = C.TRIANGLE_STRIP_SK_VERTICES_VERTEX_MODE
	SK_VERTICES_VERTEX_MODE_TRIANGLE_FAN   = C.TRIANGLE_FAN_SK_VERTICES_VERTEX_MODE
)

func NewVerticesCopy(vmode SkVerticesVertexMode, positions []Point, texs []Point, colors []SkColor, indices []uint16) *SkVertices {
	return C.sk_vertices_make_copy(vmode, cint32_t(len(positions)), (*skPoint)(internal.SliceDataPtr(positions)), (*skPoint)(internal.SliceDataPtr(texs)), (*SkColor)(internal.SliceDataPtr(colors)), cint_t(len(indices)), (*cuint16_t)(internal.SliceDataPtr(indices)))
}
