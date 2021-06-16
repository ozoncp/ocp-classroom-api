package main

import (
	"context"
	"fmt"

	"github.com/rs/zerolog/log"

	_ "github.com/jackc/pgx/stdlib"

	"github.com/ozoncp/ocp-classroom-api/internal/models"
	"github.com/ozoncp/ocp-classroom-api/internal/utils"
)

const logPrefix = "classroomRepo: "

func main() {

	ctx := context.Background()

	classroomRepo, err := utils.GetConnectedRepo(ctx)
	if err != nil {
		log.Fatal().Err(err).Msg(logPrefix + "failed to connect to repo")
	}

	for {

		var cmd string
		fmt.Print("What to do? ('l' - list, 'a' - add, 'ma' - multi add, 'd' - describe, 'u' - update, 'r' - remove):")
		fmt.Scan(&cmd)

		switch cmd {

		case "l":

			var limit uint64
			var offset uint64
			fmt.Print("Enter limit and offset: ")
			fmt.Scan(&limit, &offset)

			classrooms, err := classroomRepo.ListClassrooms(ctx, limit, offset)
			if err != nil {
				log.Error().Err(err).Msg(logPrefix + "failed to list classrooms")
			} else {

				log.Debug().Msgf(logPrefix+"listed classrooms: %v", classrooms)
			}

		case "a":

			classroom := *models.FromFmtScan()

			id, err := classroomRepo.AddClassroom(ctx, classroom)
			if err != nil {
				log.Error().Err(err).Msg(logPrefix + "failed to add classrooms")
			} else {

				log.Debug().Msgf(logPrefix+"added classroom, its id: %v", id)
			}

		case "ma":

			var count int
			fmt.Print("Enter count: ")
			fmt.Scan(&count)

			var classrooms []models.Classroom
			for i := 0; i < count; i++ {

				classroom := *models.FromFmtScan()
				classrooms = append(classrooms, classroom)
			}

			added_count, err := classroomRepo.MultiAddClassroom(ctx, classrooms)
			if err != nil {
				log.Error().Err(err).Msg(logPrefix + "failed to add classrooms")
			} else {

				log.Debug().Msgf(logPrefix+"added classrooms count: %v", added_count)
			}

		case "d":

			var classroom_id uint64
			fmt.Print("Enter classroom_id: ")
			fmt.Scan(&classroom_id)

			classroom, err := classroomRepo.DescribeClassroom(ctx, classroom_id)
			if err != nil {
				log.Error().Err(err).Msgf(logPrefix+"failed to describe classroom with id: %v", classroom_id)
			} else {

				log.Debug().Msgf(logPrefix+"Described classroom: %v", classroom)
			}

		case "u":

			classroom := *models.FromFmtScan()

			found, err := classroomRepo.UpdateClassroom(ctx, classroom)
			if err != nil {
				log.Error().Err(err).Msgf(logPrefix+"failed to update classroom with id: %v", classroom.Id)
			} else {

				log.Debug().Msgf(logPrefix+"updated classroom with id: %v %v", classroom, found)
			}

		case "r":

			var classroom_id uint64
			fmt.Print("Enter classroom_id: ")
			fmt.Scan(&classroom_id)

			found, err := classroomRepo.RemoveClassroom(ctx, classroom_id)
			if err != nil {
				log.Error().Err(err).Msgf(logPrefix+"failed to remove classroom with id: %v", classroom_id)
			} else {

				log.Debug().Msgf(logPrefix+"removed classroom with id: %v %v", classroom_id, found)
			}
		}
	}
}
