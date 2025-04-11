package server

import (
	"context"
	"fmt"
	"net/http"
	"portfolio/api/config"
	router "portfolio/api/router"
	"time"
)

const ShutdownTimeout = 10 * time.Second

func Init(environment string) *http.Server {
	port := config.GetEnvInt("PORT")
	if port == 0 {
		port = 8080
	}

	rt := router.NewRouter(environment)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: rt,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("Error starting server: %v\n", err)
		}
	}()

	return server
}

func Shutdown(server *http.Server) {
	ctx, cancel := context.WithTimeout(context.Background(), ShutdownTimeout)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		fmt.Printf("Error shutting down server: %v\n", err)
	}

	fmt.Println("Server shut down gracefully")
}
