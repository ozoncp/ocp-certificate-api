package flusher_test

import (
	"errors"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/ozoncp/ocp-certificate-api/internal/flusher"
	"github.com/ozoncp/ocp-certificate-api/internal/mocks"
	"github.com/ozoncp/ocp-certificate-api/internal/model"
	"time"
)

var mockError = errors.New("error")

var _ = Describe("Flusher", func() {
	now := time.Now()
	var (
		ctrl         *gomock.Controller
		mockRepo     *mocks.MockRepo
		certificates []model.Certificate
		results      []model.Certificate
		f            flusher.Flusher
		chunkSize    int
	)
	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		mockRepo = mocks.NewMockRepo(ctrl)

		certificates = []model.Certificate{
			0: {1.0, 1.0, now, "http://link"},
			1: {2.0, 2.0, now, "http://link"},
			2: {3.0, 3.0, now, "http://link"},
		}

		chunkSize = 2
	})

	JustBeforeEach(func() {
		f = flusher.NewFlusher(chunkSize, mockRepo)
		results = f.Flush(certificates)
	})

	Context("Save all in repository", func() {
		BeforeEach(func() {
			mockRepo.EXPECT().AddCertificates(gomock.Any()).Return(nil).AnyTimes()
		})

		It("", func() {
			Expect(results).Should(BeNil())
		})
	})

	Context("Error when saving in repository", func() {
		BeforeEach(func() {
			mockRepo.EXPECT().AddCertificates(gomock.Any()).Return(mockError)
		})

		It("", func() {
			Expect(len(results)).Should(BeEquivalentTo(len(certificates[:chunkSize])))
			Expect(results).Should(BeEquivalentTo(certificates[:chunkSize]))
		})
	})

	Context("Partial saving to repo", func() {
		BeforeEach(func() {
			mockRepo.EXPECT().AddCertificates(gomock.Any()).Return(mockError)
			mockRepo.EXPECT().AddCertificates(gomock.Any()).Return(nil).Times(1)
		})

		It("", func() {
			Expect(results).Should(BeEquivalentTo(certificates[:chunkSize]))
		})
	})
})
