package gnum

import (
	"fmt"
	"github.com/joelboim/gnum/infra"
	"reflect"
	"strings"
)

var globalConfig = &config{}

type config struct {
	caseInsensitive bool
	parseCallback   func(value string) string
	stringCallback  func(value string) string
}

type enumMetadata struct {
	enumNameLoweredToEnumValue map[string]int
	enumNameToEnumValue        map[string]int
	enumValueToEnumName        map[int]string
	enumValueToEnumString      map[int]string
	joinedEnumNames            string
	sortedEnumNames            []string
	sortedEnumStrings          []string
	sortedEnumValues           []int
}

// Option callback function that sets specific value on an *config instance.
type Option func(config *config)

// SetOptions Sets multiple Option on the global scope.
func SetOptions(options ...Option) {
	for _, option := range options {
		option(globalConfig)
	}
}

// StringCallback will be applied for each Enum.String call and Enum.Strings.
func StringCallback(callback func(enumName string) string) Option {
	return func(c *config) {
		c.stringCallback = callback
	}
}

// CaseInsensitive - when set to true, Enum.Parse and Enum.UnmarshalText,
// will ignore the string case when parsing.
func CaseInsensitive(caseInsensitive bool) Option {
	return func(c *config) {
		c.caseInsensitive = caseInsensitive
	}
}

// ParseCallback will be applied for each Enum.Parse call and Enum.UnmarshalText.
func ParseCallback(callback func(value string) string) Option {
	return func(c *config) {
		c.parseCallback = callback
	}
}

// newEnumMetadata return a new *enumMetadata, based on the provided T
// and applies the globalConfig.
func newEnumMetadata[T any]() *enumMetadata {
	enumNameToEnumValue := getEnumNameToEnumValue[T]()
	metadata := &enumMetadata{
		enumNameLoweredToEnumValue: make(map[string]int),
		enumNameToEnumValue:        enumNameToEnumValue,
		enumValueToEnumName:        make(map[int]string),
		enumValueToEnumString:      make(map[int]string),
		sortedEnumNames:            make([]string, 0, len(enumNameToEnumValue)),
		sortedEnumStrings:          make([]string, 0, len(enumNameToEnumValue)),
		sortedEnumValues:           make([]int, 0, len(enumNameToEnumValue)),
	}

	for enumName, enumValue := range enumNameToEnumValue {
		if duplicateEnumName, ok := metadata.enumValueToEnumName[enumValue]; ok {
			panic(fmt.Sprintf(
				"`%s` and `%s` have the same value",
				duplicateEnumName,
				enumName))
		}

		enumString := enumName
		if globalConfig.stringCallback != nil {
			enumString = globalConfig.stringCallback(enumName)
		}

		metadata.enumValueToEnumName[enumValue] = enumName
		metadata.enumValueToEnumString[enumValue] = enumString

		metadata.enumNameLoweredToEnumValue[strings.ToLower(enumName)] = enumValue

		sortedIndex := infra.InsertToSortedSlice(
			&metadata.sortedEnumValues,
			enumValue)
		infra.InsertToSliceByIndex(
			&metadata.sortedEnumNames,
			enumName,
			sortedIndex)
		infra.InsertToSliceByIndex(
			&metadata.sortedEnumStrings,
			enumString,
			sortedIndex)
	}

	metadata.joinedEnumNames = strings.Join(metadata.sortedEnumNames, ", ")

	return metadata
}

// getEnumNameToEnumValue crates a mapping of enum names to enum values based on the T and its tags.
func getEnumNameToEnumValue[T any]() map[string]int {
	enumNameToEnumValue := make(map[string]int)
	nextEnumValue := 0
	for _, field := range reflect.VisibleFields(reflect.TypeOf(*new(T))) {
		enumTag := newEnumTag(field)
		if enumTag == nil {
			enumNameToEnumValue[field.Name] = nextEnumValue
			nextEnumValue += 1
			continue
		}

		enumName := infra.GetPointerValue(enumTag.Name, field.Name)
		if _, ok := enumNameToEnumValue[enumName]; ok {
			panic(fmt.Sprintf("duplicate enum name - `%s`", enumName))
		}

		if enumTag.Value == nil {
			enumNameToEnumValue[enumName] = nextEnumValue
			nextEnumValue += 1
			continue
		}

		enumNameToEnumValue[enumName] = *enumTag.Value
		nextEnumValue += *enumTag.Value
	}

	return enumNameToEnumValue
}
