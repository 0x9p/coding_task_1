package main

import (
	"fmt"
	"github.com/0x9p/coding_task_1/internal"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func init() {
	env := "development"

	if err := godotenv.Load("./config/.env." + env); err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	config := internal.NewConfig()
	app := internal.NewApp(config)
	app.Init()

	exitCh := make(chan os.Signal)

	signal.Notify(exitCh,
		syscall.SIGTERM, // terminate: stopped by `kill -9 PID`
		syscall.SIGINT,  // interrupt: stopped by Ctrl + C
	)

	go func() {
		defer func() {
			exitCh <- syscall.SIGTERM // send terminate signal when application stop naturally
		}()

		err := http.ListenAndServe("127.0.0.1:2106", app)

		if err != nil {
			fmt.Println(err)
		}
	}()

	<-exitCh       // blocking until receive exit signal
	app.Shutdown() // stop the application
}
