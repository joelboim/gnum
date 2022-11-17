package main

import (
	"fmt"
	"github.com/joelboim/gnum"
	"github.com/joelboim/gnum/.examples/enums"
)

func foo[T gnum.Enumer[T]](enum T) {
	fmt.Println(
		enum.Names(),
		enum.Type())
}

func main() {
	foo(enums.Dog) // [dog cat cow] animal
	foo(enums.Red) // [red Blue GREEN] color
}
