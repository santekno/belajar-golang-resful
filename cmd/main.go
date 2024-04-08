package main

import (
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
	httpHandler "github.com/santekno/belajar-golang-restful/internal/delivery/http"
	"github.com/santekno/belajar-golang-restful/internal/middleware"
	mysqlRepository "github.com/santekno/belajar-golang-restful/internal/repository/mysql"
	articleUsecase "github.com/santekno/belajar-golang-restful/internal/usecase/article"
	"github.com/santekno/belajar-golang-restful/pkg/database"
	middleware_chain "github.com/santekno/belajar-golang-restful/pkg/middleware-chain"
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

	// inisialisasi database
	db := database.New()

	// inisialisasi repository
	repository := mysqlRepository.New(db)
	// inisialisasi usecase
	articleUsecase := articleUsecase.New(repository)
	// inisialisasi handler
	articleHandler := httpHandler.New(articleUsecase)
	// // inisialisasi new router
	router := NewRouter(articleHandler)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	err = server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func NewRouter(articleHandler *httpHandler.Delivery) *httprouter.Router {
	// inisialisasi http router
	router := httprouter.New()

	// inisialisasi chain middleware
	m := middleware_chain.New(
		middleware.AuthenticationBasic,
	)

	// entrypoint
	router.GET("/api/articles", m.Then(articleHandler.GetAll))
	router.GET("/api/articles/:article_id", m.Then(articleHandler.GetByID))
	router.POST("/api/articles/", m.Then(articleHandler.Store))
	router.PUT("/api/articles/:article_id", m.Then(articleHandler.Update))
	router.DELETE("/api/articles/:article_id", m.Then(articleHandler.Delete))

	return router
}
