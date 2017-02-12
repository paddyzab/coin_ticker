package main

import (
	"github.com/jroimartin/gocui"
	"fmt"
)

// UI Widget used for displaying value in label
type PriceWidget struct {
	name string
	x, y int
	w int
	label string
}

// UI Widget capable of executing functions after a trigger
type ButtonWidget struct {
	name string
	x, y int
	w int
	label string
	handler func(g *gocui.Gui, w *gocui.View) error
	handler2 func(g *gocui.Gui, w *gocui.View) error
}

// Creates new PriceWidget
func NewPriceWidget(name string, x, y int, w int, label string) *PriceWidget {
	return &PriceWidget{name: name, x: x, y: y, w: w, label: label}
}

// Creates new ButtonWidget
func NewButtonWidget(name string, x, y int, label string, handler func(g *gocui.Gui, v *gocui.View) error, handler2 func(g *gocui.Gui, v *gocui.View) error) *ButtonWidget {
	return &ButtonWidget{name: name, x: x, y: y, w: len(label) + 1, label: label, handler: handler, handler2: handler2}
}

// Modifies label of a PriceWidget
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

		if err := g.SetKeybinding(w.name, gocui.KeyEnter, gocui.ModNone, w.handler2); err != nil {
			return err
		}
		fmt.Fprint(v, w.label)
	}
	return nil
}


