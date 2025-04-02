package structures

import (
	"math"
	"testing"
)


func TestArea(t *testing.T){
	t.Run("area of a Rectangle", func(t * testing.T){
		rect := Rectangle{
			Width: 50,
			Height: 4,
		}

		got := Area(rect)
		want := 200.00

		checkIfCorrect(t, got, want)

	})
	t.Run("area of a circle", func(t * testing.T){
		circ := Circle{13}

		got := Area(circ)
		want := math.Pi*circ.radius*circ.radius

		checkIfCorrect(t, got, want)

	})
}

func TestPerimeter(t *testing.T) {

	t.Run("perimeter of rectangle", func(t *testing.T){
		rect := Rectangle{
			Width: 10,
			Height: 40,
		}

		got := Perimeter(rect)
		want := 2 * (rect.Height + rect.Width)

		checkIfCorrect(t, got, want)
	})
	t.Run("perimeter of circle", func(t *testing.T){
		circ := Circle{
			radius: 10,
		}

		got := Perimeter(circ)
		want := 2 * math.Pi * circ.radius

		checkIfCorrect(t, got, want)
	})
}

func checkIfCorrect(t *testing.T, got, want float64){
	t.Helper()
	if(got != want) {
		t.Errorf("got %.2f, want %.2f", got, want)
	}
}