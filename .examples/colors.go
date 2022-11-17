package main

import (
	"encoding/json"
	"fmt"
	"github.com/joelboim/gnum"
	"github.com/joelboim/gnum/.examples/enums"
)

func main() {
	fmt.Println(enums.Red, enums.Blue, enums.Green) // red Blue GREEN

	fmt.Println(gnum.Names[enums.Color]()) // [red Blue GREEN]

	fmt.Println(fmt.Sprintf("%T", gnum.Enums[enums.Color]())) // []gnum.Enum[gnum.color]

	red, _ := gnum.Parse[enums.Color]("red")
	fmt.Println(red) // red

	colorJson, _ := json.Marshal(struct{ Color enums.Color }{enums.Blue})
	fmt.Println(string(colorJson)) // {"Color":"Blue"}
}
