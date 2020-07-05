package skia

/*
#include "skia.h"
*/
import "C"
import (
	"unsafe"

	"github.com/uiez/skia-go/internal"
)

type skCanvas = C.sk_canvas_t
type SkCanvas = skCanvas

func NewCanvasFromBitmap(bitmap *SkBitmap) *SkCanvas {
	return C.sk_canvas_new_from_bitmap(bitmap)
}

func (v *SkCanvas) Destroy() {
	C.sk_canvas_destroy(v)
}

func (v *SkCanvas) Save() int32 {
	return int32(C.sk_canvas_save(v))
}

func (v *SkCanvas) SaveLayer(bounds *Rect, paint *SkPaint) int32 {
	return int32(C.sk_canvas_save_layer(v, bounds.native(), paint))
}

type skSaveLayerFlags = C.sk_save_layer_flags_t
type SkSaveLayerFlags = skSaveLayerFlags

const (
	SK_CANVAS_SAVE_LAYER_FLAGS_INIT_WITH_PREVIOUS = C.INIT_WITH_PREVIOUS_SK_CANVAS_SAVE_LAYER_FLAGS
)

type skSaveLayerRec = C.sk_save_layer_rec_t
type SaveLayerRec struct {
	Bounds   Rect
	Paint    *SkPaint
	Backdrop *SkImagefilter
	Flags    SkSaveLayerFlags
}

func fromNativeSaveLayerRec(v *skSaveLayerRec) *SaveLayerRec {
	return (*SaveLayerRec)(unsafe.Pointer(v))
}
func (v *SaveLayerRec) native() *skSaveLayerRec {
	return (*skSaveLayerRec)(unsafe.Pointer(v))
}

func (v *SkCanvas) SaveLayerRec(rec *SaveLayerRec) int32 {
	return int32(C.sk_canvas_save_layer_rec(v, rec.native()))
}

func (v *SkCanvas) Restore() {
	C.sk_canvas_restore(v)
}

func (v *SkCanvas) Translate(Dx float32, Dy float32) {
	cDx := (C.float)(Dx)
	cDy := (C.float)(Dy)
	C.sk_canvas_translate(v, cDx, cDy)
}

func (v *SkCanvas) Scale(Sx float32, Sy float32) {
	cSx := (C.float)(Sx)
	cSy := (C.float)(Sy)
	C.sk_canvas_scale(v, cSx, cSy)
}

func (v *SkCanvas) RotateDegrees(Degrees float32) {
	cDegrees := (C.float)(Degrees)
	C.sk_canvas_rotate_degrees(v, cDegrees)
}

func (v *SkCanvas) RotateRadians(Radians float32) {
	cRadians := (C.float)(Radians)
	C.sk_canvas_rotate_radians(v, cRadians)
}

func (v *SkCanvas) Skew(Sx float32, Sy float32) {
	cSx := (C.float)(Sx)
	cSy := (C.float)(Sy)
	C.sk_canvas_skew(v, cSx, cSy)
}

func (v *SkCanvas) Concat(matrix *Matrix) {
	C.sk_canvas_concat(v, matrix.native())
}

func (v *SkCanvas) QuickReject(rect *Rect) bool {
	return bool(C.sk_canvas_quick_reject(v, rect.native()))
}

func (v *SkCanvas) ClipRegion(region *SkRegion, op SkClipop) {
	C.sk_canvas_clip_region(v, region, op)
}

func (v *SkCanvas) DrawPaint(paint *SkPaint) {
	C.sk_canvas_draw_paint(v, paint)
}

func (v *SkCanvas) DrawRect(rect *Rect, paint *SkPaint) {
	C.sk_canvas_draw_rect(v, rect.native(), paint)
}

func (v *SkCanvas) DrawRrect(rrect *SkRrect, paint *SkPaint) {
	C.sk_canvas_draw_rrect(v, rrect, paint)
}

func (v *SkCanvas) DrawRegion(region *SkRegion, paint *SkPaint) {
	C.sk_canvas_draw_region(v, region, paint)
}

func (v *SkCanvas) DrawCircle(cx, cy, rad float32, paint *SkPaint) {
	C.sk_canvas_draw_circle(v, cfloat_t(cx), cfloat_t(cy), cfloat_t(rad), paint)
}

func (v *SkCanvas) DrawOval(rect *Rect, paint *SkPaint) {
	C.sk_canvas_draw_oval(v, rect.native(), paint)
}

func (v *SkCanvas) DrawPath(path *SkPath, paint *SkPaint) {
	C.sk_canvas_draw_path(v, path, paint)
}

func (v *SkCanvas) DrawImage(img *SkImage, x, y float32, paint *SkPaint) {
	C.sk_canvas_draw_image(v, img, cfloat_t(x), cfloat_t(y), paint)
}

