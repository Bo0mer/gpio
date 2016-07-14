package gpio

// Light represents a hardware LED connected on a GPIO pin.
type Light struct {
	pin Pin
}

// NewLight creates a Light on the specified GPIO pin.
// It is caller's responsiblity to open the GPIO pin with Output mode.
func NewLight(pin Pin) (*Light, error) {
	pin.SetMode(ModeOutput)
	return &Light{
		pin: pin,
	}, nil
}

// TurnOn lights the light.
func (l *Light) TurnOn() {
	l.pin.Set()
}

// TurnOff extinguishes the light.
func (l *Light) TurnOff() {
	l.pin.Clear()
}

// Close frees all reosuces allocated by the light.
func (l *Light) Close() error {
	return l.pin.Close()
}
