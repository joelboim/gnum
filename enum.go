package gnum

import (
	"fmt"
	"strings"
)

type Enumer[T ~int] interface {
	~int
	Enums() []T
	Names() []string
	Parse(name string) (T, error)
	String() string
	Config() *Config
}

func Enums[T Enumer[T]]() []T {
	return T.Enums(-1)
}

type enumConfig[T ~int] interface {
	~int
	Config() *Config
}

type Enum[T enumConfig[T]] int

func Names[T Enumer[T]]() []string {
	return T.Names(-1)
}

func Parse[T Enumer[T]](name string) (T, error) {
	enum, err := T.Parse(-1, name)
	if err != nil {
		return -1, err
	}

	return enum, nil
}

func (e Enum[T]) Enums() []Enum[T] {
	var values []Enum[T]
	for _, value := range T.Config(-1).sortedValues {
		values = append(values, Enum[T](value))
	}
	return values
}

func (e Enum[T]) MarshalText() ([]byte, error) {
	return []byte(e.String()), nil
}

func (e *Enum[T]) UnmarshalText(text []byte) error {
	name := string(text)
	enum, err := e.Parse(name)
	if err != nil {
		return err
	}

	*e = enum
	return nil
}

func (e Enum[T]) Names() []string {
	return T.Config(-1).sortedNames
}

func (e Enum[T]) Parse(name string) (Enum[T], error) {
	config := T.Config(-1)
	normalizedName := config.parseNormalizationCallback(name)

	var (
		value int
		ok    bool
	)
	switch config.parseCaseInsensitive {
	case true:
		value, ok = config.caseInsensitiveNameToValue[strings.ToLower(normalizedName)]
	case false:
		value, ok = config.nameToValue[normalizedName]
	}

	if !ok {
		return -1, fmt.Errorf("`%s` isn't part of (%s)",
			normalizedName,
			config.allEnumsString)
	}

	return Enum[T](value), nil
}

func (e Enum[T]) String() string {
	name, ok := T.Config(-1).valueToName[int(e)]
	if !ok {
		panic(fmt.Sprintf("`%d` isn't part of `%T` mapping", e, e))
	}

	return name
}

func (e Enum[T]) Config() *Config {
	return T.Config(-1)
}
