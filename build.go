package skia

/*
#cgo darwin,!ios LDFLAGS: -lskia -framework Cocoa -framework Foundation -framework CoreGraphics
#cgo linux,!android LDFLAGS: -lskia -lfreetype -lfontconfig -lm -ldl
#cgo windows LDFLAGS: -lskia -luser32
#cgo linux,android LDFLAGS: -lskia -lm -ldl -llog -lz
#cgo CFLAGS: -Iinternal
*/
/*
#include "skia.h"
*/
import "C"
