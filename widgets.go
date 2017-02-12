package main

import (
	"github.com/jroimartin/gocui"
	"fmt"
)

type PriceWidget struct {
	name string
	x, y int
	w int
	label string
}

type ButtonWidget struct {
	name string
	x, y int
	w int
	label string
	handler func(g *gocui.Gui, w *gocui.View) error
}

func NewPriceWidget(name string, x, y int, w int, label string) *PriceWidget {
	return &PriceWidget{name: name, x: x, y: y, w: w, label: label}
}

func NewButtonWidget(name string, x, y int, label string, handler func(g *gocui.Gui, v *gocui.View) error) *ButtonWidget {
	return &ButtonWidget{name: name, x: x, y: y, w: len(label), label: label, handler: handler}
}

func (w *PriceWidget) SetVal(val string) error {
	w.label = val
	return nil
}

func (w *PriceWidget) Layout(g *gocui.Gui) error {
	v, err := g.SetView(w.name, w.x, w.y, w.x+w.w, w.y+2)
	if err != nil && err != gocui.ErrUnknownView {
		return err
	}
	v.Clear()

	fmt.Fprint(v,  w.label)
	return nil
}

func (w *ButtonWidget) Layout(g *gocui.Gui) error {
	v, err := g.SetView(w.name, w.x, w.y, w.x+w.w, w.y+2)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		if _, err := g.SetCurrentView(w.name); err != nil {
			return err
		}
		if err := g.SetKeybinding(w.name, gocui.KeyEnter, gocui.ModNone, w.handler); err != nil {
			return err
		}
		fmt.Fprint(v, w.label)
	}
	return nil
}


