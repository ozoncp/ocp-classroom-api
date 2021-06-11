package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"time"

	"github.com/rs/zerolog/log"

	"github.com/golang/mock/gomock"
	"github.com/onsi/ginkgo"

	"github.com/ozoncp/ocp-classroom-api/internal/api"
	"github.com/ozoncp/ocp-classroom-api/internal/flusher"
	"github.com/ozoncp/ocp-classroom-api/internal/mocks"
	"github.com/ozoncp/ocp-classroom-api/internal/models"
	"github.com/ozoncp/ocp-classroom-api/internal/saver"

	desc "github.com/ozoncp/ocp-classroom-api/pkg/ocp-classroom-api"
	"google.golang.org/grpc"
)

func main() {

	flag.Parse()

	introduce()

	cmd := 0
	fmt.Print("What to call? (0 - concurrency, 1 - file, 2 - grpc): ")
	fmt.Scan(&cmd)
	fmt.Println()

	log.Debug().Int("cmd", cmd).Send()

	if cmd == 0 {

		doConcurrencyWork()

	} else if cmd == 1 {

		doFileWork()

	} else if cmd == 2 {

		doGrpcWork()
	}
}

func introduce() {
	fmt.Println("Hello World! I'm ocp-classroom-api package by Aleksandr Kuzminykh.")
}

func doFileWork() {

	log.Debug().Msg("doFileWork...")

	openReadCloseFile := func(i int) {

		file, err := os.Open("hello.txt")

		if err != nil {
			log.Error().Err(err).Msg("Unable to create file")
			os.Exit(1)
		}

		defer func() {
			file.Close()
			log.Info().Msg("Closing file for " + fmt.Sprint(i+1) + " th time")
		}()

		var bytes []byte = make([]byte, 1024)
		var bytesCount int

		bytesCount, err = file.Read(bytes)

		log.Info().Str("File", string(bytes[:bytesCount])).Msg("Reading file for" + fmt.Sprint(i+1) + "th time")

		if err != nil {
			log.Error().Err(err).Msg("Unable to write to file")
			os.Exit(1)
		}
	}

	for i := 0; i < 10; i++ {

		openReadCloseFile(i)

		time.Sleep(1 * time.Second)
	}
}

func doConcurrencyWork() {

	log.Debug().Msg("doConcurrencyWork...")

	ctrl := gomock.NewController(ginkgo.GinkgoT())
	mockRepo := mocks.NewMockRepo(ctrl)

	saver, err := saver.New(5, saver.Policy_DropAll, time.Second*15, flusher.New(mockRepo, 3))

	if err != nil {
		fmt.Println("Can not get new Saver instance:", err)
		os.Exit(0)
	}

	mockRepo.EXPECT().AddClassrooms(gomock.Any()).AnyTimes().Return(nil)

	saver.Init()

	var classroomId uint64 = 0

	for {

		fmt.Print("Enter the command ('s' - save, 'x' - exit): ")

		var cmd string
		fmt.Scan(&cmd)

		log.Debug().Str("cmd", cmd).Send()

		if cmd == "s" {

			classroomId++

			saver.Save(models.Classroom{Id: classroomId, TenantId: classroomId, CalendarId: classroomId})

		} else if cmd == "x" {

			break
		}

		time.Sleep(time.Millisecond * 100)
	}

	saver.Close()
}

func doGrpcWork() {

	log.Debug().Msg("doGrpcWork...")

	const grpcPort = ":7002"
	var grpcEndpoint = *flag.String("grpc-server-endpoint", "0.0.0.0"+grpcPort, "gRPC server endpoint")

	listen, err := net.Listen("tcp", grpcEndpoint)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to listen")
		os.Exit(1)
	}

	s := grpc.NewServer()
	desc.RegisterOcpClassroomApiServer(s, api.NewOcpClassroomApi())

	log.Info().Str("gRPC server endpoint", grpcEndpoint).Msg("Server listening")
	if err := s.Serve(listen); err != nil {
		log.Fatal().Err(err).Msg("Failed to serve")
		os.Exit(1)
	}
}
