package skia

/*
#include "skia.h"
*/
import "C"
import (
	"unsafe"

	"github.com/uiez/skia-go/internal"
)

type skBitmap = C.sk_bitmap_t
type SkBitmap = skBitmap

func NewBitmap() *SkBitmap {
	return C.sk_bitmap_new()
}

func (v *SkBitmap) Destroy() {
	C.sk_bitmap_destructor(v)
}

func (v *SkBitmap) GetInfo() Imageinfo {
	var info Imageinfo
	C.sk_bitmap_get_info(v, info.native())
	return info
}

func (v *SkBitmap) GetPixels() []byte {
	var length csize_t
	data := C.sk_bitmap_get_pixels(v, &length)
	if length <= 0 {
		return nil
	}
	return internal.GoBytes(data, int(length))
}

func (v *SkBitmap) GetRowBytes() uint {
	return uint(C.sk_bitmap_get_row_bytes(v))
}

func (v *SkBitmap) GetByteCount() uint {
	return uint(C.sk_bitmap_get_byte_count(v))
}

func (v *SkBitmap) Reset() {
	C.sk_bitmap_reset(v)
}

func (v *SkBitmap) IsNull() bool {
	return bool(C.sk_bitmap_is_null(v))
}

func (v *SkBitmap) IsImmutable() bool {
	return bool(C.sk_bitmap_is_immutable(v))
}

func (v *SkBitmap) SetImmutable() {
	C.sk_bitmap_set_immutable(v)
}

func (v *SkBitmap) IsVolatile() bool {
	return bool(C.sk_bitmap_is_volatile(v))
}

func (v *SkBitmap) SetVolatile(val bool) {
	C.sk_bitmap_set_volatile(v, cbool_t(val))
}

func (v *SkBitmap) Erase(c SkColor) {
	C.sk_bitmap_erase(v, c)
}

func (v *SkBitmap) EraseRect(color SkColor, rect *Irect) {
	C.sk_bitmap_erase_rect(v, color, rect.native())
}

func (v *SkBitmap) GetAddr8(x, y int) *byte {
	return (*byte)(unsafe.Pointer(C.sk_bitmap_get_addr_8(v, cint_t(x), cint_t(y))))
}

func (v *SkBitmap) GetAddr16(x, y int) *uint16 {
	return (*uint16)(unsafe.Pointer(C.sk_bitmap_get_addr_16(v, cint_t(x), cint_t(y))))
}

func (v *SkBitmap) GetAddr32(x, y int) *uint32 {
	return (*uint32)(unsafe.Pointer(C.sk_bitmap_get_addr_32(v, cint_t(x), cint_t(y))))
}

func (v *SkBitmap) GetAddr(x, y int) unsafe.Pointer {
	return C.sk_bitmap_get_addr(v, cint_t(x), cint_t(y))
}

func (v *SkBitmap) GetPixelColor(x, y int) SkColor {
	return SkColor(C.sk_bitmap_get_pixel_color(v, cint_t(x), cint_t(y)))
}

func (v *SkBitmap) ReadyToDraw() bool {
	return bool(C.sk_bitmap_ready_to_draw(v))
}

type skBitmapAllocflags = C.sk_bitmap_allocflags_t
type SkBitmapAllocflags = skBitmapAllocflags

const (
	SK_BITMAP_ALLOC_FLAGS_NONE        = C.NONE_SK_BITMAP_ALLOC_FLAGS
	SK_BITMAP_ALLOC_FLAGS_ZERO_PIXELS = C.ZERO_PIXELS_SK_BITMAP_ALLOC_FLAGS
)

func (v *SkBitmap) InstallPixels(info *Imageinfo, pixels unsafe.Pointer, rowBytes uint, releaseProc SkBitmapReleaseProc, context interface{}) bool {
	crelease, cctx := skiaBitmapReleaseProcConvert(releaseProc, context)
	return bool(C.sk_bitmap_install_pixels(v, info.native(), pixels, csize_t(rowBytes), crelease, cctx))
}

func (v *SkBitmap) InstallPixelsWithPixmap(pixmap *SkPixmap) bool {
	return bool(C.sk_bitmap_install_pixels_with_pixmap(v, pixmap))
}

func (v *SkBitmap) InstallMaskPixels(mask *Mask) bool {
	return bool(C.sk_bitmap_install_mask_pixels(v, mask.native()))
}

func (v *SkBitmap) TryAllocPixels(info *Imageinfo, rowBytes uint) bool {
	return bool(C.sk_bitmap_try_alloc_pixels(v, info.native(), csize_t(rowBytes)))
}

func (v *SkBitmap) TryAllocPixelsWithFlags(info *Imageinfo, flags uint32) bool {
	return bool(C.sk_bitmap_try_alloc_pixels_with_flags(v, info.native(), cuint32_t(flags)))
}

func (v *SkBitmap) SetPixels(pixels unsafe.Pointer) {
	C.sk_bitmap_set_pixels(v, pixels)
}

func (v *SkBitmap) PeekPixels(pixmap *SkPixmap) bool {
	return bool(C.sk_bitmap_peek_pixels(v, pixmap))
}

func (v *SkBitmap) ExtractSubset(dst *SkBitmap, subset *Irect) bool {
	return bool(C.sk_bitmap_extract_subset(v, dst, subset.native()))
}

func (v *SkBitmap) ExtractAlpha(dst *SkBitmap, paint *SkPaint, offset *Ipoint) bool {
	return bool(C.sk_bitmap_extract_alpha(v, dst, paint, offset.native()))
}

func (v *SkBitmap) NotifyPixelsChanged() {
	C.sk_bitmap_notify_pixels_changed(v)
}

func (v *SkBitmap) Swap(Cother *SkBitmap) {
	C.sk_bitmap_swap(v, Cother)
}

func (v *SkBitmap) MakeShader(modeX, modeY SkShaderTilemode, matrix *Matrix) *SkShader {
	return C.sk_bitmap_make_shader(v, modeX, modeY, matrix.native())
}
