package gpio

import "time"

const statePressed = false
const stateReleased = true

// Button represents a hardware button connected on a GPIO pin.
type Button struct {
	pin Pin

	lastChange time.Time
	state      bool

	onPress   func()
	onRelease func()
}

// NewButton creates a button on the specified GPIO pin.
// It is caller's responsiblity to open the GPIO pin with Input mode.
func NewButton(pin Pin, onPress, onRelease func()) (*Button, error) {
	if pin.Mode() != ModeInput {
		return nil, &UnsupportedModeError{pin.Mode()}
	}

	b := &Button{
		pin:        pin,
		lastChange: time.Now(),
		state:      pin.Get(),
		onPress:    onPress,
		onRelease:  onRelease,
	}

	pin.BeginWatch(EdgeBoth, func() {
		state := pin.Get()
		now := time.Now()
		if b.state != state && now.Sub(b.lastChange) > time.Millisecond*50 {
			b.state = state
			b.lastChange = now
			if b.state == statePressed && b.onPress != nil {
				b.onPress()
			}
			if b.state == stateReleased && b.onRelease != nil {
				b.onRelease()
			}
		}
	})

	return b, nil
}

// Close frees all resoruces allocated by the button.
func (b *Button) Close() error {
	if err := b.pin.EndWatch(); err != nil {
		return err
	}
	return b.pin.Close()
}
