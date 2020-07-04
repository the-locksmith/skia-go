package skia

/*
#include "skia.h"
*/
import "C"
import (
	"unsafe"
)

type skMaskFormat = C.sk_mask_format_t
type SkMaskFormat = skMaskFormat

const (
	SK_MASK_FORMAT_BW      = C.BW_SK_MASK_FORMAT
	SK_MASK_FORMAT_A8      = C.A8_SK_MASK_FORMAT
	SK_MASK_FORMAT_THREE_D = C.THREE_D_SK_MASK_FORMAT
	SK_MASK_FORMAT_ARGB32  = C.ARGB32_SK_MASK_FORMAT
	SK_MASK_FORMAT_LCD16   = C.LCD16_SK_MASK_FORMAT
	SK_MASK_FORMAT_SDF     = C.SDF_SK_MASK_FORMAT
)

type skMask = C.sk_mask_t
type Mask struct {
	FImage    uintptr
	FBounds   Irect
	FRowBytes uint32
	FFormat   SkMaskFormat
}

func (v *Mask) native() *skMask {
	return (*skMask)(unsafe.Pointer(v))
}

func AllocMaskImage(bytes uint) unsafe.Pointer {
	return unsafe.Pointer(C.sk_mask_alloc_image(csize_t(bytes)))
}

func FreeMaskImage(image unsafe.Pointer) {
	C.sk_mask_free_image(image)
}

func (v *Mask) IsEmpty() bool {
	return bool(C.sk_mask_is_empty(v.native()))
}

func (v *Mask) ComputeImageSize() uint {
	return uint(C.sk_mask_compute_image_size(v.native()))
}

func (v *Mask) ComputeTotalImageSize() uint {
	return uint(C.sk_mask_compute_total_image_size(v.native()))
}

func (v *Mask) GetAddr1(x int32, y int32) *byte {
	return (*byte)(unsafe.Pointer(C.sk_mask_get_addr_1(v.native(), cint32_t(x), cint32_t(y))))
}

func (v *Mask) GetAddr8(x int32, y int32) *byte {
	return (*byte)(unsafe.Pointer(C.sk_mask_get_addr_8(v.native(), cint32_t(x), cint32_t(y))))
}

func (v *Mask) GetAddrLcd16(x int32, y int32) *uint16 {
	return (*uint16)(unsafe.Pointer(C.sk_mask_get_addr_lcd_16(v.native(), cint32_t(x), cint32_t(y))))
}

func (v *Mask) GetAddr32(x int32, y int32) *uint32 {
	return (*uint32)(unsafe.Pointer(C.sk_mask_get_addr_32(v.native(), cint32_t(x), cint32_t(y))))
}

func (v *Mask) GetAddr(x int32, y int32) unsafe.Pointer {
	return (unsafe.Pointer)(unsafe.Pointer(C.sk_mask_get_addr(v.native(), cint32_t(x), cint32_t(y))))
}
