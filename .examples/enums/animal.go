package enums

import "github.com/joelboim/gnum"

type (
	Color = gnum.Enum[color]
	color int
)

const (
	Red Color = iota
	Blue
	Green
)

var configColor = gnum.NewConfig(
	map[string]Color{
		"red":   Red,
		"Blue":  Blue,
		"GREEN": Green,
	})

func (color) Config() *gnum.Config { return configColor }
