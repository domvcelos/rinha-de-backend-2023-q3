package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/domvcelos/rinha-de-backend-2023-q3/configs"
	"github.com/domvcelos/rinha-de-backend-2023-q3/internal/people"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	_ "github.com/lib/pq"
)

func main() {
	config, err := configs.LoadConfig(".")
	if err != nil {
		log.Fatalf("Unable to read configuration file: %v", err)
	}
	dataSourceName := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.DBHost, config.DBPort, config.DBUser, config.DBPassword, config.DBName)
	db, err := sql.Open(config.DBDriver, dataSourceName)
	if err != nil {
		log.Fatalf("Unable to connect with database: %v", err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		log.Fatalf("Unable ping database: %v", err)
	}
	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Heartbeat("/health-check"))
	router.Use(middleware.Timeout(60 * time.Second))

	peopleRepository := people.NewPostgres(db)
	peopleService := people.NewService(peopleRepository)
	peopleHandler := people.NewHandler(peopleService)

	router.Route("/pessoas", func(r chi.Router) {
		r.Post("/", peopleHandler.Create)
		r.Get("/{peopleID}", peopleHandler.FindById)
		r.Get("/count", peopleHandler.Count)
		r.Get("/", peopleHandler.Find)
	})
	fmt.Println("Starting server at port: 8080...")
	err = http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatalf("Unable start a server: %v", err)
	}
}
