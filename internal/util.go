package internal

/*
#include <stdlib.h>
#include "skia.h"
*/
import "C"
import (
	"reflect"
	"unsafe"

	"github.com/mattn/go-pointer"
)

type CommonProcData struct {
	Proc    interface{}
	Context interface{}
}

func PushReleaseProcData(v interface{}) unsafe.Pointer {
	return pointer.Save(reflect.ValueOf(v))
}

func PopReleaseProcData(ptr unsafe.Pointer) {
	pointer.Unref(ptr)
}

func CallReleaseProcData(ptr unsafe.Pointer, fn interface{}) (interface{}, bool) {
	d := pointer.Restore(ptr)
	if d == nil {
		return nil, false
	}

	res := reflect.ValueOf(fn).Call([]reflect.Value{d.(reflect.Value)})
	if len(res) > 0 {
		return res[0].Interface(), true
	}
	return nil, false
}

func Cstring(s string) (unsafe.Pointer, func()) {
	if s == "" {
		return nil, func() {}
	}
	v := unsafe.Pointer(C.CString(s))
	return v, func() {
		C.free(v)
	}
}

func SliceDataPtr(v interface{}) unsafe.Pointer {
	return unsafe.Pointer(reflect.ValueOf(v).Pointer())
}

func StringDataPtr(v string) unsafe.Pointer {
	hdr := *(*reflect.StringHeader)(unsafe.Pointer(&v))
	return unsafe.Pointer(hdr.Data)
}

func GoBytes(ptr unsafe.Pointer, length int) []byte {
	if length <= 0 {
		return nil
	}
	var b []byte
	hdr := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	hdr.Data = uintptr(ptr)
	hdr.Len = length
	hdr.Cap = hdr.Len
	return b
}
