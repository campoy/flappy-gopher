package sdl

/*
#include "sdl_wrapper.h"

#if !(SDL_VERSION_ATLEAST(2,0,4))
#pragma message("SDL_GameControllerFromInstanceID is not supported before SDL 2.0.4")
static SDL_GameController* SDL_GameControllerFromInstanceID(SDL_JoystickID joyid)
{
	return NULL;
}
#endif

#if !(SDL_VERSION_ATLEAST(2,0,6))
#pragma message("SDL_GameControllerGetVendor is not supported before SDL 2.0.6")
static Uint16 SDL_GameControllerGetVendor(SDL_GameController* gamecontroller)
{
	return 0;
}

#pragma message("SDL_GameControllerGetProduct is not supported before SDL 2.0.6")
static Uint16 SDL_GameControllerGetProduct(SDL_GameController* gamecontroller)
{
	return 0;
}

#pragma message("SDL_GameControllerGetProductVersion is not supported before SDL 2.0.6")
static Uint16 SDL_GameControllerGetProductVersion(SDL_GameController* gamecontroller)
{
	return 0;
}

#pragma message("SDL_GameControllerNumMappings is not supported before SDL 2.0.6")
static int SDL_GameControllerNumMappings(void)
{
	return 0;
}

#pragma message("SDL_GameControllerMappingForIndex is not supported before SDL 2.0.6")
static char* SDL_GameControllerMappingForIndex(int mapping_index)
{
	return NULL;
}
#endif
*/
import "C"
import "unsafe"
import "encoding/binary"

// Types of game controller inputs.
const (
	CONTROLLER_BINDTYPE_NONE   = C.SDL_CONTROLLER_BINDTYPE_NONE
	CONTROLLER_BINDTYPE_BUTTON = C.SDL_CONTROLLER_BINDTYPE_BUTTON
	CONTROLLER_BINDTYPE_AXIS   = C.SDL_CONTROLLER_BINDTYPE_AXIS
	CONTROLLER_BINDTYPE_HAT    = C.SDL_CONTROLLER_BINDTYPE_HAT
)

// An enumeration of axes available from a controller.
// (https://wiki.libsdl.org/SDL_GameControllerAxis)
const (
	CONTROLLER_AXIS_INVALID      = C.SDL_CONTROLLER_AXIS_INVALID
	CONTROLLER_AXIS_LEFTX        = C.SDL_CONTROLLER_AXIS_LEFTX
	CONTROLLER_AXIS_LEFTY        = C.SDL_CONTROLLER_AXIS_LEFTY
	CONTROLLER_AXIS_RIGHTX       = C.SDL_CONTROLLER_AXIS_RIGHTX
	CONTROLLER_AXIS_RIGHTY       = C.SDL_CONTROLLER_AXIS_RIGHTY
	CONTROLLER_AXIS_TRIGGERLEFT  = C.SDL_CONTROLLER_AXIS_TRIGGERLEFT
	CONTROLLER_AXIS_TRIGGERRIGHT = C.SDL_CONTROLLER_AXIS_TRIGGERRIGHT
	CONTROLLER_AXIS_MAX          = C.SDL_CONTROLLER_AXIS_MAX
)

