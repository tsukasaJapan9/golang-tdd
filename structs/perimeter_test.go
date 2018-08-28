package perimeter

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	assertMsg := func(t *testing.T, got, want float64) {
		t.Helper()
		if got != want {
			t.Errorf("got:%.2f, want:%.2f", got, want)
		}
	}

	t.Run("rect perimeter", func(t *testing.T) {
		rect := Rect{10.0, 10.0}
		got := Perimeter(rect)
		want := 40.0
		assertMsg(t, got, want)
	})
}

func TestArea(t *testing.T) {
	areaTest := []struct{
		shape Shape
		want float64
	} {
		{shape:Rect{12, 6}, want:72.0},
		{shape:Circle{10}, want:314.1592653589793},
		{shape:Tri{12, 6}, want:36.0},
	}

	for _, tt := range areaTest {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("got:%f, want:%f", got, tt.want)
		}
	}
}
