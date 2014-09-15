package geometry

import "fmt"

// Plotter maps points, lines and curves to a geometric construction
type Plotter interface {
	Move(p Point)
	Line(start, end Point)
	Curve(start, c1, c2, end Point)
}

type Instruction interface {
	Render(p Plotter)
}

type Point vector

func (pt Point) Render(p Plotter) {
	p.Move(pt)
}

type Line struct {
	Start, End Point
}

func (l Line) Render(p Plotter) {
	p.Line(l.Start, l.End)
}

type Curve struct {
	Start, C1, C2, End Point
}

func (c Curve) Render(p Plotter) {
	p.Curve(c.Start, c.C1, c.C2, c.End)
}

type TextPlotter struct {
	Instructions []string
}

func (t *TextPlotter) Move(p Point) {
	t.Instructions = append(t.Instructions, fmt.Sprintf("Move to (%f, %f)", p.x, p.y))
}

func (t *TextPlotter) Line(start, end Point) {
	t.Instructions = append(t.Instructions,
		fmt.Sprintf("Line from (%f, %f) to (%f, %f)", start.x, start.y, end.x, end.y))
}

func (t *TextPlotter) Curve(start, c1, c2, end Point) {
	t.Instructions = append(t.Instructions,
		fmt.Sprintf("Curve from (%f, %f) to (%f, %f) with controls at (%f, %f) and (%f, %f)", start.x, start.y, end.x, end.y))
}
