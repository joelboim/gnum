package gnum

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

const (
	red testColor = iota
	blue
	yellow
	green testColor = -1
)

type color int

type testColor = Enum[color]

var testColorConfig = NewConfig(
	map[string]testColor{
		"RedSuffix":    red,
		"BlueSuffix":   blue,
		"YellowSuffix": yellow,
		"GreenSuffix":  green},
	ParseCallback(func(value string) string {
		return value + "Suffix"
	}),
	StringCallback(func(enumName string) string {
		return "Prefix" + enumName
	}))

func (color) Config() *Config {
	return testColorConfig
}

func TestNewConfig_OnValidConfig_ThenReturnConfig(t *testing.T) {
	// Arrange
	// Act
	// Assert
	assert.NotPanics(t, func() {
		NewConfig(
			map[string]testColor{
				"RedSuffix":    red,
				"BlueSuffix":   blue,
				"YellowSuffix": yellow,
				"GreenSuffix":  green})
	})
}

func TestNewConfig_OnDuplicateEnumValues_ThenPanic(t *testing.T) {
	// Arrange
	// Act
	// Assert
	assert.Panics(t, func() {
		NewConfig(
			map[string]testColor{
				"RedSuffix":    red,
				"BlueSuffix":   red,
				"YellowSuffix": yellow,
				"GreenSuffix":  green})
	})
}

func TestReceiverString_OnCustomConfig_ThenReturnString(t *testing.T) {
	// Arrange
	// Act
	actualString := red.String()

	// Assert
	assert.Equal(t, "PrefixRedSuffix", actualString)
}

func TestReceiverStrings_OnCustomConfig_ThenReturnStrings(t *testing.T) {
	// Arrange
	// Act
	actualStrings := red.Strings()

	// Assert
	assert.Equal(t, []string{"PrefixGreenSuffix", "PrefixRedSuffix", "PrefixBlueSuffix", "PrefixYellowSuffix"}, actualStrings)
}

func TestReceiverName_OnCustomConfig_ThenReturnName(t *testing.T) {
	// Arrange
	// Act
	actualName := red.Name()

	// Assert
	assert.Equal(t, "RedSuffix", actualName)
}

func TestReceiverNames_OnCustomConfigAndMultipleEnums_ThenReturnDifferentNames(t *testing.T) {
	// Arrange
	// Act
	actualNames := red.Names()

	// Assert
	assert.Equal(t, []string{"GreenSuffix", "RedSuffix", "BlueSuffix", "YellowSuffix"}, actualNames)
}

func TestReceiverParse_OnCustomConfigAndExistingEnumName_ThenReturnEnum(t *testing.T) {
	// Arrange
	// Act
	actualEnum, err := red.Parse("Blue")
	require.NoError(t, err)

	// Assert
	assert.Equal(t, blue, actualEnum)
}

func TestReceiverParse_OnCustomConfigAndNonExistingEnumName_ThenReturnError(t *testing.T) {
	// Arrange
	// Act
	_, err := red.Parse("blue")

	// Assert
	assert.Error(t, err)
}

func TestReceiverParse_OnCaseInsensitiveAndExists_ThenReturnEnum(t *testing.T) {
	// Arrange
	defer func(revertedConfig *Config) {
		testColorConfig = revertedConfig
	}(testColorConfig)

	testColorConfig = NewConfig(
		map[string]testColor{
			"Red":    red,
			"Blue":   blue,
			"Yellow": yellow,
			"Green":  green},
		CaseInsensitive(true))

	// Act
	actualEnum, err := red.Parse("BLUE")
	require.NoError(t, err)

	// Assert
	assert.Equal(t, blue, actualEnum)
}

func TestReceiverParse_OnCaseInsensitiveAndMissing_ThenReturnError(t *testing.T) {
	// Arrange
	defer func(revertedConfig *Config) {
		testColorConfig = revertedConfig
	}(testColorConfig)

	testColorConfig = NewConfig(
		map[string]testColor{
			"Red":    red,
			"Blue":   blue,
			"Yellow": yellow,
			"Green":  green},
		CaseInsensitive(true))

	// Act
	_, err := red.Parse("NOT_COLOR")

	// Assert
	assert.Error(t, err)
}

func TestReceiverEnums_OnCustomConfigAndMultipleEnums_ThenReturnAll(t *testing.T) {
	// Arrange
	// Act
	actualEnums := red.Enums()

	// Assert
	assert.Equal(t, []testColor{green, red, blue, yellow}, actualEnums)
}

func TestReceiverMarshalText_OnCustomConfig_ThenReturnName(t *testing.T) {
	// Arrange
	// Act
	actualTextBytes, err := red.MarshalText()
	require.NoError(t, err)

	// Assert
	assert.Equal(t, []byte("RedSuffix"), actualTextBytes)
}

func TestReceiverMarshalText_OnCustomConfigAndJsonMarshal_ThenReturnJsonEncoded(t *testing.T) {
	// Arrange
	// Act
	actualJsonBytes, err := json.Marshal(red)
	require.NoError(t, err)

	// Assert
	assert.Equal(t, []byte("\"RedSuffix\""), actualJsonBytes)
}

