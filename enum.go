package gnum

import (
	"fmt"
	"reflect"
	"strings"
)

const enumValueNotExistsErrorFormat = "`%d` isn't part of `%T` mapping"

// Enum uses T interface.Config for it's mapping of enum name to value.
type Enum[T interface {
	~int
	Config() *Config
}] int

// Config returns *Config, that in public scope, is useless,
// but it's a part of the Enumer interface for mapping and caching operation internally.
func (e Enum[T]) Config() *Config {
	return T.Config(-1)
}

// Enums returns a list of all Enum[T] declarations mapped to T
func (e Enum[T]) Enums() []Enum[T] {
	var values []Enum[T]
	for _, value := range T.Config(-1).sortedEnumValues {
		values = append(values, Enum[T](value))
	}
	return values
}

// Name returns the Enum[T] programmatic string representation.
func (e Enum[T]) Name() string {
	name, ok := T.Config(-1).enumValueToEnumName[int(e)]
	if !ok {
		panic(fmt.Sprintf(enumValueNotExistsErrorFormat, e, e))
	}

	return name
}

// Names returns all the Enum[T] programmatic string representations sorted by the enum values.
func (e Enum[T]) Names() []string {
	return T.Config(-1).sortedEnumNames
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
	config := T.Config(-1)
	processedName := config.parseCallback(name)

	var (
		value int
		ok    bool
	)
	if config.caseInsensitive {
		value, ok = config.enumNameLoweredToEnumValue[strings.ToLower(processedName)]
	} else {
		value, ok = config.enumNameToEnumValue[processedName]
	}

	if !ok {
		return -1, fmt.Errorf("`%s` isn't part of (%s)",
			processedName,
			config.joinedEnumNames)
	}

	return Enum[T](value), nil
}

// String returns the string representation of an Enum[T] value.
func (e Enum[T]) String() string {
	enumString, ok := T.Config(-1).enumValueToEnumString[int(e)]
	if !ok {
		panic(fmt.Sprintf(enumValueNotExistsErrorFormat, e, e))
	}

	return enumString
}

// Strings returns all the Enum[T] string representations sorted by the enum values.
func (e Enum[T]) Strings() []string {
	return T.Config(-1).sortedEnumStrings
}

// Type returns the underline T type.
func (e Enum[T]) Type() string {
	return reflect.TypeOf(*new(T)).Name()
}
