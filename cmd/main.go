package main

import (
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
	httpHandler "github.com/santekno/belajar-golang-restful/delivery/http"
	"github.com/santekno/belajar-golang-restful/middleware"
	"github.com/santekno/belajar-golang-restful/pkg/database"
	mysqlRepository "github.com/santekno/belajar-golang-restful/repository/mysql"
	articleUsecase "github.com/santekno/belajar-golang-restful/usecase/article"
)

func main() {
	fileEnv := ".env"
	if os.Getenv("environment") == "development" {
		fileEnv = "../.env"
	}

	err := godotenv.Load(fileEnv)
	if err != nil {
		log.Fatalf("error loading .env file")
	}

	// inisialisasi http router
	router := httprouter.New()

	// inisialisasi database
	db := database.New()

	// inisialisasi repository
	repository := mysqlRepository.New(db)
	// inisialisasi usecase
	articleUsecase := articleUsecase.New(repository)
	// inisialisasi handler
	articleHandler := httpHandler.New(articleUsecase)

	// inisialisasi chain middleware
	m := middleware.New(
		middleware.AuthenticationBasic,
	)

	// entrypoint
	router.GET("/api/articles", m.Then(articleHandler.GetAll))
	router.GET("/api/articles/:article_id", m.Then(articleHandler.GetByID))
	router.POST("/api/articles/", m.Then(articleHandler.Store))
	router.PUT("/api/articles/:article_id", m.Then(articleHandler.Update))
	router.DELETE("/api/articles/:article_id", m.Then(articleHandler.Delete))

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	err = server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
