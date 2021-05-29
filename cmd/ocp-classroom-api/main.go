package main

import (
	"fmt"
	"os"
	"time"
)

func main() {

	fmt.Println("Hello World!\r\nI'm ocp-classroom-api package by Aleksandr Kuzminykh.")

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
