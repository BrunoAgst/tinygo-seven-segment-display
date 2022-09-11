package main

import (
	"machine"
	"time"
)

const (
	a = machine.D2
	b = machine.D3
	c = machine.D4
	d = machine.D5
	e = machine.D6
	f = machine.D7
	g = machine.D8
)

func main() {
	a.Configure(machine.PinConfig{Mode: machine.PinOutput})
	b.Configure(machine.PinConfig{Mode: machine.PinOutput})
	c.Configure(machine.PinConfig{Mode: machine.PinOutput})
	d.Configure(machine.PinConfig{Mode: machine.PinOutput})
	e.Configure(machine.PinConfig{Mode: machine.PinOutput})
	f.Configure(machine.PinConfig{Mode: machine.PinOutput})
	g.Configure(machine.PinConfig{Mode: machine.PinOutput})

	for {
		zero()
		one()
		two()
		three()
		four()
		five()
		six()
		seven()
		eight()
		nine()
	}
}

func zero() {
	a.High()
	b.High()
	c.High()
	d.High()
	e.High()
	f.High()
	g.Low()
	time.Sleep(time.Second * 1)
}

func one() {
	a.Low()
	b.High()
	c.High()
	d.Low()
	e.Low()
	f.Low()
	g.Low()
	time.Sleep(time.Second * 1)
}

func two() {
	a.High()
	b.High()
	c.Low()
	d.High()
	e.High()
	f.Low()
	g.High()
	time.Sleep(time.Second * 1)
}

func three() {
	a.High()
	b.High()
	c.High()
	d.High()
	e.Low()
	f.Low()
	g.High()
	time.Sleep(time.Second * 1)
}

func four() {
	a.Low()
	b.High()
	c.High()
	d.Low()
	e.Low()
	f.High()
	g.High()
	time.Sleep(time.Second * 1)
}

func five() {
	a.High()
	b.Low()
	c.High()
	d.High()
	e.Low()
	f.High()
	g.High()
	time.Sleep(time.Second * 1)
}

func six() {
	a.High()
	b.Low()
	c.High()
	d.High()
	e.High()
	f.High()
	g.High()
	time.Sleep(time.Second * 1)
}

func seven() {
	a.High()
	b.High()
	c.High()
	d.Low()
	e.Low()
	f.Low()
	g.Low()
	time.Sleep(time.Second * 1)
}

func eight() {
	a.High()
	b.High()
	c.High()
	d.High()
	e.High()
	f.High()
	g.High()
	time.Sleep(time.Second * 1)
}

func nine() {
	a.High()
	b.High()
	c.High()
	d.High()
	e.Low()
	f.High()
	g.High()
	time.Sleep(time.Second * 1)
}
