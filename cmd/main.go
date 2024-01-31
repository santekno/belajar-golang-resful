package main

import (
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
	httpHandler "github.com/santekno/belajar-golang-restful/delivery/http"
	"github.com/santekno/belajar-golang-restful/pkg/database"
	mysqlRepository "github.com/santekno/belajar-golang-restful/repository/mysql"
	articleUsecase "github.com/santekno/belajar-golang-restful/usecase/article"
)

func main() {
	err := godotenv.Load(".env")
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

	// entrypoint
	router.GET("/api/articles", articleHandler.GetAll)
	router.GET("/api/articles/:article_id", articleHandler.GetByID)
	router.POST("/api/articles/", articleHandler.Store)
	router.PUT("/api/articles/:article_id", articleHandler.Update)
	router.DELETE("/api/articles/:article_id", articleHandler.Delete)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	err = server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
