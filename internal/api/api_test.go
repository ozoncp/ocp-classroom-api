package api_test

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/DATA-DOG/go-sqlmock"

	"github.com/ozoncp/ocp-classroom-api/internal/api"
	"github.com/ozoncp/ocp-classroom-api/internal/mocks"
	"github.com/ozoncp/ocp-classroom-api/internal/repo"
	grpcApi "github.com/ozoncp/ocp-classroom-api/pkg/ocp-classroom-api"
)

var _ = Describe("Api", func() {

	var (
		ctx context.Context

		db   *sql.DB
		mock sqlmock.Sqlmock

		classroomsRepo repo.Repo

		ctrl         *gomock.Controller
		mockProducer *mocks.MockLogProducer

		apiServer grpcApi.OcpClassroomApiServer
	)

	BeforeEach(func() {

		ctx = context.Background()

		var err error
		db, mock, err = sqlmock.New()
		if err != nil {
			Fail("can not create mocksql")
		}

		classroomsRepo = repo.New(db)

		ctrl = gomock.NewController(GinkgoT())
		mockProducer = mocks.NewMockLogProducer(ctrl)

		mockProducer.EXPECT().Send(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(nil)

		apiServer = api.NewOcpClassroomApi(classroomsRepo, mockProducer)
	})

	AfterEach(func() {

		db.Close()
	})

	Describe("ListClassrooms call", func() {

		When("parameters are valid", func() {

			It("returns classrooms if query executes", func() {

				request := &grpcApi.ListClassroomsV1Request{
					Limit:  5,
					Offset: 0,
				}

				rows := sqlmock.NewRows([]string{"id", "tenant_id", "calendar_id"}).
					AddRow(1, 1, 1).
					AddRow(2, 2, 2).
					AddRow(3, 3, 3).
					AddRow(4, 4, 4).
					AddRow(5, 5, 5)

				mock.ExpectQuery("SELECT id, tenant_id, calendar_id FROM classrooms LIMIT " +
					fmt.Sprint(request.Limit) + " OFFSET " + fmt.Sprint(request.Offset)).
					WillReturnRows(rows)

				res, err := apiServer.ListClassroomsV1(ctx, request)

				Expect(res).ShouldNot(BeNil())
				Expect(res.Classrooms).Should(HaveLen(5))
				Expect(err).ShouldNot(HaveOccurred())

				if err := mock.ExpectationsWereMet(); err != nil {
					Fail("there were unfulfilled expectations: " + err.Error())
				}
			})

			It("returns error if query returns error", func() {

				request := &grpcApi.ListClassroomsV1Request{
					Limit:  5,
					Offset: 0,
				}

				mock.ExpectQuery("SELECT id, tenant_id, calendar_id FROM classrooms LIMIT " +
					fmt.Sprint(request.Limit) + " OFFSET " + fmt.Sprint(request.Offset)).
					WillReturnError(errors.New(""))

				res, err := apiServer.ListClassroomsV1(ctx, request)

				Expect(res).Should(BeNil())
				Expect(err).Should(HaveOccurred())

				if err := mock.ExpectationsWereMet(); err != nil {
					Fail("there were unfulfilled expectations: " + err.Error())
				}
			})
		})

		When("parameters are not valid", func() {

			It("returns error if limit is 0 or offset is 0", func() {

				request := &grpcApi.ListClassroomsV1Request{
					Limit:  0,
					Offset: 0,
				}

				res, err := apiServer.ListClassroomsV1(ctx, request)

				Expect(res).Should(BeNil())
				Expect(err).Should(HaveOccurred())
			})

			It("panics if request is nil", func() {

				Expect(func() {

					_, err := apiServer.ListClassroomsV1(ctx, nil)
					if err != nil {
						Fail("error occured")
					}

				}).Should(Panic())
			})
		})
	})

	Describe("DescribeClassroom call", func() {

		When("parameters are valid", func() {

			It("returns classroom if query executes", func() {

				request := &grpcApi.DescribeClassroomV1Request{
					ClassroomId: 1,
				}

				rows := sqlmock.NewRows([]string{"id", "tenant_id", "calendar_id"}).
					AddRow(request.ClassroomId, 1, 1)

				mock.ExpectQuery("SELECT id, tenant_id, calendar_id FROM classrooms WHERE").
					WithArgs(request.ClassroomId).
					WillReturnRows(rows)

				res, err := apiServer.DescribeClassroomV1(ctx, request)

				Expect(res).ShouldNot(BeNil())
				Expect(res.Classroom.ClassroomId).Should(BeEquivalentTo(request.ClassroomId))
				Expect(err).ShouldNot(HaveOccurred())

				if err := mock.ExpectationsWereMet(); err != nil {
					Fail("there were unfulfilled expectations: " + err.Error())
				}
			})

			It("returns error if query returns error", func() {

				request := &grpcApi.DescribeClassroomV1Request{
					ClassroomId: 1,
				}

				mock.ExpectQuery("SELECT id, tenant_id, calendar_id FROM classrooms WHERE").
					WithArgs(request.ClassroomId).
					WillReturnError(errors.New(""))

				res, err := apiServer.DescribeClassroomV1(ctx, request)

				Expect(res).Should(BeNil())
				Expect(err).Should(HaveOccurred())

				if err := mock.ExpectationsWereMet(); err != nil {
					Fail("there were unfulfilled expectations: " + err.Error())
				}
			})
		})

		When("parameters are not valid", func() {

			It("returns error if classroom_id is 0", func() {

				request := &grpcApi.DescribeClassroomV1Request{
					ClassroomId: 0,
				}

				res, err := apiServer.DescribeClassroomV1(ctx, request)

				Expect(res).Should(BeNil())
				Expect(err).Should(HaveOccurred())
			})

			It("panics if request is nil", func() {

				Expect(func() {

					_, err := apiServer.DescribeClassroomV1(ctx, nil)
					if err != nil {
						Fail("error occured")
					}

				}).Should(Panic())
			})
		})
	})

	Describe("CreateClassroom call", func() {

		When("parameters are valid", func() {

			It("returns classroom_id if query executes", func() {

				request := &grpcApi.CreateClassroomV1Request{
					TenantId:   1,
					CalendarId: 1,
				}

				expectedId := 5

				mock.ExpectQuery("INSERT INTO classrooms").
					WithArgs(request.TenantId, request.CalendarId).
					WillReturnRows(sqlmock.NewRows([]string{"id"}).
						AddRow(expectedId))

				res, err := apiServer.CreateClassroomV1(ctx, request)

				Expect(res).ShouldNot(BeNil())
				Expect(res.ClassroomId).Should(BeEquivalentTo(expectedId))
				Expect(err).ShouldNot(HaveOccurred())

				if err := mock.ExpectationsWereMet(); err != nil {
					Fail("there were unfulfilled expectations: " + err.Error())
				}
			})

			It("returns error if query returns error", func() {

				request := &grpcApi.CreateClassroomV1Request{
					TenantId:   1,
					CalendarId: 1,
				}

				mock.ExpectQuery("INSERT INTO classrooms").
					WithArgs(request.TenantId, request.CalendarId).
					WillReturnError(errors.New(""))

				res, err := apiServer.CreateClassroomV1(ctx, request)

				Expect(res).Should(BeNil())
				Expect(err).Should(HaveOccurred())

				if err := mock.ExpectationsWereMet(); err != nil {
					Fail("there were unfulfilled expectations: " + err.Error())
				}
			})
		})

		When("parameters are not valid", func() {

			It("returns error if tenant_id is 0 or calendar_id is 0", func() {

				request := &grpcApi.CreateClassroomV1Request{
					TenantId:   0,
					CalendarId: 0,
				}

				res, err := apiServer.CreateClassroomV1(ctx, request)

				Expect(res).Should(BeNil())
				Expect(err).Should(HaveOccurred())
			})

			It("panics if request is nil", func() {

				Expect(func() {

					_, err := apiServer.CreateClassroomV1(ctx, nil)
					if err != nil {
						Fail("error occured")
					}

				}).Should(Panic())
			})
		})
	})

	Describe("MultiCreateClassroom call", func() {

		When("parameters are valid", func() {

			It("returns created_count if query executes", func() {

				request := &grpcApi.MultiCreateClassroomV1Request{
					Classrooms: []*grpcApi.CreateClassroomV1Request{
						{TenantId: 1, CalendarId: 1},
						{TenantId: 2, CalendarId: 2},
					},
				}

				mock.ExpectExec("INSERT INTO classrooms").
					WithArgs(request.Classrooms[0].TenantId, request.Classrooms[0].CalendarId,
						request.Classrooms[1].TenantId, request.Classrooms[1].CalendarId).
					WillReturnResult(sqlmock.NewResult(2, 2))

				res, err := apiServer.MultiCreateClassroomV1(ctx, request)

				Expect(res).ShouldNot(BeNil())
				Expect(res.CreatedCount).Should(BeEquivalentTo(2))
				Expect(err).ShouldNot(HaveOccurred())

				if err := mock.ExpectationsWereMet(); err != nil {
					Fail("there were unfulfilled expectations: " + err.Error())
				}
			})

			It("returns error if query returns error", func() {

				request := &grpcApi.MultiCreateClassroomV1Request{
					Classrooms: []*grpcApi.CreateClassroomV1Request{
						{TenantId: 1, CalendarId: 1},
						{TenantId: 2, CalendarId: 2},
					},
				}

				mock.ExpectExec("INSERT INTO classrooms").
					WithArgs(request.Classrooms[0].TenantId, request.Classrooms[0].CalendarId,
						request.Classrooms[1].TenantId, request.Classrooms[1].CalendarId).
					WillReturnError(errors.New(""))

				res, err := apiServer.MultiCreateClassroomV1(ctx, request)

				Expect(res).Should(BeNil())
				Expect(err).Should(HaveOccurred())

				if err := mock.ExpectationsWereMet(); err != nil {
					Fail("there were unfulfilled expectations: " + err.Error())
				}
			})
		})

		When("parameters are not valid", func() {

			It("returns error if tenant_id is 0 or calendar_id is 0", func() {

				for i := uint64(1); i < 3; i++ {

					request := &grpcApi.MultiCreateClassroomV1Request{
						Classrooms: []*grpcApi.CreateClassroomV1Request{
							{TenantId: i - 1, CalendarId: i % 2},
						}}

					res, err := apiServer.MultiCreateClassroomV1(ctx, request)

					Expect(res).Should(BeNil())
					Expect(err).Should(HaveOccurred())
				}
			})

			It("panics if request is nil", func() {

				Expect(func() {

					_, err := apiServer.MultiCreateClassroomV1(ctx, nil)
					if err != nil {
						Fail("error occured")
					}

				}).Should(Panic())
			})
		})
	})

	Describe("UpdateClassroom call", func() {

		When("parameters are valid", func() {

			It("returns found as true if query executes with affected rows", func() {

				request := &grpcApi.UpdateClassroomV1Request{
					Classroom: &grpcApi.Classroom{
						ClassroomId: 1,
						TenantId:    1,
						CalendarId:  1,
					},
				}

				mock.ExpectExec("UPDATE classrooms SET").
					WithArgs(request.Classroom.TenantId, request.Classroom.CalendarId, request.Classroom.ClassroomId).
					WillReturnResult(sqlmock.NewResult(1, 1))

				res, err := apiServer.UpdateClassroomV1(ctx, request)

				Expect(res).ShouldNot(BeNil())
				Expect(res.Found).Should(BeEquivalentTo(true))
				Expect(err).ShouldNot(HaveOccurred())

				if err := mock.ExpectationsWereMet(); err != nil {
					Fail("there were unfulfilled expectations: " + err.Error())
				}
			})

			It("returns found as false if query executes without affected rows", func() {

				request := &grpcApi.UpdateClassroomV1Request{
					Classroom: &grpcApi.Classroom{
						ClassroomId: 1,
						TenantId:    1,
						CalendarId:  1,
					},
				}

				mock.ExpectExec("UPDATE classrooms SET").
					WithArgs(request.Classroom.TenantId, request.Classroom.CalendarId, request.Classroom.ClassroomId).
					WillReturnResult(sqlmock.NewResult(1, 0))

				res, err := apiServer.UpdateClassroomV1(ctx, request)

				Expect(res).ShouldNot(BeNil())
				Expect(res.Found).Should(BeEquivalentTo(false))
				Expect(err).ShouldNot(HaveOccurred())

				if err := mock.ExpectationsWereMet(); err != nil {
					Fail("there were unfulfilled expectations: " + err.Error())
				}
			})

			It("returns error if query returns error", func() {

				request := &grpcApi.UpdateClassroomV1Request{
					Classroom: &grpcApi.Classroom{
						ClassroomId: 1,
						TenantId:    1,
						CalendarId:  1,
					},
				}

				mock.ExpectExec("UPDATE classrooms SET").
					WithArgs(request.Classroom.TenantId, request.Classroom.CalendarId, request.Classroom.ClassroomId).
					WillReturnError(errors.New(""))

				res, err := apiServer.UpdateClassroomV1(ctx, request)

				Expect(res).Should(BeNil())
				Expect(err).Should(HaveOccurred())

				if err := mock.ExpectationsWereMet(); err != nil {
					Fail("there were unfulfilled expectations: " + err.Error())
				}
			})
		})

		When("parameters are not valid", func() {

			It("returns error if any param is 0", func() {

				for i := uint64(1); i < 4; i++ {

					request := &grpcApi.UpdateClassroomV1Request{
						Classroom: &grpcApi.Classroom{
							ClassroomId: i - 1,
							TenantId:    i % 2,
							CalendarId:  i % 3,
						},
					}

					res, err := apiServer.UpdateClassroomV1(ctx, request)

					Expect(res).Should(BeNil())
					Expect(err).Should(HaveOccurred())
				}
			})

			It("panics if request is nil", func() {

				Expect(func() {

					_, err := apiServer.UpdateClassroomV1(ctx, nil)
					if err != nil {
						Fail("error occured")
					}

				}).Should(Panic())
			})
		})
	})

	Describe("RemoveClassroom call", func() {

		When("parameters are valid", func() {

			It("returns found as true if query executes with affected rows", func() {

				request := &grpcApi.RemoveClassroomV1Request{
					ClassroomId: 1,
				}

				mock.ExpectExec("DELETE FROM classrooms WHERE").
					WithArgs(request.ClassroomId).
					WillReturnResult(sqlmock.NewResult(1, 1))

				res, err := apiServer.RemoveClassroomV1(ctx, request)

				Expect(res).ShouldNot(BeNil())
				Expect(res.Found).Should(BeEquivalentTo(true))
				Expect(err).ShouldNot(HaveOccurred())

				if err := mock.ExpectationsWereMet(); err != nil {
					Fail("there were unfulfilled expectations: " + err.Error())
				}
			})

			It("returns found as false if query executes without affected rows", func() {

				request := &grpcApi.RemoveClassroomV1Request{
					ClassroomId: 1,
				}

				mock.ExpectExec("DELETE FROM classrooms WHERE").
					WithArgs(request.ClassroomId).
					WillReturnResult(sqlmock.NewResult(1, 0))

				res, err := apiServer.RemoveClassroomV1(ctx, request)

				Expect(res).ShouldNot(BeNil())
				Expect(res.Found).Should(BeEquivalentTo(false))
				Expect(err).ShouldNot(HaveOccurred())

				if err := mock.ExpectationsWereMet(); err != nil {
					Fail("there were unfulfilled expectations: " + err.Error())
				}
			})

			It("returns error if query returns error", func() {

				request := &grpcApi.RemoveClassroomV1Request{
					ClassroomId: 1,
				}

				mock.ExpectExec("DELETE FROM classrooms WHERE").
					WithArgs(request.ClassroomId).
					WillReturnError(errors.New(""))

				res, err := apiServer.RemoveClassroomV1(ctx, request)

				Expect(res).Should(BeNil())
				Expect(err).Should(HaveOccurred())

				if err := mock.ExpectationsWereMet(); err != nil {
					Fail("there were unfulfilled expectations: " + err.Error())
				}
			})
		})

		When("parameters are not valid", func() {

			It("returns error if classroom_id is 0", func() {

				request := &grpcApi.RemoveClassroomV1Request{
					ClassroomId: 0,
				}

				res, err := apiServer.RemoveClassroomV1(ctx, request)

				Expect(res).Should(BeNil())
				Expect(err).Should(HaveOccurred())
			})

			It("panics if request is nil", func() {

				Expect(func() {

					_, err := apiServer.RemoveClassroomV1(ctx, nil)
					if err != nil {
						Fail("error occured")
					}

				}).Should(Panic())
			})
		})
	})
})
