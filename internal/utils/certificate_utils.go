package utils

import (
	"errors"
	"fmt"
	"github.com/ozoncp/ocp-certificate-api/internal/model"
	"io"
	"log"
	"os"
)

// SplitToBulks - splits the certificate slice into certificate slices with the specified batchSize size
func SplitToBulks(certificate []model.Certificate, batchSize int) [][]model.Certificate {

	if len(certificate) == 0 || batchSize <= 0 {
		return [][]model.Certificate{}
	}

	result := make([][]model.Certificate, (len(certificate)-1)/batchSize+1)

	for index := range result {
		first, last := index*batchSize, (index+1)*batchSize

		if last < len(certificate) {
			result[index] = certificate[first:last]
			continue
		}

		result[index] = certificate[first:]
	}

	return result
}

// SliceToMap - convert a slice certificates in map, where the key is the certificate
// identifier and the value is the certificate itself
func SliceToMap(certificate []model.Certificate) (map[uint64]model.Certificate, error) {
	if len(certificate) == 0 {
		return nil, errors.New("The swapSlice size cannot be zero.")
	}

	result := make(map[uint64]model.Certificate, len(certificate))

	for _, value := range certificate {
		if _, found := result[value.Id]; found {
			return nil, errors.New("Certificate Id is not unique. Error: duplicate value")
		}
		result[value.Id] = value
	}

	return result, nil
}

// ReadFileByCount - filtering on the list - which filter the input slice
// by the criterion of the absence of an element in the list
func ReadFileByCount(filePath string, count int) {
	readFile := func() ([]byte, error) {
		file, err := os.OpenFile(filePath, os.O_RDONLY, 0666)
		if err != nil {
			return nil, err
		}

		defer func() {
			if err = file.Close(); err != nil {
				log.Fatal(err)
			} else {
				fmt.Printf("File %s is successfully closed \n", file.Name())
			}
		}()

		data, err := io.ReadAll(file)
		if err != nil {
			return nil, err
		}

		return data, nil
	}

	for i := 0; i < count; i++ {
		if data, err := readFile(); err == nil {
			fmt.Printf(string(data))
		} else {
			panic(err)
		}
	}
}
