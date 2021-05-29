package flusher_test

import (
	"errors"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/ozoncp/ocp-classroom-api/internal/flusher"
	"github.com/ozoncp/ocp-classroom-api/internal/mocks"
	"github.com/ozoncp/ocp-classroom-api/internal/models"
)

var _ = Describe("Flusher", func() {

	var (
		ctrl     *gomock.Controller
		mockRepo *mocks.MockRepo

		fl flusher.Flusher
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		mockRepo = mocks.NewMockRepo(ctrl)

		fl = flusher.New(mockRepo)
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	Context("Flush call", func() {

		It("can flush", func() {

			mockRepo.EXPECT().AddClassrooms(gomock.Any()).Return(nil)

			classrooms := []models.Classroom{}
			remainingClassrooms := fl.Flush(classrooms)

			Expect(remainingClassrooms).Should(BeNil())
		})

		It("can not flush", func() {

			mockRepo.EXPECT().AddClassrooms(gomock.Any()).DoAndReturn(
				func(classrooms []models.Classroom) error {
					return errors.New("can not add classrooms")
				},
			)

			classrooms := []models.Classroom{}
			remainingClassrooms := fl.Flush(classrooms)

			Expect(remainingClassrooms).Should(BeEquivalentTo(classrooms))
		})
	})
})
