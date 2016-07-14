package main

import (
	"log"
	"time"

	"github.com/Bo0mer/gpio"
	dgpio "github.com/davecheney/gpio"
)

type Pin struct {
	pin dgpio.Pin
}

func (p *Pin) Mode() gpio.Mode {
	return gpio.Mode(p.pin.Mode())
}

func (p *Pin) SetMode(mode gpio.Mode) {
	p.pin.SetMode(dgpio.Mode(mode))
}

func (p *Pin) Get() bool {
	return p.pin.Get()
}

func (p *Pin) Set() {
	p.pin.Set()
}

func (p *Pin) Clear() {
	p.pin.Clear()
}

func (p *Pin) Close() error {
	return p.pin.Close()
}

func (p *Pin) BeginWatch(edge gpio.Edge, e gpio.Event) error {
	return p.pin.BeginWatch(dgpio.Edge(edge), dgpio.IRQEvent(e))
}

func (p *Pin) EndWatch() error {
	return p.pin.EndWatch()
}

func OnPin(n int, mode dgpio.Mode) *Pin {
	p, err := dgpio.OpenPin(n, mode)
	if err != nil {
		log.Fatalf("main: error open pin %d: %v\n", n, err)
	}
	return &Pin{p}
}

func main() {
	blink, err := gpio.NewBlinker(OnPin(dgpio.GPIO17, dgpio.ModeOutput), time.Second)
	if err != nil {
		log.Fatalln(err)
	}
	defer blink.Close()
	defer blink.Stop()
	blink.Start()

	select {}
}
