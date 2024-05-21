package main

import (
	"context"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"port-service/internal/config"
	"port-service/internal/core/repository"
	"port-service/internal/core/service"
	"port-service/internal/data-source/storage/inmem"
	"port-service/internal/transport/handler"
	"syscall"
	"time"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
	os.Exit(0)
}

func run() error {
	//read config from env
	cfg := config.Read()

	// create storage
	storage := inmem.New()

	// create port repository
	portStoreRepo := repository.NewPortRepository(storage)

	// create port service
	portService := service.NewPortService(portStoreRepo)

	// create handler server with application injected
	httpServer := handler.NewHttpServer(portService)

	// create handler router
	router := mux.NewRouter()
	router.HandleFunc("/ports", httpServer.FetchPorts).Methods("POST")

	srv := &http.Server{
		Addr:    cfg.HttpAddr,
		Handler: router,
	}

	// listen to OS signals and gracefully shutdown HTTP server
	stopped := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
		<-sigint
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := srv.Shutdown(ctx); err != nil {
			log.Printf("HTTP Server Shutdown Error: %v", err)
		}
		close(stopped)
	}()

	log.Printf("Starting HTTP server on %s", cfg.HttpAddr)

	// start HTTP server
	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("HTTP server ListenAndServe Error: %v", err)
	}

	<-stopped

	log.Print(storage.GetMap())

	log.Printf("Server stopped")

	return nil
}
