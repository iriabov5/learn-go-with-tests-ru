package geometry

import (
	"fmt"
	"testing"
)

func TestRectangle(t *testing.T) {
	rectangle := Rectangle{Width: 10, Height: 5}
	area := rectangle.Area()
	expectedArea := 50.0

	if area != expectedArea {
		t.Errorf("Expected area %f, but got %f", expectedArea, area)
	}

	perimeter := rectangle.Perimeter()
	expectedPerimeter := 30.0

	if perimeter != expectedPerimeter {
		t.Errorf("Expected perimeter %f, but got %f", expectedPerimeter, perimeter)
	}
}

func TestCircle(t *testing.T) {
	circle := Circle{Radius: 10}
	area := circle.Area()
	expectedArea := 314.1592653589793

	if area != expectedArea {
		t.Errorf("Expected area %f, but got %f", expectedArea, area)
	}

	perimeter := circle.Perimeter()
	expectedPerimeter := 62.83185307179586

	if perimeter != expectedPerimeter {
		t.Errorf("Expected perimeter %f, but got %f", expectedPerimeter, perimeter)
	}
}

func TestTriangle(t *testing.T) {
	triangle := Triangle{Base: 10, Height: 5}
	area := triangle.Area()
	expectedArea := 25.0

	if area != expectedArea {
		t.Errorf("Expected area %f, but got %f", expectedArea, area)
	}

	perimeter := triangle.Perimeter()
	expectedPerimeter := 26.18033988749895 // Base + Height + Hypotenuse

	if perimeter != expectedPerimeter {
		t.Errorf("Expected perimeter %f, but got %f", expectedPerimeter, perimeter)
	}
}

func TestAreaWithInterface(t *testing.T) {
	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, hasArea: 72.0},
		{name: "Circle", shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, hasArea: 36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			area := tt.shape.Area()
			if area != tt.hasArea {
				t.Errorf("%#v: Expected area %f, but got %f", tt.shape, tt.hasArea, area)
			}
		})
	}
}

func ExampleRectangle_Area() {
	rect := Rectangle{Width: 3, Height: 4}
	area := rect.Area()
	fmt.Println(area)
	// Output: 12
}