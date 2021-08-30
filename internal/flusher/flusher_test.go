package flusher_test

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/ozoncp/ocp-certificate-api/internal/flusher"
	"github.com/ozoncp/ocp-certificate-api/internal/mocks"
	"github.com/ozoncp/ocp-certificate-api/internal/model"
	"time"
)

var errNew = errors.New("error")

var _ = Describe("Flusher", func() {
	now := time.Now()
	var (
		ctrl         *gomock.Controller
		ctx          context.Context
		mockRepo     *mocks.MockRepo
		certificates []model.Certificate
		results      []model.Certificate
		f            flusher.Flusher
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		mockRepo = mocks.NewMockRepo(ctrl)
		ctx = context.Background()
		link := "https://link.ru"

		certificates = []model.Certificate{
			{1.0, 1.0, now, link, false},
			{2.0, 2.0, now, link, false},
			{3.0, 3.0, now, link, false},
			{4.0, 4.0, now, link, false},
		}
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	Context("Save all in repository", func() {
		chunkSize := 2
		BeforeEach(func() {
			mockRepo.EXPECT().MultiCreateCertificates(ctx, gomock.Any()).Return([]uint64{4}, nil).MinTimes(1)
		})

		It("", func() {
			f = flusher.NewFlusher(chunkSize, mockRepo)
			results = f.Flush(ctx, certificates)
			Expect(results).Should(BeNil())
		})
	})

	Context("Error when saving in repository", func() {
		chunkSize := 2
		BeforeEach(func() {
			mockRepo.EXPECT().MultiCreateCertificates(ctx, gomock.Any()).Return(nil, errNew).MinTimes(1)
		})

		It("", func() {
			f = flusher.NewFlusher(chunkSize, mockRepo)
			results = f.Flush(ctx, certificates)
			Expect(len(results)).Should(BeEquivalentTo(len(certificates)))
			Expect(results).Should(BeEquivalentTo(certificates))
		})
	})

	Context("Partial saving to repo", func() {
		chunkSize := 2
		BeforeEach(func() {
			gomock.InOrder(
				mockRepo.EXPECT().MultiCreateCertificates(ctx, gomock.Any()).Return([]uint64{4}, nil),
				mockRepo.EXPECT().MultiCreateCertificates(ctx, gomock.Any()).Return(nil, errNew),
			)
		})

		It("", func() {
			f = flusher.NewFlusher(chunkSize, mockRepo)
			results = f.Flush(ctx, certificates)
			Expect(results).Should(BeEquivalentTo(certificates[chunkSize:]))
		})
	})
})
