package main

import "fmt"

func main() {
	rc := &Rectangle{2, 3}
	UseIt(rc)

	sq := NewSquare(5)
	UseIt(sq)

	sq2 := Square2{5}
	final := sq2.Rectangle()
	UseIt(final)

}

type Sized interface {
	SetWidth(width int)
	GetWidth() int
	SetHeight(height int)
	GetHeight() int
}

type Rectangle struct {
	width  int
	height int
}

type Square struct {
	Rectangle
}

func (r *Rectangle) SetWidth(width int) {
	r.width = width
}

func (r *Rectangle) GetWidth() int {
	return r.width
}

func (r *Rectangle) SetHeight(height int) {
	r.height = height
}

func (r *Rectangle) GetHeight() int {
	return r.height
}

//------------------------------------------------------------

func NewSquare(value int) Sized {
	sq := &Square{}
	sq.width = value
	sq.height = value
	return sq
}

func (s *Square) SetWidth(width int) { // It breaks lsp
	s.width = width
	s.height = width
}

func (s *Square) SetHeight(height int) { //It breaks lsp
	s.height = height
	s.width = height
}

func (s *Square) GetHeight() int {
	return s.height
}

func (s *Square) GetWidth() int {
	return s.width
}

type Square2 struct {
	size int
}

func (s *Square2) Rectangle() *Rectangle {
	return &Rectangle{s.size, s.size}
}

func UseIt(sized Sized) {
	width := sized.GetWidth()
	sized.SetHeight(10)
	expectedArea := 10 * width
	actualArea := sized.GetWidth() * sized.GetHeight()
	fmt.Print("Expected an area of ", expectedArea,
		", but got ", actualArea, "\n")
}
