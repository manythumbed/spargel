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
