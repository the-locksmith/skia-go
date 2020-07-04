package skia

/*
#include "skia.h"

void goBitmapReleaseProc(void *addr, void *context);
void goDataReleaseProc(void *ptr, void *context);
void goImageRasterReleaseProc(void *addr, void *context);
void goImageTextureReleaseProc(void *context);
void goSurfaceRasterReleaseProc(void *addr, void *context);
void goGlyphPathProc(sk_path_t *path, sk_matrix_t *matrix, void *context);
*/
import "C"
import (
	"unsafe"

	"github.com/zhuah/skia-go/internal"
)

type SkBitmapReleaseProc func(addr unsafe.Pointer, context interface{})

type SkDataReleaseProc func(ptr unsafe.Pointer, context interface{})

type SkImageRasterReleaseProc func(addr unsafe.Pointer, context interface{})

type SkImageTextureReleaseProc func(context interface{})

type SkSurfaceRasterReleaseProc func(addr unsafe.Pointer, context interface{})

type SkGlyphPathProc func(pathOrNull *SkPath, matrix *Matrix, context interface{})

type cSkBitmapReleaseProc = C.sk_bitmap_release_proc

type cSkDataReleaseProc = C.sk_data_release_proc

type cSkImageRasterReleaseProc = C.sk_image_raster_release_proc

type cSkImageTextureReleaseProc = C.sk_image_texture_release_proc

type cSkSurfaceRasterReleaseProc = C.sk_surface_raster_release_proc

type cSkGlyphPathProc = C.sk_glyph_path_proc

//export goBitmapReleaseProc
func goBitmapReleaseProc(addr unsafe.Pointer, context unsafe.Pointer) {
	internal.CallReleaseProcData(context, func(data internal.CommonProcData) {
		data.Proc.(SkBitmapReleaseProc)(addr, data.Context)
	})
	internal.PopReleaseProcData(context)
}

func skiaBitmapReleaseProcConvert(proc SkBitmapReleaseProc, context interface{}) (cSkBitmapReleaseProc, unsafe.Pointer) {
	ctx := internal.PushReleaseProcData(internal.CommonProcData{proc, context})
	return cSkBitmapReleaseProc(C.goBitmapReleaseProc), ctx
}

//export goDataReleaseProc
func goDataReleaseProc(ptr unsafe.Pointer, context unsafe.Pointer) {
	internal.CallReleaseProcData(context, func(data internal.CommonProcData) {
		data.Proc.(SkDataReleaseProc)(ptr, data.Context)
	})
	internal.PopReleaseProcData(context)
}

func skiaDataReleaseProcConvert(proc SkDataReleaseProc, context interface{}) (cSkDataReleaseProc, unsafe.Pointer) {
	ctx := internal.PushReleaseProcData(internal.CommonProcData{proc, context})
	return cSkDataReleaseProc(C.goDataReleaseProc), ctx
}

//export goImageRasterReleaseProc
func goImageRasterReleaseProc(addr unsafe.Pointer, context unsafe.Pointer) {
	internal.CallReleaseProcData(context, func(data internal.CommonProcData) {
		data.Proc.(SkImageRasterReleaseProc)(addr, data.Context)
	})
	internal.PopReleaseProcData(context)
}
func skiaImageRasterReleaseProcConvert(proc SkImageRasterReleaseProc, context interface{}) (cSkImageRasterReleaseProc, unsafe.Pointer) {
	ctx := internal.PushReleaseProcData(internal.CommonProcData{proc, context})
	return cSkImageRasterReleaseProc(C.goImageRasterReleaseProc), ctx
}

//export goImageTextureReleaseProc
func goImageTextureReleaseProc(context unsafe.Pointer) {
	internal.CallReleaseProcData(context, func(data internal.CommonProcData) {
		data.Proc.(SkImageTextureReleaseProc)(data.Context)
	})
	internal.PopReleaseProcData(context)
}
func skiaImageTextureReleaseProcConvert(proc SkImageTextureReleaseProc, context interface{}) (cSkImageTextureReleaseProc, unsafe.Pointer) {
	ctx := internal.PushReleaseProcData(internal.CommonProcData{proc, context})
	return cSkImageTextureReleaseProc(C.goImageTextureReleaseProc), ctx
}

//export goSurfaceRasterReleaseProc
func goSurfaceRasterReleaseProc(addr unsafe.Pointer, context unsafe.Pointer) {
	internal.CallReleaseProcData(context, func(data internal.CommonProcData) {
		data.Proc.(SkSurfaceRasterReleaseProc)(addr, data.Context)
	})
	internal.PopReleaseProcData(context)
}
func skiaSurfaceRasterReleaseProcConvert(proc SkSurfaceRasterReleaseProc, context interface{}) (cSkSurfaceRasterReleaseProc, unsafe.Pointer) {
	ctx := internal.PushReleaseProcData(internal.CommonProcData{proc, context})
	return cSkSurfaceRasterReleaseProc(C.goSurfaceRasterReleaseProc), ctx
}

//export goGlyphPathProc
func goGlyphPathProc(pathOrNull *SkPath, matrix *skMatrix, context unsafe.Pointer) {
	internal.CallReleaseProcData(context, func(data internal.CommonProcData) {
		data.Proc.(SkGlyphPathProc)(pathOrNull, fromNativeMatrix(matrix), data.Context)
	})
}

func skiaGlyphPathProcConvert(proc SkGlyphPathProc, context interface{}) (cSkGlyphPathProc, unsafe.Pointer, func()) {
	ctx := internal.PushReleaseProcData(internal.CommonProcData{proc, context})
	return cSkGlyphPathProc(C.goGlyphPathProc), ctx, func() {
		internal.PopReleaseProcData(ctx)
	}
}
