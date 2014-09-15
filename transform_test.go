package geometry

import "testing"

func TestTranslation(t *testing.T) {
	var data = []struct {
		in   Point
		t    Translation
		want Point
	}{
		{Point{0, 0}, Translation{Direction{0, 0}}, Point{0, 0}},
		{Point{1, -1}, Translation{Direction{0, 0}}, Point{1, -1}},
		{Point{1, -1}, Translation{Direction{-1, 1}}, Point{0, 0}},
		{Point{1, -1}, Translation{Direction{-2, 3}}, Point{-1, 2}},
	}

	for _, value := range data {
		out := value.t.Transform(value.in)
		if out != value.want {
			t.Errorf("%v.Transform(%v) = %v, want %v", value.t, value.in, out, value.want)
		}
	}
}

func TestScaling(t *testing.T) {
	var data = []struct {
		in   Point
		s    Scaling
		want Point
	}{
		{Point{0, 0}, Scaling{Magnitude(0)}, Point{0, 0}},
		{Point{1, -1}, Scaling{Magnitude(0)}, Point{0, 0}},
		{Point{1, -1}, Scaling{Magnitude(-1)}, Point{-1, 1}},
		{Point{1, -1}, Scaling{Magnitude(3)}, Point{3, -3}},
	}

	for _, value := range data {
		out := value.s.Transform(value.in)
		if out != value.want {
			t.Errorf("%v.Transform(%v) = %v, want %v", value.s, value.in, out, value.want)
		}
	}
}
