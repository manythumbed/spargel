package geometry

import (
	"bytes"
	"fmt"
)

// Plotter maps points, lines and curves to a geometric construction
type Plotter interface {
	Move(p Point)
	Line(start, end Point)
	Curve(start, c1, c2, end Point)
}

type Instruction interface {
	Render(p Plotter)
}

type Construction struct {
	Title        string
	Instructions []Instruction
}

func (c Construction) Plot(p Plotter) {
	for _, instruction := range c.Instructions {
		instruction.Render(p)
	}
}

func (c *Construction) Add(i Instruction) {
	c.Instructions = append(c.Instructions, i)
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
		fmt.Sprintf("Curve from (%f, %f) to (%f, %f) with controls at (%f, %f) and (%f, %f)", start.x, start.y, end.x, end.y, c1.x, c1.y, c2.x, c2.y))
}

func (t TextPlotter) String() string {
	buffer := &bytes.Buffer{}
	for _, value := range t.Instructions {
		fmt.Fprintf(buffer, "%s\n", value)
	}
	return buffer.String()
}
