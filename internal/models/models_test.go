package models_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/ozoncp/ocp-classroom-api/internal/models"
	grpcApi "github.com/ozoncp/ocp-classroom-api/pkg/ocp-classroom-api"
)

var _ = Describe("Models", func() {

	var (
		classroom      models.Classroom
		protoClassroom *grpcApi.Classroom
	)

	BeforeEach(func() {

		classroom = models.Classroom{Id: 1, TenantId: 1, CalendarId: 1}
		protoClassroom = &grpcApi.Classroom{ClassroomId: 1, TenantId: 1, CalendarId: 1}
	})

	Describe("String call", func() {

		It("returns string", func() {

			Expect(classroom.String()).Should(
				BeEquivalentTo("Classroom = { Id: 1, TenantId: 1, CalendarId: 1 }"))
		})
	})

	Describe("ToProtoClassroom call", func() {

		It("returns protoClassroom", func() {

			Expect(classroom.ToProtoClassroom()).Should(
				BeEquivalentTo(protoClassroom))
		})
	})

	Describe("FromProtoClassroom call", func() {

		It("return Classroom", func() {

			Expect(*models.FromProtoClassroom(protoClassroom)).Should(
				BeEquivalentTo(classroom))
		})
	})

	Describe("FromFmtScan call", func() {

		It("returns classroom", func() {

			// It is impossible to test
		})
	})

})
