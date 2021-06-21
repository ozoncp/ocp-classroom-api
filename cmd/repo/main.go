package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/rs/zerolog/log"

	_ "github.com/jackc/pgx/stdlib"

	"github.com/ozoncp/ocp-classroom-api/internal/models"
	"github.com/ozoncp/ocp-classroom-api/internal/repo"
)

const logPrefix = "classroomRepo: "

func main() {

	var repoArgs = repo.RepoArgs{}
	repo.SetRepoArgsFromCommandLine(&repoArgs)

	flag.Parse()

	ctx := context.Background()

	classroomRepo, err := repo.GetConnectedRepo(ctx, &repoArgs)
	if err != nil {
		log.Fatal().Err(err).Msg(logPrefix + "failed to connect to repo")
	}

	for {

		switch getCommandFromUserInput() {

		case "l":

			limit, offset := getLimitOffsetFromUserInput()

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

			addedCount, err := classroomRepo.MultiAddClassroom(ctx, getClassroomsFromUserInput())
			if err != nil {
				log.Error().Err(err).Msg(logPrefix + "failed to add classrooms")
			} else {

				log.Debug().Msgf(logPrefix+"added classrooms count: %v", addedCount)
			}

		case "d":

			classroomId := getClassroomIdFromUserInput()

			classroom, err := classroomRepo.DescribeClassroom(ctx, classroomId)
			if err != nil {
				log.Error().Err(err).Msgf(logPrefix+"failed to describe classroom with id: %v", classroomId)
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

			classroomId := getClassroomIdFromUserInput()

			found, err := classroomRepo.RemoveClassroom(ctx, classroomId)
			if err != nil {
				log.Error().Err(err).Msgf(logPrefix+"failed to remove classroom with id: %v", classroomId)
			} else {

				log.Debug().Msgf(logPrefix+"removed classroom with id: %v %v", classroomId, found)
			}

		case "x":

			return
		}
	}
}

func getCommandFromUserInput() (cmd string) {

	for {
		fmt.Print("What to do? (",
			"'l' - list,",
			" 'a' - add,",
			" 'ma' - multi add,",
			" 'd' - describe,",
			" 'u' - update,",
			" 'r' - remove,",
			" 'x' - exit): ")

		if _, err := fmt.Scan(&cmd); err != nil {
			fmt.Println("Error occurred", err, ". Try again")
			continue
		}

		break
	}

	return
}

func getLimitOffsetFromUserInput() (limit, offset uint64) {

	for {
		fmt.Print("Enter limit and offset: ")

		if _, err := fmt.Scan(&limit, &offset); err != nil {
			fmt.Println("Error occurred", err, ". Try again")
			continue
		}

		break
	}

	return
}

func getClassroomsFromUserInput() []models.Classroom {

	var count int
	for {
		fmt.Print("Enter count: ")

		if _, err := fmt.Scan(&count); err != nil {
			fmt.Println("Error occurred", err, ". Try again")
			continue
		}

		if count < 1 {
			fmt.Println("Count can not be less 1")
			continue
		}

		break
	}

	var classrooms []models.Classroom
	for i := 0; i < count; i++ {

		classroom := *models.FromFmtScan()
		classrooms = append(classrooms, classroom)
	}

	return classrooms
}

func getClassroomIdFromUserInput() (classroomId uint64) {

	for {
		fmt.Print("Enter classroomId: ")

		if _, err := fmt.Scan(&classroomId); err != nil {
			fmt.Println("Error occurred", err, ". Try again")
			continue
		}

		break
	}

	return
}
