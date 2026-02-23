package rgfw

/*
	#include "RFGW_impl.h"
*/
import "C"
import (
	"unsafe"
)

const WINDOW_CENTER = 0x40
const WINDOW_NO_RESIZE = 0x2

type Window struct {
	cPtr *C.RGFW_window
}

func CreateWindow(name string, x, y, w, h int, flags uint) *Window {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	winPtr := C.RGFW_createWindow(cName, C.i32(x), C.i32(y), C.i32(w), C.i32(h), C.uint(flags))

	return &Window{cPtr: winPtr}
}

func (win *Window) GetSize() (int, int) {
	var w, h C.i32
	C.RGFW_window_getSize(win.cPtr, &w, &h)
	return int(w), int(h)
}

func (win *Window) ShouldClose() bool {
	return C.RGFW_window_shouldClose(win.cPtr) == C.RGFW_TRUE
}

func (win *Window) Close() {
	C.RGFW_window_close(win.cPtr)
}