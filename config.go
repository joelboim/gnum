package gnum

import (
	"fmt"
	"github.com/joelboim/gnum/infra"
	"strings"
)

type Config struct {
	nameToValue                map[string]int
	caseInsensitiveNameToValue map[string]int
	valueToName                map[int]string
	sortedNames                []string
	sortedValues               []int
	allEnumsString             string
	stringCallback             func(value string) string
	parseNormalizationCallback func(value string) string
	parseCaseInsensitive       bool
}

type ConfigOption func(*Config)

// NewConfig return a new *Config instance with all of its private fields populated based on the provided nameToValue
// and all the options applied to it.
// e.g. NewConfig(map[string]Color{"red": Red}, OptionParseCaseInsensitive(true))
func NewConfig[T ~int](nameToValue map[string]T, options ...ConfigOption) *Config {
	config := &Config{
		make(map[string]int),
		make(map[string]int),
		make(map[int]string),
		make([]string, 0, len(nameToValue)),
		make([]int, 0),
		"",
		func(value string) string { return value },
		func(value string) string { return value },
		false,
	}

	for _, option := range options {
		option(config)
	}

	for name, value := range nameToValue {
		valueConverted := int(value)
		nameConverted := config.stringCallback(name)

		config.nameToValue[name] = valueConverted

		if duplicateName, ok := config.valueToName[valueConverted]; ok {
			panic(fmt.Sprintf("`%s` and `%s` have the same value", duplicateName, name))
		}

		config.valueToName[valueConverted] = nameConverted
		insertedSortedIndex := infra.InsertToSortedSlice(
			&config.sortedValues,
			valueConverted)
		infra.InsertToSliceByIndex(
			&config.sortedNames,
			name,
			uint64(insertedSortedIndex))
		config.caseInsensitiveNameToValue[strings.ToLower(name)] = valueConverted
	}

	config.allEnumsString = strings.Join(config.sortedNames, ", ")

	return config
}

// OptionParseCaseInsensitive - is set to true, Enum.Parse and Enum.UnmarshalText respectively,
// will ignore the input case when parsing.
func OptionParseCaseInsensitive(caseInsensitive bool) func(c *Config) {
	return func(c *Config) {
		c.parseCaseInsensitive = caseInsensitive
	}
}

// OptionParseNormalizationCallback will be applied for each Enum.Parse call and Enum.UnmarshalText respectively.
func OptionParseNormalizationCallback(callback func(value string) string) func(c *Config) {
	return func(c *Config) {
		c.parseNormalizationCallback = callback
	}
}
