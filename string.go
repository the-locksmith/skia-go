package skia

/*
#include "skia.h"
*/
import "C"
import (
	"unsafe"

	"github.com/uiez/skia-go/internal"
)

type skString = C.sk_string_t
type SkString = skString

func NewStringEmpty() *SkString {
	return C.sk_string_new_empty()
}

func NewStringWithCopy(src string) *SkString {
	ptr := internal.StringDataPtr(src)
	return C.sk_string_new_with_copy((*cchar_t)(ptr), csize_t(len(src)))
}

func (v *SkString) Destroy() {
	C.sk_string_destructor(v)
}

func (v *SkString) GetSize() uint {
	return uint(C.sk_string_get_size(v))
}

func (v *SkString) GetData() []byte {
	size := int(v.GetSize())
	ptr := C.sk_string_get_c_str(v)

	return internal.GoBytes(unsafe.Pointer(ptr), size)
}
