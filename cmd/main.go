package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/santekno/learn-golang-restful/pkg/wire"
)

func main() {
	// initialization server using wire generator
	server, err := wire.InitializedServer()
	if err != nil {
		panic(err)
	}

	err = server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
