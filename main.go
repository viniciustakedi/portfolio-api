package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"portfolio/api/config"
	"portfolio/api/infra/db"
	"portfolio/api/server"
	"syscall"
)

func main() {
	environment := flag.String("e", "development", "Environment to run the application in (development, staging, production)")

	flag.Usage = func() {
		log.Fatalf(
			"Usage: %s [options]\nOptions:\n  -env string\n\tEnvironment to run the application in (development, staging, production)",
			flag.CommandLine.Name(),
		)

		os.Exit(1)
	}

	flag.Parse()

	config.Init(*environment)

	if err := db.InitMongoDB(); err != nil {
		log.Fatalf("Error to init mongodb client: ", err.Error())
		os.Exit(1)
	}

	httpServer := server.Init(*environment)

	fmt.Printf("Server started in %s mode and running on port %s\n", *environment, httpServer.Addr)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	fmt.Println("\nShutting down server...")
	server.Shutdown(httpServer)
}
