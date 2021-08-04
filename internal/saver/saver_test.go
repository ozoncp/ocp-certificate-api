package saver_test

import (
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/ozoncp/ocp-certificate-api/internal/mocks"
	"github.com/ozoncp/ocp-certificate-api/internal/saver"
	"time"
)

var _ = Describe("Saver", func() {
	const capacity = 12
	ticker := time.NewTicker(200 * time.Millisecond)

	var (
		ctrl        *gomock.Controller
		mockFlusher *mocks.MockFlusher
		svr         saver.Saver
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		mockFlusher = mocks.NewMockFlusher(ctrl)

		svr = saver.NewSaver(capacity, mockFlusher, *ticker)
	})

	JustBeforeEach(func() {
		svr.Init()
	})

	AfterEach(func() {
		svr.Close()
		ctrl.Finish()
	})

	Context("Save all in repository", func() {
		BeforeEach(func() {
			mockFlusher.EXPECT().Flush(gomock.Any()).Return(nil).AnyTimes()
		})

		It("", func() {
			Expect(svr).ShouldNot(BeNil())
		})
	})
})
