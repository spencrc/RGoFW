package rgfw

/*
	#include "RFGW_impl.h"
	#include "event.h"
*/
import "C"

type Event struct {
	Type EventType
	Common CommonEvent
	Button MouseButtonEvent
	Scroll MouseScrollEvent
	Mouse MousePosEvent
	Key KeyEvent
	KeyChar KeyCharEvent
	Drop DataDropEvent
	Drag DataDragEvent
	Scale ScaleUpdatedEvent
	Monitor MonitorEvent
}

func (win *Window) CheckEvent() (Event, bool) {
	var cEvent C.RGFW_event
	eventFound := C.RGFW_window_checkEvent(win.cPtr, &cEvent)
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
	case KeyPressed, KeyReleased:
		event.Key = KeyEvent{
			Type: eventType,
			Window: win,
			Value: Key(data.keyValue),
			Repeat: data.keyRepeat == C.RGFW_TRUE,
			Mod: Keymod(data.keyMod),
		}
	case MouseButtonPressed, MouseButtonReleased:
		event.Button = MouseButtonEvent{
            Type: eventType,
            Window: win,
            Value: MouseButton(data.mouseButtonValue),
        }
	}

	return event, true
}