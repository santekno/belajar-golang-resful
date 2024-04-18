//go:build wireinject
// +build wireinject

package wire

import (
	"net/http"

	"github.com/google/wire"
	httpHandler "github.com/santekno/learn-golang-restful/internal/delivery/http"
	mysqlRepository "github.com/santekno/learn-golang-restful/internal/repository/mysql"
	articleUsecase "github.com/santekno/learn-golang-restful/internal/usecase/article"
	"github.com/santekno/learn-golang-restful/pkg/router"
	"github.com/santekno/learn-golang-restfulkg/database"
	"github.com/santekno/learn-golang-restfulkg/environment"
	"github.com/santekno/learn-golang-restfulnternal/repository"
	"github.com/santekno/learn-golang-restfulnternal/usecase"
)

var articleSet = wire.NewSet(
	mysqlRepository.New,
	wire.Bind(new(repository.ArticleRepository), new(*mysqlRepository.ArticleStore)),
	articleUsecase.New,
	wire.Bind(new(usecase.ArticleUsecase), new(*articleUsecase.Usecase)),
	httpHandler.New,
)

func InitializedServer() (*http.Server, error) {
	wire.Build(
		environment.Load,
		database.New,
		articleSet,
		router.NewRouter,
		router.NewServer,
	)
	return nil, nil
}