// An enumeration of buttons available from a controller.
// (https://wiki.libsdl.org/SDL_GameControllerButton)
const (
	CONTROLLER_BUTTON_INVALID       = C.SDL_CONTROLLER_BUTTON_INVALID
	CONTROLLER_BUTTON_A             = C.SDL_CONTROLLER_BUTTON_A
	CONTROLLER_BUTTON_B             = C.SDL_CONTROLLER_BUTTON_B
	CONTROLLER_BUTTON_X             = C.SDL_CONTROLLER_BUTTON_X
	CONTROLLER_BUTTON_Y             = C.SDL_CONTROLLER_BUTTON_Y
	CONTROLLER_BUTTON_BACK          = C.SDL_CONTROLLER_BUTTON_BACK
	CONTROLLER_BUTTON_GUIDE         = C.SDL_CONTROLLER_BUTTON_GUIDE
	CONTROLLER_BUTTON_START         = C.SDL_CONTROLLER_BUTTON_START
	CONTROLLER_BUTTON_LEFTSTICK     = C.SDL_CONTROLLER_BUTTON_LEFTSTICK
	CONTROLLER_BUTTON_RIGHTSTICK    = C.SDL_CONTROLLER_BUTTON_RIGHTSTICK
	CONTROLLER_BUTTON_LEFTSHOULDER  = C.SDL_CONTROLLER_BUTTON_LEFTSHOULDER
	CONTROLLER_BUTTON_RIGHTSHOULDER = C.SDL_CONTROLLER_BUTTON_RIGHTSHOULDER
	CONTROLLER_BUTTON_DPAD_UP       = C.SDL_CONTROLLER_BUTTON_DPAD_UP
	CONTROLLER_BUTTON_DPAD_DOWN     = C.SDL_CONTROLLER_BUTTON_DPAD_DOWN
	CONTROLLER_BUTTON_DPAD_LEFT     = C.SDL_CONTROLLER_BUTTON_DPAD_LEFT
	CONTROLLER_BUTTON_DPAD_RIGHT    = C.SDL_CONTROLLER_BUTTON_DPAD_RIGHT
	CONTROLLER_BUTTON_MAX           = C.SDL_CONTROLLER_BUTTON_MAX
)

// GameControllerBindType is a type of game controller input.
type GameControllerBindType C.SDL_GameControllerBindType

// GameControllerAxis is an axis on a game controller.
// (https://wiki.libsdl.org/SDL_GameControllerAxis)
type GameControllerAxis C.SDL_GameControllerAxis

// GameControllerButton is a button on a game controller.
// (https://wiki.libsdl.org/SDL_GameControllerButton)
type GameControllerButton C.SDL_GameControllerButton

// GameController used to identify an SDL game controller.
type GameController C.SDL_GameController

// GameControllerButtonBind SDL joystick layer binding for controller button/axis mapping.
type GameControllerButtonBind C.SDL_GameControllerButtonBind

func (ctrl *GameController) cptr() *C.SDL_GameController {
	return (*C.SDL_GameController)(unsafe.Pointer(ctrl))
}

func (axis GameControllerAxis) c() C.SDL_GameControllerAxis {
	return C.SDL_GameControllerAxis(axis)
}

func (btn GameControllerButton) c() C.SDL_GameControllerButton {
	return C.SDL_GameControllerButton(btn)
}

// GameControllerAddMapping adds support for controllers that SDL is unaware of or to cause an existing controller to have a different binding.
// (https://wiki.libsdl.org/SDL_GameControllerAddMapping)
func GameControllerAddMapping(mappingString string) int {
	_mappingString := C.CString(mappingString)
	defer C.free(unsafe.Pointer(_mappingString))
	return int(C.SDL_GameControllerAddMapping(_mappingString))
}

// GameControllerNumMappings returns the number of mappings installed.
func GameControllerNumMappings() int {
	return int(C.SDL_GameControllerNumMappings())
}

// GameControllerMappingForIndex returns the game controller mapping string at a particular index.
func GameControllerMappingForIndex(index int) string {
	mappingString := C.SDL_GameControllerMappingForIndex(C.int(index))
	defer C.free(unsafe.Pointer(mappingString))
	return C.GoString(mappingString)
}

// GameControllerMappingForGUID returns the game controller mapping string for a given GUID.
// (https://wiki.libsdl.org/SDL_GameControllerMappingForGUID)
func GameControllerMappingForGUID(guid JoystickGUID) string {
	mappingString := C.SDL_GameControllerMappingForGUID(guid.c())
	defer C.free(unsafe.Pointer(mappingString))
	return C.GoString(mappingString)
}

