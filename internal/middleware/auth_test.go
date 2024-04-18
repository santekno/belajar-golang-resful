package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	middleware_chain "github.com/santekno/learn-golang-restful/pkg/middleware-chain"
	"github.com/stretchr/testify/assert"
)

func TestAuthenticationBasic(t *testing.T) {
	router := httprouter.New()
	mv := middleware_chain.New(
		AuthenticationBasic,
	)
	router.GET("/onlytest", mv.Then(func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		w.WriteHeader(http.StatusOK)
	}))

	type args struct {
		next httprouter.Handle
	}
	tests := []struct {
		name     string
		args     args
		want     httprouter.Handle
		wantCode int
		mock     func(*http.Request)
	}{
		{
			name: "success",
			args: args{
				next: func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {},
			},
			want:     func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {},
			wantCode: http.StatusOK,
			mock: func(r *http.Request) {
				r.Header.Set(XApiKey, Secret)
			},
		},
		{
			name: "unauthorized",
			args: args{
				next: func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {},
			},
			want:     func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {},
			wantCode: http.StatusUnauthorized,
			mock:     func(r *http.Request) {},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			recorder := httptest.NewRecorder()
			requestTest, _ := http.NewRequest("GET", "/onlytest", nil)
			tt.mock(requestTest)
			router.ServeHTTP(recorder, requestTest)

			got := AuthenticationBasic(tt.args.next)
			assert.ObjectsAreEqual(tt.want, got)
			assert.Equal(t, tt.wantCode, recorder.Result().StatusCode)
		})
	}
}
