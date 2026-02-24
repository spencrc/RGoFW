package rgfw

/*
	#include "RFGW_impl.h"
*/
import "C"

type WindowMoveCallback func(win *C.RGFW_window, x, y C.i32)

type WindowResizedCallback func(win *C.RGFW_window, w, h C.i32)

type WindowRestoredCallback func(win *C.RGFW_window, x, y, w, h C.i32)

type WindowMaximizedCallback func(win *C.RGFW_window, x, y, w, h C.i32)

type WindowMinimizedCallback func(win *C.RGFW_window)

type WindowQuitCallback func(win *C.RGFW_window)

type FocusCallback func(win *C.RGFW_window, inFocus C.RGFW_bool)

type MouseNotifyCallback func(win *C.RGFW_window, x, y C.i32, status C.RGFW_bool)

type MousePosCallback func(win *C.RGFW_window, x, y C.i32, vecX, vecY C.float)

type DataDragCallback func(win *C.RGFW_window, x, y C.i32)

type WindowRefreshCallback func(win *C.RGFW_window)

type KeyCallback func(win *C.RGFW_window, key, sym byte, mod C.RGFW_keymod, repeat, pressed C.RGFW_bool)

type MouseButtonCallback func(win *C.RGFW_window, button C.RGFW_mouseButton, pressed C.RGFW_bool)

type MouseScrollCallback func(win *C.RGFW_window, x, y C.float)

type DataDropCallback func(win *C.RGFW_window, files *C.char)

type ScaleUpdatedCallback func(win *C.RGFW_window, scaleX, scaleY C.float)