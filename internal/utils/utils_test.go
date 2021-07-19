package utils

import (
	"errors"
	"reflect"
	"testing"
)

type Certificate struct {
	id int
}

func TestSplitSliceSuccess(t *testing.T) {
	actual, _ := SplitSlice([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}, 4)
	expected := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11}}
	assertDeapEqual(t, actual, expected)
}

func TestSplitSliceError(t *testing.T) {
	_, actual := SplitSlice([]int{}, 4)
	expected := errors.New("The splitSlice size cannot be zero.")
	assertDeapEqual(t, actual, expected)
}

func TestSplitSliceBatchError(t *testing.T) {
	_, actual := SplitSlice([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}, -2)
	expected := errors.New("The batch size cannot be zero or negative.")
	assertDeapEqual(t, actual, expected)
}

func TestFilterSliceSuccess(t *testing.T) {
	actual, _ := FilterSlice([]int{1, 6, 8, 6, 3, 10, -4, 6, 7, 6, 9, 6, 6, 6},
		[]int{2, 4, 6, -3, 10, 8})
	expected := []int{1, 3, -4, 7, 9}
	assertDeapEqual(t, actual, expected)
}

func TestFilterSliceError(t *testing.T) {
	_, actual := FilterSlice([]int{}, []int{})
	expected := errors.New("The filterSlice size cannot be zero.")
	assertDeapEqual(t, actual, expected)
}

func TestFilterSliceFilteredError(t *testing.T) {
	_, actual := FilterSlice([]int{1, 6, 8, 6, 3, 10, -4, 6, 7, 6, 9, 6, 6, 6}, []int{})
	expected := errors.New("The filter size cannot be zero.")
	assertDeapEqual(t, actual, expected)
}

func TestSwapMapSuccess(t *testing.T) {
	actual, _ := SwapMap(map[int]string{1: "11", 2: "22"})
	expected := map[string]int{"11": 1, "22": 2}
	assertDeapEqual(t, actual, expected)
}

func TestSwapMapError(t *testing.T) {
	_, actual := SwapMap(map[int]string{})
	expected := errors.New("The swapSlice size cannot be zero.")
	assertDeapEqual(t, actual, expected)
}

func assertDeapEqual(t *testing.T, actual, expected interface{}) {
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v, but got %v", expected, actual)
		return
	}
}
