package main

import (
	"context"
	"fmt"

	"github.com/rs/zerolog/log"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"

	"github.com/ozoncp/ocp-classroom-api/internal/models"
	"github.com/ozoncp/ocp-classroom-api/internal/repo"
)

// login: postgres
// password : postgres
// port: 5432

// user: aleksandr
// password: ozon_course

func main() {

	const dbName = "ozon"
	const address = "postgres://postgres:postgres@localhost:5432/" + dbName + "?sslmode=disable"
	//const address = "user=postgres dbname=exampledb sslmode=verify-full password=postgres"

	db, err := sqlx.Connect("pgx", address)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to postgres")
	}

	log.Debug().Msgf("Connected to DB %v", dbName)

	classroomRepo := repo.New(db)

	ctx := context.Background()

	for {

		var cmd string
		fmt.Println("What to do? ('l' - list, 'a' - add, 'am' - add many, 'd' - describe, 'r' - remove):")
		fmt.Scan(&cmd)

		switch cmd {

		case "l":
			classrooms, err := classroomRepo.ListClassrooms(ctx, 5, 0)
			if err != nil {
				log.Error().Err(err).Msg("Failed to list classrooms")
			} else {

				log.Debug().Msgf("Listed classrooms: %v", classrooms)
			}

		case "a":
			var classroom = models.Classroom{TenantId: 2, CalendarId: 2}
			id, err := classroomRepo.AddClassroom(ctx, classroom)
			if err != nil {
				log.Error().Err(err).Msg("Failed to add classrooms")
			} else {

				log.Debug().Msgf("Added classroom, its id: %v", id)
			}

		case "am":
			var classrooms = []models.Classroom{{TenantId: 3, CalendarId: 3}, {TenantId: 4, CalendarId: 4}}
			err := classroomRepo.AddClassrooms(ctx, classrooms)
			if err != nil {
				log.Error().Err(err).Msg("Failed to add classrooms")
			} else {

				log.Debug().Msgf("Added classrooms: %v", classrooms)
			}

		case "d":

			var classroomId uint64 = 3

			classroom, err := classroomRepo.DescribeClassroom(ctx, classroomId)
			if err != nil {
				log.Error().Err(err).Msgf("Failed to describe classroom with id: %v", classroomId)
			} else {

				log.Debug().Msgf("Described classroom: %v", classroom)
			}

		case "r":
			var classroomId uint64 = 2
			found, err := classroomRepo.RemoveClassroom(ctx, classroomId)
			if err != nil {
				log.Error().Err(err).Msgf("Failed to remove classroom with id: %v", classroomId)
			} else {

				log.Debug().Msgf("Removed classroom with id: %v %v", classroomId, found)
			}

		}
	}
}
