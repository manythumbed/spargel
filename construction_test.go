package geometry

import "testing"

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
		t.Errorf("len(plotter.Instructions) = %v, want %v", len(plotter.Instructions), 3)
	}

}
