package rgfw

/*
	#include "RFGW_impl.h"
*/
import "C"
import (
	"unsafe"
)

// Retrieves the X11 Window handle for the window
// Returns the X11 Window handle, or 0 if not on X11
func (win *Window) GetWindowX11() uint32 {
	return uint32(C.RGFW_window_getWindow_X11(win.cPtr))
}

// Retrieves the HWND handle for the window
// Returns a pointer to the Windows HWND handle, or nil if not on Windows
func (win *Window) GetHWND() unsafe.Pointer {
	return unsafe.Pointer(C.RGFW_window_getHWND(win.cPtr))
}

func (win *Window) GetWindowWayland() unsafe.Pointer {
	return unsafe.Pointer(C.RGFW_window_getWindow_Wayland(win.cPtr))
}

// Retrieves the current X11 display connection.
// Returns a pointer to the X11 display, or nil if the platform is not in use.
func GetDisplayX11() unsafe.Pointer {
	return unsafe.Pointer(C.RGFW_getDisplay_X11())
}

// Retrieves the current Wayland display connection.
// Returns a pointer to the Wayland display, or nil if the platform is not in use.
func GetDisplayWayland() unsafe.Pointer {
	return unsafe.Pointer(C.RGFW_getDisplay_Wayland())
}