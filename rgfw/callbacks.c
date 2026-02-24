#include "_cgo_export.h"
#include "../third_party/RGFW.h"

void RFGW_setWindowMovedCallbackCB() {
    RGFW_setWindowMovedCallback((RGFW_windowMovedfunc) goWindowMovedCB);
}