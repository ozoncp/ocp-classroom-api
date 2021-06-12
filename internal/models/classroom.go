package models

import (
	"fmt"
)

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
