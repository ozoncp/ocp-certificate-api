package utils

import (
	"errors"
	"github.com/ozoncp/ocp-certificate-api/internal/model"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestSplitToBulksSuccess(t *testing.T) {
	now := time.Now()
	slice := []model.Certificate{
		*model.NewCertificate(1.0, 1.0, now, "http://link"),
		*model.NewCertificate(2.0, 2.0, now, "http://link"),
		*model.NewCertificate(3.0, 3.0, now, "http://link"),
		*model.NewCertificate(4.0, 4.0, now, "http://link"),
		*model.NewCertificate(5.0, 5.0, now, "http://link"),
	}

	want := [][]model.Certificate{
		{*model.NewCertificate(1.0, 1.0, now, "http://link"),
			*model.NewCertificate(2.0, 2.0, now, "http://link")},
		{*model.NewCertificate(3.0, 3.0, now, "http://link"),
			*model.NewCertificate(4.0, 4.0, now, "http://link")},
		{*model.NewCertificate(5.0, 5.0, now, "http://link")},
	}

	got := SplitToBulks(slice, 2)

	assert.Equal(t, len(want), len(got))
}

func TestSplitToBulksEmpty(t *testing.T) {
	var slice []model.Certificate

	var want [][]model.Certificate

	got := SplitToBulks(slice, 2)

	assert.Equal(t, len(want), len(got))
}

func TestSliceToMapSuccess(t *testing.T) {
	now := time.Now()
	slice := []model.Certificate{
		*model.NewCertificate(1.0, 1.0, now, "http://link"),
		*model.NewCertificate(2.0, 2.0, now, "http://link"),
		*model.NewCertificate(3.0, 3.0, now, "http://link"),
		*model.NewCertificate(4.0, 4.0, now, "http://link"),
		*model.NewCertificate(5.0, 5.0, now, "http://link"),
	}

	want := map[uint64]model.Certificate{
		1: *model.NewCertificate(1.0, 1.0, now, "http://link"),
		2: *model.NewCertificate(1.0, 2.0, now, "http://link"),
		3: *model.NewCertificate(2.0, 3.0, now, "http://link"),
		4: *model.NewCertificate(3.0, 4.0, now, "http://link"),
		5: *model.NewCertificate(4.0, 5.0, now, "http://link"),
	}

	got, _ := SliceToMap(slice)

	assert.Equal(t, len(want), len(got))
}

func TestSliceToMapError(t *testing.T) {
	var slice []model.Certificate

	want := errors.New("The swapSlice size cannot be zero.")

	_, got := SliceToMap(slice)

	assert.Equal(t, want, got)
}
