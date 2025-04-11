package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"portfolio/api/config"
	"portfolio/api/server"
	"syscall"
)

func main() {
	environment := flag.String("env", "development", "Environment to run the application in (development, staging, production)")

	flag.Usage = func() {
		log.Fatalf(
			"Usage: %s [options]\nOptions:\n  -env string\n\tEnvironment to run the application in (development, staging, production)",
			flag.CommandLine.Name(),
		)

		os.Exit(1)
	}

	config.Init(*environment)
	httpServer := server.Init(*environment)

	fmt.Printf("Server started in %s mode and running on port %s\n", *environment, httpServer.Addr)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	fmt.Println("\nShutting down server...")
	server.Shutdown(httpServer)
}
