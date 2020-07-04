package skia

/*
#include "skia.h"
*/
import "C"
import (
	"unsafe"
)

type skPoint = C.sk_point_t
type Point struct {
	X float32
	Y float32
}

func fromNativePoint(v *skPoint) *Point {
	return (*Point)(unsafe.Pointer(v))
}
func (v *Point) native() *skPoint {
	return (*skPoint)(unsafe.Pointer(v))
}

type skVector = C.sk_vector_t
type Vector struct {
	X float32
	Y float32
}

func fromNativeVector(v *skVector) *Vector {
	return (*Vector)(unsafe.Pointer(v))
}
func (v *Vector) native() *skVector {
	return (*skVector)(unsafe.Pointer(v))
}

type skIrect = C.sk_irect_t
type Irect struct {
	Left   int32
	Top    int32
	Right  int32
	Bottom int32
}

func fromNativeIrect(v *skIrect) *Irect {
	return (*Irect)(unsafe.Pointer(v))
}
func (v *Irect) native() *skIrect {
	return (*skIrect)(unsafe.Pointer(v))
}

type skRect = C.sk_rect_t
type Rect struct {
	Left   float32
	Top    float32
	Right  float32
	Bottom float32
}

func fromNativeRect(v *skRect) *Rect {
	return (*Rect)(unsafe.Pointer(v))
}
func (v *Rect) native() *skRect {
	return (*skRect)(unsafe.Pointer(v))
}

type skPoint3 = C.sk_point3_t
type Point3 struct {
	X float32
	Y float32
	Z float32
}

func fromNativePoint3(v *skPoint3) *Point3 {
	return (*Point3)(unsafe.Pointer(v))
}
func (v *Point3) native() *skPoint3 {
	return (*skPoint3)(unsafe.Pointer(v))
}

type skIpoint = C.sk_ipoint_t
type Ipoint struct {
	X int32
	Y int32
}

func fromNativeIpoint(v *skIpoint) *Ipoint {
	return (*Ipoint)(unsafe.Pointer(v))
}
func (v *Ipoint) native() *skIpoint {
	return (*skIpoint)(unsafe.Pointer(v))
}

type skSize = C.sk_size_t
type Size struct {
	W float32
	H float32
}

func fromNativeSize(v *skSize) *Size {
	return (*Size)(unsafe.Pointer(v))
}
func (v *Size) native() *skSize {
	return (*skSize)(unsafe.Pointer(v))
}

type skIsize = C.sk_isize_t
type Isize struct {
	W int32
	H int32
}

func fromNativeIsize(v *skIsize) *Isize {
	return (*Isize)(unsafe.Pointer(v))
}
func (v *Isize) native() *skIsize {
	return (*skIsize)(unsafe.Pointer(v))
}

type skRsxform = C.sk_rsxform_t
type Rsxform struct {
	FSCos float32
	FSSin float32
	FTX   float32
	FTY   float32
}

func fromNativeRsxform(v *skRsxform) *Rsxform {
	return (*Rsxform)(unsafe.Pointer(v))
}
func (v *Rsxform) native() *skRsxform {
	return (*skRsxform)(unsafe.Pointer(v))
}
