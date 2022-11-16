package main

import (
	"encoding/json"
	"fmt"
	"github.com/joelboim/gnum"
)

type (
	Color = gnum.Enum[color]
	color int
)

const (
	Red Color = iota
	Blue
	Green
)

var config = gnum.NewConfig(
	map[string]Color{
		"red":   Red,
		"Blue":  Blue,
		"GREEN": Green,
	})

func (color) Config() *gnum.Config { return config }

func main() {
	fmt.Println(Red, Blue, Green) // red Blue GREEN

	fmt.Println(gnum.Names[Color]()) // [red Blue GREEN]

	fmt.Println(fmt.Sprintf("%T", gnum.Enums[Color]())) // []gnum.Enum[gnum.color]

	red, _ := gnum.Parse[Color]("red")
	fmt.Println(red) // red

	colorJson, _ := json.Marshal(struct{ Color Color }{Blue})
	fmt.Println(string(colorJson)) // {"Color":"Blue"}
}
