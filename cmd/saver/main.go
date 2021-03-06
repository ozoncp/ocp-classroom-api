package main

import (
	"context"
	"flag"
	"fmt"
	"time"

	"github.com/ozoncp/ocp-classroom-api/internal/flusher"
	"github.com/ozoncp/ocp-classroom-api/internal/models"
	"github.com/ozoncp/ocp-classroom-api/internal/repo"
	"github.com/ozoncp/ocp-classroom-api/internal/saver"
	"github.com/rs/zerolog/log"
)

func main() {

	var repoArgs = repo.RepoArgs{}
	repo.SetRepoArgsFromCommandLine(&repoArgs)

	flag.Parse()

	const logPrefix = "ConcurrencyWork: "

	log.Debug().Msg(logPrefix + "started")

	ctx := context.Background()

	repo, err := repo.GetConnectedRepo(ctx, &repoArgs)
	if err != nil {
		log.Fatal().Err(err).Msg(logPrefix + "failed to connect to repo")
	}

	saver, err := saver.New(5, saver.Policy_DropAll, time.Second*15, flusher.New(repo, 3))
	if err != nil {
		log.Fatal().Err(err).Msg(logPrefix + "failed to get new Saver instance")
	}

	saver.Init(ctx)
	defer saver.Close()

	for {

		var cmd string
		for {
			fmt.Print("Enter the command ('s' - save, 'x' - exit): ")

			if _, err := fmt.Scan(&cmd); err != nil {
				fmt.Println("Error occurred", err, ". Try again")
				continue
			}

			break
		}

		if cmd == "s" {

			classroom := *models.FromFmtScan()

			saver.Save(classroom)

			log.Debug().Msg(logPrefix + "saved")

		} else if cmd == "x" {

			log.Debug().Msg(logPrefix + "finished")

			break
		}

		time.Sleep(time.Millisecond * 100)
	}
}
