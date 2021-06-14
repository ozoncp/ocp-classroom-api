package repo

import (
	"context"
	"database/sql"

	sq "github.com/Masterminds/squirrel"

	"github.com/ozoncp/ocp-classroom-api/internal/models"
)

// Repo is interface of classrooms' storage
type Repo interface {
	ListClassrooms(ctx context.Context, limit, offset uint64) ([]models.Classroom, error)
	DescribeClassroom(ctx context.Context, classroomId uint64) (*models.Classroom, error)
	AddClassroom(ctx context.Context, classroom models.Classroom) (uint64, error)
	MultiAddClassroom(ctx context.Context, classrooms []models.Classroom) (uint64, error)
	UpdateClassroom(ctx context.Context, classroom models.Classroom) (bool, error)
	RemoveClassroom(ctx context.Context, classroomId uint64) (bool, error)
}

// tableName is name of table in PostgreSQL DB
const tableName = "classrooms"

// classroomRepo is implementation of Repo interface that uses DB for storage
type classroomRepo struct {
	db *sql.DB
}

// New returs Repo instance which uses DB for classrooms' storage
func New(db *sql.DB) Repo {

	return &classroomRepo{db: db}
}

// ListClassrooms returs list of classrooms from DB by passed limit and offset
func (cr *classroomRepo) ListClassrooms(ctx context.Context, limit, offset uint64) ([]models.Classroom, error) {

	query := sq.Select("id", "tenant_id", "calendar_id").
		From(tableName).
		RunWith(cr.db).
		Limit(limit).
		Offset(offset).
		PlaceholderFormat(sq.Dollar)

	var classrooms []models.Classroom

	rows, err := query.QueryContext(ctx)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var classroom models.Classroom
		err = rows.Scan(&classroom.Id, &classroom.TenantId, &classroom.CalendarId)

		if err != nil {
			return nil, err
		}

		classrooms = append(classrooms, classroom)
	}

	return classrooms, nil
}

// DescribeClassroom returns classroom from DB by passed id
func (cr *classroomRepo) DescribeClassroom(ctx context.Context, classroomId uint64) (*models.Classroom, error) {

	query := sq.Select("id", "tenant_id", "calendar_id").
		From(tableName).
		Where(sq.Eq{"id": classroomId}).
		RunWith(cr.db).
		PlaceholderFormat(sq.Dollar)

	var classroom models.Classroom

	if err := query.QueryRowContext(ctx).Scan(&classroom.Id, &classroom.TenantId, &classroom.CalendarId); err != nil {
		return nil, err
	}

	return &classroom, nil
}

// AddClassroom creates new classroom in DB and returs his id
func (cr *classroomRepo) AddClassroom(ctx context.Context, classroom models.Classroom) (uint64, error) {

	query := sq.Insert(tableName).
		Columns("tenant_id", "calendar_id").
		Values(classroom.TenantId, classroom.CalendarId).
		Suffix("RETURNING \"id\"").
		RunWith(cr.db).
		PlaceholderFormat(sq.Dollar)

	err := query.QueryRowContext(ctx).Scan(&classroom.Id)
	if err != nil {
		return 0, err
	}

	return classroom.Id, nil
}

// MultiAddClassroom creates new classrooms in DB and returs count of created classroms
func (cr *classroomRepo) MultiAddClassroom(ctx context.Context, classrooms []models.Classroom) (uint64, error) {

	query := sq.Insert(tableName).
		Columns("tenant_id", "calendar_id").
		RunWith(cr.db).
		PlaceholderFormat(sq.Dollar)

	for _, classroom := range classrooms {
		query = query.Values(classroom.TenantId, classroom.CalendarId)
	}

	result, err := query.ExecContext(ctx)
	if err != nil {
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return uint64(rowsAffected), nil
}

// UpdateClassroom changes classroom in DB by passed id and new tenant_id and new calendar_id
func (cr *classroomRepo) UpdateClassroom(ctx context.Context, classroom models.Classroom) (bool, error) {

	query := sq.Update(tableName).
		Set("tenant_id", classroom.TenantId).
		Set("calendar_id", classroom.CalendarId).
		Where(sq.Eq{"id": classroom.Id}).
		RunWith(cr.db).
		PlaceholderFormat(sq.Dollar)

	result, err := query.ExecContext(ctx)
	if err != nil {
		return false, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}

	return rowsAffected > 0, nil
}

// RemoveClassroom removes classroom in DB by passed id
func (cr *classroomRepo) RemoveClassroom(ctx context.Context, classroomId uint64) (bool, error) {

	query := sq.Delete(tableName).
		Where(sq.Eq{"id": classroomId}).
		RunWith(cr.db).
		PlaceholderFormat(sq.Dollar)

	result, err := query.ExecContext(ctx)
	if err != nil {
		return false, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}

	return rowsAffected > 0, nil
}
