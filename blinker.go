package gpio

import (
	"time"

	"github.com/davecheney/gpio"
)

type Blinker struct {
	pin      gpio.Pin
	interval time.Duration
	stop     chan struct{}
}

func NewBlinker(n int, interval time.Duration) (*Blinker, error) {
	pin, err := gpio.OpenPin(n, gpio.ModeOutput)
	if err != nil {
		return nil, err
	}

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
