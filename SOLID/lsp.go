package main

import "fmt"

type Sized interface {
	GetWidth() int
	SetWidth(width int)
	GetHeight() int
	SetHeight(height int)
}

type Rectangle struct {
	width, height int
}

//     vvv !! POINTER
func (r *Rectangle) GetWidth() int {
	return r.width
}

func (r *Rectangle) SetWidth(width int) {
	r.width = width
}

func (r *Rectangle) GetHeight() int {
	return r.height
}

func (r *Rectangle) SetHeight(height int) {
	r.height = height
}

// modified LSP
// If a function(CalculateArea) takes an interface and works with a type T(Rectangle) that implements this interface
// any other structure(Square) that aggregates T should also be usable in that function.
// If not means you are violate the LSP
type Square struct {
	Rectangle
}

func (s *Square) SetWidth(width int) {
	s.width = width
	s.height = width
}

func (s *Square) SetHeight(height int) {
	s.width = height
	s.height = height
}

func CalculateArea(sized Sized) {
	width := sized.GetWidth()
	sized.SetHeight(10)
	expectedArea := 10 * width
	actualArea := sized.GetWidth() * sized.GetHeight()
	fmt.Print("Expected an area of ", expectedArea, ", but got ", actualArea, "\n")
}

// instead of aggregating a rectangle like Square type did, new Square2 have its own member size
type Square2 struct {
	size int // width, height
}

// represent the square as a rectangle
func (s *Square2) Rectangle() *Rectangle {
	return &Rectangle{s.size, s.size}
}

func main() {
	rc := &Rectangle{2, 3}
	CalculateArea(rc)

	// violate the lsp, composity Square type didn't work with interface function
	// because our SetHeight set both width and height but rectangle only set Height
	// athrough logically is correct
	sq := &Square{Rectangle{5, 5}}
	CalculateArea(sq)

	sq2 := &Square2{5}
	rc = sq2.Rectangle()
	CalculateArea(rc)
}
