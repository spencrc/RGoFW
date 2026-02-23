package rgfw

/*
	#include "RFGW_impl.h"
*/
import "C"
import (
	"unsafe"
)

// Returns true if Wayland is currently being used
func UsingWayland() bool {
	return C.RGFW_usingWayland() == C.RGFW_TRUE
}

func GetLayerOSX() unsafe.Pointer {
	return unsafe.Pointer(C.RGFW_getLayer_OSX())
}

// Retrieves the current X11 display connection.
// Returns a pointer to the X11 display, or nil if the platform is not in use
func GetDisplayX11() unsafe.Pointer {
	return unsafe.Pointer(C.RGFW_getDisplay_X11())
}

// Retrieves the current Wayland display connection.
// Returns a pointer to the Wayland display, or nil if the platform is not in use
func GetDisplayWayland() unsafe.Pointer {
	return unsafe.Pointer(C.RGFW_getDisplay_Wayland())
}

func (win *Window) SetLayerOSX(layer unsafe.Pointer) {
    C.RGFW_window_setLayer_OSX(win.ref, layer)
}

func (win *Window) GetViewOSX() unsafe.Pointer {
    return unsafe.Pointer(C.RGFW_window_getView_OSX(win.ref))
}

// Retrieves the HWND handle for the window.
// Returns a pointer to the Windows HWND handle, or nil if not on Windows
func (win *Window) GetHWND() unsafe.Pointer {
	return unsafe.Pointer(C.RGFW_window_getHWND(win.ref))
}

func (win *Window) GetHDC() unsafe.Pointer {
    return unsafe.Pointer(C.RGFW_window_getHDC(win.ref))
}

// Retrieves the X11 Window handle for the window.
// Returns the X11 Window handle, or 0 if not on X11
func (win *Window) GetWindowX11() uint32 {
	return uint32(C.RGFW_window_getWindow_X11(win.ref))
}

func (win *Window) GetWindowWayland() unsafe.Pointer {
	return unsafe.Pointer(C.RGFW_window_getWindow_Wayland(win.ref))
}