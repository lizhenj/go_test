package main

import "fmt"

type Shaper interface {
	Area() float32
}

type Square struct {
	side float32
}

func (sq Square) Area() float32 {
	return sq.side * sq.side
}

type Rectangke struct {
	length, width float32
}

func (r Rectangke) Area() float32 {
	return r.length * r.width
}

func main13() {
	r := Rectangke{5, 3}
	q := Square{5}
	shapes := []Shaper{&r, &q}
	for n := range shapes {
		fmt.Println(shapes[n].Area())
	}
}
