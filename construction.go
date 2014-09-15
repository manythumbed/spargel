package geometry

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
type Direction vector

func (pt Point) Render(p Plotter) {
	p.Move(pt)
}

func (pt Point) Project(d Direction, scale float64) Point {
	v := vector(pt).add(vector(d).scale(scale))

	return Point{v.x, v.y}
}

func (pt Point) Left(scale float64) Point {
	return pt.Project(Direction(left), scale)
}

func (pt Point) Right(scale float64) Point {
	return pt.Project(Direction(right), scale)
}

func (pt Point) Up(scale float64) Point {
	return pt.Project(Direction(up), scale)
}

func (pt Point) Down(scale float64) Point {
	return pt.Project(Direction(down), scale)
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
