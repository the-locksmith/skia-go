package skia

/*
#include "skia.h"
*/
import "C"

type sk3dview = C.sk_3dview_t
type Sk3dview = sk3dview

func New3dview() *Sk3dview {
	return C.sk_3dview_new()
}

func (v *Sk3dview) Destroy() {
	C.sk_3dview_destroy(v)
}

func (v *Sk3dview) Save() {
	C.sk_3dview_save(v)
}

func (v *Sk3dview) Restore() {
	C.sk_3dview_restore(v)
}

func (v *Sk3dview) Translate(x, y, z float32) {
	C.sk_3dview_translate(v, cfloat_t(x), cfloat_t(y), cfloat_t(z))
}

func (v *Sk3dview) RotateXDegrees(degrees float32) {
	C.sk_3dview_rotate_x_degrees(v, cfloat_t(degrees))
}

func (v *Sk3dview) RotateYDegrees(degrees float32) {
	C.sk_3dview_rotate_y_degrees(v, cfloat_t(degrees))
}

func (v *Sk3dview) RotateZDegrees(degrees float32) {
	C.sk_3dview_rotate_z_degrees(v, cfloat_t(degrees))
}

func (v *Sk3dview) RotateXRadians(radians float32) {
	C.sk_3dview_rotate_x_radians(v, cfloat_t(radians))
}

func (v *Sk3dview) RotateYRadians(radians float32) {
	C.sk_3dview_rotate_y_radians(v, cfloat_t(radians))
}

func (v *Sk3dview) RotateZRadians(radians float32) {
	C.sk_3dview_rotate_z_radians(v, cfloat_t(radians))
}

func (v *Sk3dview) GetMatrix(matrix *Matrix) {
	C.sk_3dview_get_matrix(v, matrix.native())
}

func (v *Sk3dview) ApplyToCanvas(canvas *SkCanvas) {
	C.sk_3dview_apply_to_canvas(v, canvas)
}

func (v *Sk3dview) DotWithNormal(dx, dy, dz float32) float32 {
	return float32(C.sk_3dview_dot_with_normal(v, cfloat_t(dx), cfloat_t(dy), cfloat_t(dz)))
}
