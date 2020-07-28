package main

import (
	"log"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"time"
	"os"
	"os/signal"
	"context"
	"syscall"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("Hello World!!")))
}

func main() {
    greeting := "Hello World!!\n"
    log.Printf(greeting)

    r := mux.NewRouter()
    r.HandleFunc("/", handler)

    srv := &http.Server{
        Handler:      r,
        Addr:         ":8080",
        ReadTimeout:  12 * time.Second,
        WriteTimeout: 12 * time.Second,
    }

    // Start Server
    go func() {
        log.Println("Starting Server")
        if err := srv.ListenAndServe(); err != nil {
            log.Fatal(err)
        }
    }()

    // Graceful Shutdown
    waitForShutdown(srv)
}

func waitForShutdown(srv *http.Server) {
	interruptChan := make(chan os.Signal, 1)
	signal.Notify(interruptChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// Block until we receive our signal.
	<-interruptChan

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	srv.Shutdown(ctx)

	log.Println("Shutting down")
	os.Exit(0)
}