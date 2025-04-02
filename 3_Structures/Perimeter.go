package structures

import "math"

type Shapes2 interface{
	Perimeter() float64
}

func (rect Rectangle) Perimeter() float64{
	return 2*(rect.Height + rect.Width);
}

func (circ Circle) Perimeter() float64{
	return 2 * math.Pi * circ.radius;
}

func Perimeter(s Shapes2) float64{
	return s.Perimeter();
}

