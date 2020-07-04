package skia

/*
#include "skia.h"
*/
import "C"
import (
	"unsafe"
)

type skImage = C.sk_image_t
type SkImage = skImage

func NewImageRasterCopy(info *Imageinfo, pixels unsafe.Pointer, rowBytes uint) *SkImage {
	return C.sk_image_new_raster_copy(info.native(), pixels, csize_t(rowBytes))
}

func NewImageRasterCopyWithPixmap(pixmap *SkPixmap) *SkImage {
	return C.sk_image_new_raster_copy_with_pixmap(pixmap)
}

func NewImageRasterData(info *Imageinfo, pixels *SkData, rowBytes uint) *SkImage {
	return C.sk_image_new_raster_data(info.native(), pixels, csize_t(rowBytes))
}

func NewImageRaster(pixmap *SkPixmap, releaseProc SkImageRasterReleaseProc, context interface{}) *SkImage {
	crelease, ccontext := skiaImageRasterReleaseProcConvert(releaseProc, context)
	return C.sk_image_new_raster(pixmap, crelease, ccontext)
}

func NewImageFromBitmap(bitmap *SkBitmap) *SkImage {
	return C.sk_image_new_from_bitmap(bitmap)
}

func NewImageFromEncoded(encoded *SkData, subset *Irect) *SkImage {
	return C.sk_image_new_from_encoded(encoded, subset.native())
}

func NewImageFromTexture(context *GrContext, texture *GrBackendtexture, origin GrSurfaceorigin, colorType SkColortype, alpha SkAlphatype, colorSpace *SkColorspace, releaseProc SkImageTextureReleaseProc, releaseContext interface{}) *SkImage {
	crelease, ccontext := skiaImageTextureReleaseProcConvert(releaseProc, releaseContext)
	return C.sk_image_new_from_texture(context, texture, origin, colorType, alpha, colorSpace, crelease, ccontext)
}

func NewImageFromAdoptedTexture(context *GrContext, texture *GrBackendtexture, origin GrSurfaceorigin, colorType SkColortype, alpha SkAlphatype, colorSpace *SkColorspace) *SkImage {
	return C.sk_image_new_from_adopted_texture(context, texture, origin, colorType, alpha, colorSpace)
}

func NewImageFromPicture(picture *SkPicture, dimensions *Isize, matrix *Matrix, paint *SkPaint) *SkImage {
	return C.sk_image_new_from_picture(picture, dimensions.native(), matrix.native(), paint)
}

func (v *SkImage) MakeSubset(subset *Irect) *SkImage {
	return C.sk_image_make_subset(v, subset.native())
}

func (v *SkImage) MakeTextureImage(context *GrContext, mipmapped bool) *SkImage {
	return C.sk_image_make_texture_image(v, context, cbool_t(mipmapped))
}

func (v *SkImage) MakeNonTextureImage() *SkImage {
	return C.sk_image_make_non_texture_image(v)
}

func (v *SkImage) MakeRasterImage() *SkImage {
	return C.sk_image_make_raster_image(v)
}

func (v *SkImage) MakeWithFilter(filter *SkImagefilter, subset *Irect, clipBounds *Irect, outSubset *Irect, outOffset *Ipoint) *SkImage {
	return C.sk_image_make_with_filter(v, filter, subset.native(), clipBounds.native(), outSubset.native(), outOffset.native())
}

func (v *SkImage) Ref() {
	C.sk_image_ref(v)
}

func (v *SkImage) Unref() {
	C.sk_image_unref(v)
}

func (v *SkImage) GetWidth() int32 {
	return int32(C.sk_image_get_width(v))
}

func (v *SkImage) GetHeight() int32 {
	return int32(C.sk_image_get_height(v))
}

func (v *SkImage) GetUniqueId() uint32 {
	return uint32(C.sk_image_get_unique_id(v))
}

func (v *SkImage) GetAlphaType() SkAlphatype {
	return C.sk_image_get_alpha_type(v)
}

func (v *SkImage) GetColorType() SkColortype {
	return C.sk_image_get_color_type(v)
}

