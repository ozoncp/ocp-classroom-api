package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/onsi/ginkgo"
	"google.golang.org/grpc"

	"github.com/ozoncp/ocp-classroom-api/internal/flusher"
	"github.com/ozoncp/ocp-classroom-api/internal/mocks"
	"github.com/ozoncp/ocp-classroom-api/internal/models"
	"github.com/ozoncp/ocp-classroom-api/internal/saver"

	"github.com/ozoncp/ocp-classroom-api/internal/api"
	desc "github.com/ozoncp/ocp-classroom-api/pkg/ocp-classroom-api"
)

func main() {

	introduce()

	cmd := 0
	fmt.Print("What to call? (0 - concurrency, 1 - file, 2 - grpc): ")
	fmt.Scan(&cmd)
	fmt.Println()

	if cmd == 0 {

		doConcurrencyWork()

	} else if cmd == 1 {

		doFileWork()

	} else if cmd == 2 {

		doGrpcWork()
	}
}

func introduce() {
	fmt.Println("Hello World!\r\nI'm ocp-classroom-api package by Aleksandr Kuzminykh.")
	fmt.Println()
}

func doFileWork() {

	fmt.Println("doFileWork...")

	openReadCloseFile := func(i int) {

		file, err := os.Open("hello.txt")

		if err != nil {
			fmt.Println("Unable to create file:", err)
			os.Exit(1)
		}

		defer func() {
			file.Close()
			fmt.Println("Closing file for", i+1, "th time.")
		}()

		var bytes []byte = make([]byte, 1024)

		_, err = file.Read(bytes)

		fmt.Println("Reading file for", i+1, "th time:", string(bytes))

		if err != nil {
			fmt.Println("Unable to write to file:", err)
			os.Exit(1)
		}
	}

	for i := 0; i < 10; i++ {

		openReadCloseFile(i)

		time.Sleep(1 * time.Second)
	}
}

func doConcurrencyWork() {

	fmt.Println("doConcurrencyWork...")

	ctrl := gomock.NewController(ginkgo.GinkgoT())
	mockRepo := mocks.NewMockRepo(ctrl)

	saver := saver.NewSaver(5, saver.Policy_DropAll, time.Second*15, flusher.New(mockRepo, 3))

	mockRepo.EXPECT().AddClassrooms(gomock.Any()).AnyTimes().Return(nil)

	err := saver.Init()

	if err != nil {
		fmt.Println("Can not Init saver:", err)
		os.Exit(0)
	}

	var classroomId uint64 = 0

	for {

		fmt.Print("Enter the command ('s' - save, 'x' - exit): ")

		var cmd string
		fmt.Scan(&cmd)

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

	fmt.Println("doGrpcWork...")

	const grpcPort = ":7002"

	var grpcEndpoint = flag.String("grpc-server-endpoint", "0.0.0.0"+grpcPort, "gRPC server endpoint")

	listen, err := net.Listen("tcp", grpcPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	desc.RegisterOcpClassroomApiServer(s, api.NewOcpClassroomApi())

	fmt.Printf("Server listening on %s\n", *grpcEndpoint)
	if err := s.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
