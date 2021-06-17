package flusher_test

import (
	"context"
	"errors"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/ozoncp/ocp-classroom-api/internal/flusher"
	"github.com/ozoncp/ocp-classroom-api/internal/mocks"
	"github.com/ozoncp/ocp-classroom-api/internal/models"
)

var _ = Describe("Flusher", func() {

	Describe("Flush call", func() {

		const chunkSize int = 2

		var (
			ctrl     *gomock.Controller
			mockRepo *mocks.MockRepo

			classrooms []models.Classroom

			ctx context.Context
			fl  flusher.Flusher
		)

		BeforeEach(func() {
			ctrl = gomock.NewController(GinkgoT())
			mockRepo = mocks.NewMockRepo(ctrl)

			classrooms = []models.Classroom{
				{Id: 1, TenantId: 1, CalendarId: 1},
				{Id: 2, TenantId: 2, CalendarId: 2},
				{Id: 3, TenantId: 3, CalendarId: 3},
				{Id: 4, TenantId: 4, CalendarId: 4},
				{Id: 5, TenantId: 5, CalendarId: 5},
			}

			ctx = context.Background()
		})

		AfterEach(func() {
			ctrl.Finish()
		})

		When("parameters are not valid", func() {

			It("should return whole slice", func() {

				By("receiving wrong chunkSize")

				fl = flusher.New(mockRepo, -1)
				remainingClassrooms := fl.Flush(ctx, classrooms)

				Expect(remainingClassrooms).Should(BeEquivalentTo(classrooms))

				By("receiving nil slice of classroom")

				classrooms = nil

				fl = flusher.New(mockRepo, chunkSize)
				remainingClassrooms = fl.Flush(ctx, classrooms)

				Expect(remainingClassrooms).Should(BeEquivalentTo(classrooms))
			})
		})

		When("parameters are valid", func() {

			BeforeEach(func() {

				fl = flusher.New(mockRepo, chunkSize)
			})

			It("flushes successfully", func() {

				mockRepo.EXPECT().AddClassrooms(ctx, gomock.Any()).Times(3).Return(nil)

				remainingClassrooms := fl.Flush(ctx, classrooms)

				Expect(remainingClassrooms).Should(BeNil())
			})

			It("can not flush fully", func() {

				gomock.InOrder(
					mockRepo.EXPECT().AddClassrooms(ctx, gomock.Any()).Return(nil),
					mockRepo.EXPECT().AddClassrooms(ctx, gomock.Any()).Return(errors.New("can not add classrooms")),
				)

				remainingClassrooms := fl.Flush(ctx, classrooms)

				Expect(remainingClassrooms).Should(BeEquivalentTo(classrooms[chunkSize:]))
			})

			It("can not flush anything", func() {

				mockRepo.EXPECT().AddClassrooms(ctx, gomock.Any()).Return(errors.New("can not add classrooms"))

				remainingClassrooms := fl.Flush(ctx, classrooms)

				Expect(remainingClassrooms).Should(BeEquivalentTo(classrooms))
			})
		})
	})
})
