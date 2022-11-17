package gnum

import (
	"fmt"
	"reflect"
	"strings"
)

// Enumer is an interface for using Enum instances as generics,
// e.g, `func foo[T Enummer[T]](enum T)` could do any Enum operations
// while preserving the original Enum type (T)
type Enumer[T ~int] interface {
	~int
	Enums() []T
	Names() []string
	Parse(name string) (T, error)
	String() string
	Config() *Config
	Type() string
}

type enumConfig[T ~int] interface {
	~int
	Config() *Config
}

// Enum provides all the enum functionalities for T
// using the underline T.Config() mapping.
type Enum[T enumConfig[T]] int

// Enums is a static function to handel all enums that implements Enumer[T] interface.
// It returns a list of all Enum[T] declarations mapped to T.
func Enums[T Enumer[T]]() []T {
	return T.Enums(-1)
}

// Names is a static function to handel all enums that implements Enumer[T] interface.
// It returns list of all Enum[T] string representations.
func Names[T Enumer[T]]() []string {
	return T.Names(-1)
}

// Type is a static function to handel all enums that implements Enumer[T] interface.
// It returns the underline type name.
func Type[T Enumer[T]]() string {
	return T.Type(-1)
}

// Parse is a static function to handel all enums that implements Enumer[T] interface.
// It will try to parse the given name with the underline Enum.Parse implementation.
func Parse[T Enumer[T]](name string) (T, error) {
	enum, err := T.Parse(-1, name)
	if err != nil {
		return -1, err
	}

	return enum, nil
}

// Enums returns a list of all Enum[T] declarations mapped to T
func (e Enum[T]) Enums() []Enum[T] {
	var values []Enum[T]
	for _, value := range T.Config(-1).sortedValues {
		values = append(values, Enum[T](value))
	}
	return values
}

// MarshalText implements the TextMarshaler interface for T.
func (e Enum[T]) MarshalText() ([]byte, error) {
	return []byte(e.String()), nil
}

// UnmarshalText implements the TextUnmarshaler interface for T.
func (e *Enum[T]) UnmarshalText(text []byte) error {
	enum, err := e.Parse(string(text))
	if err != nil {
		return err
	}

	*e = enum
	return nil
}

// Names returns all the Enum[T] string representations sorted by the enum values.
func (e Enum[T]) Names() []string {
	return T.Config(-1).sortedNames
}

// Parse tries to parse a name based on the underline T.Config mapping.
// If OptionParseCaseInsensitive(true) is set, Parse will use the lowered case mapping instead.
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

// String returns the string representation of Enum[T] value.
func (e Enum[T]) String() string {
	name, ok := T.Config(-1).valueToName[int(e)]
	if !ok {
		panic(fmt.Sprintf("`%d` isn't part of `%T` mapping", e, e))
	}

	return name
}

// Config returns *Config, that in public scope, is useless,
// but it's a part of the Enumer interface for mapping and caching operation internally.
func (e Enum[T]) Config() *Config {
	return T.Config(-1)
}

// Type returns the underline T type.
func (e Enum[T]) Type() string {
	return reflect.TypeOf(*new(T)).Name()
}
