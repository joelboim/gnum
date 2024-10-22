package gnum

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNames_OnMultipleEnums_ThenReturnDifferentNames(t *testing.T) {
	// Arrange
	// Act
	actualNames := Names[testAnimal]()

	// Assert
	assert.Equal(t, []string{"Chic\tken", "Dog", "Cat", "Cow"}, actualNames)
}

func TestStrings_OnMultipleEnums_ThenReturnDifferentStrings(t *testing.T) {
	// Arrange
	// Act
	actualStrings := Names[testAnimal]()

	// Assert
	assert.Equal(t, []string{"Chic\tken", "Dog", "Cat", "Cow"}, actualStrings)
}

func TestParse_OnExistingEnumName_ThenReturnEnum(t *testing.T) {
	// Arrange
	// Act
	actualEnum, err := Parse[testAnimal]("Cat")
	require.NoError(t, err)

	// Assert
	assert.Equal(t, cat, actualEnum)
}

func TestParse_OnNonExistingEnumName_ThenReturnError(t *testing.T) {
	// Arrange
	// Act
	_, err := Parse[testAnimal]("nop")

	// Assert
	assert.Error(t, err)
}

func TestEnums_OnMultipleEnums_ThenReturnAll(t *testing.T) {
	// Arrange
	// Act
	actualEnums := Enums[testAnimal]()

	// Assert
	assert.Equal(t, []testAnimal{chicken, dog, cat, cow}, actualEnums)
}

func TestType_OnEnumWithUnderlineNamedType_ThenReturnTypeName(t *testing.T) {
	// Arrange
	// Act
	actualType := Type[testAnimal]()

	// Assert
	assert.Equal(t, "animal", actualType)
}

func TestValues_OnMultipleEnums_ThenReturnInts(t *testing.T) {
	// Arrange
	// Act
	// Assert
	assert.Equal(t, []int{-1, 0, 1, 2}, Values[testAnimal]())
}
