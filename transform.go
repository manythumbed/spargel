package geometry

type Transformer interface {
	Transform(p Point) Point
}

type Translation struct {
	Direction
}

func (t Translation) Transform(p Point) Point {
	v := vector(p).add(vector(t.Direction))
	return Point(v)
}

type Scaling struct {
	Magnitude
}

func (s Scaling) Transform(p Point) Point {
	v := vector(p).scale(float64(s.Magnitude))

	return Point(v)
}

type Transformations []Transformer

func (t Transformations) Transform(p Point) Point {
	for _, value := range t {
		p = value.Transform(p)
	}

	return p
}
