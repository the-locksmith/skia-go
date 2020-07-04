package skia

/*
#include "skia.h"
*/
import "C"
import (
	"unsafe"

	"github.com/zhuah/skia-go/internal"
)

type skData = C.sk_data_t
type SkData = skData

func NewDataEmpty() *SkData {
	return C.sk_data_new_empty()
}

func NewDataWithCopy(src unsafe.Pointer, length uint) *SkData {
	return C.sk_data_new_with_copy(src, csize_t(length))
}

func NewDataSubset(src *SkData, offset uint, length uint) *SkData {
	return C.sk_data_new_subset(src, csize_t(offset), csize_t(length))
}
func NewDataFromFile(path string) *SkData {
	cs, free := cstring(path)
	data := C.sk_data_new_from_file(cs)
	free()
	return data
}

func NewDataWithProc(ptr unsafe.Pointer, length uint, releaseProc SkDataReleaseProc, ctx interface{}) *SkData {
	crelease, cctx := skiaDataReleaseProcConvert(releaseProc, ctx)
	return C.sk_data_new_with_proc(ptr, csize_t(length), crelease, cctx)
}

func NewDataUninitialized(size uint) *SkData {
	return C.sk_data_new_uninitialized(csize_t(size))
}

func (v *SkData) Ref() {
	C.sk_data_ref(v)
}

func (v *SkData) Unref() {
	C.sk_data_unref(v)
}

func (v *SkData) GetSize() uint {
	return uint(C.sk_data_get_size(v))
}

func (v *SkData) GetData() []byte {
	data := C.sk_data_get_data(v)
	return internal.GoBytes(data, int(v.GetSize()))
}

func (v *SkData) GetBytes() []byte {
	data := C.sk_data_get_bytes(v)
	return internal.GoBytes(unsafe.Pointer(data), int(v.GetSize()))
}
