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

	ticker := time.NewTicker(200 * time.Millisecond)
	now := time.Now()

	var (
		ctrl               *gomock.Controller
		mockFlusher        *mocks.MockFlusher
		svr                saver.Saver
		certificateChannel chan model.Certificate
		certificates       []model.Certificate
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		mockFlusher = mocks.NewMockFlusher(ctrl)

		certificateChannel = make(chan model.Certificate, buffer)
		svr = saver.NewSaver(capacity, mockFlusher, *ticker)

		certificates = []model.Certificate{
			0: {1.0, 1.0, now, "https://link"},
			1: {2.0, 2.0, now, "https://link"},
			2: {3.0, 3.0, now, "https://link"},
		}
	})

	JustBeforeEach(func() {
		svr.Init()
	})

	AfterEach(func() {
		svr.Close()
		ctrl.Finish()
	})

	Context("Run tests", func() {
		BeforeEach(func() {
			mockFlusher.EXPECT().Flush(gomock.Any()).MinTimes(1).Return([]model.Certificate{{}})
		})

		It("Saving", func() {
			saver := saver.NewSaver(capacity, mockFlusher, *ticker)
			saver.Init()

			for _, cert := range certificates {
				certificateChannel <- cert
			}

			for _, cert := range certificates {
				saver.Save(cert)
			}

			saver.Close()
			Expect(len(certificateChannel)).Should(BeEquivalentTo(3))
		})

		It("Panic", func() {
			mockFlusher.Flush(nil)
			saver := saver.NewSaver(capacity, mockFlusher, *ticker)
			saver.Init()

			save := func() {
				saver.Save(certificates[0])
			}

			saver.Close()
			Expect(save).Should(Panic())
		})
	})
})
