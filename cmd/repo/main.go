package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/rs/zerolog/log"

	_ "github.com/jackc/pgx/stdlib"

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

	ctx := context.Background()

	db, err := sql.Open("pgx", address)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to open postgres")
	}

	defer db.Close()

	if err := db.PingContext(ctx); err != nil {
		log.Fatal().Err(err).Msg("Failed to ping postgres")
	}

	log.Debug().Msgf("Connected to DB %v", dbName)

	classroomRepo := repo.New(db)

	for {

		var cmd string
		fmt.Println("What to do? ('l' - list, 'a' - add, 'ma' - multi add, 'd' - describe, 'u' - update, 'r' - remove):")
		fmt.Scan(&cmd)

		switch cmd {

		case "l":
			// TODO: add user input
			classrooms, err := classroomRepo.ListClassrooms(ctx, 5, 0)
			if err != nil {
				log.Error().Err(err).Msg("Failed to list classrooms")
			} else {

				log.Debug().Msgf("Listed classrooms: %v", classrooms)
			}

		case "a":
			// TODO: add user input
			var classroom = models.Classroom{TenantId: 2, CalendarId: 2}
			id, err := classroomRepo.AddClassroom(ctx, classroom)
			if err != nil {
				log.Error().Err(err).Msg("Failed to add classrooms")
			} else {

				log.Debug().Msgf("Added classroom, its id: %v", id)
			}

		case "ma":
			// TODO: add user input
			var classrooms = []models.Classroom{{TenantId: 3, CalendarId: 3}, {TenantId: 4, CalendarId: 4}}
			added_count, err := classroomRepo.MultiAddClassroom(ctx, classrooms)
			if err != nil {
				log.Error().Err(err).Msg("Failed to add classrooms")
			} else {

				log.Debug().Msgf("Added classrooms count: %v", added_count)
			}

		case "d":
			// TODO: add user input
			var classroomId uint64 = 3

			classroom, err := classroomRepo.DescribeClassroom(ctx, classroomId)
			if err != nil {
				log.Error().Err(err).Msgf("Failed to describe classroom with id: %v", classroomId)
			} else {

				log.Debug().Msgf("Described classroom: %v", classroom)
			}

		case "u":
			// TODO: add user input
			classroom := models.Classroom{Id: 8, TenantId: 66, CalendarId: 66}
			found, err := classroomRepo.UpdateClassroom(ctx, classroom)
			if err != nil {
				log.Error().Err(err).Msgf("Failed to update classroom with id: %v", classroom.Id)
			} else {

				log.Debug().Msgf("Updated classroom with id: %v %v", classroom, found)
			}

		case "r":
			// TODO: add user input
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
