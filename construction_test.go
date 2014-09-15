package geometry

import (
	"fmt"
	"testing"
)

func TestTextPlotter(t *testing.T) {
	plotter := &TextPlotter{}

	origin := Point{0, 0}
	destination := Point{1, 1}
	l := Line{origin, destination}
	c := Curve{origin, Point(up), Point(right), destination}

	origin.Render(plotter)
	destination.Render(plotter)
	l.Render(plotter)
	c.Render(plotter)

	if 4 != len(plotter.Instructions) {
		t.Errorf("len(plotter.Instructions) = %v, want %v", len(plotter.Instructions), 4)
	}

	text := fmt.Sprintf("%s", plotter)
	if len(text) <= 0 {
		t.Errorf("String representation of plotter should be greater than 0")
	}

}

func TestConstruction(t *testing.T) {
	p := &TextPlotter{}
	c := Construction{"test 1", nil}

	p1 := Point{0, 0}
	p2 := Point{0, 1}
	p3 := Point{1, 1}
	p4 := Point{1, 0}

	c.Add(p1)
	c.Add(Line{p1, p2})
	c.Add(Line{p2, p3})
	c.Add(Line{p3, p4})
	c.Add(Line{p4, p1})

	c.Plot(p)

	if len(p.Instructions) != 5 {
		t.Errorf("Should have 5 instructions")
	}
}
