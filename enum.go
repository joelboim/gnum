package gnum

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

const enumValueNotExistsErrorFormat = "`%d` isn't part of `%T` mapping"

var cache = newEnumCache()

// Enum uses T struct definition for it's mapping of enum name to value.
type Enum[T any] int

// Enums returns a list of all Enum[T] declarations mapped to T
func (e Enum[T]) Enums() []Enum[T] {
	var values []Enum[T]
	for _, value := range e.getConfig().sortedEnumValues {
		values = append(values, Enum[T](value))
	}

	return values
}

// Name returns the Enum[T] programmatic string representation.
func (e Enum[T]) Name() string {
	name, ok := e.getConfig().enumValueToEnumName[int(e)]
	if !ok {
		panic(fmt.Sprintf(enumValueNotExistsErrorFormat, e, e))
	}

	return name
}

// Names returns all the Enum[T] programmatic string representations sorted by the enum values.
func (e Enum[T]) Names() []string {
	return e.getConfig().sortedEnumNames
}

// MarshalText implements the TextMarshaler interface for T.
func (e Enum[T]) MarshalText() ([]byte, error) {
	return []byte(e.Name()), nil
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

// Parse tries to parse an enum name based on the underline enum name to enum value mapping.
// If CaseInsensitive(true) is set, Parse will use the lowered case name to value mapping instead.
func (e Enum[T]) Parse(name string) (Enum[T], error) {
	config := e.getConfig()

	if globalConfig.parseCallback != nil {
		name = globalConfig.parseCallback(name)
	}

	var (
		value int
		ok    bool
	)
	if globalConfig.caseInsensitive {
		value, ok = config.enumNameLoweredToEnumValue[strings.ToLower(name)]
	} else {
		value, ok = config.enumNameToEnumValue[name]
	}

	if !ok {
		return -1, errors.New("`" + name + "`" + " isn't part of [" + config.joinedEnumNames + "]")
	}

	return Enum[T](value), nil
}

// String returns the string representation of an Enum[T] value.
func (e Enum[T]) String() string {
	enumString, ok := e.getConfig().enumValueToEnumString[int(e)]
	if !ok {
		panic(fmt.Sprintf(enumValueNotExistsErrorFormat, e, e))
	}

	return enumString
}

// Strings returns all the Enum[T] string representations sorted by the enum values.
func (e Enum[T]) Strings() []string {
	return e.getConfig().sortedEnumStrings
}

// Type returns the underline T type.
func (e Enum[T]) Type() string {
	return reflect.TypeOf(*new(T)).Field(0).Type.Name()
}

func (e Enum[T]) getConfig() *enumMetadata {
	enumType := reflect.TypeOf(e)
	if config_, ok := cache.Get(enumType); ok {
		return config_
	}

	newConfig := newEnumMetadata[T]()
	cache.Set(enumType, newConfig)

	return newConfig
}
