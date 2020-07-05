package skia

/*
#include "skia.h"
*/
import "C"
import (
	"reflect"
	"unsafe"

	"github.com/uiez/skia-go/internal"
)

type skColortable = C.sk_colortable_t
type SkColortable = skColortable

func (v *SkColortable) Unref() {
	C.sk_colortable_unref(v)
}

func NewColortable(colors []SkPmcolor) *SkColortable {
	return C.sk_colortable_new((*SkPmcolor)(internal.SliceDataPtr(colors)), cint_t(len(colors)))
}

func (v *SkColortable) Count() int32 {
	return int32(C.sk_colortable_count(v))
}

func (v *SkColortable) ReadColors() []SkPmcolor {
	count := v.Count()
	var colors *SkPmcolor
	C.sk_colortable_read_colors(v, &colors)

	var slice []SkPmcolor
	ptr := (*reflect.SliceHeader)(unsafe.Pointer(&slice))
	ptr.Len = int(count)
	ptr.Cap = int(count)
	ptr.Data = uintptr(unsafe.Pointer(colors))
	return slice
}
