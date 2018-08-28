package perimeter

import "math"

// Shape is interface
type Shape interface {
	Area() float64
}

// Rect is one of impl of shape
type Rect struct {
	width float64
	height float64
}

// Circle is one of impl of shape
type Circle struct {
	radius float64
}

// Tri is one of impl of shape
type Tri struct {
	base float64
	height float64
}

// Perimeter is calc function
func Perimeter(r Rect) float64 {
	return 2 * (r.width + r.height)
}

// Area is calc function
func (r Rect)Area() float64 {
	return r.width * r.height
}

// Area is calc function
func (c Circle)Area() float64 {
	return math.Pi * c.radius * c.radius
}

// Area is calc function
func (t Tri)Area() float64 {
	return t.base * t.height / 2.0
}