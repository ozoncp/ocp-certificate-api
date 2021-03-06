package utils

import "errors"

// SplitSlice - splits the slice into slices with the specified batchSize size
func SplitSlice(splitSlice []int, batchSize int) ([][]int, error) {
	if len(splitSlice) == 0 {
		return nil, errors.New("The splitSlice size cannot be zero.")
	}

	if batchSize <= 0 {
		return nil, errors.New("The batch size cannot be zero or negative.")
	}

	result := make([][]int, (len(splitSlice)-1)/batchSize+1)

	for index := range result {
		first, last := index*batchSize, (index+1)*batchSize

		if last < len(splitSlice) {
			result[index] = splitSlice[first:last]
			continue
		}

		result[index] = splitSlice[first:]
	}

	return result, nil
}

// FilterSlice - filtering the list items using a parameter filter,
// where items will be returned that do not intersect with the filter with a parameter
func FilterSlice(filterSlice []int, filter []int) ([]int, error) {
	if len(filterSlice) == 0 {
		return nil, errors.New("The filterSlice size cannot be zero.")
	}

	if len(filter) == 0 {
		return nil, errors.New("The filter size cannot be zero.")
	}

	var filtered = map[int]struct{}{}
	for _, f := range filter {
		filtered[f] = struct{}{}
	}

	result := make([]int, 0)
	for _, value := range filterSlice {
		if _, found := filtered[value]; !found {
			result = append(result, value)
		}
	}

	return result, nil
}

// SwapMap - swaps the key and value
func SwapMap(swapMap map[int]string) (map[string]int, error) {
	if len(swapMap) == 0 {
		return nil, errors.New("The swapSlice size cannot be zero.")
	}

	result := make(map[string]int, len(swapMap))

	for key, value := range swapMap {
		result[value] = key
	}

	return result, nil
}
