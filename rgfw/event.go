package rgfw

/*
	#include "RFGW_impl.h"
	#include "event.h"
*/
import "C"

type CommonEvent struct {
	Type EventType
	Window *Window
}

type MouseButtonEvent struct {
	Type EventType
	Window *Window
	Value  MouseButton 
}

type MouseScrollEvent struct {
	Type EventType
	Window *Window
	X, Y   float32
}

type MousePosEvent struct {
	Type EventType
	Window     *Window
	X, Y       int32
	VecX, VecY float32
}

type KeyEvent struct {
	Type EventType
	Window *Window
	Value  Key
	Repeat bool   
	Mod    Keymod 
}

// type KeyCharEvent struct {
// 	Type EventType
// 	Window *Window
// 	Value  uint
// }

type DataDropEvent struct {
	Type EventType
	Window *Window
	Files  []string // Converted from char** for Go-friendliness
	Count  uint64   // size_t
}

type DataDragEvent struct {
	Type EventType
	Window *Window
	X, Y   int32
}

type ScaleUpdatedEvent struct {
	Type EventType
	Window *Window
	X, Y   float32
}

// type MonitorEvent struct {
// 	Type EventType
// 	Window  *Window
// 	Monitor unsafe.Pointer // Or a specific Monitor struct if you've wrapped it
// }

type Event struct {
	Type EventType
	Common CommonEvent
	Button MouseButtonEvent
	Scroll MouseScrollEvent
	Mouse MousePosEvent
	Key KeyEvent
	// KeyChar KeyCharEvent
	Drop DataDropEvent
	Drag DataDragEvent
	Scale ScaleUpdatedEvent
	// Monitor MonitorEvent
}

// Polls and pops the next event with the matching target window in event queue, pushes back events that don't match
//
// NOTE: Using this function without a loop may cause event lag.
// For multi-threaded systems, use RGFW_pollEvents combined with RGFW_window_checkQueuedEvent.
//  
// Example:
// 
//  event, eventFound := win.CheckEvent()
//  for ;eventFound; event, eventFound = win.CheckEvent() {
// 		// handle event
//  }
func (win *Window) CheckEvent() (Event, bool) {
	var cEvent C.RGFW_event
	eventFound := C.RGFW_window_checkEvent(win.ref, &cEvent)
	if eventFound == C.RGFW_FALSE {
		return Event{}, false
	}

	data := C.wrapper_getEventData(&cEvent)
	eventType := EventType(data.eventType)

	common := CommonEvent{
		Type: eventType,
		Window: win,
	}

	event := Event{Type: eventType, Common: common}

	switch eventType {
	case EventKeyPressed, EventKeyReleased:
		event.Key = KeyEvent{
			Type: eventType,
			Window: win,
			Value: Key(data.keyValue),
			Repeat: data.keyRepeat == C.RGFW_TRUE,
			Mod: Keymod(data.keyMod),
		}
	case EventMouseButtonPressed, EventMouseButtonReleased:
		event.Button = MouseButtonEvent{
            Type: eventType,
            Window: win,
            Value: MouseButton(data.mouseButtonValue),
        }
	}

	return event, true
}

// Pops the first queued event with the matching target window, pushes back events that don't match
// RGFW_window_checkQueuedEvent