func TestReceiverConfig_OnCustomConfig_ThenReturnOnCustomConfig(t *testing.T) {
	// Arrange
	defer func(revertedConfig *Config) {
		testColorConfig = revertedConfig
	}(testColorConfig)

	testColorConfig = NewConfig(
		map[string]testColor{
			"PrefixRed":    red,
			"PrefixBlue":   blue,
			"PrefixYellow": yellow,
			"PrefixGreen":  green},
		CaseInsensitive(true),
		ParseCallback(func(value string) string {
			return "Prefix" + value
		}),
		StringCallback(func(enumName string) string {
			return enumName + "Suffix"
		}))

	// Act
	actualConfig := red.Config()
	actualConfig.parseCallback = nil
	actualConfig.stringCallback = nil

	// Assert
	assert.Equal(
		t,
		&Config{
			caseInsensitive: true,
			enumNameLoweredToEnumValue: map[string]int{
				"prefixred":    0,
				"prefixblue":   1,
				"prefixyellow": 2,
				"prefixgreen":  -1},
			enumNameToEnumValue: map[string]int{
				"PrefixRed":    0,
				"PrefixBlue":   1,
				"PrefixYellow": 2,
				"PrefixGreen":  -1},
			enumValueToEnumName: map[int]string{
				0:  "PrefixRed",
				1:  "PrefixBlue",
				2:  "PrefixYellow",
				-1: "PrefixGreen",
			},
			enumValueToEnumString: map[int]string{
				0:  "PrefixRedSuffix",
				1:  "PrefixBlueSuffix",
				2:  "PrefixYellowSuffix",
				-1: "PrefixGreenSuffix"},
			joinedEnumNames:   "PrefixGreen, PrefixRed, PrefixBlue, PrefixYellow",
			sortedEnumNames:   []string{"PrefixGreen", "PrefixRed", "PrefixBlue", "PrefixYellow"},
			sortedEnumStrings: []string{"PrefixGreenSuffix", "PrefixRedSuffix", "PrefixBlueSuffix", "PrefixYellowSuffix"},
			sortedEnumValues:  []int{-1, 0, 1, 2},
		},
		actualConfig)
}

func TestStrings_OnCustomConfigAndMultipleEnums_ThenReturnDifferentStrings(t *testing.T) {
	// Arrange
	// Act
	actualNames := Strings[testColor]()

	// Assert
	assert.Equal(t, []string{"PrefixGreenSuffix", "PrefixRedSuffix", "PrefixBlueSuffix", "PrefixYellowSuffix"}, actualNames)
}

func TestNames_OnCustomConfigAndMultipleEnums_ThenReturnDifferentNames(t *testing.T) {
	// Arrange
	// Act
	actualNames := Names[testColor]()

	// Assert
	assert.Equal(t, []string{"GreenSuffix", "RedSuffix", "BlueSuffix", "YellowSuffix"}, actualNames)
}

func TestParse_OnCustomConfigAndExistingEnumName_ThenReturnEnum(t *testing.T) {
	// Arrange
	// Act
	actualEnum, err := Parse[testColor]("Red")
	require.NoError(t, err)

	// Assert
	assert.Equal(t, red, actualEnum)
}

func TestParse_OnCustomConfigAndNonExistingEnumName_ThenReturnError(t *testing.T) {
	// Arrange
	// Act
	_, err := Parse[testColor]("RedSuffix")

	// Assert
	assert.Error(t, err)
}

func TestParse_OnCaseInsensitiveAndExists_ThenReturnEnum(t *testing.T) {
	// Arrange
	defer func(revertedConfig *Config) {
		testColorConfig = revertedConfig
	}(testColorConfig)

	testColorConfig = NewConfig(
		map[string]testColor{
			"Red":    red,
			"Blue":   blue,
			"Yellow": yellow,
			"Green":  green},
		CaseInsensitive(true))

	// Act
	actualEnum, err := Parse[testColor]("BLUE")
	require.NoError(t, err)

	// Assert
	assert.Equal(t, blue, actualEnum)
}

func TestParse_OnCaseInsensitiveAndMissing_ThenReturnError(t *testing.T) {
	// Arrange
	defer func(revertedConfig *Config) {
		testColorConfig = revertedConfig
	}(testColorConfig)

	testColorConfig = NewConfig(
		map[string]testColor{
			"Red":    red,
			"Blue":   blue,
			"Yellow": yellow,
			"Green":  green},
		CaseInsensitive(true))

	// Act
	_, err := Parse[testColor]("NOT_COLOR")

	// Assert
	assert.Error(t, err)
}

func TestEnums_OnCustomConfigAndMultipleEnums_ThenReturnAll(t *testing.T) {
	// Arrange
	// Act
	actualEnums := Enums[testColor]()

	// Assert
	assert.Equal(t, []testColor{green, red, blue, yellow}, actualEnums)
}
