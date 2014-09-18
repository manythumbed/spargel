// Package geometry provides simple primitives for 2D geometry.
package geometry

import "math"

type vector struct {
	x, y float64
}

func (a vector) add(b vector) vector {
	return vector{a.x + b.x, a.y + b.y}
}

func (a vector) subtract(b vector) vector {
	return vector{a.x - b.x, a.y - b.y}
}

func (a vector) scale(b float64) vector {
	return vector{a.x * b, a.y * b}
}

func (a vector) normalLeft() vector {
	return vector{-1 * a.y, a.x}
}

func (a vector) normalRight() vector {
	return vector{a.y, -1 * a.x}
}

func (a vector) magnitude() float64 {
	return math.Hypot(a.x, a.y)
}

func (a vector) unit() vector {
	magnitude := a.magnitude()
	return vector{a.x / magnitude, a.y / magnitude}
}

var up = vector{0, 1}
var down = vector{0, -1}
var left = vector{-1, 0}
var right = vector{1, 0}
