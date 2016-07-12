package gpio

import "github.com/davecheney/gpio"

func NewLight(n int) (*Light, error) {
	pin, err := gpio.OpenPin(n, gpio.ModeOutput)
	if err != nil {
		return nil, err
	}

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
