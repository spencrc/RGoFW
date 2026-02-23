package rgfw

/*
	#include "RFGW_impl.h"
*/
import "C"
import (
	"unsafe"
)

type WindowFlag uint

const (
    WindowFlagNoBorder          WindowFlag = C.RGFW_windowNoBorder
    WindowFlagNoResize          WindowFlag = C.RGFW_windowNoResize
    WindowFlagAllowDND          WindowFlag = C.RGFW_windowAllowDND
    WindowFlagHideMouse         WindowFlag = C.RGFW_windowHideMouse
    WindowFlagFullscreen        WindowFlag = C.RGFW_windowFullscreen
    WindowFlagTransparent       WindowFlag = C.RGFW_windowTransparent
    WindowFlagCenter            WindowFlag = C.RGFW_windowCenter
    // WindowFlagRawMouse          WindowFlag = C.RGFW_windowRawMouse
    WindowFlagScaleToMonitor    WindowFlag = C.RGFW_windowScaleToMonitor
    WindowFlagHide              WindowFlag = C.RGFW_windowHide
    WindowFlagMaximize          WindowFlag = C.RGFW_windowMaximize
    WindowFlagCenterCursor      WindowFlag = C.RGFW_windowCenterCursor
    WindowFlagFloating          WindowFlag = C.RGFW_windowFloating
    WindowFlagFocusOnShow       WindowFlag = C.RGFW_windowFocusOnShow
    WindowFlagMinimize          WindowFlag = C.RGFW_windowMinimize
    WindowFlagFocus             WindowFlag = C.RGFW_windowFocus
    // WindowFlagCaptureMouse      WindowFlag = C.RGFW_windowCaptureMouse
    WindowFlagOpenGL            WindowFlag = C.RGFW_windowOpenGL
    WindowFlagEGL               WindowFlag = C.RGFW_windowEGL
    // WindowFlagNoDeinitOnClose   WindowFlag = C.RGFW_noDeinitOnClose
    WindowFlagWindowedFullscreen   WindowFlag = C.RGFW_windowedFullscreen
    // WindowFlagCaptureRawMouse   WindowFlag = C.RGFW_windowCaptureRawMouse
)

// type FlashRequest uint8

// const (
// 	FlashCancel FlashRequest = C.RGFW_flashCancel
// 	FlashBriefly FlashRequest = C.RGFW_flashBriefly
// 	FlashUntilFocused FlashRequest = C.RGFW_flashUntilFocused
// )

type Window struct {
	ref *C.RGFW_window
}

// Creates a new window.
func CreateWindow(name string, x, y, w, h int, flags WindowFlag) *Window {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	winPtr := C.RGFW_createWindow(cName, C.i32(x), C.i32(y), C.i32(w), C.i32(h), C.uint(flags))

	return &Window{ref: winPtr}
}

// INTENTIONALLY UNIMPLEMENTED. WILL NOT ADD
// Creates a new window using a pre-allocated window structure.
// RGFW_createWindowPtr

// Creates a new surface structure
// RGFW_window_createSurface

// INTENTIONALLY UNIMPLEMENTED. WILL NOT ADD
// Creates a new surface structure using a pre-allocated surface structure
// RGFW_window_createSurfacePtr

// Set the function/callback used for converting surface data between formats
// RGFW_surface_setConvertFunc

// Blits a surface stucture to the window
// RGFW_window_blitSurface

// Gets the position of the window
func (win *Window) GetPosition() (int, int) {
    var x, y C.i32
	C.RGFW_window_getPosition(win.ref, &x, &y)
	return int(x), int(y)
}

// Gets the size of the window
func (win *Window) GetSize() (int, int) {
	var w, h C.i32
	C.RGFW_window_getSize(win.ref, &w, &h)
	return int(w), int(h)
}

// Gets the size of the window in exact pixels
// func (win *Window) GetSizeInPixels() (int, int) {
// 	var w, h C.i32
// 	C.RGFW_window_getSizeInPixels(win.ref, &w, &h)
// 	return int(w), int(h)
// }

