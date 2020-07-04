package skia

/*
#include "skia.h"
*/
import "C"

type skDrawable = C.sk_drawable_t
type SkDrawable = skDrawable

func (v *SkDrawable) Unref() {
	C.sk_drawable_unref(v)
}

func (v *SkDrawable) GetGenerationId() uint32 {
	return uint32(C.sk_drawable_get_generation_id(v))
}

func (v *SkDrawable) GetBounds() Rect {
	var rect Rect
	C.sk_drawable_get_bounds(v, rect.native())
	return rect
}

func (v *SkDrawable) Draw(cvs *SkCanvas, matrix *Matrix) {
	C.sk_drawable_draw(v, cvs, matrix.native())
}

func (v *SkDrawable) NewPictureSnapshot() *SkPicture {
	return C.sk_drawable_new_picture_snapshot(v)
}

func (v *SkDrawable) NotifyDrawingChanged() {
	C.sk_drawable_notify_drawing_changed(v)
}
