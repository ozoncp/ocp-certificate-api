package saver_test

import (
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/ozoncp/ocp-certificate-api/internal/mocks"
	"github.com/ozoncp/ocp-certificate-api/internal/model"
	"github.com/ozoncp/ocp-certificate-api/internal/saver"
	"time"
)

var _ = Describe("Saver", func() {
	const capacity = 10
	const buffer = 10

	var (
		ctrl               *gomock.Controller
		mockFlusher        *mocks.MockFlusher
		certificateChannel chan model.Certificate
		certificates       []model.Certificate
		s                  saver.Saver
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		mockFlusher = mocks.NewMockFlusher(ctrl)
		certificateChannel = make(chan model.Certificate, buffer)

		now := time.Now()
		certificates = []model.Certificate{
			{1.0, 1.0, now, "https://link"},
			{2.0, 2.0, now, "https://link"},
			{3.0, 3.0, now, "https://link"},
		}
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	Context("Saving success", func() {
		expected := 3
		BeforeEach(func() {
			mockFlusher.EXPECT().Flush(gomock.Any()).Times(1).Return(certificates)
		})

		JustBeforeEach(func() {
			ticker := time.NewTicker(time.Millisecond * 100)
			s = saver.NewSaver(capacity, mockFlusher, *ticker)
			s.Init()
		})

		It("Save and close", func() {
			defer s.Close()
			for _, cert := range certificates {
				s.Save(cert)
				certificateChannel <- cert
			}

			time.Sleep(time.Millisecond * 120)
			Expect(len(certificateChannel)).Should(BeEquivalentTo(expected))
		})
	})

	Context("Saving success", func() {
		BeforeEach(func() {
			mockFlusher.EXPECT().Flush(gomock.Any()).Times(1).Return(certificates)
		})

		JustBeforeEach(func() {
			ticker := time.NewTicker(time.Millisecond * 100)
			s = saver.NewSaver(capacity, mockFlusher, *ticker)
			s.Init()
		})

		It("Save after tick", func() {
			defer s.Close()
			s.Save(certificates[0])
			time.Sleep(time.Millisecond * 120)
		})
	})

	Context("Error when try save", func() {
		BeforeEach(func() {
			mockFlusher.EXPECT().Flush(gomock.Any()).AnyTimes().Return(nil)
		})

		JustBeforeEach(func() {
			ticker := time.NewTicker(time.Millisecond * 100)
			s = saver.NewSaver(capacity, mockFlusher, *ticker)
			s.Init()
		})

		It("Panic, because close channel before Save() ", func() {
			s.Close()
			save := func() {
				s.Save(certificates[0])
			}

			Expect(save).Should(Panic())
		})

		It("Panic, because close channel before Init() ", func() {
			s.Close()
			save := func() {
				s.Save(certificates[0])
				s.Init()
			}

			Expect(save).Should(Panic())
		})
	})
})
