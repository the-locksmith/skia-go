package skia

/*
#include "skia.h"
*/
import "C"
import (
	"unsafe"

	"github.com/zhuah/skia-go/internal"
)

type skColorspace = C.sk_colorspace_t
type SkColorspace = skColorspace

func (c *SkColorspace) Ref() {
	C.sk_colorspace_ref(c)
}

func (c *SkColorspace) Unref() {
	C.sk_colorspace_unref(c)
}

func NewSrgbColorspace() *SkColorspace {
	return C.sk_colorspace_new_srgb()

}

func NewSrgbLinearColorspace() *SkColorspace {
	return C.sk_colorspace_new_srgb_linear()

}

type skColorspaceXyz = C.sk_colorspace_xyz_t
type ColorspaceXyz struct {
	FM00 float32
	FM01 float32
	FM02 float32
	FM10 float32
	FM11 float32
	FM12 float32
	FM20 float32
	FM21 float32
	FM22 float32
}

func fromNativeColorspaceXyz(v *skColorspaceXyz) *ColorspaceXyz {
	return (*ColorspaceXyz)(unsafe.Pointer(v))
}
func (v *ColorspaceXyz) native() *skColorspaceXyz {
	return (*skColorspaceXyz)(unsafe.Pointer(v))
}

type skColorspaceTransferFn = C.sk_colorspace_transfer_fn_t
type ColorspaceTransferFn struct {
	FG float32
	FA float32
	FB float32
	FC float32
	FD float32
	FE float32
	FF float32
}

func fromNativeColorspaceTransferFn(v *skColorspaceTransferFn) *ColorspaceTransferFn {
	return (*ColorspaceTransferFn)(unsafe.Pointer(v))
}
func (v *ColorspaceTransferFn) native() *skColorspaceTransferFn {
	return (*skColorspaceTransferFn)(unsafe.Pointer(v))
}

type skColorspacePrimaries = C.sk_colorspace_primaries_t
type ColorspacePrimaries struct {
	FRX float32
	FRY float32
	FGX float32
	FGY float32
	FBX float32
	FBY float32
	FWX float32
	FWY float32
}

func fromNativeColorspacePrimaries(v *skColorspacePrimaries) *ColorspacePrimaries {
	return (*ColorspacePrimaries)(unsafe.Pointer(v))
}
func (v *ColorspacePrimaries) native() *skColorspacePrimaries {
	return (*skColorspacePrimaries)(unsafe.Pointer(v))
}

func NewRgbColorspace(transferFn *ColorspaceTransferFn, toXYZD50 *ColorspaceXyz) *SkColorspace {
	return C.sk_colorspace_new_rgb(transferFn.native(), toXYZD50.native())

}

func NewIccColorspace(profile *SkColorspaceIccProfile) *SkColorspace {
	return C.sk_colorspace_new_icc(profile)

}

func (c *SkColorspace) ToProfile(profile *SkColorspaceIccProfile) {
	C.sk_colorspace_to_profile(c, profile)
}

func (c *SkColorspace) GammaCloseToSrgb() bool {
	return bool(C.sk_colorspace_gamma_close_to_srgb(c))
}

func (c *SkColorspace) GammaIsLinear() bool {
	return bool(C.sk_colorspace_gamma_is_linear(c))
}

func (c *SkColorspace) IsNumericalTransferFn(transferFn *ColorspaceTransferFn) bool {
	return bool(C.sk_colorspace_is_numerical_transfer_fn(c, transferFn.native()))
}

func (c *SkColorspace) ToXyzd50(toXYZD50 *ColorspaceXyz) bool {
	return bool(C.sk_colorspace_to_xyzd50(c, toXYZD50.native()))
}

func (c *SkColorspace) MakeLinearGamma() *SkColorspace {
	return C.sk_colorspace_make_linear_gamma(c)
}

func (c *SkColorspace) MakeSrgbGamma() *SkColorspace {
	return C.sk_colorspace_make_srgb_gamma(c)
}

func (c *SkColorspace) IsSrgb() bool {
	return bool(C.sk_colorspace_is_srgb(c))
}

func (c *SkColorspace) Equals(src *SkColorspace, dst *SkColorspace) bool {
	return bool(C.sk_colorspace_equals(src, dst))
}

func (c *SkColorspace) TransferFnNamedSrgb(transferFn *ColorspaceTransferFn) {
	C.sk_colorspace_transfer_fn_named_srgb(transferFn.native())
}

