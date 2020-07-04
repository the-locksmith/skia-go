package skia

/*
#include "skia.h"
*/
import "C"
import (
	"unsafe"
)

type skAlphatype = C.sk_alphatype_t
type SkAlphatype = skAlphatype

const (
	SK_ALPHATYPE_UNKNOWN  = C.UNKNOWN_SK_ALPHATYPE
	SK_ALPHATYPE_OPAQUE   = C.OPAQUE_SK_ALPHATYPE
	SK_ALPHATYPE_PREMUL   = C.PREMUL_SK_ALPHATYPE
	SK_ALPHATYPE_UNPREMUL = C.UNPREMUL_SK_ALPHATYPE
)

type skImageinfo = C.sk_imageinfo_t
type Imageinfo struct {
	Colorspace *SkColorspace
	Width      int32
	Height     int32
	ColorType  SkColortype
	AlphaType  SkAlphatype
}

func fromNativeImageinfo(v *skImageinfo) *Imageinfo {
	return (*Imageinfo)(unsafe.Pointer(v))
}
func (v *Imageinfo) native() *skImageinfo {
	return (*skImageinfo)(unsafe.Pointer(v))
}
