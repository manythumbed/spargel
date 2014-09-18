package geometry

import (
	"fmt"
	"math"
	"testing"
)

func TestNormals(t *testing.T) {
	checkVector(up, up.normalLeft(), left, ".normalLeft()", t)
	checkVector(down, down.normalLeft(), right, ".normalLeft()", t)
	checkVector(left, left.normalLeft(), down, ".normalLeft()", t)
	checkVector(right, right.normalLeft(), up, ".normalLeft()", t)
	checkVector(up, up.normalRight(), right, ".normalRight()", t)
	checkVector(down, down.normalRight(), left, ".normalRight()", t)
	checkVector(left, left.normalRight(), up, ".normalRight()", t)
	checkVector(right, right.normalRight(), down, ".normalRight()", t)
}

func TestUnit(t *testing.T) {
	root2 := 1 / math.Sqrt(2)
	checkVector(vector{1, 1}, vector{1, 1}.unit(), vector{root2, root2}, ".unit()", t)
	b := up.scale(4)
	checkVector(b, b.unit(), up, ".unit()", t)
}

func TestAdd(t *testing.T) {
	checkVector(up, up.add(right), vector{1, 1}, fmt.Sprintf(".add(%v)", right), t)
	checkVector(right, right.add(up), vector{1, 1}, fmt.Sprintf(".add(%v)", up), t)
	checkVector(up, up.add(left), vector{-1, 1}, fmt.Sprintf(".add(%v)", left), t)
	checkVector(left, left.add(up), vector{-1, 1}, fmt.Sprintf(".add(%v)", up), t)
}

func TestSubtract(t *testing.T) {
	checkVector(up, up.subtract(right), vector{-1, 1}, fmt.Sprintf(".subtract(%v)", right), t)
	checkVector(right, right.subtract(up), vector{1, -1}, fmt.Sprintf(".subtract(%v)", up), t)
	checkVector(up, up.subtract(left), vector{1, 1}, fmt.Sprintf(".subtract(%v)", left), t)
	checkVector(left, left.subtract(up), vector{-1, -1}, fmt.Sprintf(".subtract(%v)", up), t)
}

func checkVector(in, out, want vector, label string, t *testing.T) {
	if out != want {
		t.Errorf("%v%s = %v, want %v", in, label, out, want)
	}
}
