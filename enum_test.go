package gnum

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

const (
	dog testAnimal = iota
	cat
	cow
	chicken testAnimal = -1
)

var testAnimalConfig = NewConfig(
	map[string]testAnimal{
		"Dog":       dog,
		"Cat":       cat,
		"Cow":       cow,
		"Chic\tken": chicken},
)

type animal int

type testAnimal = Enum[animal]

func (animal) Config() *Config {
	return testAnimalConfig
}

func TestReceiverString_OnDefaultConfig_ThenReturnString(t *testing.T) {
	// Arrange
	// Act
	actualString := dog.String()

	// Assert
	assert.Equal(t, "Dog", actualString)
}

func TestReceiverString_OnDifferentEnum_ThenReturnDifferentName(t *testing.T) {
	// Arrange
	// Act
	actualString := cow.String()

	// Assert
	assert.Equal(t, "Cow", actualString)
}

func TestReceiverString_OnEnumNotRegisteredInConfig_ThenPanic(t *testing.T) {
	// Arrange
	const notRegisteredEnum testAnimal = 10

	// Act
	// Assert
	assert.Panics(t, func() {
		_ = notRegisteredEnum.String()
	})
}

func TestReceiverStrings_OnDefaultConfig_ThenReturnStrings(t *testing.T) {
	// Arrange
	// Act
	actualStrings := dog.Strings()

	// Assert
	assert.Equal(t, []string{"Chic\tken", "Dog", "Cat", "Cow"}, actualStrings)
}

func TestReceiverName_OnDefaultConfig_ThenReturnName(t *testing.T) {
	// Arrange
	// Act
	actualString := dog.Name()

	// Assert
	assert.Equal(t, "Dog", actualString)
}

func TestReceiverNames_OnMultipleEnums_ThenReturnDifferentNames(t *testing.T) {
	// Arrange
	// Act
	actualNames := cow.Names()

	// Assert
	assert.Equal(t, []string{"Chic\tken", "Dog", "Cat", "Cow"}, actualNames)
}

func TestReceiverParse_OnExistingEnumName_ThenReturnEnum(t *testing.T) {
	// Arrange
	// Act
	actualEnum, err := cow.Parse("Cat")
	require.NoError(t, err)

	// Assert
	assert.Equal(t, cat, actualEnum)
}

func TestReceiverParse_OnNonExistingEnumName_ThenReturnError(t *testing.T) {
	// Arrange
	// Act
	_, err := cow.Parse("nop")

	// Assert
	assert.Error(t, err)
}

func TestReceiverEnums_OnMultipleEnums_ThenReturnAll(t *testing.T) {
	// Arrange
	// Act
	actualEnums := dog.Enums()

	// Assert
	assert.Equal(t, []testAnimal{chicken, dog, cat, cow}, actualEnums)
}

func TestReceiverMarshalText_OnDefaultConfig_ThenReturnName(t *testing.T) {
	// Arrange
	// Act
	actualTextBytes, err := dog.MarshalText()
	require.NoError(t, err)

	// Assert
	assert.Equal(t, []byte("Dog"), actualTextBytes)
}

func TestReceiverMarshalText_OnJsonMarshal_ThenReturnJsonEncoded(t *testing.T) {
	// Arrange
	// Act
	actualJsonBytes, err := json.Marshal(dog)
	require.NoError(t, err)

	// Assert
	assert.Equal(t, []byte("\"Dog\""), actualJsonBytes)
}

func TestReceiverUnmarshalText_OnDefaultConfig_ThenReturnName(t *testing.T) {
	// Arrange
	actualEnum := new(testAnimal)

	// Act
	err := actualEnum.UnmarshalText([]byte("Dog"))
	require.NoError(t, err)

	// Assert
	assert.Equal(t, dog, *actualEnum)
}

func TestReceiverUnmarshalText_OnMarshelText_ThenReturnEnum(t *testing.T) {
	// Arrange
	enumMarsheled, err := dog.MarshalText()
	require.NoError(t, err)
	actualEnum := new(testAnimal)

	// Act
	err = actualEnum.UnmarshalText(enumMarsheled)
	require.NoError(t, err)

	// Assert
	assert.Equal(t, dog, *actualEnum)
}

func TestReceiverUnmarshalText_OnInvalidEnum_ThenReturnError(t *testing.T) {
	// Arrange
	actualEnum := new(testAnimal)

	// Act
	err := actualEnum.UnmarshalText([]byte("not enum"))

	// Assert
	assert.Error(t, err)
}

func TestReceiverUnmarshalText_OnJsonUnmarshal_ThenReturnStructWithEnum(t *testing.T) {
	// Arrange
	actual := &struct {
		Animal testAnimal
	}{}

	// Act
	err := json.Unmarshal([]byte("{\"Animal\":\"Dog\"}"), actual)
	require.NoError(t, err)

	// Assert
	assert.Equal(t, &struct{ Animal testAnimal }{dog}, actual)
}

func TestReceiverType_OnEnumWithUnderlineNamedType_ThenReturnTypeName(t *testing.T) {
	// Arrange
	// Act
	actualType := dog.Type()

	// Assert
	assert.Equal(t, "animal", actualType)
}
