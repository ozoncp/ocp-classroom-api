package repo

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"

	"github.com/ozoncp/ocp-classroom-api/internal/models"
)

type Repo interface {
	ListClassrooms(ctx context.Context, limit, offset uint64) ([]models.Classroom, error)
	DescribeClassroom(ctx context.Context, classroomId uint64) (*models.Classroom, error)
	AddClassroom(ctx context.Context, classroom models.Classroom) (uint64, error)
	AddClassrooms(ctx context.Context, classrooms []models.Classroom) error
	RemoveClassroom(ctx context.Context, classroomId uint64) (bool, error)
}

const tableName = "classrooms"

type classroomRepo struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) Repo {

	return &classroomRepo{db: db}
}

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
			continue
		}

		classrooms = append(classrooms, classroom)
	}

	return classrooms, nil
}

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

func (cr *classroomRepo) AddClassrooms(ctx context.Context, classrooms []models.Classroom) error {

	query := sq.Insert(tableName).
		Columns("tenant_id", "calendar_id").
		RunWith(cr.db).
		PlaceholderFormat(sq.Dollar)

	for _, classroom := range classrooms {
		query = query.Values(classroom.TenantId, classroom.CalendarId)
	}

	_, err := query.ExecContext(ctx)

	return err
}

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
