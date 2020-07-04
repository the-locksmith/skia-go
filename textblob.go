package skia

/*
#include "skia.h"
*/
import "C"
import (
	"reflect"
	"unsafe"

	"github.com/zhuah/skia-go/internal"
)

type skTextblob = C.sk_textblob_t
type SkTextblob = skTextblob

type skTextblobBuilder = C.sk_textblob_builder_t
type SkTextblobBuilder = skTextblobBuilder

func (v *SkTextblob) Ref() {
	C.sk_textblob_ref(v)
}

func (v *SkTextblob) Unref() {
	C.sk_textblob_unref(v)
}

func (v *SkTextblob) GetUniqueId() uint32 {
	return uint32(C.sk_textblob_get_unique_id(v))
}

func (v *SkTextblob) GetBounds() Rect {
	var bounds Rect
	C.sk_textblob_get_bounds(v, bounds.native())
	return bounds
}

func (v *SkTextblob) GetIntercepts(bounds [2]float32, paint *SkPaint) []float32 {
	n := C.sk_textblob_get_intercepts(v, (*cfloat_t)(unsafe.Pointer(&bounds)), nil, paint)
	if n <= 0 {
		return nil
	}
	intervals := make([]float32, n)
	C.sk_textblob_get_intercepts(v, (*cfloat_t)(unsafe.Pointer(&bounds)), (*cfloat_t)(internal.SliceDataPtr(intervals)), paint)
	return intervals
}

func NewTextblobBuilder() *SkTextblobBuilder {
	return C.sk_textblob_builder_new()
}

func (v *SkTextblobBuilder) Destroy() {
	C.sk_textblob_builder_delete(v)
}

func (v *SkTextblobBuilder) Make() *SkTextblob {
	return C.sk_textblob_builder_make(v)
}

type skTextblobBuilderRunbuffer = C.sk_textblob_builder_runbuffer_t
type TextblobBuilderRunPosbuffer struct {
	Glyphs []SkGlyphId
	Pos    []Point
}

func (v *SkTextblobBuilder) AllocRunPos(font *SkFont, count int32, bounds *Rect) TextblobBuilderRunPosbuffer {
	var buf TextblobBuilderRunPosbuffer

	var cRunbuffer skTextblobBuilderRunbuffer
	C.sk_textblob_builder_alloc_run_pos(v, font, cint_t(count), bounds.native(), &cRunbuffer)

	glyphPtr := (*reflect.SliceHeader)(unsafe.Pointer(&buf.Glyphs))
	glyphPtr.Data = uintptr(cRunbuffer.glyphs)
	glyphPtr.Len = int(count)
	glyphPtr.Cap = int(count)

	posPtr := (*reflect.SliceHeader)(unsafe.Pointer(&buf.Pos))
	posPtr.Data = uintptr(cRunbuffer.pos)
	posPtr.Len = int(count)
	posPtr.Cap = int(count)
	return buf
}
