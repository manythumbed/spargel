package geometry

import (
	"fmt"
	"math"
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

	p1 := Point{0, -1}.Left(1)
	p2 := p1.Up(1)
	p3 := p2.Right(1)
	p4 := p3.Down(1)

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

func TestProjection(t *testing.T) {
	p1 := Point{2, 2}
	p2 := p1.Up(2)

	if p2.x != 2 && p2.y != 3 {
		t.Errorf("BOOM")
	}
}

func TestTo(t *testing.T) {
	a := Point{1, 1}
	b := Point{4, 4}

	root2 := 1 / math.Sqrt(2)
	d := Direction{root2, root2}
	if a.To(b) != d {
		t.Errorf("%v.To(%v) = %v, want %v", a, b, a.To(b), d)
	}

	d1 := a.Between(b)
	b1 := a.Project(d, d1)
	if b != b1 {
		t.Errorf("%v.Project(%v, %v) = %v, want %v", a, d, d1, b1, b)
	}
}
