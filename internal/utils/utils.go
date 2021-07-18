package utils

import "errors"

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

func FilterSlice(filterSlice []int) ([]int, error) {
	if len(filterSlice) == 0 {
		return nil, errors.New("The filterSlice size cannot be zero.")
	}

	filter := []int{2, 4, 6, -3, 10, 8}
	result := make([]int, 0)
	isAppendValue := true

	for _, vSlice := range filterSlice {

		isAppendValue = true
		for _, vFilter := range filter {
			if vSlice == vFilter {
				isAppendValue = false
				break
			}
		}

		if isAppendValue == true {
			result = append(result, vSlice)
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
