package util

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/santekno/learn-golang-restful/internal/models"
)

func TestResponse(t *testing.T) {
	recorder := httptest.NewRecorder()
	type args struct {
		w          http.ResponseWriter
		response   interface{}
		statusCode int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "success http ok",
			args: args{
				w:          recorder,
				response:   models.HeaderResponse{},
				statusCode: http.StatusOK,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Response(tt.args.w, tt.args.response, tt.args.statusCode)
		})
	}
}
