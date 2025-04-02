package structures

import "math"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width float64
	Height float64
}

type Circle struct {
	radius float64
}


func (rect Rectangle) Area() float64 {
	return rect.Width * rect.Height;
}

func (c Circle) Area() float64{
	return math.Pi * c.radius * c.radius;
}

func Area(s Shape) float64 {
	return s.Area();
}
