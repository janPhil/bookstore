package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/janPhil/bookstore/internal/handlers"
	"github.com/janPhil/bookstore/internal/storage"

	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
)

var bindAddress = flag.String("BIND_ADDRESS", ":9090", "Bind address for the server")

func main() {
	flag.Parse()
	l := log.New(os.Stdout, "bookstore ", log.LstdFlags)

	storage, err := storage.NewBookStorage()
	if err != nil {
		log.Fatal("Failed to create storage")
	}

	booksHandler := handlers.NewBooksHandler(storage, l)
	sm := mux.NewRouter()

	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/books", booksHandler.ListAll)
	getRouter.HandleFunc("/books/{id}", booksHandler.ListOne)

	postRouter := sm.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/book", booksHandler.Create)

	deleteRouter := sm.Methods(http.MethodDelete).Subrouter()
	deleteRouter.HandleFunc("/books/{id}", booksHandler.Delete)

	// handler for documentation
	opts := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
	sh := middleware.Redoc(opts, nil)

	getRouter.Handle("/docs", sh)
	getRouter.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))

	s := http.Server{
		Addr:         *bindAddress,
		Handler:      sm,
		ErrorLog:     l,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	go func() {
		l.Println("Starting server on port 9090")

		err := s.ListenAndServe()
		if err != nil {
			l.Printf("Error starting server: %s\n", err)
			os.Exit(1)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	sig := <-c
	log.Println("Got signal:", sig)

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(ctx)

}