// Gets the flags of the window
func (win *Window) GetFlags() WindowFlag {
    return WindowFlag(C.RGFW_window_getFlags(win.ref))
}

// Returns the exit key assigned to the window
func (win *Window) GetExitKey() Key {
    return Key(C.RGFW_window_getExitKey(win.ref))
}

// Sets the exit key for the window
func (win *Window) SetExitKey(key Key) {
    C.RGFW_window_setExitKey(win.ref, C.RGFW_key(key))
}

// Sets the types of events you want the window to receive
func (win *Window) SetEnabledEvents(events EventFlag) {
    C.RGFW_window_setEnabledEvents(win.ref, C.RGFW_eventFlag(events))
}

// Gets the currently enabled events for the window
func (win *Window) GetEnabledEvents() EventFlag {
    return EventFlag(C.RGFW_window_getEnabledEvents(win.ref))
}

// Enables all events and disables selected ones
func (win *Window) SetDisabledEvents(events EventFlag) {
    C.RGFW_window_setDisabledEvents(win.ref, C.RGFW_eventFlag(events))
}

// Directly enables or disables a specific event or group of events
func (win *Window) SetEventState(event EventFlag, state bool) {
    C.RGFW_window_setEventState(win.ref, C.RGFW_eventFlag(event), C.RGFW_bool(boolToInt(state)))
}

// Gets the user pointer associated with the window
// func (win *Window) GetUserPtr() unsafe.Pointer {
//     return C.RGFW_window_getUserPtr(win.ref)
// }

// Sets a user pointer for the window
// func (win *Window) SetUserPtr(ptr unsafe.Pointer) {
//     C.RGFW_window_setUserPtr(win.ref, ptr)
// }

// INTENTIONALLY UNIMPLEMENTED. WILL NOT ADD
// Retrieves the platform-specific window source pointer
// RGFW_window_getSrc

// Set the window flags (will undo flags if they don't match the old ones)
func (win *Window) SetFlags(flags WindowFlag) {
    C.RGFW_window_setFlags(win.ref, C.RGFW_windowFlags(flags))
}

// RGFW_window_checkEvent can be found in event.go!

// RGFW_window_checkQueuedEvent can be found in event.go!

// Checks if a key was pressed while the window is in focus.
// key is the key code to check
func (win *Window) IsKeyPressed(key Key) bool {
    return C.RGFW_window_isKeyPressed(win.ref, C.RGFW_key(key)) == C.RGFW_TRUE
}

// Checks if a key is currently being held down.
// key is the key code to check
func (win *Window) IsKeyDown(key Key) bool {
    return C.RGFW_window_isKeyDown(win.ref, C.RGFW_key(key)) == C.RGFW_TRUE
}

// Checks if a key was released.
// key is the key code to check
func (win *Window) IsKeyReleased(key Key) bool {
    return C.RGFW_window_isKeyReleased(win.ref, C.RGFW_key(key)) == C.RGFW_TRUE
}

// Checks if a mouse button was pressed.
// button is the mouse button code to check
func (win *Window) IsMousePressed(button MouseButton) bool {
    return C.RGFW_window_isMousePressed(win.ref, C.RGFW_mouseButton(button)) == C.RGFW_TRUE
}

// Checks if a mouse button is currently held down.
// button is the mouse button code to check
func (win *Window) IsMouseDown(button MouseButton) bool {
    return C.RGFW_window_isMouseDown(win.ref, C.RGFW_mouseButton(button)) == C.RGFW_TRUE
}

// Checks if a mouse button was released.
// button is the mouse button code to check
func (win *Window) IsMouseReleased(button MouseButton) bool {
    return C.RGFW_window_isMouseReleased(win.ref, C.RGFW_mouseButton(button)) == C.RGFW_TRUE
}

// Checks if the mouse left the window (true only for the first frame)
func (win *Window) DidMouseLeave() bool {
    return C.RGFW_window_didMouseLeave(win.ref) == C.RGFW_TRUE
}

// Checks if the mouse entered the window (true only for the first frame)
func (win *Window) DidMouseEnter() bool {
    return C.RGFW_window_didMouseEnter(win.ref) == C.RGFW_TRUE
}

