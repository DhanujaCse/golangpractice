package coding1

import (
	"fmt"
)

type Shape interface {
	Area() int
	Perimeter() int
}
type Square struct {
	A int
}

func (s Square) Area() int {
	area := s.A * s.A
	fmt.Printf("Area of Square is %v\n", area)
	return s.A * s.A
}
func (s Square) Perimeter() int {

	peri := 4 * s.A
	fmt.Printf("Perimeter of Square is %v\n", peri)
	return 4 * s.A
}

type Rectangle struct {
	Length, Width int
}

func (r Rectangle) Area() int {
	area := r.Length * r.Width
	fmt.Printf("Area of rectangle is %v\n", area)
	return r.Length * r.Width
}

func (r Rectangle) Perimeter() int {
	peri := 2 * (r.Length + r.Width)
	fmt.Printf("Perimeter of Reactangle is %v\n", peri)
	return 2 * (r.Length + r.Width)

}

func PrintAreaAndPerimeter(s Shape) {
	s.Area()
	s.Perimeter()

}
