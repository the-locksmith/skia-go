package skia

/*
#include <stdint.h>
#include <stdbool.h>
*/
import "C"
import (
	"github.com/zhuah/skia-go/internal"
)

type csize_t = C.size_t
type cbool_t = C._Bool
type cuint8_t = C.uint8_t
type cuint16_t = C.uint16_t
type cuint32_t = C.uint32_t
type cuint64_t = C.uint64_t
type cint8_t = C.int8_t
type cint16_t = C.int16_t
type cshort_t = C.short
type cushort_t = C.ushort
type cint32_t = C.int32_t
type cint64_t = C.int64_t
type clong_t = C.long
type culong_t = C.ulong
type cfloat_t = C.float
type cdouble_t = C.double
type cint_t = C.int
type cuint_t = C.uint
type cchar_t = C.char
type cuchar_t = C.uchar

func cstring(str string) (*cchar_t, func()) {
	ptr, free := internal.Cstring(str)
	return (*cchar_t)(ptr), free
}
