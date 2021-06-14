package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/rs/zerolog/log"

	"github.com/ozoncp/ocp-classroom-api/internal/flusher"
	"github.com/ozoncp/ocp-classroom-api/internal/models"
	"github.com/ozoncp/ocp-classroom-api/internal/saver"
	"github.com/ozoncp/ocp-classroom-api/internal/utils"
)

func main() {

	introduce()

	cmd := 0
	fmt.Print("What to call? (0 - concurrency, 1 - file): ")
	fmt.Scan(&cmd)
	fmt.Println()

	switch cmd {
	case 0:
		doConcurrencyWork()
	case 1:
		doFileWork()
	}
}

func introduce() {
	fmt.Println("Hello World! I'm ocp-classroom-api package by Aleksandr Kuzminykh.")
}

func doFileWork() {

	const logPrefix = "FileWork: "

	log.Debug().Msg(logPrefix + "started")

	openReadCloseFile := func(i int) {

		file, err := os.Open("hello.txt")

		if err != nil {
			log.Fatal().Err(err).Msg(logPrefix + "failed to create file")
		}

		defer func() {
			file.Close()
			log.Debug().Msgf(logPrefix+"is closing file for %vth time", i+1)
		}()

		var bytes []byte = make([]byte, 1024)
		var bytesCount int

		bytesCount, err = file.Read(bytes)

		log.Debug().Str("File", string(bytes[:bytesCount])).Msgf(logPrefix+"is reading file for %vth time", i+1)

		if err != nil {
			log.Fatal().Err(err).Msg(logPrefix + "failed to write to file")
		}
	}

	for i := 0; i < 10; i++ {

		openReadCloseFile(i)

		time.Sleep(1 * time.Second)
	}
}

func doConcurrencyWork() {

	const logPrefix = "ConcurrencyWork: "

	log.Debug().Msg(logPrefix + "started")

	ctx := context.Background()

	saver, err := saver.New(5, saver.Policy_DropAll, time.Second*15, flusher.New(*utils.GetConnectedRepo(ctx), 3))
	if err != nil {
		log.Fatal().Err(err).Msg(logPrefix + "Failed to get new Saver instance")
	}

	saver.Init(ctx)
	defer saver.Close()

	for {

		var cmd string
		fmt.Print("Enter the command ('s' - save, 'x' - exit): ")
		fmt.Scan(&cmd)

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
