package gnum

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
)

type (
	shape           int
	shapeDefinition struct {
		Square,
		Triangle,
		Circle shape
	}
	Shape = Enum[shapeDefinition]
)

func TestSet_OnEmptyCache_ThenReturnOneItem(t *testing.T) {
	// Arrange
	cache := newEnumCache()
	enumType := reflect.TypeOf(*new(Shape))
	metadata := newEnumMetadata[shapeDefinition]()

	// Act
	cache.Set(enumType, metadata)
	actualMetadata, ok := cache.Get(enumType)
	require.True(t, ok)

	// Assert
	assert.Equal(
		t,
		&enumMetadata{
			enumNameLoweredToEnumValue: map[string]int{
				"square":   0,
				"triangle": 1,
				"circle":   2,
			},
			enumNameToEnumValue: map[string]int{
				"Square":   0,
				"Triangle": 1,
				"Circle":   2,
			},
			enumValueToEnumName: map[int]string{
				0: "Square",
				1: "Triangle",
				2: "Circle",
			},
			enumValueToEnumString: map[int]string{
				0: "Square",
				1: "Triangle",
				2: "Circle",
			},
			joinedEnumNames: "Square, Triangle, Circle",
			sortedEnumNames: []string{
				"Square",
				"Triangle",
				"Circle",
			},
			sortedEnumStrings: []string{
				"Square",
				"Triangle",
				"Circle",
			},
			sortedEnumValues: []int{0, 1, 2},
		},
		actualMetadata)
}

func TestSet_OnInsertTwoItems_ThenTwoItemsExists(t *testing.T) {
	// Arrange
	type (
		shape2           int
		shapeDefinition2 struct {
			Ellipsis,
			Star,
			Hexagon shape2
		}
		Shape2 = Enum[shapeDefinition2]
	)

	cache := newEnumCache()

	enumType1 := reflect.TypeOf(*new(Shape))
	enumMetadata1 := newEnumMetadata[shapeDefinition]()

	enumType2 := reflect.TypeOf(*new(Shape2))
	enumMetadata2 := newEnumMetadata[shapeDefinition2]()

	// Act
	cache.Set(enumType1, enumMetadata1)
	cache.Set(enumType2, enumMetadata2)

	actualMetadata1, ok := cache.Get(enumType1)
	require.True(t, ok)

	actualMetadata2, ok := cache.Get(enumType2)
	require.True(t, ok)

	// Assert
	assert.Equal(
		t,
		[]*enumMetadata{
			{
				enumNameLoweredToEnumValue: map[string]int{
					"square":   0,
					"triangle": 1,
					"circle":   2,
				},
				enumNameToEnumValue: map[string]int{
					"Square":   0,
					"Triangle": 1,
					"Circle":   2,
				},
				enumValueToEnumName: map[int]string{
					0: "Square",
					1: "Triangle",
					2: "Circle",
				},
				enumValueToEnumString: map[int]string{
					0: "Square",
					1: "Triangle",
					2: "Circle",
				},
				joinedEnumNames: "Square, Triangle, Circle",
				sortedEnumNames: []string{
					"Square",
					"Triangle",
					"Circle",
				},
				sortedEnumStrings: []string{
					"Square",
					"Triangle",
					"Circle",
				},
				sortedEnumValues: []int{0, 1, 2},
			},
			{
				enumNameLoweredToEnumValue: map[string]int{
					"ellipsis": 0,
					"star":     1,
					"hexagon":  2,
				},
				enumNameToEnumValue: map[string]int{
					"Ellipsis": 0,
					"Star":     1,
					"Hexagon":  2,
				},
				enumValueToEnumName: map[int]string{
					0: "Ellipsis",
					1: "Star",
					2: "Hexagon",
				},
				enumValueToEnumString: map[int]string{
					0: "Ellipsis",
					1: "Star",
					2: "Hexagon",
				},
				joinedEnumNames: "Ellipsis, Star, Hexagon",
				sortedEnumNames: []string{
					"Ellipsis",
					"Star",
					"Hexagon",
				},
				sortedEnumStrings: []string{
					"Ellipsis",
					"Star",
					"Hexagon",
				},
				sortedEnumValues: []int{0, 1, 2},
			},
		},
		[]*enumMetadata{
			actualMetadata1,
			actualMetadata2,
		})
}

