package gpio

type Light struct {
	pin Pin
}

func NewLight(pin Pin) (*Light, error) {
	pin.SetMode(ModeOutput)
	return &Light{
		pin: pin,
	}, nil
}

func (l *Light) TurnOn() {
	l.pin.Set()
}

func (l *Light) TurnOff() {
	l.pin.Clear()
}

func (l *Light) Close() error {
	return l.pin.Close()
}