func (c *SkColorspace) TransferFnNamed2dot2(transferFn *ColorspaceTransferFn) {
	C.sk_colorspace_transfer_fn_named_2dot2(transferFn.native())
}

func (c *SkColorspace) TransferFnNamedLinear(transferFn *ColorspaceTransferFn) {
	C.sk_colorspace_transfer_fn_named_linear(transferFn.native())
}

func (c *SkColorspace) TransferFnNamedRec2020(transferFn *ColorspaceTransferFn) {
	C.sk_colorspace_transfer_fn_named_rec2020(transferFn.native())
}

func (c *SkColorspace) TransferFnNamedPq(transferFn *ColorspaceTransferFn) {
	C.sk_colorspace_transfer_fn_named_pq(transferFn.native())
}

func (c *SkColorspace) TransferFnNamedHlg(transferFn *ColorspaceTransferFn) {
	C.sk_colorspace_transfer_fn_named_hlg(transferFn.native())
}

func (c *SkColorspace) TransferFnEval(transferFn *ColorspaceTransferFn, x float32) float32 {
	return float32(C.sk_colorspace_transfer_fn_eval(transferFn.native(), cfloat_t(x)))
}

func (c *SkColorspace) TransferFnInvert(src *ColorspaceTransferFn, dst *ColorspaceTransferFn) bool {
	return bool(C.sk_colorspace_transfer_fn_invert(src.native(), dst.native()))
}

func (c *SkColorspace) PrimariesToXyzd50(primaries *ColorspacePrimaries, toXYZD50 *ColorspaceXyz) bool {
	return bool(C.sk_colorspace_primaries_to_xyzd50(primaries.native(), toXYZD50.native()))
}

func (c *SkColorspace) XyzNamedSrgb(xyz *ColorspaceXyz) {
	C.sk_colorspace_xyz_named_srgb(xyz.native())
}

func (c *SkColorspace) XyzNamedAdobeRgb(xyz *ColorspaceXyz) {
	C.sk_colorspace_xyz_named_adobe_rgb(xyz.native())
}

func (c *SkColorspace) XyzNamedDcip3(xyz *ColorspaceXyz) {
	C.sk_colorspace_xyz_named_dcip3(xyz.native())
}

func (c *SkColorspace) XyzNamedRec2020(xyz *ColorspaceXyz) {
	C.sk_colorspace_xyz_named_rec2020(xyz.native())
}

func (c *SkColorspace) XyzNamedXyz(xyz *ColorspaceXyz) {
	C.sk_colorspace_xyz_named_xyz(xyz.native())
}

func (v *SkColorspace) XyzInvert(src *ColorspaceXyz, dst *ColorspaceXyz) bool {
	return bool(C.sk_colorspace_xyz_invert(src.native(), dst.native()))
}

func (v *SkColorspace) XyzConcat(a *ColorspaceXyz, b *ColorspaceXyz, result *ColorspaceXyz) {
	C.sk_colorspace_xyz_concat(a.native(), b.native(), result.native())
}

type skColorspaceIccProfile = C.sk_colorspace_icc_profile_t
type SkColorspaceIccProfile = skColorspaceIccProfile

func (c *SkColorspaceIccProfile) Destroy() {
	C.sk_colorspace_icc_profile_delete(c)
}

func NewColorspaceIccProfile() *SkColorspaceIccProfile {
	return C.sk_colorspace_icc_profile_new()
}

func (c *SkColorspaceIccProfile) Parse(buffer unsafe.Pointer, length uint, profile *SkColorspaceIccProfile) bool {
	return bool(C.sk_colorspace_icc_profile_parse(buffer, csize_t(length), profile))
}

func (c *SkColorspaceIccProfile) GetBuffer(profile *SkColorspaceIccProfile) []byte {
	var size cuint32_t
	buf := C.sk_colorspace_icc_profile_get_buffer(profile, &size)
	if size <= 0 {
		return nil
	}

	return internal.GoBytes(unsafe.Pointer(buf), int(size))
}

func (c *SkColorspaceIccProfile) GetToXyzd50(profile *SkColorspaceIccProfile, toXYZD50 *ColorspaceXyz) bool {
	return bool(C.sk_colorspace_icc_profile_get_to_xyzd50(profile, toXYZD50.native()))
}
