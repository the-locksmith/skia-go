package skia

/*
#include "skia.h"
*/
import "C"

type grContext = C.gr_context_t
type GrContext = grContext

type grBackendrendertarget = C.gr_backendrendertarget_t
type GrBackendrendertarget = grBackendrendertarget

type grBackendtexture = C.gr_backendtexture_t
type GrBackendtexture = grBackendtexture

type grBackendobject = C.gr_backendobject_t
type GrBackendobject = grBackendobject

type grBackendcontext = C.gr_backendcontext_t
type GrBackendcontext = grBackendcontext

func (v *GrContext) Unref() {
	C.gr_context_unref(v)
}

func (v *GrContext) AbandonContext() {
	C.gr_context_abandon_context(v)
}

func (v *GrContext) ReleaseResourcesAndAbandonContext() {
	C.gr_context_release_resources_and_abandon_context(v)
}

func (v *GrContext) GetResourceCacheLimits() (maxResources int32, maxResourceBytes uint) {
	var cmaxResources cint32_t
	var cmaxResourceBytes csize_t
	C.gr_context_get_resource_cache_limits(v, &cmaxResources, &cmaxResourceBytes)
	return int32(cmaxResources), uint(cmaxResourceBytes)
}

func (v *GrContext) SetResourceCacheLimits(maxResources int32, maxResourceBytes uint) {
	C.gr_context_set_resource_cache_limits(v, cint32_t(maxResources), csize_t(maxResourceBytes))
}

func (v *GrContext) GetResourceCacheUsage() (maxResources int32, maxResourceBytes uint) {
	var cmaxResources cint32_t
	var cmaxResourceBytes csize_t
	C.gr_context_get_resource_cache_usage(v, &cmaxResources, &cmaxResourceBytes)
	return int32(cmaxResources), uint(cmaxResourceBytes)
}

func (v *GrContext) GetMaxSurfaceSampleCountForColorType(colorType SkColortype) int32 {
	return int32(C.gr_context_get_max_surface_sample_count_for_color_type(v, colorType))
}

func (v *GrContext) Flush() {
	C.gr_context_flush(v)
}

func (v *GrContext) ResetContext(state uint32) {
	C.gr_context_reset_context(v, cuint32_t(state))
}

type grBackend = C.gr_backend_t
type GrBackend = grBackend

const (
	GR_BACKEND_OPENGL = C.OPENGL_GR_BACKEND
	GR_BACKEND_VULKAN = C.VULKAN_GR_BACKEND
	GR_BACKEND_METAL  = C.METAL_GR_BACKEND
	GR_BACKEND_DAWN   = C.DAWN_GR_BACKEND
)

func (v *GrContext) GetBackend() GrBackend {
	return C.gr_context_get_backend(v)
}

func (v *GrBackendtexture) Destroy() {
	C.gr_backendtexture_delete(v)
}

func (v *GrBackendtexture) IsValid() bool {
	return bool(C.gr_backendtexture_is_valid(v))
}

func (v *GrBackendtexture) GetWidth() int32 {
	return int32(C.gr_backendtexture_get_width(v))
}

func (v *GrBackendtexture) GetHeight() int32 {
	return int32(C.gr_backendtexture_get_height(v))
}

func (v *GrBackendtexture) HasMipmaps() bool {
	return bool(C.gr_backendtexture_has_mipmaps(v))
}

func (v *GrBackendtexture) GetBackend() GrBackend {
	return C.gr_backendtexture_get_backend(v)
}

func (v *GrBackendrendertarget) Destroy() {
	C.gr_backendrendertarget_delete(v)
}

func (v *GrBackendrendertarget) IsValid() bool {
	return bool(C.gr_backendrendertarget_is_valid(v))
}

func (v *GrBackendrendertarget) GetWidth() int32 {
	return int32(C.gr_backendrendertarget_get_width(v))
}

func (v *GrBackendrendertarget) GetHeight() int32 {
	return int32(C.gr_backendrendertarget_get_height(v))
}

func (v *GrBackendrendertarget) GetSamples() int32 {
	return int32(C.gr_backendrendertarget_get_samples(v))
}

func (v *GrBackendrendertarget) GetStencils() int32 {
	return int32(C.gr_backendrendertarget_get_stencils(v))
}

func (v *GrBackendrendertarget) GetBackend() GrBackend {
	return C.gr_backendrendertarget_get_backend(v)
}
