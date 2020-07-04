package skia

/*
#include "skia.h"
*/
import "C"

type skPicture = C.sk_picture_t
type SkPicture = skPicture

type skPictureRecorder = C.sk_picture_recorder_t
type SkPictureRecorder = skPictureRecorder

func NewPictureRecorder() *SkPictureRecorder {
	return C.sk_picture_recorder_new()
}

func (v *SkPictureRecorder) Destroy() {
	C.sk_picture_recorder_delete(v)
}

func (v *SkPictureRecorder) BeginRecording(rect *Rect) *SkCanvas {
	return C.sk_picture_recorder_begin_recording(v, rect.native())
}

func (v *SkPictureRecorder) EndRecording() *SkPicture {
	return C.sk_picture_recorder_end_recording(v)
}

func (v *SkPictureRecorder) EndRecordingAsDrawable() *SkDrawable {
	return C.sk_picture_recorder_end_recording_as_drawable(v)
}

func (v *SkPictureRecorder) GetRecordingCanvas() *SkCanvas {
	return C.sk_picture_get_recording_canvas(v)
}

func (v *SkPicture) Ref() {
	C.sk_picture_ref(v)
}

func (v *SkPicture) Unref() {
	C.sk_picture_unref(v)
}

func (v *SkPicture) GetUniqueId() uint32 {
	return uint32(C.sk_picture_get_unique_id(v))
}

func (v *SkPicture) GetCullRect() Rect {
	var rec Rect
	C.sk_picture_get_cull_rect(v, rec.native())
	return rec
}

func (v *SkPicture) MakeShader(tmx SkShaderTilemode, tmy SkShaderTilemode, localMatrix *Matrix, tile *Rect) *SkShader {
	return C.sk_picture_make_shader(v, tmx, tmy, localMatrix.native(), tile.native())
}
