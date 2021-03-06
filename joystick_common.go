// Package joystick implements a Polled API to read the state of an attached joystick.
// currently Windows & Linux are supported.
// Package is pure go and requires no external dependencies
//
// Installation:
//   go get github.com/simulatedsimian/joystick
//
// Example:
//   js, err := joystick.Open(jsid)
//   if err != nil {
//     panic(err)
//   }
//
//   fmt.Printf("Joystick Name: %s", js.Name())
//   fmt.Printf("   Axis Count: %d", js.AxisCount())
//   fmt.Printf(" Button Count: %d", js.ButtonCount())
//
//   state, err := joystick.Read()
//   if err != nil {
//      panic(err)
//   }
//
//   fmt.Printf("Axis Data: %v", state.AxisData)
//   js.Close()
//
package joystick

// State holds the current state of the joystick
type State struct {
	// Value of each axis as an integer in the range -32767 to 32768
	AxisData []int
	// The state of each button as a bit in a 32 bit integer. 1 = pressed, 0 = not pressed
	Buttons uint32
}

// Interface Joystick provides access to the Joystick opened with the Open() function
type Joystick interface {
	// AxisCount returns the number of Axis supported by this Joystick
	AxisCount() int
	// ButtonCount returns the number of buttons supported by this Joystick
	ButtonCount() int
	// Name returns the string name of this Joystick
	Name() string
	// Read returns the current State of the joystick.
	// On an error condition (for example, joystick has been unplugged) error is not nil
	Read() (State, error)
	// Close releases this joystick resource
	Close()
}

// Event is a raw event from gamepad serial port device
type Event struct {
	Time   uint32 /* event timestamp in milliseconds */
	Value  int16  /* value */
	Type   uint8  /* event type */
	Number uint8  /* axis/button number */
}

// JoystickChannelled Interface provides access to the Joystick opened with the Open() function and make events channel to out
type JoystickChannelled interface {
	Joystick
	// Events receives the events read from the joystick input.
	// If an Event is not consumed before the next Event is ready, then
	// the new event will be dropped. Events can be used to wait for state changes,
	// after which the full state can be retrieved with Read.
	Events() <-chan Event
}