// Checks if the mouse is currently inside the window bounds
func (win *Window) IsMouseInside() bool {
    return C.RGFW_window_isMouseInside(win.ref) == C.RGFW_TRUE
}

// Checks if there is data being dragged into or within the window
func (win *Window) IsDataDragging() bool {
    return C.RGFW_window_isDataDragging(win.ref) == C.RGFW_TRUE
}

// Gets the position of a data drag
// RGFW_window_getDataDrag

// Checks if a data drop occurred in the window (first frame only)
// RGFW_window_didDataDrop

// Retrieves files from a data drop (drag and drop)
// RGFW_window_getDataDrop

// Closes the window and frees its associated structure
func (win *Window) Close() {
	C.RGFW_window_close(win.ref)
}

// INTENTIONALLY UNIMPLEMENTED. WILL NOT ADD
// Closes the window without freeing its structure
// RGFW_window_closePtr

// Moves the window to a new position on the screen
func (win *Window) Move(x, y int) {
    C.RGFW_window_move(win.ref, C.i32(x), C.i32(y))
}

// Moves the window to a specific monitor
// RGFW_window_moveToMonitor

// Resizes the window to the given dimensions
func (win *Window) Resize(w, h int) {
    C.RGFW_window_resize(win.ref, C.i32(w), C.i32(h))
}

// Sets the aspect ratio of the window
func (win *Window) SetAspectRatio(w, h int) {
    C.RGFW_window_setAspectRatio(win.ref, C.i32(w), C.i32(h))
}

// Sets the minimum size of the window
func (win *Window) SetMinSize(w, h int) {
    C.RGFW_window_setMinSize(win.ref, C.i32(w), C.i32(h))
}

// Sets the maximum size of the window
func (win *Window) SetMaxSize(w, h int) {
    C.RGFW_window_setMaxSize(win.ref, C.i32(w), C.i32(h))
}

// Sets focus to the window
func (win *Window) Focus() {
    C.RGFW_window_focus(win.ref)
}

// Checks if the window is currently in focus
func (win *Window) IsInFocus() bool {
    return C.RGFW_window_isInFocus(win.ref) == C.RGFW_TRUE
}

// Raises the window to the top of the stack
func (win *Window) Raise() {
    C.RGFW_window_raise(win.ref)
}

// Maximizes the window
func (win *Window) Maximize() {
    C.RGFW_window_maximize(win.ref)
}

// Toggles fullscreen mode for the window
func (win *Window) SetFullscreen(fullscreen bool) {
    C.RGFW_window_setFullscreen(win.ref, C.RGFW_bool(boolToInt(fullscreen)))
}

// Centers the window on the screen
func (win *Window) Center() {
    C.RGFW_window_center(win.ref)
}

// Minimizes the window
func (win *Window) Minimize() {
    C.RGFW_window_minimize(win.ref)
}

// Restores the window from minimized state
func (win *Window) Restore() {
    C.RGFW_window_restore(win.ref)
}

// Makes the window a floating window
func (win *Window) SetFloating(floating bool) {
    C.RGFW_window_setFloating(win.ref, C.RGFW_bool(boolToInt(floating)))
}

// Sets the opacity level of the window
func (win *Window) SetOpacity(opacity uint8) {
    C.RGFW_window_setOpacity(win.ref, C.u8(opacity))
}

// Toggles window borders
func (win *Window) SetBorder(border bool) {
    C.RGFW_window_setBorder(win.ref, C.RGFW_bool(boolToInt(border)))
}

// Checks if the window is borderless
func (win *Window) Borderless() bool {
    return C.RGFW_window_borderless(win.ref) == C.RGFW_TRUE
}

// Toggles drag-and-drop (DND) support for the window
func (win *Window) SetDND(allow bool) {
    C.RGFW_window_setDND(win.ref, C.RGFW_bool(boolToInt(allow)))
}

// Checks if drag-and-drop (DND) is allowed
func (win *Window) AllowsDND() bool {
    return C.RGFW_window_allowsDND(win.ref) == C.RGFW_TRUE
}

