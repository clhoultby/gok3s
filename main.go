package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	fmt.Println("starting server")
	go (func() {
		if r := recover(); r != nil {
			fmt.Println("Recover", r)
		}
	})()
	
	mux := http.NewServeMux();

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("request received")
		io.WriteString(w, "Hello World")
	})

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("health check request received")
		w.WriteHeader(http.StatusOK)
	})

	mux.HandleFunc("/readiness", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("health check request received")
		w.WriteHeader(http.StatusOK)
	})

	// create the server instance passing our ready to go mux
	server := &http.Server{
		Addr: ":8080",
		Handler: mux,
	}

	go func ()  {
		server.ListenAndServe()
	}()

	interruptChan := make(chan os.Signal, 1)
	signal.Notify(interruptChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// Block until we receive our signal.
	<-interruptChan

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	server.Shutdown(ctx)

	fmt.Println("Shutting down")
	os.Exit(0)
}