func (v *SkCanvas) DrawImageRect(img *SkImage, src, dst *Rect, paint *SkPaint) {
	C.sk_canvas_draw_image_rect(v, img, src.native(), dst.native(), paint)
}

func (v *SkCanvas) DrawPicture(pic *SkPicture, matrix *Matrix, paint *SkPaint) {
	C.sk_canvas_draw_picture(v, pic, matrix.native(), paint)
}

func (v *SkCanvas) DrawDrawable(d *SkDrawable, matrix *Matrix) {
	C.sk_canvas_draw_drawable(v, d, matrix.native())
}

func (v *SkCanvas) Clear(c SkColor) {
	C.sk_canvas_clear(v, c)
}

func (v *SkCanvas) Discard() {
	C.sk_canvas_discard(v)
}

func (v *SkCanvas) GetSaveCount() int32 {
	return int32(C.sk_canvas_get_save_count(v))
}

func (v *SkCanvas) RestoreToCount(SaveCount int32) {
	cSaveCount := (C.int)(SaveCount)
	C.sk_canvas_restore_to_count(v, cSaveCount)
}

func (v *SkCanvas) DrawColor(color SkColor, mode SkBlendmode) {
	C.sk_canvas_draw_color(v, color, mode)
}

type skPointMode = C.sk_point_mode_t
type SkPointMode = skPointMode

const (
	SK_POINT_MODE_POINTS  = C.POINTS_SK_POINT_MODE
	SK_POINT_MODE_LINES   = C.LINES_SK_POINT_MODE
	SK_POINT_MODE_POLYGON = C.POLYGON_SK_POINT_MODE
)

func (v *SkCanvas) DrawPoints(mode SkPointMode, points []Point, paint *SkPaint) {
	C.sk_canvas_draw_points(v, mode, csize_t(len(points)), (*skPoint)(internal.SliceDataPtr(points)), paint)
}

func (v *SkCanvas) DrawPoint(x, y float32, paint *SkPaint) {
	C.sk_canvas_draw_point(v, cfloat_t(x), cfloat_t(y), paint)
}

func (v *SkCanvas) DrawLine(x0 float32, y0 float32, x1 float32, y1 float32, paint *SkPaint) {
	C.sk_canvas_draw_line(v, cfloat_t(x0), cfloat_t(y0), cfloat_t(x1), cfloat_t(y1), paint)
}

func (v *SkCanvas) DrawTextBlob(text *SkTextblob, x float32, y float32, paint *SkPaint) {
	C.sk_canvas_draw_text_blob(v, text, cfloat_t(x), cfloat_t(y), paint)
}

func (v *SkCanvas) DrawBitmap(bitmap *SkBitmap, Left float32, Top float32, paint *SkPaint) {
	cLeft := (C.float)(Left)
	cTop := (C.float)(Top)
	C.sk_canvas_draw_bitmap(v, bitmap, cLeft, cTop, paint)
}

func (v *SkCanvas) DrawBitmapRect(bitmap *SkBitmap, src *Rect, dst *Rect, paint *SkPaint) {
	C.sk_canvas_draw_bitmap_rect(v, bitmap, src.native(), dst.native(), paint)
}

func (v *SkCanvas) ResetMatrix() {
	C.sk_canvas_reset_matrix(v)
}

func (v *SkCanvas) SetMatrix(matrix *Matrix) {
	C.sk_canvas_set_matrix(v, matrix.native())
}

func (v *SkCanvas) GetTotalMatrix(matrix *Matrix) {
	C.sk_canvas_get_total_matrix(v, matrix.native())
}

func (v *SkCanvas) DrawRoundRect(rect *Rect, rx float32, ry float32, paint *SkPaint) {
	C.sk_canvas_draw_round_rect(v, rect.native(), cfloat_t(rx), cfloat_t(ry), paint)
}

type skClipop = C.sk_clipop_t
type SkClipop = skClipop

const (
	SK_CLIPOP_DIFFERENCE = C.DIFFERENCE_SK_CLIPOP
	SK_CLIPOP_INTERSECT  = C.INTERSECT_SK_CLIPOP
)

func (v *SkCanvas) ClipRectWithOperation(rect *Rect, op SkClipop, doAA bool) {
	C.sk_canvas_clip_rect_with_operation(v, rect.native(), op, cbool_t(doAA))
}

func (v *SkCanvas) ClipPathWithOperation(path *SkPath, op SkClipop, doAA bool) {
	C.sk_canvas_clip_path_with_operation(v, path, op, cbool_t(doAA))
}

func (v *SkCanvas) ClipRrectWithOperation(rrect *SkRrect, op SkClipop, doAA bool) {
	C.sk_canvas_clip_rrect_with_operation(v, rrect, op, cbool_t(doAA))
}

