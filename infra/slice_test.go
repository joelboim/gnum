package infra

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestInsertToSortedSlice_OnNegativeNumber_ThenReturnSortedValue(t *testing.T) {
	// Arrange
	values := []int{1, 2, 3}

	// Act
	insertedIndex := InsertToSortedSlice(&values, -1)
	require.Equal(t, uint64(0), insertedIndex)

	// Assert
	assert.Equal(t, []int{-1, 1, 2, 3}, values)
}

func TestInsertToSortedSlice_OnPositiveNumber_ThenReturnSortedValue(t *testing.T) {
	// Arrange
	values := []int{1, 2, 3}

	// Act
	insertedIndex := InsertToSortedSlice(&values, 4)
	require.Equal(t, uint64(3), insertedIndex)

	// Assert
	assert.Equal(t, []int{1, 2, 3, 4}, values)
}

func TestInsertToSortedSlice_OnFloat_ThenReturnSortedValue(t *testing.T) {
	// Arrange
	values := []float64{1, 2, 3}

	// Act
	insertedIndex := InsertToSortedSlice(&values, 1.5)
	require.Equal(t, uint64(1), insertedIndex)

	// Assert
	assert.Equal(t, []float64{1, 1.5, 2, 3}, values)
}

func TestInsertToSliceIndex_OnTheMiddleIndex_ThenPutInTheMiddle(t *testing.T) {
	// Arrange
	values := []int{1, 2, 3}

	// Act
	InsertToSliceByIndex(&values, 4, 1)

	// Assert
	assert.Equal(t, []int{1, 4, 2, 3}, values)
}

func TestInsertToSliceIndex_OnTheEndIndex_ThenPutAtTheEnd(t *testing.T) {
	// Arrange
	values := []int{1, 2, 3}

	// Act
	InsertToSliceByIndex(&values, 4, 3)

	// Assert
	assert.Equal(t, []int{1, 2, 3, 4}, values)
}

func TestInsertToSliceIndex_OnOverflowIndex_ThenPanic(t *testing.T) {
	// Arrange
	values := []int{1, 2, 3}

	// Act
	// Assert
	assert.Panics(t,
		func() { InsertToSliceByIndex(&values, 4, 5) },
	)
}
