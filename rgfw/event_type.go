package rgfw

/*
	#include "RFGW_impl.h"
*/
import "C"

type EventType uint8

const (
    EventNone                EventType = C.RGFW_eventNone
    EventKeyPressed          EventType = C.RGFW_keyPressed
    EventKeyReleased         EventType = C.RGFW_keyReleased
    // EventKeyChar             EventType = C.RGFW_keyChar
    EventMouseButtonPressed  EventType = C.RGFW_mouseButtonPressed
    EventMouseButtonReleased EventType = C.RGFW_mouseButtonReleased
    EventMouseScroll         EventType = C.RGFW_mouseScroll
    EventMousePosChanged     EventType = C.RGFW_mousePosChanged
    EventWindowMoved         EventType = C.RGFW_windowMoved
    EventWindowResized       EventType = C.RGFW_windowResized
    EventFocusIn             EventType = C.RGFW_focusIn
    EventFocusOut            EventType = C.RGFW_focusOut
    EventMouseEnter          EventType = C.RGFW_mouseEnter
    EventMouseLeave          EventType = C.RGFW_mouseLeave
    EventWindowRefresh       EventType = C.RGFW_windowRefresh
    EventQuit                EventType = C.RGFW_quit
    EventDataDrop            EventType = C.RGFW_dataDrop
    EventDataDrag            EventType = C.RGFW_dataDrag
    EventWindowMaximized     EventType = C.RGFW_windowMaximized
    EventWindowMinimized     EventType = C.RGFW_windowMinimized
    EventWindowRestored      EventType = C.RGFW_windowRestored
    EventScaleUpdated        EventType = C.RGFW_scaleUpdated
    // EventMonitorConnected    EventType = C.RGFW_monitorConnected
    // EventMonitorDisconnected EventType = C.RGFW_monitorDisconnected
)

type EventFlag uint

const (
    EventFlagKeyPressed          EventFlag = C.RGFW_keyPressedFlag
    EventFlagKeyReleased         EventFlag = C.RGFW_keyReleasedFlag
    // EventFlagKeyChar             EventFlag = C.RGFW_keyCharFlag
    EventFlagMouseScroll         EventFlag = C.RGFW_mouseScrollFlag
    EventFlagMouseButtonPressed  EventFlag = C.RGFW_mouseButtonPressedFlag
    EventFlagMouseButtonReleased EventFlag = C.RGFW_mouseButtonReleasedFlag
    EventFlagMousePosChanged     EventFlag = C.RGFW_mousePosChangedFlag
    EventFlagMouseEnter          EventFlag = C.RGFW_mouseEnterFlag
    EventFlagMouseLeave          EventFlag = C.RGFW_mouseLeaveFlag
    EventFlagWindowMoved         EventFlag = C.RGFW_windowMovedFlag
    EventFlagWindowResized       EventFlag = C.RGFW_windowResizedFlag
    EventFlagFocusIn             EventFlag = C.RGFW_focusInFlag
    EventFlagFocusOut            EventFlag = C.RGFW_focusOutFlag
    EventFlagWindowRefresh       EventFlag = C.RGFW_windowRefreshFlag
    EventFlagWindowMaximized     EventFlag = C.RGFW_windowMaximizedFlag
    EventFlagWindowMinimized     EventFlag = C.RGFW_windowMinimizedFlag
    EventFlagWindowRestored      EventFlag = C.RGFW_windowRestoredFlag
    EventFlagScaleUpdated        EventFlag = C.RGFW_scaleUpdatedFlag
    EventFlagQuit                EventFlag = C.RGFW_quitFlag
    EventFlagDataDrop            EventFlag = C.RGFW_dataDropFlag
    EventFlagDataDrag            EventFlag = C.RGFW_dataDragFlag
    // EventFlagMonitorConnected    EventFlag = C.RGFW_monitorConnectedFlag
    // EventFlagMonitorDisconnected EventFlag = C.RGFW_monitorDisconnectedFlag
    EventFlagKeyEvents           EventFlag = C.RGFW_keyEventsFlag
    EventFlagMouseEvents         EventFlag = C.RGFW_mouseEventsFlag
    EventFlagWindowEvents        EventFlag = C.RGFW_windowEventsFlag
    EventFlagFocusEvents         EventFlag = C.RGFW_focusEventsFlag
    EventFlagDataDropEvents      EventFlag = C.RGFW_dataDropEventsFlag
    // EventFlagMonitorEvents       EventFlag = C.RGFW_monitorEventsFlag
    EventFlagAll                 EventFlag = C.RGFW_allEventFlags
)