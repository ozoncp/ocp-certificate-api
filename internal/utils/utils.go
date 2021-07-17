package utils

import "errors"

func SplitSlice(splitSlice []uint, batchSize int) ([][]uint, error) {
	if batchSize == 0 {
		return nil, errors.New("The batch size cannot be zero.")
	}

	result := make([][]uint, (len(splitSlice)-1)/batchSize+1)

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

func FilterSlice(filterSlice []uint, filter uint) ([]uint, error) {
	if len(filterSlice) == 0 {
		return nil, errors.New("The filterSlice size cannot be zero.")
	}

	result := make([]uint, 0, len(filterSlice))

	contains := func(value uint) bool {
		return value == filter
	}

	for index := range filterSlice {
		if !contains(filterSlice[index]) {
			result = append(result, filterSlice[index])
		}
	}

	return result, nil
}

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