func (v *SkImage) GetColorspace() *SkColorspace {
	return C.sk_image_get_colorspace(v)
}

func (v *SkImage) IsAlphaOnly() bool {
	return bool(C.sk_image_is_alpha_only(v))
}

func (v *SkImage) MakeShader(tileX SkShaderTilemode, tileY SkShaderTilemode, localMatrix *Matrix) *SkShader {
	return C.sk_image_make_shader(v, tileX, tileY, localMatrix.native())
}

func (v *SkImage) PeekPixels(pixmap *SkPixmap) bool {
	return bool(C.sk_image_peek_pixels(v, pixmap))
}

func (v *SkImage) IsTextureBacked() bool {
	return bool(C.sk_image_is_texture_backed(v))
}

func (v *SkImage) IsLazyGenerated() bool {
	return bool(C.sk_image_is_lazy_generated(v))
}

func (v *SkImage) IsValid(context *GrContext) bool {
	return bool(C.sk_image_is_valid(v, context))
}

func (v *SkImage) ReadPixels(dstInfo *Imageinfo, dstPixels unsafe.Pointer, dstRowBytes uint, srcX int32, srcY int32, cachingHint SkImageCachingHint) bool {
	return bool(C.sk_image_read_pixels(v, dstInfo.native(), dstPixels, csize_t(dstRowBytes), cint32_t(srcX), cint32_t(srcY), cachingHint))
}

type skImageCachingHint = C.sk_image_caching_hint_t
type SkImageCachingHint = skImageCachingHint

const (
	SK_IMAGE_CACHING_HINT_ALLOW    = C.ALLOW_SK_IMAGE_CACHING_HINT
	SK_IMAGE_CACHING_HINT_DISALLOW = C.DISALLOW_SK_IMAGE_CACHING_HINT
)

func (v *SkImage) ReadPixelsIntoPixmap(dst *SkPixmap, srcX int32, srcY int32, cachingHint SkImageCachingHint) bool {
	return bool(C.sk_image_read_pixels_into_pixmap(v, dst, cint32_t(srcX), cint32_t(srcY), cachingHint))
}

func (v *SkImage) ScalePixels(dst *SkPixmap, quality SkFilterQuality, cachingHint SkImageCachingHint) bool {
	return bool(C.sk_image_scale_pixels(v, dst, quality, cachingHint))
}

func (v *SkImage) RefEncoded() *SkData {
	return C.sk_image_ref_encoded(v)
}

func (v *SkImage) Encode() *SkData {
	return C.sk_image_encode(v)
}

type skEncodedImageFormat = C.sk_encoded_image_format_t
type SkEncodedImageFormat = skEncodedImageFormat

const (
	SK_ENCODED_FORMAT_BMP  = C.BMP_SK_ENCODED_FORMAT
	SK_ENCODED_FORMAT_GIF  = C.GIF_SK_ENCODED_FORMAT
	SK_ENCODED_FORMAT_ICO  = C.ICO_SK_ENCODED_FORMAT
	SK_ENCODED_FORMAT_JPEG = C.JPEG_SK_ENCODED_FORMAT
	SK_ENCODED_FORMAT_PNG  = C.PNG_SK_ENCODED_FORMAT
	SK_ENCODED_FORMAT_WBMP = C.WBMP_SK_ENCODED_FORMAT
	SK_ENCODED_FORMAT_WEBP = C.WEBP_SK_ENCODED_FORMAT
	SK_ENCODED_FORMAT_PKM  = C.PKM_SK_ENCODED_FORMAT
	SK_ENCODED_FORMAT_KTX  = C.KTX_SK_ENCODED_FORMAT
	SK_ENCODED_FORMAT_ASTC = C.ASTC_SK_ENCODED_FORMAT
	SK_ENCODED_FORMAT_DNG  = C.DNG_SK_ENCODED_FORMAT
	SK_ENCODED_FORMAT_HEIF = C.HEIF_SK_ENCODED_FORMAT
)

func (v *SkImage) EncodeSpecific(encoder SkEncodedImageFormat, quality int) *SkData {
	return C.sk_image_encode_specific(v, encoder, cint_t(quality))
}
