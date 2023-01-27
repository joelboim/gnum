package main

import (
	"fmt"
	"github.com/joelboim/gnum"
)

type (
	Color = gnum.Enum[struct {
		Green color
	}]
	color int

	Shape = gnum.Enum[struct {
		Square shape
	}]
	shape int
)

const (
	Red    Color = iota
	Square Shape = iota
)

func foo[T gnum.Enumer[T]](enum T) {
	fmt.Println(
		enum.Names(),
		enum.Type())
}

func main() {
	foo(Red)    // [dog cat cow] animal
	foo(Square) // [Square] shape
}
