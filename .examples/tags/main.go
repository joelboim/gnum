package main

import (
	"fmt"
	"github.com/joelboim/gnum"
)

type (
	Color = gnum.Enum[struct {
		Red    color
		Blue   color `gnum:"value=3,name=b_l_u_e"`
		Green  color
		Yellow color
	}]
	color int
)

const (
	Red  Color = iota
	Blue Color = iota + 2
	Green
	Yellow
)

func main() {
	fmt.Println(gnum.Enums[Color]()) // [Red b_l_u_e Green Yellow]
	fmt.Printf("%d\n", Red)          // 0
	fmt.Printf("%d\n", Blue)         // 3
	fmt.Printf("%d\n", Yellow)       // 5
}
