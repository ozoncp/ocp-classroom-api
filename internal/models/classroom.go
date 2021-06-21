package models

import (
	"fmt"

	grpcApi "github.com/ozoncp/ocp-classroom-api/pkg/ocp-classroom-api"
)

// TODO: add tests

type Classroom struct {
	Id uint64 `db:"id"`

	TenantId   uint64 `db:"tenant_id"`
	CalendarId uint64 `db:"calendar_id"`
}

func (cr *Classroom) String() (str string) {

	str = fmt.Sprintf("Classroom = { id: %v, tenantId: %v, calendarId: %v }",
		cr.Id, cr.TenantId, cr.CalendarId)

	return
}

func (cr *Classroom) ToProtoClassroom() *grpcApi.Classroom {

	return &grpcApi.Classroom{
		ClassroomId: cr.Id,
		TenantId:    cr.TenantId,
		CalendarId:  cr.CalendarId,
	}
}

func FromProtoClassroom(protoClassroom *grpcApi.Classroom) *Classroom {

	return &Classroom{
		Id:         protoClassroom.ClassroomId,
		TenantId:   protoClassroom.TenantId,
		CalendarId: protoClassroom.CalendarId,
	}
}

func FromFmtScan() *Classroom {

	var tenantId uint64
	var calendarId uint64

	for {
		fmt.Print("Enter tenantId and calendarId: ")

		if _, err := fmt.Scan(&tenantId, &calendarId); err != nil {
			fmt.Println("Error occurred", err, ". Try again")
			continue
		}

		break
	}

	return &Classroom{TenantId: tenantId, CalendarId: calendarId}
}
