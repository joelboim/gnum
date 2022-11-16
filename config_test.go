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
	OptionStringCallback(func(value string) string {
		return "Prefix" + value
	}),
	OptionParseNormalizationCallback(func(value string) string {
		return value + "Suffix"
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

func TestReceiverString_OnCustomConfig_ThenReturnName(t *testing.T) {
	// Arrange
	// Act
	actualName := red.String()

	// Assert
	assert.Equal(t, "PrefixRedSuffix", actualName)
}

func TestReceiverNames_OnCustomConfigAndMultipleEnums_ThenReturnDifferentNames(t *testing.T) {
	// Arrange
	// Act
	actualNames := red.Names()

	// AssertRedSuffix
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

func TestReceiverParse_OnParseCaseInsensitiveAndExists_ThenReturnEnum(t *testing.T) {
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
		OptionParseCaseInsensitive(true))

	// Act
	actualEnum, err := red.Parse("BLUE")
	require.NoError(t, err)

	// Assert
	assert.Equal(t, blue, actualEnum)
}

func TestReceiverParse_OnParseCaseInsensitiveAndMissing_ThenReturnError(t *testing.T) {
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
		OptionParseCaseInsensitive(true))

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
	assert.Equal(t, []byte("PrefixRedSuffix"), actualTextBytes)
}

func TestReceiverMarshalText_OnCustomConfigAndJsonMarshal_ThenReturnJsonEncoded(t *testing.T) {
	// Arrange
	// Act
	actualJsonBytes, err := json.Marshal(red)
	require.NoError(t, err)

	// Assert
	assert.Equal(t, []byte("\"PrefixRedSuffix\""), actualJsonBytes)
}

func TestReceiverConfig_OnCustomConfig_ThenReturnOnCustomConfig(t *testing.T) {
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
		OptionParseCaseInsensitive(true))

	// Act
	actualConfig := red.Config()
	actualConfig.parseNormalizationCallback = nil
	actualConfig.stringCallback = nil

	// Assert
	assert.Equal(
		t,
		&Config{
			nameToValue: map[string]int{
				"Red":    0,
				"Blue":   1,
				"Yellow": 2,
				"Green":  -1},
			caseInsensitiveNameToValue: map[string]int{
				"red":    0,
				"blue":   1,
				"yellow": 2,
				"green":  -1},
			valueToName: map[int]string{
				0:  "Red",
				1:  "Blue",
				2:  "Yellow",
				-1: "Green"},
			sortedNames:          []string{"Green", "Red", "Blue", "Yellow"},
			sortedValues:         []int{-1, 0, 1, 2},
			allEnumsString:       "Green, Red, Blue, Yellow",
			parseCaseInsensitive: true},
		actualConfig)
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

func TestParse_OnParseCaseInsensitiveAndExists_ThenReturnEnum(t *testing.T) {
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
		OptionParseCaseInsensitive(true))

	// Act
	actualEnum, err := Parse[testColor]("BLUE")
	require.NoError(t, err)

	// Assert
	assert.Equal(t, blue, actualEnum)
}

func TestParse_OnParseCaseInsensitiveAndMissing_ThenReturnError(t *testing.T) {
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
		OptionParseCaseInsensitive(true))

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
