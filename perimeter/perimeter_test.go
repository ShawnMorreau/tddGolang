package perimeter

import "testing"

func TestPerimeter(t *testing.T) {
	perimeterTests := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{5, 5}, 20.0},
		{Circle{5}, 31.400000000000002},
		{Triangle{5, 2}, 15},
	}

	for _, test := range perimeterTests {
		got := test.shape.Perimeter()
		if got != test.want {
			t.Errorf("want %g, got %g", test.want, got)
		}
	}
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, want: 72.0},
		{name: "Circle", shape: Circle{Radius: 10}, want: 314},
		{name: "Triangle", shape: Triangle{Base: 5, Height: 2}, want: 5},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {

			got := tt.shape.Area()
			if got != tt.want {
				t.Errorf("%v got %g want %g", tt.shape, got, tt.want)
			}
		})

	}
}
