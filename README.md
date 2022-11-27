[![GitHub tag (latest by date)](https://img.shields.io/github/v/tag/joelboim/gnum)](https://github.com/joelboim/gnum/tags)
![Test](https://github.com/joelboim/gnum/actions/workflows/test.yml/badge.svg)
![go version](https://img.shields.io/badge/go-%3E%3D18-blue)
[![Go Reference](https://pkg.go.dev/badge/github.com/joelboim/gnum.svg)](https://pkg.go.dev/github.com/joelboim/gnum)
[![GoReportCard](https://goreportcard.com/badge/github.com/joelboim/gnum)](https://goreportcard.com/report/github.com/joelboim/gnum)

![Fancy logo](docs/dark.png#gh-dark-mode-only)
![Fancy logo](docs/light.png#gh-light-mode-only)

Enum for GO. **without** code generation. works with **const**s.

## Y Use gnum:grey_question:

* You don't need code generation.
* You can assign it to const.
* It's fast.
* You can use generics for processing any Enum.

## Benchmarks:dash:

| name   \     time/op | gnum        | go-enum      | enumer       |
|----------------------|-------------|--------------|--------------|
| Parse                | 1.92µs ± 6% | 1.82µs ±16%  |              |
| ParseCaseInsensitive | 1.60µs ± 7% | 1.67µs ±10%  | 1.39µs ± 2%  |
| String               | 11.4ns ± 9% | 52.9ns ±14%  | 52.2ns ±11%  |
| Names                | 96.4ns ± 7% | 172.7ns ±17% | 167.3ns ± 6% |
| MarshalText          | 84.2ns ± 3% | 75.0ns ±13%  | 17.9ns ±11%  |
| Enums                | 94.8ns ±18% |              | 35.5ns ± 3%  |

## Getting Started

```bash
go get github.com/joelboim/gnum
````

# Example

First lets declare an enum type:

```go
package main

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

var config = gnum.NewConfig(
	map[string]Color{
		"red":   Red,
		"Blue":  Blue,
		"GREEN": Green,
	})

func (color) Config() *gnum.Config { return config }
```

Now we can use it like other languages Enums:

```go 
func main() {
	fmt.Println(Red, Blue, Green) // red Blue GREEN

	fmt.Println(gnum.Names[Color]()) // [red Blue GREEN]

	fmt.Println(fmt.Sprintf("%T", gnum.Enums[Color]())) // []gnum.Enum[gnum.color]

	red, _ := gnum.Parse[Color]("red")
	fmt.Println(red) // red

	colorJson, _ := json.Marshal(struct{ Color Color }{Blue})
	fmt.Println(string(colorJson)) // {"Color":"Blue"}
}
```

Can be also used in a generic context:

```go
func foo[T gnum.Enumer[T]](enum T) {
	fmt.Println(
		enum.Names(),
		enum.Type())
}

func main() {
	foo(Dog) // [dog cat cow] animal
	foo(Red) // [red Blue GREEN] color
}
```

