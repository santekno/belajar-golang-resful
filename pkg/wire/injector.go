//go:build wireinject
// +build wireinject

package wire

import (
	"net/http"

	"github.com/google/wire"
	httpHandler "github.com/santekno/belajar-golang-restful/internal/delivery/http"
	"github.com/santekno/belajar-golang-restful/internal/repository"
	mysqlRepository "github.com/santekno/belajar-golang-restful/internal/repository/mysql"
	"github.com/santekno/belajar-golang-restful/internal/usecase"
	articleUsecase "github.com/santekno/belajar-golang-restful/internal/usecase/article"
	"github.com/santekno/belajar-golang-restful/pkg/database"
	"github.com/santekno/belajar-golang-restful/pkg/environment"
	"github.com/santekno/belajar-golang-restful/pkg/router"
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
