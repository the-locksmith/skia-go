package skia

/*
#include "skia.h"
*/
import "C"
import (
	"unsafe"
)

type skMatrix44 = C.sk_matrix44_t
type SkMatrix44 = skMatrix44

type skMatrix44TypeMask = C.sk_matrix44_type_mask_t
type SkMatrix44TypeMask = skMatrix44TypeMask

const (
	SK_MATRIX44_TYPE_MASK_IDENTITY    = C.IDENTITY_SK_MATRIX44_TYPE_MASK
	SK_MATRIX44_TYPE_MASK_TRANSLATE   = C.TRANSLATE_SK_MATRIX44_TYPE_MASK
	SK_MATRIX44_TYPE_MASK_SCALE       = C.SCALE_SK_MATRIX44_TYPE_MASK
	SK_MATRIX44_TYPE_MASK_AFFINE      = C.AFFINE_SK_MATRIX44_TYPE_MASK
	SK_MATRIX44_TYPE_MASK_PERSPECTIVE = C.PERSPECTIVE_SK_MATRIX44_TYPE_MASK
)

func NewMatrix44() *SkMatrix44 {
	return C.sk_matrix44_new()
}

func NewMatrix44Identity() *SkMatrix44 {
	return C.sk_matrix44_new_identity()
}

func NewMatrix44Copy(src *SkMatrix44) *SkMatrix44 {
	return C.sk_matrix44_new_copy(src)
}

func NewMatrix44Concat(a *SkMatrix44, b *SkMatrix44) *SkMatrix44 {
	return C.sk_matrix44_new_concat(a, b)
}

func NewMatrix44FromMatrix3(src *Matrix) *SkMatrix44 {
	return C.sk_matrix44_new_matrix(src.native())
}

func (v *SkMatrix44) Destroy() {
	C.sk_matrix44_destroy(v)
}

func (v *SkMatrix44) Equals(other *SkMatrix44) bool {
	return bool(C.sk_matrix44_equals(v, other))
}

func (v *SkMatrix44) ToMatrix() *Matrix {
	var m Matrix
	C.sk_matrix44_to_matrix(v, m.native())
	return &m
}

func (v *SkMatrix44) GetType() SkMatrix44TypeMask {
	return C.sk_matrix44_get_type(v)
}

func (v *SkMatrix44) SetIdentity() {
	C.sk_matrix44_set_identity(v)
}

func (v *SkMatrix44) Get(row int32, col int32) float32 {
	return float32(C.sk_matrix44_get(v, cint32_t(row), cint32_t(col)))
}

func (v *SkMatrix44) Set(row int32, col int32, value float32) {
	C.sk_matrix44_set(v, cint32_t(row), cint32_t(col), cfloat_t(value))
}

func (v *SkMatrix44) SetTranslate(dx float32, dy float32, dz float32) {
	C.sk_matrix44_set_translate(v, cfloat_t(dx), cfloat_t(dy), cfloat_t(dz))
}

func (v *SkMatrix44) PreTranslate(dx float32, dy float32, dz float32) {
	C.sk_matrix44_pre_translate(v, cfloat_t(dx), cfloat_t(dy), cfloat_t(dz))
}

func (v *SkMatrix44) PostTranslate(dx float32, dy float32, dz float32) {
	C.sk_matrix44_post_translate(v, cfloat_t(dx), cfloat_t(dy), cfloat_t(dz))
}

func (v *SkMatrix44) SetScale(sx float32, sy float32, sz float32) {
	C.sk_matrix44_set_scale(v, cfloat_t(sx), cfloat_t(sy), cfloat_t(sz))
}

func (v *SkMatrix44) PreScale(sx float32, sy float32, sz float32) {
	C.sk_matrix44_pre_scale(v, cfloat_t(sx), cfloat_t(sy), cfloat_t(sz))
}

func (v *SkMatrix44) PostScale(sx float32, sy float32, sz float32) {
	C.sk_matrix44_post_scale(v, cfloat_t(sx), cfloat_t(sy), cfloat_t(sz))
}

func (v *SkMatrix44) SetRotateAboutDegrees(x float32, y float32, z float32, degrees float32) {
	C.sk_matrix44_set_rotate_about_degrees(v, cfloat_t(x), cfloat_t(y), cfloat_t(z), cfloat_t(degrees))
}

func (v *SkMatrix44) SetRotateAboutRadians(x float32, y float32, z float32, radians float32) {
	C.sk_matrix44_set_rotate_about_radians(v, cfloat_t(x), cfloat_t(y), cfloat_t(z), cfloat_t(radians))
}

func (v *SkMatrix44) SetRotateAboutRadiansUnit(x float32, y float32, z float32, radians float32) {
	C.sk_matrix44_set_rotate_about_radians_unit(v, cfloat_t(x), cfloat_t(y), cfloat_t(z), cfloat_t(radians))
}

func (v *SkMatrix44) SetConcat(a *SkMatrix44, b *SkMatrix44) {
	C.sk_matrix44_set_concat(v, a, b)
}

func (v *SkMatrix44) PreConcat(m *SkMatrix44) {
	C.sk_matrix44_pre_concat(v, m)
}

func (v *SkMatrix44) PostConcat(m *SkMatrix44) {
	C.sk_matrix44_post_concat(v, m)
}
func (v *SkMatrix44) CanInvert() bool {
	return bool(C.sk_matrix44_invert(v, nil))
}
func (v *SkMatrix44) TryInvert(inverse *SkMatrix44) bool {
	return bool(C.sk_matrix44_invert(v, inverse))
}

func (v *SkMatrix44) Transpose() {
	C.sk_matrix44_transpose(v)
}

func (v *SkMatrix44) MapScalars(src *[4]float32) {
	ptr := (*cfloat_t)(unsafe.Pointer(src))
	C.sk_matrix44_map_scalars(v, ptr, ptr)
}

func (v *SkMatrix44) Preserves2dAxisAlignment(epsilon float32) bool {
	return bool(C.sk_matrix44_preserves_2d_axis_alignment(v, cfloat_t(epsilon)))
}

func (v *SkMatrix44) Determinant() float64 {
	return float64(C.sk_matrix44_determinant(v))
}
