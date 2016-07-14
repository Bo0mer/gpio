package gpio

import "time"

// Blinker represents a hardware Light connected on a GPIO pin that blinks.
type Blinker struct {
	pin      Pin
	interval time.Duration
	stop     chan struct{}
}

// NewBlinker creates a Blinker on the specified GPIO pin.
// It is caller's responsiblity to open the GPIO pin with Output mode.
func NewBlinker(pin Pin, interval time.Duration) (*Blinker, error) {
	if pin.Mode() != ModeOutput {
		return nil, &UnsupportedModeError{pin.Mode()}
	}

	return &Blinker{
		pin:      pin,
		interval: interval,
	}, nil
}

// Start starts the blinker.
func (b *Blinker) Start() {
	b.stop = make(chan struct{})
	go b.blink()
}

// Stop stops the blinker.
func (b *Blinker) Stop() {
	close(b.stop)
}

// Close frees all resources allocated by the blinker.
func (b *Blinker) Close() error {
	return b.pin.Close()
}

func (b *Blinker) blink() {
	tick := time.Tick(b.interval)
	for {
		select {
		case <-tick:
			if on := b.pin.Get(); on {
				b.pin.Clear()
				continue
			}
			b.pin.Set()
		case <-b.stop:
			return
		}
	}
}
