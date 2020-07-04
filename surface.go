package skia

/*
#include "skia.h"
*/
import "C"
import (
	"unsafe"
)

type skSurfaceprops = C.sk_surfaceprops_t
type SkSurfaceprops = skSurfaceprops

type skSurface = C.sk_surface_t
type SkSurface = skSurface

func NewSurfaceNull(width int32, height int32) *SkSurface {
	return C.sk_surface_new_null(cint32_t(width), cint32_t(height))
}

func NewSurfaceRaster(info *Imageinfo, rowBytes uint, props *SkSurfaceprops) *SkSurface {
	return C.sk_surface_new_raster(info.native(), csize_t(rowBytes), props)
}

func NewSurfaceRasterDirect(info *Imageinfo, pixels unsafe.Pointer, rowBytes uint, releaseProc SkSurfaceRasterReleaseProc, context interface{}, props *SkSurfaceprops) *SkSurface {
	crelease, cctx := skiaSurfaceRasterReleaseProcConvert(releaseProc, context)
	return C.sk_surface_new_raster_direct(info.native(), pixels, csize_t(rowBytes), crelease, cctx, props)
}

type GrSurfaceorigin = C.gr_surfaceorigin_t

const (
	GR_SURFACE_ORIGIN_TOP_LEFT    = C.TOP_LEFT_GR_SURFACE_ORIGIN
	GR_SURFACE_ORIGIN_BOTTOM_LEFT = C.BOTTOM_LEFT_GR_SURFACE_ORIGIN
)

func NewSurfaceBackendTexture(context *GrContext, texture *GrBackendtexture, origin GrSurfaceorigin, samples int32, colorType SkColortype, colorspace *SkColorspace, props *SkSurfaceprops) *SkSurface {
	return C.sk_surface_new_backend_texture(context, texture, origin, cint32_t(samples), colorType, colorspace, props)
}

func NewSurfaceBackendRenderTarget(context *GrContext, target *GrBackendrendertarget, origin GrSurfaceorigin, colorType SkColortype, colorspace *SkColorspace, props *SkSurfaceprops) *SkSurface {
	return C.sk_surface_new_backend_render_target(context, target, origin, colorType, colorspace, props)
}

func NewSurfaceBackendTextureAsRenderTarget(context *GrContext, texture *GrBackendtexture, origin GrSurfaceorigin, samples int32, colorType SkColortype, colorspace *SkColorspace, props *SkSurfaceprops) *SkSurface {
	return C.sk_surface_new_backend_texture_as_render_target(context, texture, origin, cint32_t(samples), colorType, colorspace, props)
}

func NewSurfaceRenderTarget(context *GrContext, budgeted bool, info *Imageinfo, sampleCount int32, origin GrSurfaceorigin, props *SkSurfaceprops, shouldCreateWithMips bool) *SkSurface {
	return C.sk_surface_new_render_target(context, cbool_t(budgeted), info.native(), cint32_t(sampleCount), origin, props, cbool_t(shouldCreateWithMips))
}

func (v *SkSurface) Unref() {
	C.sk_surface_unref(v)
}

func (v *SkSurface) GetCanvas() *SkCanvas {
	return C.sk_surface_get_canvas(v)
}

func NewSurfaceImageSnapshot(v *SkSurface) *SkImage {
	return C.sk_surface_new_image_snapshot(v)
}

func (v *SkSurface) Draw(canvas *SkCanvas, x float32, y float32, paint *SkPaint) {
	C.sk_surface_draw(v, canvas, cfloat_t(x), cfloat_t(y), paint)
}

func (v *SkSurface) PeekPixels(pixmap *SkPixmap) bool {
	return bool(C.sk_surface_peek_pixels(v, pixmap))
}

func (v *SkSurface) ReadPixels(dstInfo *Imageinfo, dstPixels unsafe.Pointer, dstRowBytes uint, srcX int32, srcY int32) bool {
	return bool(C.sk_surface_read_pixels(v, dstInfo.native(), dstPixels, csize_t(dstRowBytes), cint32_t(srcX), cint32_t(srcY)))
}

func (v *SkSurface) GetProps() *SkSurfaceprops {
	return C.sk_surface_get_props(v)
}

type skPixelgeometry = C.sk_pixelgeometry_t
type SkPixelgeometry = skPixelgeometry

const (
	SK_PIXELGEOMETRY_UNKNOWN = C.UNKNOWN_SK_PIXELGEOMETRY
	SK_PIXELGEOMETRY_RGB_H   = C.RGB_H_SK_PIXELGEOMETRY
	SK_PIXELGEOMETRY_BGR_H   = C.BGR_H_SK_PIXELGEOMETRY
	SK_PIXELGEOMETRY_RGB_V   = C.RGB_V_SK_PIXELGEOMETRY
	SK_PIXELGEOMETRY_BGR_V   = C.BGR_V_SK_PIXELGEOMETRY
)

func NewSurfaceprops(flags uint32, geometry SkPixelgeometry) *SkSurfaceprops {
	return C.sk_surfaceprops_new(cuint32_t(flags), geometry)
}

func (v *SkSurfaceprops) Destroy() {
	C.sk_surfaceprops_delete(v)
}

func (v *SkSurfaceprops) GetFlags() uint32 {
	return uint32(C.sk_surfaceprops_get_flags(v))
}

func (v *SkSurfaceprops) GetPixelGeometry() SkPixelgeometry {
	return C.sk_surfaceprops_get_pixel_geometry(v)
}
