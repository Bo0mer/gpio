package gpio

import "time"

type Blinker struct {
	pin      Pin
	interval time.Duration
	stop     chan struct{}
}

func NewBlinker(pin Pin, interval time.Duration) (*Blinker, error) {
	pin.SetMode(ModeOutput)
	return &Blinker{
		pin:      pin,
		interval: interval,
	}, nil
}

func (b *Blinker) Start() {
	b.stop = make(chan struct{})
	go b.blink()
}

func (b *Blinker) Stop() {
	close(b.stop)
}

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
