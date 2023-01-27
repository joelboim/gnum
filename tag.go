package gnum

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
)

var (
	enumTagValuePattern = regexp.MustCompile(`value=(?P<value>-?\d+)(,|$)`)
	enumTagNamePattern  = regexp.MustCompile(`name=(?P<name>.+)(,|$)`)
)

type enumTag struct {
	Name  *string
	Value *int
}

func newEnumTag(field reflect.StructField) *enumTag {
	rawFieldTag, ok := field.Tag.Lookup("gnum")
	if !ok {
		return nil
	}

	enumTag := &enumTag{
		Name:  getEnumName(rawFieldTag),
		Value: getEnumValue(rawFieldTag),
	}

	if enumTag.Name == nil &&
		enumTag.Value == nil {
		panic(fmt.Sprintf("enum definition not found - `%s`", rawFieldTag))
	}

	return enumTag

}

func getEnumName(rawFieldTag string) *string {
	enumName := getTagValue(
		enumTagNamePattern,
		rawFieldTag)
	if enumName == nil {
		return nil
	}

	if *enumName == "" {
		panic(fmt.Sprintf("enum name can't be empty - `%s`", rawFieldTag))
	}

	return enumName

}

func getEnumValue(rawFieldTag string) *int {
	enumValue := getTagValue(
		enumTagValuePattern,
		rawFieldTag)
	if enumValue == nil {
		return nil
	}

	if *enumValue == "" {
		panic(fmt.Sprintf("enum value can't be empty - `%s`", rawFieldTag))
	}

	enumValueInt, err := strconv.Atoi(*enumValue)
	if err != nil {
		panic(err)
	}

	return &enumValueInt
}

func getTagValue(pattern *regexp.Regexp, rawFieldTag string) *string {
	submatches := pattern.FindAllStringSubmatch(rawFieldTag, -1)
	if len(submatches) == 0 {
		return nil
	}

	if len(submatches) > 1 {
		panic(
			fmt.Sprintf(
				"invalid number of enum %ss - `%s`",
				pattern.SubexpNames()[0],
				rawFieldTag))
	}

	return &submatches[0][1]
}
