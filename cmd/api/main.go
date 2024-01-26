package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/domvcelos/rinha-de-backend-2023-q3/internal/people"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/redis/go-redis/v9"
)

func main() {
	if os.Getenv("ENV") != "CONTAINER" {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatalf("Some error occured. Err: %s", err)
		}
	}
	dataSourceName := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))
	db, err := sql.Open(os.Getenv("DB_DRIVER"), dataSourceName)
	if err != nil {
		log.Fatalf("Unable to connect with database: %v", err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		log.Fatalf("Unable ping database: %v", err)
	}
	rdb := redis.NewClient(&redis.Options{
		Addr: os.Getenv("CACHE_HOST") + ":" + os.Getenv("CACHE_PORT"),
		DB:   0,
	})
	defer rdb.Close()
	ctx := context.Background()
	status, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatalln("Redis connection was refused")
	}
	fmt.Println("Redis Status: " + status)
	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Heartbeat("/health-check"))
	router.Use(middleware.Timeout(60 * time.Second))
	createPeopleChan := make(chan *people.People)
	peopleRepository := people.NewPostgres(db, rdb)
	QtdWorkers := 100
	for i := 0; i < QtdWorkers; i++ {
		go peopleRepository.Create(ctx, createPeopleChan)
	}
	peopleService := people.NewService(peopleRepository, rdb, createPeopleChan)
	peopleHandler := people.NewHandler(peopleService)
	router.Get("/contagem-pessoas", peopleHandler.Count)
	router.Route("/pessoas", func(r chi.Router) {
		r.Post("/", peopleHandler.Create)
		r.Get("/{peopleID}", peopleHandler.FindById)
		r.Get("/", peopleHandler.Find)
	})
	fmt.Println("Starting server at port: " + os.Getenv("SERVER_PORT"))
	err = http.ListenAndServe(":"+os.Getenv("SERVER_PORT"), router)
	if err != nil {
		log.Fatalf("Unable start a server: %v", err)
	}
}