// IsGameController reports whether the given joystick is supported by the game controller interface.
// (https://wiki.libsdl.org/SDL_IsGameController)
func IsGameController(index int) bool {
	return C.SDL_IsGameController(C.int(index)) == C.SDL_TRUE
}

// GameControllerNameForIndex returns the implementation dependent name for the game controller.
// (https://wiki.libsdl.org/SDL_GameControllerNameForIndex)
func GameControllerNameForIndex(index int) string {
	return C.GoString(C.SDL_GameControllerNameForIndex(C.int(index)))
}

// GameControllerOpen opens a gamecontroller for use.
// (https://wiki.libsdl.org/SDL_GameControllerOpen)
func GameControllerOpen(index int) *GameController {
	return (*GameController)(C.SDL_GameControllerOpen(C.int(index)))
}

// GameControllerFromInstanceID returns the GameController associated with an instance id.
// (https://wiki.libsdl.org/SDL_GameControllerFromInstanceID)
func GameControllerFromInstanceID(joyid JoystickID) *GameController {
	return (*GameController)(C.SDL_GameControllerFromInstanceID(joyid.c()))
}

// Name returns the implementation dependent name for an opened game controller.
// (https://wiki.libsdl.org/SDL_GameControllerName)
func (ctrl *GameController) Name() string {
	return C.GoString(C.SDL_GameControllerName(ctrl.cptr()))
}

// Vendor returns the USB vendor ID of an opened controller, if available, 0 otherwise.
func (ctrl *GameController) Vendor() int {
	return int(C.SDL_GameControllerGetVendor(ctrl.cptr()))
}

// Product returns the USB product ID of an opened controller, if available, 0 otherwise.
func (ctrl *GameController) Product() int {
	return int(C.SDL_GameControllerGetProduct(ctrl.cptr()))
}

// ProductVersion returns the product version of an opened controller, if available, 0 otherwise.
func (ctrl *GameController) ProductVersion() int {
	return int(C.SDL_GameControllerGetProductVersion(ctrl.cptr()))
}

// Attached reports whether a controller has been opened and is currently connected.
// (https://wiki.libsdl.org/SDL_GameControllerGetAttached)
func (ctrl *GameController) Attached() bool {
	return C.SDL_GameControllerGetAttached(ctrl.cptr()) == C.SDL_TRUE
}

// Mapping returns the current mapping of a Game Controller.
// (https://wiki.libsdl.org/SDL_GameControllerMapping)
func (ctrl *GameController) Mapping() string {
	mappingString := C.SDL_GameControllerMapping(ctrl.cptr())
	defer C.free(unsafe.Pointer(mappingString))
	return C.GoString(mappingString)
}

// Joystick returns the Joystick ID from a Game Controller. The game controller builds on the Joystick API, but to be able to use the Joystick's functions with a gamepad, you need to use this first to get the joystick object.
// (https://wiki.libsdl.org/SDL_GameControllerGetJoystick)
func (ctrl *GameController) Joystick() *Joystick {
	return (*Joystick)(unsafe.Pointer(C.SDL_GameControllerGetJoystick(ctrl.cptr())))
}

// GameControllerEventState returns the current state of, enable, or disable events dealing with Game Controllers. This will not disable Joystick events, which can also be fired by a controller (see https://wiki.libsdl.org/SDL_JoystickEventState).
// (https://wiki.libsdl.org/SDL_GameControllerEventState)
func GameControllerEventState(state int) int {
	return int(C.SDL_GameControllerEventState(C.int(state)))
}

// GameControllerUpdate manually pumps game controller updates if not using the loop.
// (https://wiki.libsdl.org/SDL_GameControllerUpdate)
func GameControllerUpdate() {
	C.SDL_GameControllerUpdate()
}

