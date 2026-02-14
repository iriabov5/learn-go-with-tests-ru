package geometry

import "math"

// Rectangle представляет собой прямоугольник с шириной и высотой
type Rectangle struct {
	Width  float64
	Height float64
}

// Area вычисляет площадь прямоугольника
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Perimeter вычисляет периметр прямоугольника
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Circle представляет собой круг с радиусом
type Circle struct {
	Radius float64
}

// Area вычисляет площадь круга
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Perimeter вычисляет периметр круга
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// Triangle представляет собой треугольник с основанием и высотой
type Triangle struct {
	Base   float64
	Height float64
}

// Area вычисляет площадь треугольника
func (t Triangle) Area() float64 {
	return 0.5 * t.Base * t.Height
}

// Perimeter вычисляет периметр треугольника
func (t Triangle) Perimeter() float64 {
	// Для простоты мы просто возвращаем основание + высота + гипотенуза
	// предполагая, что это прямоугольный треугольник
	hypotenuse := math.Sqrt(t.Base*t.Base + t.Height*t.Height)
	return t.Base + t.Height + hypotenuse
}

// Shape - это интерфейс для геометрических фигур
type Shape interface {
	Area() float64
	Perimeter() float64
}