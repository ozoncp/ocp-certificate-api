package utils

import (
	"github.com/ozoncp/ocp-certificate-api/internal/model"
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
