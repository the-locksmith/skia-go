package skia

/*
#include "skia.h"
*/
import "C"
import (
	"unsafe"
)

type skPixmap = C.sk_pixmap_t
type SkPixmap = skPixmap

func NewPixmap() *SkPixmap {
	return C.sk_pixmap_new()
}

func NewPixmapWithParams(info *Imageinfo, addr unsafe.Pointer, rowBytes uint) *SkPixmap {
	return C.sk_pixmap_new_with_params(info.native(), addr, csize_t(rowBytes))
}

func (v *SkPixmap) Destroy() {
	C.sk_pixmap_destructor(v)
}

func (v *SkPixmap) Reset() {
	C.sk_pixmap_reset(v)
}

func (v *SkPixmap) ResetWithParams(cinfo *Imageinfo, addr unsafe.Pointer, rowBytes uint) {
	C.sk_pixmap_reset_with_params(v, cinfo.native(), addr, csize_t(rowBytes))
}

func (v *SkPixmap) GetInfo() Imageinfo {
	var info Imageinfo
	C.sk_pixmap_get_info(v, info.native())
	return info
}

func (v *SkPixmap) GetRowBytes() uint {
	return uint(C.sk_pixmap_get_row_bytes(v))
}

func (v *SkPixmap) GetWritableAddr() unsafe.Pointer {
	return C.sk_pixmap_get_writable_addr(v)
}

func (v *SkPixmap) GetPixels() unsafe.Pointer {
	return C.sk_pixmap_get_pixels(v)
}

func (v *SkPixmap) GetPixelsWithXy(x int32, y int32) unsafe.Pointer {
	return C.sk_pixmap_get_pixels_with_xy(v, cint32_t(x), cint32_t(y))
}

func (v *SkPixmap) GetPixelColor(x int32, y int32) SkColor {
	return C.sk_pixmap_get_pixel_color(v, cint32_t(x), cint32_t(y))
}

func (v *SkPixmap) ReadPixels(dstInfo *Imageinfo, dstPixels unsafe.Pointer, dstRowBytes uint, srcX int32, srcY int32) bool {
	return bool(C.sk_pixmap_read_pixels(v, dstInfo.native(), dstPixels, csize_t(dstRowBytes), cint32_t(srcX), cint32_t(srcY)))
}

func (v *SkPixmap) ScalePixels(dst *SkPixmap, quality SkFilterQuality) bool {
	return bool(C.sk_pixmap_scale_pixels(v, dst, quality))
}

func (v *SkPixmap) ExtractSubset(dst *SkPixmap, subset *Irect) bool {
	return bool(C.sk_pixmap_extract_subset(v, dst, subset.native()))
}

func (v *SkPixmap) EraseColor(color SkColor, subset *Irect) bool {
	return bool(C.sk_pixmap_erase_color(v, color, subset.native()))
}

func (v *SkPixmap) EraseColor4f(color *Color4f, subset *Irect) bool {
	return bool(C.sk_pixmap_erase_color4f(v, color.native(), subset.native()))
}
