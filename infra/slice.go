package infra

import (
	"golang.org/x/exp/constraints"
	"sort"
)

func InsertToSliceByIndex[T constraints.Ordered](
	values *[]T,
	value T,
	index uint64) {

	if len(*values) == int(index) {
		*values = append(*values, value)
		return
	}

	*values = append((*values)[:index+1], (*values)[index:]...)
	(*values)[index] = value
}

func InsertToSortedSlice[T constraints.Ordered](values *[]T, value T) int {
	index := sort.Search(
		len(*values),
		func(i int) bool { return (*values)[i] >= value })

	*values = append(*values, *new(T))
	copy((*values)[index+1:], (*values)[index:])
	(*values)[index] = value

	return index
}
