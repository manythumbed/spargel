package geometry

import "testing"

func TestNormals(t *testing.T)	{
	checkVector(up, up.normalLeft(), left, ".normalLeft()", t)
	checkVector(down, down.normalLeft(), right, ".normalLeft()", t)
	checkVector(left, left.normalLeft(), down, ".normalLeft()", t)
	checkVector(right, right.normalLeft(), up, ".normalLeft()", t)
	checkVector(up, up.normalRight(), right, ".normalRight()", t)
	checkVector(down, down.normalRight(), left, ".normalRight()", t)
	checkVector(left, left.normalRight(), up, ".normalRight()", t)
	checkVector(right, right.normalRight(), down, ".normalRight()", t)
}

func checkVector(in, out, want vector, label string, t *testing.T)	{
	if(out != want)	{
		t.Errorf("%v%s = %v, want %v", in, label, out, want)
	}
}
