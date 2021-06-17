package saver_test

import (
	"context"
	"sync"
	"time"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/ozoncp/ocp-classroom-api/internal/flusher"
	"github.com/ozoncp/ocp-classroom-api/internal/mocks"
	"github.com/ozoncp/ocp-classroom-api/internal/models"
	"github.com/ozoncp/ocp-classroom-api/internal/saver"
)

var _ = Describe("Saver", func() {

	const capacity = 12
	const chunkSize = 3

	var (
		ctrl     *gomock.Controller
		mockRepo *mocks.MockRepo

		ctx context.Context
		fl  flusher.Flusher
	)

	BeforeEach(func() {

		ctrl = gomock.NewController(GinkgoT())
		mockRepo = mocks.NewMockRepo(ctrl)

		ctx = context.Background()
		fl = flusher.New(mockRepo, chunkSize)
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	Describe("New call", func() {

		When("parameters are valid", func() {

			It("returns Saver instance", func() {

				svr, err := saver.New(capacity, saver.Policy_DropAll, time.Second, fl)

				Expect(svr).ShouldNot(BeNil())
				Expect(err).ShouldNot(HaveOccurred())
			})
		})

		When("parameters are not valid", func() {

			It("returns nil if capacity is not valid", func() {

				svr, err := saver.New(0, saver.Policy_DropAll, time.Second, fl)

				Expect(svr).Should(BeNil())
				Expect(err).Should(HaveOccurred())
			})

			It("returns nil if interval is not valid", func() {

				svr, err := saver.New(capacity, saver.Policy_DropAll, 0, fl)

				Expect(svr).Should(BeNil())
				Expect(err).Should(HaveOccurred())
			})

			It("returns nil if flusher is nil", func() {

				svr, err := saver.New(capacity, saver.Policy_DropAll, time.Second, nil)

				Expect(svr).Should(BeNil())
				Expect(err).Should(HaveOccurred())
			})
		})
	})

	Describe("Init goroutine work", func() {

		When("Save is called", func() {

			When("capacity is not reached", func() {

				It("flushes all saved classrooms", func() {

					svr, err := saver.New(capacity, saver.Policy_DropAll, time.Second, fl)

					Expect(svr).ShouldNot(BeNil())
					Expect(err).ShouldNot(HaveOccurred())

					svr.Init(ctx)

					var wg sync.WaitGroup

					classroomsLen := capacity - 1

					wg.Add((classroomsLen + 1) / chunkSize)

					mockRepo.EXPECT().AddClassrooms(ctx, gomock.Any()).
						AnyTimes().
						Do(func(ctx context.Context, cr []models.Classroom) {

							wg.Done()

						}).Return(nil)

					for i := 0; i < classroomsLen; i++ {

						var id uint64 = uint64(i)

						svr.Save(models.Classroom{Id: id, TenantId: id, CalendarId: id})
					}

					wg.Wait()
					svr.Close()
				})
			})

			When("capacity is reached", func() {

				It("flushes only new classrooms after dropping all", func() {

					svr, err := saver.New(capacity, saver.Policy_DropAll, time.Second, fl)

					Expect(svr).ShouldNot(BeNil())
					Expect(err).ShouldNot(HaveOccurred())

					svr.Init(ctx)

					var wg sync.WaitGroup

					classroomsLen := capacity + 3

					wg.Add((classroomsLen - capacity + 1) / chunkSize)

					mockRepo.EXPECT().AddClassrooms(ctx, gomock.Any()).
						AnyTimes().
						Do(func(ctx context.Context, cr []models.Classroom) {

							wg.Done()

						}).Return(nil)

					for i := 0; i < classroomsLen; i++ {

						var id uint64 = uint64(i)

						svr.Save(models.Classroom{Id: id, TenantId: id, CalendarId: id})
					}

					wg.Wait()
					svr.Close()
				})

				It("flushes classrooms after dropping first", func() {

					svr, err := saver.New(capacity, saver.Policy_DropFirst, time.Second, fl)

					Expect(svr).ShouldNot(BeNil())
					Expect(err).ShouldNot(HaveOccurred())

					svr.Init(ctx)

					var wg sync.WaitGroup

					classroomsLen := capacity + 3

					wg.Add((capacity + 1) / chunkSize)

					mockRepo.EXPECT().AddClassrooms(ctx, gomock.Any()).
						AnyTimes().
						Do(func(ctx context.Context, cr []models.Classroom) {

							wg.Done()

						}).Return(nil)

					for i := 0; i < classroomsLen; i++ {

						var id uint64 = uint64(i)

						svr.Save(models.Classroom{Id: id, TenantId: id, CalendarId: id})
					}

					wg.Wait()
					svr.Close()
				})
			})
		})

		When("Close is called", func() {

			When("there were not flushes by ticker", func() {

				It("flushes all classrooms", func() {

					svr, err := saver.New(capacity, saver.Policy_DropAll, time.Minute, fl)

					Expect(svr).ShouldNot(BeNil())
					Expect(err).ShouldNot(HaveOccurred())

					svr.Init(ctx)

					var wg sync.WaitGroup

					classroomsLen := capacity - 1

					wg.Add((classroomsLen + 1) / chunkSize)

					mockRepo.EXPECT().AddClassrooms(ctx, gomock.Any()).
						AnyTimes().
						Do(func(ctx context.Context, cr []models.Classroom) {

							wg.Done()

						}).Return(nil)

					for i := 0; i < classroomsLen; i++ {

						var id uint64 = uint64(i)

						svr.Save(models.Classroom{Id: id, TenantId: id, CalendarId: id})
					}

					svr.Close()
					wg.Wait()
				})
			})
		})
	})
})