func (v *SkCanvas) GetLocalClipBounds(bounds *Rect) bool {
	return bool(C.sk_canvas_get_local_clip_bounds(v, bounds.native()))
}

func (v *SkCanvas) GetDeviceClipBounds(bounds *Irect) bool {
	return bool(C.sk_canvas_get_device_clip_bounds(v, bounds.native()))
}

func (v *SkCanvas) Flush() {
	C.sk_canvas_flush(v)
}

func (v *SkCanvas) DrawAnnotation(rect *Rect, key string, value *SkData) {
	ckey, free := cstring(key)
	C.sk_canvas_draw_annotation(v, rect.native(), ckey, value)
	free()
}

func (v *SkCanvas) DrawUrlAnnotation(rect *Rect, value *SkData) {
	C.sk_canvas_draw_url_annotation(v, rect.native(), value)
}

func (v *SkCanvas) DrawNamedDestinationAnnotation(point *Point, value *SkData) {
	C.sk_canvas_draw_named_destination_annotation(v, point.native(), value)
}

func (v *SkCanvas) DrawLinkDestinationAnnotation(rect *Rect, value *SkData) {
	C.sk_canvas_draw_link_destination_annotation(v, rect.native(), value)
}

func (v *SkCanvas) DrawImageNine(img *SkImage, center *Irect, dst *Rect, paint *SkPaint) {
	C.sk_canvas_draw_image_nine(v, img, center.native(), dst.native(), paint)
}

func (v *SkCanvas) DrawVertices(vertices *SkVertices, mode SkBlendmode, paint *SkPaint) {
	C.sk_canvas_draw_vertices(v, vertices, mode, paint)
}

func (v *SkCanvas) DrawArc(oval *Rect, startAngle float32, sweepAngle float32, useCenter bool, paint *SkPaint) {
	C.sk_canvas_draw_arc(v, oval.native(), cfloat_t(startAngle), cfloat_t(sweepAngle), cbool_t(useCenter), paint)
}

func (v *SkCanvas) DrawDrrect(outer, inner *SkRrect, paint *SkPaint) {
	C.sk_canvas_draw_drrect(v, outer, inner, paint)
}

func (v *SkCanvas) DrawAtlas(atlas *SkImage, xform *Rsxform, tex *Rect, colors []SkColor, mode SkBlendmode, cullRect *Rect, paint *SkPaint) {
	C.sk_canvas_draw_atlas(v, atlas, xform.native(), tex.native(), (*SkColor)(internal.SliceDataPtr(colors)), cint_t(len(colors)), mode, cullRect.native(), paint)
}

//
//func (v *SkCanvas) DrawPatch(cubics *Point, color *SkColor, TexCoords *SkPoint, mode SkBlendmode, paint *SkPaint) {
//	C.sk_canvas_draw_patch(v, cCubics, cColors, cTexCoords, cMode, paint)
//}

func (v *SkCanvas) IsClipEmpty() bool {
	return bool(C.sk_canvas_is_clip_empty(v))
}

func (v *SkCanvas) IsClipRect() bool {
	return bool(C.sk_canvas_is_clip_rect(v))
}

type skNodrawCanvas = C.sk_nodraw_canvas_t
type SkNodrawCanvas = skNodrawCanvas

func NewNodrawCanvas(width, height int) *SkNodrawCanvas {
	return C.sk_nodraw_canvas_new(cint_t(width), cint_t(height))
}

func (v *SkNodrawCanvas) Destroy() {
	C.sk_nodraw_canvas_destroy(v)
}

type skNwayCanvas = C.sk_nway_canvas_t
type SkNwayCanvas = skNwayCanvas

func NewNwayCanvas(width, height int) *SkNwayCanvas {
	return C.sk_nway_canvas_new(cint_t(width), cint_t(height))
}

func (v *SkNwayCanvas) Destroy() {
	C.sk_nway_canvas_destroy(v)
}

func (v *SkNwayCanvas) AddCanvas(canvas *SkCanvas) {
	C.sk_nway_canvas_add_canvas(v, canvas)
}

func (v *SkNwayCanvas) RemoveCanvas(canvas *SkCanvas) {
	C.sk_nway_canvas_remove_canvas(v, canvas)
}

func (v *SkNwayCanvas) RemoveAll() {
	C.sk_nway_canvas_remove_all(v)
}

type skOverdrawCanvas = C.sk_overdraw_canvas_t
type SkOverdrawCanvas = skOverdrawCanvas

func NewOverdrawCanvas(canvas *SkCanvas) *SkOverdrawCanvas {
	return C.sk_overdraw_canvas_new(canvas)
}

func (v *SkOverdrawCanvas) Destroy() {
	C.sk_overdraw_canvas_destroy(v)
}
