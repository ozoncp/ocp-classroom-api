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

	Describe("NewSaver call", func() {

		When("parameters are valid", func() {

			It("returns Saver instance", func() {

				svr := saver.NewSaver(capacity, saver.Policy_DropAll, time.Second, fl)

				Expect(svr).ShouldNot(BeNil())
			})
		})

		When("parameters are not valid", func() {

			It("returns nil if capacity is not valid", func() {

				svr := saver.NewSaver(0, saver.Policy_DropAll, time.Second, fl)

				Expect(svr).Should(BeNil())
			})

			It("returns nil if interval is not valid", func() {

				svr := saver.NewSaver(capacity, saver.Policy_DropAll, 0, fl)

				Expect(svr).Should(BeNil())
			})

			It("returns nil if flusher is nil", func() {

				svr := saver.NewSaver(capacity, saver.Policy_DropAll, time.Second, nil)

				Expect(svr).Should(BeNil())
			})
		})
	})

	Describe("Init call", func() {

		var svr saver.Saver

		BeforeEach(func() {

			svr = saver.NewSaver(capacity, saver.Policy_DropAll, time.Second, fl)
		})

		When("parameters are valid", func() {

			It("starts to work and returns nil", func() {

				err := svr.Init()

				Expect(err).Should(BeNil())

				svr.Close()
			})
		})

		When("it is already inited", func() {

			It("returns error ", func() {

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

			When("Saver still is not inited", func() {

				It("panics", func() {

					svr := saver.NewSaver(capacity, saver.Policy_DropAll, time.Second, fl)

					Expect(func() {
						classroom := models.Classroom{Id: 1, TenantId: 1, CalendarId: 1}
						svr.Save(classroom)
					}).Should(Panic())
				})
			})

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
