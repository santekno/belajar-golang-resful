package router

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	httpHandler "github.com/santekno/learn-golang-restful/internal/delivery/http"
	"github.com/santekno/learn-golang-restful/internal/middleware"
	middleware_chain "github.com/santekno/learn-golang-restful/pkg/middleware-chain"
)

func NewServer(router *httprouter.Router) *http.Server {
	return &http.Server{
		Addr:    "localhost:3000",
		Handler: router,
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
