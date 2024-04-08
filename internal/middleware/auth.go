package middleware

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/santekno/belajar-golang-restful/internal/models"
	"github.com/santekno/belajar-golang-restful/pkg/util"
)

const (
	Secret  = "s3cr3t"
	XApiKey = "X-API-Key"
)

func AuthenticationBasic(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		if r.Header.Get(XApiKey) != Secret {
			var statusCode = http.StatusUnauthorized
			var response models.HeaderResponse
			response.Code = statusCode
			response.Status = "Unauthorized"
			util.Response(w, response, statusCode)
			return
		}

		next(w, r, params)
	}
}