func TestSet_OnSameDefinitionTwice_ThenTwoItemsExists(t *testing.T) {
	// Arrange
	type (
		shapeDefinition2 struct {
			Square,
			Triangle,
			Circle shape
		}
		Shape2 = Enum[shapeDefinition2]
	)

	cache := newEnumCache()

	enumType1 := reflect.TypeOf(*new(Shape))
	enumMetadata1 := newEnumMetadata[shapeDefinition]()

	enumType2 := reflect.TypeOf(*new(Shape2))
	enumMetadata2 := newEnumMetadata[shapeDefinition2]()

	// Act
	cache.Set(enumType1, enumMetadata1)
	cache.Set(enumType2, enumMetadata2)

	actualMetadata1, ok := cache.Get(enumType1)
	require.True(t, ok)

	actualMetadata2, ok := cache.Get(enumType2)
	require.True(t, ok)

	// Assert
	assert.Equal(
		t,
		[]*enumMetadata{
			{
				enumNameLoweredToEnumValue: map[string]int{
					"square":   0,
					"triangle": 1,
					"circle":   2,
				},
				enumNameToEnumValue: map[string]int{
					"Square":   0,
					"Triangle": 1,
					"Circle":   2,
				},
				enumValueToEnumName: map[int]string{
					0: "Square",
					1: "Triangle",
					2: "Circle",
				},
				enumValueToEnumString: map[int]string{
					0: "Square",
					1: "Triangle",
					2: "Circle",
				},
				joinedEnumNames: "Square, Triangle, Circle",
				sortedEnumNames: []string{
					"Square",
					"Triangle",
					"Circle",
				},
				sortedEnumStrings: []string{
					"Square",
					"Triangle",
					"Circle",
				},
				sortedEnumValues: []int{0, 1, 2},
			},
			{
				enumNameLoweredToEnumValue: map[string]int{
					"square":   0,
					"triangle": 1,
					"circle":   2,
				},
				enumNameToEnumValue: map[string]int{
					"Square":   0,
					"Triangle": 1,
					"Circle":   2,
				},
				enumValueToEnumName: map[int]string{
					0: "Square",
					1: "Triangle",
					2: "Circle",
				},
				enumValueToEnumString: map[int]string{
					0: "Square",
					1: "Triangle",
					2: "Circle",
				},
				joinedEnumNames: "Square, Triangle, Circle",
				sortedEnumNames: []string{
					"Square",
					"Triangle",
					"Circle",
				},
				sortedEnumStrings: []string{
					"Square",
					"Triangle",
					"Circle",
				},
				sortedEnumValues: []int{0, 1, 2},
			},
		},
		[]*enumMetadata{
			actualMetadata1,
			actualMetadata2,
		})
}

func TestGet_OnEmptyCache_ThenFalse(t *testing.T) {
	// Arrange
	cache := newEnumCache()

	// Act
	actualMetadata, ok := cache.Get(reflect.TypeOf(*new(Shape)))
	require.Nil(t, actualMetadata)

	// Assert
	assert.False(t, ok)
}

func TestGet_OnTypeNotExists_ThenFalse(t *testing.T) {
	// Arrange
	cache := newEnumCache()
	enumType := reflect.TypeOf(*new(Shape))
	enumMetadata := newEnumMetadata[shapeDefinition]()
	cache.Set(enumType, enumMetadata)

	type (
		Shape2 = Enum[struct {
			Square,
			Triangle,
			Circle shape
		}]
	)

	// Act
	actualMetadata, ok := cache.Get(reflect.TypeOf(*new(Shape2)))
	require.Nil(t, actualMetadata)

	// Assert
	assert.False(t, ok)
}
