package models

import (
	"errors"
	"fmt"
)

type Classroom struct {
	id uint

	tenantsIds   []uint
	calendarsIds []uint
}

func New(id uint, tenantsIds []uint, calendarsIds []uint) (cr Classroom) {

	cr.id = id
	cr.tenantsIds = tenantsIds
	cr.calendarsIds = calendarsIds

	return
}

func (cr *Classroom) ClassroomId() uint {
	return cr.id
}

func (cr *Classroom) AddTenantId(tenantId uint) {
	cr.tenantsIds = append(cr.tenantsIds, tenantId)
}

func (cr *Classroom) TenantId(index int) (uint, error) {

	if index < 0 {
		return 0, errors.New("index less 0")
	}

	if index >= len(cr.tenantsIds) {
		return 0, errors.New("index out of range")
	}

	return cr.tenantsIds[index], nil
}

func (cr *Classroom) AddCalendarId(calendarId uint) {
	cr.calendarsIds = append(cr.calendarsIds, calendarId)
}

func (cr *Classroom) CalendarId(index int) (uint, error) {

	if index < 0 {
		return 0, errors.New("index less 0")
	}

	if index >= len(cr.calendarsIds) {
		return 0, errors.New("index out of range")
	}

	return cr.calendarsIds[index], nil
}

func (cr *Classroom) String() (str string) {
	str = fmt.Sprintf("Classroom = { id: %v, tenantsIds: { %v }, calendarsIds: { %v } }",
		cr.id, cr.tenantsIds, cr.calendarsIds)

	return
}
