package skia

/*
#include "skia.h"
*/
import "C"
import (
	"unsafe"
)

type grGlinterface = C.gr_glinterface_t
type GrGlinterface = grGlinterface

type skGrGlTextureinfo = C.gr_gl_textureinfo_t
type GrGlTextureinfo struct {
	FTarget uint32
	FID     uint32
	FFormat uint32
}

func fromNativeGrGlTextureinfo(v *skGrGlTextureinfo) *GrGlTextureinfo {
	return (*GrGlTextureinfo)(unsafe.Pointer(v))
}
func (v *GrGlTextureinfo) native() *skGrGlTextureinfo {
	return (*skGrGlTextureinfo)(unsafe.Pointer(v))
}

type skGrGlFramebufferinfo = C.gr_gl_framebufferinfo_t
type GrGlFramebufferinfo struct {
	FFBOID  uint32
	FFormat uint32
}

func fromNativeGrGlFramebufferinfo(v *skGrGlFramebufferinfo) *GrGlFramebufferinfo {
	return (*GrGlFramebufferinfo)(unsafe.Pointer(v))
}
func (v *GrGlFramebufferinfo) native() *skGrGlFramebufferinfo {
	return (*skGrGlFramebufferinfo)(unsafe.Pointer(v))
}
func NewGLGrContext(glInterface *GrGlinterface) *GrContext {
	ptr := C.gr_context_make_gl(glInterface)
	return (*GrContext)(ptr)
}

func NewNativeGrGlinterface() *GrGlinterface {
	return C.gr_glinterface_create_native_interface()
}

func NewAssembleGrGlinterface(ctx interface{}, get GrGlGetProc) *GrGlinterface {
	cget, cctx, free := skiaGrGlGetProcConvert(get, ctx)
	v := C.gr_glinterface_assemble_interface(cctx, cget)
	free()

	return v
}

func NewAssembleGlGrGlinterface(ctx interface{}, get GrGlGetProc) *GrGlinterface {
	cget, cctx, free := skiaGrGlGetProcConvert(get, ctx)
	v := C.gr_glinterface_assemble_gl_interface(cctx, cget)
	free()

	return v
}

func NewAssembleGlesGrGlinterface(ctx interface{}, get GrGlGetProc) *GrGlinterface {
	cget, cctx, free := skiaGrGlGetProcConvert(get, ctx)
	v := C.gr_glinterface_assemble_gles_interface(cctx, cget)
	free()

	return v
}

func (v *GrGlinterface) Unref() {
	C.gr_glinterface_unref(v)
}

func (v *GrGlinterface) Validate() bool {
	return bool(C.gr_glinterface_validate(v))
}

func (v *GrGlinterface) HasExtension(extension string) bool {
	cstr, free := cstring(extension)
	ok := bool(C.gr_glinterface_has_extension(v, cstr))
	free()

	return ok
}

func (v *GrGlinterface) RemoveExtension(extension string) bool {
	cstr, free := cstring(extension)
	ok := bool(C.gr_glinterface_remove_extension(v, cstr))
	free()
	return ok
}

func NewGlGrBackendtexture(width int32, height int32, mipmapped bool, glInfo *GrGlTextureinfo) *GrBackendtexture {
	ptr := C.gr_backendtexture_new_gl(cint32_t(width), cint32_t(height), cbool_t(mipmapped), glInfo.native())
	return (*GrBackendtexture)(ptr)
}

func NewGlGrBackendrendertarget(width int32, height int32, samples int32, stencils int32, glInfo *GrGlFramebufferinfo) *GrBackendrendertarget {
	ptr := C.gr_backendrendertarget_new_gl(cint32_t(width), cint32_t(height), cint32_t(samples), cint32_t(stencils), glInfo.native())
	return (*GrBackendrendertarget)(ptr)
}

func GrBackendrendertargetGetGlFramebufferinfo(rendertarget *GrBackendrendertarget, glInfo *GrGlFramebufferinfo) bool {
	return bool(C.gr_backendrendertarget_get_gl_framebufferinfo((*C.gr_backendrendertarget_t)(unsafe.Pointer(rendertarget)), glInfo.native()))
}

func GetGlTextureinfo(texture *GrBackendtexture, glInfo *GrGlTextureinfo) bool {
	return bool(C.gr_backendtexture_get_gl_textureinfo((*C.gr_backendtexture_t)(unsafe.Pointer(texture)), glInfo.native()))
}
