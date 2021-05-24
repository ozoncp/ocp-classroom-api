package models

import (
	"fmt"
)

type Classroom struct {
	id uint

	tenantId   uint
	calendarId uint
}

func New(id uint, tenantId uint, calendarId uint) (cr Classroom) {

	cr.id = id
	cr.tenantId = tenantId
	cr.calendarId = calendarId

	return
}

func (cr *Classroom) Id() uint {
	return cr.id
}

func (cr *Classroom) TenantId() uint {
	return cr.tenantId
}

func (cr *Classroom) CalendarId() uint {
	return cr.calendarId
}

func (cr *Classroom) String() (str string) {

	str = fmt.Sprintf("Classroom = { id: %v, tenantId: %v, calendarId: %v }",
		cr.id, cr.tenantId, cr.calendarId)

	return
}
