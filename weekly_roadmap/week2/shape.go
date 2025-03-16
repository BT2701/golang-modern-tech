package week2


type Shape interface {
	Area() float32
	Perimeter() float32
}

type Rectangle struct {
	Width  float32
	Height float32
}

type Circle struct {
	Radius float32
}