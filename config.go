package gnum

import (
	"fmt"
	"github.com/joelboim/gnum/infra"
	"strings"
)

type Config struct {
	caseInsensitive            bool
	enumNameLoweredToEnumValue map[string]int
	enumNameToEnumValue        map[string]int
	enumValueToEnumName        map[int]string
	enumValueToEnumString      map[int]string
	joinedEnumNames            string
	parseCallback              func(value string) string
	sortedEnumNames            []string
	sortedEnumStrings          []string
	sortedEnumValues           []int
	stringCallback             func(value string) string
}

type ConfigOption func(*Config)

// NewConfig return a new *Config instance with all of its private fields populated based on the provided enumNameToEnumValue
// and all the options applied to it.
// e.g. NewConfig(map[string]Color{"red": Red}, CaseInsensitive(true))
func NewConfig[T ~int](
	enumNameToEnumValue map[string]T,
	options ...ConfigOption) *Config {
	config := &Config{
		false,
		make(map[string]int),
		make(map[string]int),
		make(map[int]string),
		make(map[int]string),
		"",
		func(value string) string { return value },
		make([]string, 0, len(enumNameToEnumValue)),
		make([]string, 0, len(enumNameToEnumValue)),
		make([]int, 0, len(enumNameToEnumValue)),
		func(value string) string { return value },
	}

	for _, option := range options {
		option(config)
	}

	for enumName, enumValue := range enumNameToEnumValue {
		enumValueInt := int(enumValue)
		if duplicateEnumName, ok := config.enumValueToEnumName[enumValueInt]; ok {
			panic(fmt.Sprintf(
				"`%s` and `%s` have the same value",
				duplicateEnumName,
				enumName))
		}

		enumString := config.stringCallback(enumName)
		config.enumValueToEnumName[enumValueInt] = enumName
		config.enumValueToEnumString[enumValueInt] = enumString

		config.enumNameToEnumValue[enumName] = enumValueInt
		config.enumNameLoweredToEnumValue[strings.ToLower(enumName)] = enumValueInt

		sortedIndex := infra.InsertToSortedSlice(
			&config.sortedEnumValues,
			enumValueInt)
		infra.InsertToSliceByIndex(
			&config.sortedEnumNames,
			enumName,
			sortedIndex)
		infra.InsertToSliceByIndex(
			&config.sortedEnumStrings,
			enumString,
			sortedIndex)
	}

	config.joinedEnumNames = strings.Join(config.sortedEnumNames, ", ")

	return config
}

// StringCallback will be applied for each Enum.String call and Enum.Strings respectively.
func StringCallback(callback func(enumName string) string) func(c *Config) {
	return func(c *Config) {
		c.stringCallback = callback
	}
}

// CaseInsensitive - when set to true, Enum.Parse and Enum.UnmarshalText,
// will ignore the input case when parsing.
func CaseInsensitive(caseInsensitive bool) func(c *Config) {
	return func(c *Config) {
		c.caseInsensitive = caseInsensitive
	}
}

// ParseCallback will be applied for each Enum.Parse call and Enum.UnmarshalText respectively.
func ParseCallback(callback func(value string) string) func(c *Config) {
	return func(c *Config) {
		c.parseCallback = callback
	}
}
