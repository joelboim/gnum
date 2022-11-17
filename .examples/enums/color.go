package enums

import "github.com/joelboim/gnum"

type (
	Animal = gnum.Enum[animal]
	animal int
)

const (
	Dog Animal = iota
	Cat
	Cow
)

var configAnimal = gnum.NewConfig(
	map[string]Animal{
		"dog": Dog,
		"cat": Cat,
		"cow": Cow,
	})

func (animal) Config() *gnum.Config { return configAnimal }
