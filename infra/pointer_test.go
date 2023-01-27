package infra

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetPointerValue_OnPointerHasValue_ThenReturnPointerValue(t *testing.T) {
	// Arrange
	value := 123

	// Act
	actualValue := GetPointerValue(&value, 234)

	// Assert
	assert.Equal(t, 123, actualValue)
}

func TestGetPointerValue_OnPointerDoesntHaveValue_ThenReturnDefaultValue(t *testing.T) {
	// Arrange
	// Act
	actualValue := GetPointerValue(nil, 234)

	// Assert
	assert.Equal(t, 234, actualValue)
}

func TestGetPointerValue_OnPointerPointerHaveValue_ThenReturnPointerPointerValue(t *testing.T) {
	// Arrange
	innerValue := 123
	value := &innerValue
	defaultValue := 234

	// Act
	actualValue := GetPointerValue(&value, &defaultValue)

	// Assert
	expected := 123
	assert.Equal(t, &expected, actualValue)
}

func TestGetPointerValue_OnPointerPointerDoesntHaveValue_ThenReturnDefaultValue(t *testing.T) {
	// Arrange
	var value **int
	defaultValue := 234

	// Act
	actualValue := GetPointerValue(value, &defaultValue)

	// Assert
	expected := 234
	assert.Equal(t, &expected, actualValue)
}