// GameControllerGetAxisFromString converts a string into an enum representation for a GameControllerAxis.
// (https://wiki.libsdl.org/SDL_GameControllerGetAxisFromString)
func GameControllerGetAxisFromString(pchString string) GameControllerAxis {
	_pchString := C.CString(pchString)
	defer C.free(unsafe.Pointer(_pchString))
	return GameControllerAxis(C.SDL_GameControllerGetAxisFromString(_pchString))
}

// GameControllerGetStringForAxis converts from an axis enum to a string.
// (https://wiki.libsdl.org/SDL_GameControllerGetStringForAxis)
func GameControllerGetStringForAxis(axis GameControllerAxis) string {
	return C.GoString(C.SDL_GameControllerGetStringForAxis(axis.c()))
}

// BindForAxis returns the SDL joystick layer binding for a controller button mapping.
// (https://wiki.libsdl.org/SDL_GameControllerGetBindForAxis)
func (ctrl *GameController) BindForAxis(axis GameControllerAxis) GameControllerButtonBind {
	return GameControllerButtonBind(C.SDL_GameControllerGetBindForAxis(ctrl.cptr(), axis.c()))
}

// Axis returns the current state of an axis control on a game controller.
// (https://wiki.libsdl.org/SDL_GameControllerGetAxis)
func (ctrl *GameController) Axis(axis GameControllerAxis) int16 {
	return int16(C.SDL_GameControllerGetAxis(ctrl.cptr(), axis.c()))
}

// GameControllerGetButtonFromString turns a string into a button mapping.
// (https://wiki.libsdl.org/SDL_GameControllerGetButtonFromString)
func GameControllerGetButtonFromString(pchString string) GameControllerButton {
	_pchString := C.CString(pchString)
	defer C.free(unsafe.Pointer(_pchString))
	return GameControllerButton(C.SDL_GameControllerGetButtonFromString(_pchString))
}

// GameControllerGetStringForButton turns a button enum into a string mapping.
// (https://wiki.libsdl.org/SDL_GameControllerGetStringForButton)
func GameControllerGetStringForButton(btn GameControllerButton) string {
	return C.GoString(C.SDL_GameControllerGetStringForButton(btn.c()))
}

// BindForButton returns the SDL joystick layer binding for this controller button mapping.
// (https://wiki.libsdl.org/SDL_GameControllerGetBindForButton)
func (ctrl *GameController) BindForButton(btn GameControllerButton) GameControllerButtonBind {
	return GameControllerButtonBind(C.SDL_GameControllerGetBindForButton(ctrl.cptr(), btn.c()))
}

// Button returns the current state of a button on a game controller.
// (https://wiki.libsdl.org/SDL_GameControllerGetButton)
func (ctrl *GameController) Button(btn GameControllerButton) byte {
	return byte(C.SDL_GameControllerGetButton(ctrl.cptr(), btn.c()))
}

// Close closes a game controller previously opened with GameControllerOpen().
// (https://wiki.libsdl.org/SDL_GameControllerClose)
func (ctrl *GameController) Close() {
	C.SDL_GameControllerClose(ctrl.cptr())
}

// Type returns the type of game controller input for this SDL joystick layer binding.
func (bind *GameControllerButtonBind) Type() int {
	return int(bind.bindType)
}

// Button returns button mapped for this SDL joystick layer binding.
func (bind *GameControllerButtonBind) Button() int {
	val, _ := binary.Varint(bind.value[:4])
	return int(val)
}

// Axis returns axis mapped for this SDL joystick layer binding.
func (bind *GameControllerButtonBind) Axis() int {
	val, _ := binary.Varint(bind.value[:4])
	return int(val)
}

// Hat returns hat mapped for this SDL joystick layer binding.
func (bind *GameControllerButtonBind) Hat() int {
	val, _ := binary.Varint(bind.value[:4])
	return int(val)
}

// HatMask returns hat mask for this SDL joystick layer binding.
func (bind *GameControllerButtonBind) HatMask() int {
	val, _ := binary.Varint(bind.value[4:8])
	return int(val)
}
