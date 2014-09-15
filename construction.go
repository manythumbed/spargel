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
type Magnitude float64

func (pt Point) Render(p Plotter) {
	p.Move(pt)
}

func (pt Point) Project(d Direction, scale Magnitude) Point {
	a := vector(pt)
	b := vector(d).unit().scale(float64(scale))
	v := a.add(b)

	return Point{v.x, v.y}
}

func (a Point) To(b Point) Direction {
	v := vector(b).subtract(vector(a)).unit()
	return Direction{v.x, v.y}
}

func (a Point) Between(b Point) Magnitude {
	return Magnitude(vector(b).subtract(vector(a)).magnitude())
}

func (pt Point) Left(scale Magnitude) Point {
	return pt.Project(Direction(left), scale)
}

func (pt Point) Right(scale Magnitude) Point {
	return pt.Project(Direction(right), scale)
}

func (pt Point) Up(scale Magnitude) Point {
	return pt.Project(Direction(up), scale)
}

func (pt Point) Down(scale Magnitude) Point {
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
