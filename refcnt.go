package skia

/*
#include "skia.h"
*/
import "C"

type skRefcnt = C.sk_refcnt_t
type SkRefcnt = skRefcnt

type skNvrefcnt = C.sk_nvrefcnt_t
type SkNvrefcnt = skNvrefcnt

func (v *SkRefcnt) Unique() bool {
	return bool(C.sk_refcnt_unique(v))
}

func (v *SkRefcnt) GetRefCount() int32 {
	return int32(C.sk_refcnt_get_ref_count(v))
}

func (v *SkRefcnt) SafeRef() {
	C.sk_refcnt_safe_ref(v)
}

func (v *SkRefcnt) SafeUnref() {
	C.sk_refcnt_safe_unref(v)
}

func (v *SkNvrefcnt) Unique() bool {
	return bool(C.sk_nvrefcnt_unique(v))
}

func (v *SkNvrefcnt) GetRefCount() int32 {
	return int32(C.sk_nvrefcnt_get_ref_count(v))
}

func (v *SkNvrefcnt) SafeRef() {
	C.sk_nvrefcnt_safe_ref(v)
}

func (v *SkNvrefcnt) SafeUnref() {
	C.sk_nvrefcnt_safe_unref(v)
}
