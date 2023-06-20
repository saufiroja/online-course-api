package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/saufiroja/online-course-api/config"
	"github.com/saufiroja/online-course-api/infrastructure/http/server"
)

func main() {
	conf := config.NewAppConfig()
	app := server.Server()
	port := conf.Fiber.Port

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	go func() {
		if err := app.Listen(":" + port); err != nil {
			panic(err)
		}
	}()

	log.Printf("server is running on :%s", port)

	<-stop

	log.Println("server gracefully shutdown")

	if err := app.Shutdown(); err != nil {
		panic(err)
	}

	log.Println("process clean up...")
}