// Toggles mouse passthrough for the window
func (win *Window) SetMousePassthrough(passthrough bool) {
    C.RGFW_window_setMousePassthrough(win.ref, C.RGFW_bool(boolToInt(passthrough)))
}

// Renames the window
func (win *Window) SetName(name string) {
    cName := C.CString(name)
    defer C.free(unsafe.Pointer(cName))
    C.RGFW_window_setName(win.ref, cName)
}

// Sets the icon for the window and taskbar
// RGFW_window_setIcon

// Sets the icon for the window and/or taskbar
// RGFW_window_setIconEx

// Sets the mouse icon for the window using a loaded bitmap
// RGFW_window_setMouse

// Sets the mouse to a standard system cursor.
// RGFW_window_setMouseStandard

// Sets the mouse to the default cursor icon.
// RGFW_window_setMouseDefault

// Set (enable or disable) raw mouse mode only for the select window
// RGFW_window_setRawMouseMode

// Lock/unlock the cursor.
// RGFW_window_captureMouse

// Lock/unlock the cursor and enable raw mpuise mode.
// RGFW_window_captureRawMouse

// Returns true if the mouse is using raw mouse mode
// func (win *Window) IsRawMouse() bool {
// 	return C.RGFW_window_isRawMouseMode(win.ref) == C.RGFW_TRUE
// }

// Returns true if the mouse is captured
// func (win *Window) IsCaptured() bool {
// 	return C.RGFW_window_isCaptured(win.ref) == C.RGFW_TRUE
// }

// Hides the window from view.
func (win *Window) Hide() {
    C.RGFW_window_hide(win.ref)
}

// Shows the window if it was hidden.
func (win *Window) Show() {
    C.RGFW_window_show(win.ref)
}

// Request a window flash to get attention from the user
// func (win *Window) Flash(request FlashRequest) {
//     C.RGFW_window_flash(win.ref, C.RGFW_flashRequest(request))
// }

// Sets whether the window should close.
func (win *Window) SetShouldClose(shouldClose bool) {
    C.RGFW_window_setShouldClose(win.ref, C.RGFW_bool(boolToInt(shouldClose)))
}

// Retrieves the mouse position relative to the window.
// success is true if the position was successfully retrieved
// RGFW_window_getMouse

// Shows or hides the mouse cursor for the window.
func (win *Window) ShowMouse(show bool) {
    C.RGFW_window_showMouse(win.ref, C.RGFW_bool(boolToInt(show)))
}

// Checks if the mouse is currently hidden in the window.
func (win *Window) IsMouseHidden() bool {
    return C.RGFW_window_isMouseHidden(win.ref) == C.RGFW_TRUE
}

// Moves the mouse to the specified position within the window.
func (win *Window) MoveMouse(x, y int32) {
    C.RGFW_window_moveMouse(win.ref, C.i32(x), C.i32(y))
}

// Checks if the window should close.
func (win *Window) ShouldClose() bool {
	return C.RGFW_window_shouldClose(win.ref) == C.RGFW_TRUE
}

// Checks if the window is currently fullscreen.
func (win *Window) IsFullscreen() bool {
	return C.RGFW_window_isFullscreen(win.ref) == C.RGFW_TRUE
}

// Checks if the window is currently hidden.
func (win *Window) IsHidden() bool {
	return C.RGFW_window_isHidden(win.ref) == C.RGFW_TRUE
}

// Checks if the window is minimized.
func (win *Window) IsMinimized() bool {
	return C.RGFW_window_isMinimized(win.ref) == C.RGFW_TRUE
}

// Checks if the window is maximized.
func (win *Window) IsMaximized() bool {
	return C.RGFW_window_isMaximized(win.ref) == C.RGFW_TRUE
}

// Checks if the window is floating.
func (win *Window) IsFloating() bool {
	return C.RGFW_window_isFloating(win.ref) == C.RGFW_TRUE
}

// Scales the window to match its monitor’s resolution.
// RGFW_window_scaleToMonitor

// Retrieves the monitor structure associated with the window.
// RGFW_window_getMonitor