package gpio

//go:generate counterfeiter . Pin

// Mode represents a state of a GPIO pin
type Edge string

const (
	EdgeNone    Edge = "none"
	EdgeRising  Edge = "rising"
	EdgeFalling Edge = "falling"
	EdgeBoth    Edge = "both"
)

// Edge represents the edge on which a pin interrupt is triggered
type Mode string

const (
	ModeInput  Mode = "in"
	ModeOutput Mode = "out"
	ModePWM    Mode = "pwm"
)

// Event defines the callback function used to inform the caller
// of an interrupt.
type Event func()

// Pin represents a GPIO pin.
type Pin interface {
	Mode() Mode                   // gets the current pin mode
	SetMode(Mode)                 // set the current pin mode
	Set()                         // sets the pin state high
	Clear()                       // sets the pin state low
	Close() error                 // if applicable, closes the pin
	Get() bool                    // returns the current pin state
	BeginWatch(Edge, Event) error // calls the function argument when an edge trigger event occurs
	EndWatch() error              // stops watching the pin
}
