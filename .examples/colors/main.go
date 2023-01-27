package main

import (
	"encoding/json"
	"fmt"
	"github.com/joelboim/gnum"
)

type (
	Color = gnum.Enum[struct {
		Red,
		Blue,
		Green color
	}]
	color int
)

const (
	Red Color = iota
	Blue
	Green
)

func main() {
	fmt.Println(Red, Blue, Green) // Red Blue Green

	fmt.Println(gnum.Names[Color]()) // [Red Blue Green]

	fmt.Println(fmt.Sprintf("%T", gnum.Enums[Color]())) // []gnum.Enum[struct { Red main.color; Blue main.color; Green main.color }]

	_, err := gnum.Parse[Color]("red")
	fmt.Println(err) // error

	colorJson, _ := json.Marshal(struct{ Color Color }{Blue})
	fmt.Println(string(colorJson)) // {"Color":"Blue"}
}
