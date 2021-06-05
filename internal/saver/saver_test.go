package saver_test

import (
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

		fl flusher.Flusher
	)

	BeforeEach(func() {

		ctrl = gomock.NewController(GinkgoT())
		mockRepo = mocks.NewMockRepo(ctrl)

		fl = flusher.New(mockRepo, chunkSize)
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	Describe("Init call", func() {

		When("parameters are valid", func() {

			It("starts to work and returns nil", func() {

				svr := saver.NewSaver(capacity, saver.Policy_DropAll, time.Second, fl)

				err := svr.Init()

				Expect(err).Should(BeNil())

				svr.Close()
			})
		})

		When("parameters are not valid", func() {

			It("returns error if capacity is wrong", func() {

				svr := saver.NewSaver(0, saver.Policy_DropAll, time.Second, fl)

				err := svr.Init()

				Expect(err).Should(HaveOccurred())
			})

			It("returns error if interval is wrong", func() {

				svr := saver.NewSaver(capacity, saver.Policy_DropAll, 0, fl)

				err := svr.Init()

				Expect(err).Should(HaveOccurred())
			})

			It("returns error if flusher is nil", func() {

				svr := saver.NewSaver(capacity, saver.Policy_DropAll, time.Second, nil)

				err := svr.Init()

				Expect(err).Should(HaveOccurred())
			})

			It("returns error if Saver is already inited", func() {

				svr := saver.NewSaver(capacity, saver.Policy_DropAll, time.Second, fl)

				err := svr.Init()

				Expect(err).Should(BeNil())

				err = svr.Init()

				Expect(err).Should(HaveOccurred())

				svr.Close()
			})
		})
	})

	Describe("Init goroutine work", func() {

		When("Save is called", func() {

			When("capacity is not reached", func() {

				It("flushes all saved classrooms", func() {

					svr := saver.NewSaver(capacity, saver.Policy_DropAll, time.Second, fl)

					if err := svr.Init(); err != nil {
						Fail("Init call failed")
					}

					var wg sync.WaitGroup

					classroomsLen := capacity - 1

					wg.Add((classroomsLen + 1) / chunkSize)

					mockRepo.EXPECT().AddClassrooms(gomock.Any()).AnyTimes().Do(func(cr []models.Classroom) {

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

					svr := saver.NewSaver(capacity, saver.Policy_DropAll, time.Second, fl)

					if err := svr.Init(); err != nil {
						Fail("Init call failed")
					}

					var wg sync.WaitGroup

					classroomsLen := capacity + 3

					wg.Add((classroomsLen - capacity + 1) / chunkSize)

					mockRepo.EXPECT().AddClassrooms(gomock.Any()).AnyTimes().Do(func(cr []models.Classroom) {

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

					svr := saver.NewSaver(capacity, saver.Policy_DropFirst, time.Second, fl)

					if err := svr.Init(); err != nil {
						Fail("Init call failed")
					}

					var wg sync.WaitGroup

					classroomsLen := capacity + 3

					wg.Add((capacity + 1) / chunkSize)

					mockRepo.EXPECT().AddClassrooms(gomock.Any()).AnyTimes().Do(func(cr []models.Classroom) {

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

					svr := saver.NewSaver(capacity, saver.Policy_DropAll, time.Minute, fl)

					if err := svr.Init(); err != nil {
						Fail("Init call failed")
					}

					var wg sync.WaitGroup

					classroomsLen := capacity - 1

					wg.Add((classroomsLen + 1) / chunkSize)

					mockRepo.EXPECT().AddClassrooms(gomock.Any()).AnyTimes().Do(func(cr []models.Classroom) {

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
