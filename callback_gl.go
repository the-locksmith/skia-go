package skia

/*
#include "skia.h"

void *goGrGlGetProc(void *ctx, char *name);
*/
import "C"
import (
	"unsafe"

	"github.com/zhuah/skia-go/internal"
)

type GrGlGetProc func(ctx interface{}, name string) unsafe.Pointer

type cGrGlGetProc = C.gr_gl_get_proc

//export goGrGlGetProc
func goGrGlGetProc(ctx unsafe.Pointer, name *cchar_t) unsafe.Pointer {
	ptr, ok := internal.CallReleaseProcData(ctx, func(data internal.CommonProcData) unsafe.Pointer {
		return data.Proc.(GrGlGetProc)(data.Context, C.GoString(name))
	})
	if !ok {
		return nil
	}
	return ptr.(unsafe.Pointer)
}

func skiaGrGlGetProcConvert(proc GrGlGetProc, context interface{}) (cGrGlGetProc, unsafe.Pointer, func()) {
	ctx := internal.PushReleaseProcData(internal.CommonProcData{proc, context})
	return cGrGlGetProc(C.goGrGlGetProc), ctx, func() {
		internal.PopReleaseProcData(ctx)
	}
}
