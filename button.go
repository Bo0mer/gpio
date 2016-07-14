package gpio

import "time"

const statePressed = false
const stateReleased = true

type Button struct {
	pin Pin

	lastChange time.Time
	state      bool

	onPress   func()
	onRelease func()
}

func NewButton(pin Pin, onPress, onRelease func()) (*Button, error) {
	pin.SetMode(ModeInput)
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

func (b *Button) Close() error {
	if err := b.pin.EndWatch(); err != nil {
		return err
	}
	return b.pin.Close()
}
