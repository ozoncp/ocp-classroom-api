package models

import (
	"fmt"
)

type Classroom struct {
	Id uint64

	TenantId   uint64
	CalendarId uint64
}

func New(id uint64, tenantId uint64, calendarId uint64) (cr Classroom) {

	cr.Id = id
	cr.TenantId = tenantId
	cr.CalendarId = calendarId

	return
}

func (cr *Classroom) String() (str string) {

	str = fmt.Sprintf("Classroom = { id: %v, tenantId: %v, calendarId: %v }",
		cr.Id, cr.TenantId, cr.CalendarId)

	return
}
