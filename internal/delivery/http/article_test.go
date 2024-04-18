package http

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	"github.com/go-playground/validator"
	"github.com/julienschmidt/httprouter"
	"github.com/santekno/learn-golang-restful/internal/models"
	"github.com/santekno/learn-golang-restful/internal/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNew(t *testing.T) {
	mockUsecase := new(usecase.MockArticleUsecase)
	type args struct {
		articleUsecase usecase.ArticleUsecase
	}
	tests := []struct {
		name string
		args args
		want *Delivery
	}{
		{
			name: "success",
			args: args{
				articleUsecase: mockUsecase,
			},
			want: &Delivery{
				articleUsecase: mockUsecase,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			New(tt.args.articleUsecase)
		})
	}
}

func TestDelivery_GetAll(t *testing.T) {
	mockUsecase := new(usecase.MockArticleUsecase)
	timeNow := time.Now().UTC()
	type fields struct {
		articleUsecase usecase.ArticleUsecase
		validate       *validator.Validate
	}
	type args struct {
		params httprouter.Params
	}
	tests := []struct {
		name           string
		fields         fields
		args           args
		wantStatusCode int
		want           models.ArticleListResponse
		mock           func(args args)
	}{
		{
			name: "failed get all article",
			fields: fields{
				articleUsecase: mockUsecase,
			},
			mock: func(args args) {
				var res []models.ArticleResponse
				mockUsecase.On("GetAll", mock.Anything).Return(res, errors.New("got error")).Once()
			},
			wantStatusCode: http.StatusInternalServerError,
			want: models.ArticleListResponse{
				HeaderResponse: models.HeaderResponse{
					Code:   http.StatusInternalServerError,
					Status: "got error",
				},
			},
		},
		{
			name: "success get all article",
			fields: fields{
				articleUsecase: mockUsecase,
			},
			mock: func(args args) {
				var res = []models.ArticleResponse{
					{
						ID:       1,
						Title:    "test",
						Content:  "test content",
						CreateAt: timeNow,
						UpdateAt: timeNow,
					},
				}
				mockUsecase.On("GetAll", mock.Anything).Return(res, nil).Once()
			},
			wantStatusCode: http.StatusOK,
			want: models.ArticleListResponse{
				HeaderResponse: models.HeaderResponse{
					Code:   http.StatusOK,
					Status: "OK",
				},
				Data: []models.ArticleResponse{
					{
						ID:       1,
						Title:    "test",
						Content:  "test content",
						CreateAt: timeNow,
						UpdateAt: timeNow,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", "/test", nil)
			w := httptest.NewRecorder()

			tt.mock(tt.args)

			d := &Delivery{
				articleUsecase: tt.fields.articleUsecase,
				validate:       tt.fields.validate,
			}
			d.GetAll(w, req, tt.args.params)

			resp := w.Result()
			body, _ := io.ReadAll(resp.Body)

			var result models.ArticleListResponse
			err := json.Unmarshal(body, &result)
			if err != nil {
				t.Errorf("error unmarshal %v", err)
			}

			assert.Equal(t, tt.wantStatusCode, resp.StatusCode)
			assert.Equal(t, tt.want, result)
		})
	}
}

func TestDelivery_GetByID(t *testing.T) {
	mockUsecase := new(usecase.MockArticleUsecase)
	timeNow := time.Now().UTC()
	type fields struct {
		articleUsecase usecase.ArticleUsecase
		validate       *validator.Validate
	}
	type args struct {
		params httprouter.Params
	}
	tests := []struct {
		name           string
		fields         fields
		args           args
		wantStatusCode int
		wantErr        bool
		want           models.ArticleListResponse
		mock           func(args args)
	}{
		{
			name: "failed get by id article",
			fields: fields{
				articleUsecase: mockUsecase,
			},
			args: args{
				params: []httprouter.Param{{Key: "article_id", Value: "1"}},
			},
			mock: func(args args) {
				var res = models.ArticleResponse{}
				id, _ := strconv.ParseInt(args.params.ByName("article_id"), 0, 6)
				mockUsecase.On("GetByID", mock.Anything, id).Return(res, errors.New("got error")).Once()
			},
			wantStatusCode: http.StatusInternalServerError,
			want: models.ArticleListResponse{
				HeaderResponse: models.HeaderResponse{
					Code:   http.StatusInternalServerError,
					Status: "got error",
				},
			},
		},
		{
			name: "failed params not found",
			fields: fields{
				articleUsecase: mockUsecase,
			},
			args: args{
				params: []httprouter.Param{{Key: "article_id", Value: "a"}},
			},
			mock:           func(args args) {},
			wantStatusCode: http.StatusBadRequest,
			want: models.ArticleListResponse{
				HeaderResponse: models.HeaderResponse{
					Code:   http.StatusBadRequest,
					Status: "strconv.ParseInt: parsing \"a\": invalid syntax",
				},
			},
		},
		{
			name: "failed data not found",
			fields: fields{
				articleUsecase: mockUsecase,
			},
			args: args{
				params: []httprouter.Param{{Key: "article_id", Value: "1"}},
			},
			mock: func(args args) {
				var res = models.ArticleResponse{}
				id, _ := strconv.ParseInt(args.params.ByName("article_id"), 0, 64)
				mockUsecase.On("GetByID", mock.Anything, id).Return(res, nil).Once()
			},
			wantStatusCode: http.StatusNotFound,
			want: models.ArticleListResponse{
				HeaderResponse: models.HeaderResponse{
					Code:   http.StatusNotFound,
					Status: "data not found",
				},
			},
		},
		{
			name: "success get by id article",
			fields: fields{
				articleUsecase: mockUsecase,
			},
			args: args{
				params: []httprouter.Param{{Key: "article_id", Value: "1"}},
			},
			mock: func(args args) {
				var res = models.ArticleResponse{
					ID:       1,
					Title:    "test",
					Content:  "test content",
					CreateAt: timeNow,
					UpdateAt: timeNow,
				}
				id, _ := strconv.ParseInt(args.params.ByName("article_id"), 0, 64)
				mockUsecase.On("GetByID", mock.Anything, id).Return(res, nil).Once()
			},
			wantStatusCode: http.StatusOK,
			want: models.ArticleListResponse{
				HeaderResponse: models.HeaderResponse{
					Code:   http.StatusOK,
					Status: "OK",
				},
				Data: []models.ArticleResponse{
					{
						ID:       1,
						Title:    "test",
						Content:  "test content",
						CreateAt: timeNow,
						UpdateAt: timeNow,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", "/test", nil)
			w := httptest.NewRecorder()

			tt.mock(tt.args)

			d := &Delivery{
				articleUsecase: tt.fields.articleUsecase,
				validate:       tt.fields.validate,
			}
			d.GetByID(w, req, tt.args.params)
			resp := w.Result()
			body, _ := io.ReadAll(resp.Body)

			var result models.ArticleListResponse
			err := json.Unmarshal(body, &result)
			if err != nil {
				t.Errorf("error unmarshal %v", err)
			}

			assert.Equal(t, tt.wantStatusCode, resp.StatusCode)
			assert.Equal(t, tt.want, result)
		})
	}
}

func TestDelivery_Update(t *testing.T) {
	mockUsecase := new(usecase.MockArticleUsecase)
	timeNow := time.Now().UTC()
	type fields struct {
		articleUsecase usecase.ArticleUsecase
		validate       *validator.Validate
	}
	type args struct {
		params  httprouter.Params
		request models.ArticleUpdateRequest
	}
	tests := []struct {
		name           string
		fields         fields
		args           args
		mock           func(args args)
		wantStatusCode int
		wantErr        bool
		want           models.ArticleListResponse
	}{
		{
			name: "failed when article_id not support",
			fields: fields{
				articleUsecase: mockUsecase,
			},
			args: args{
				params: []httprouter.Param{{Key: "article_id", Value: "a"}},
				request: models.ArticleUpdateRequest{
					ID:       1,
					Title:    "test data beneran",
					Content:  "test test data beneran",
					UpdateAt: timeNow,
				},
			},
			mock: func(args args) {

			},
			wantErr:        true,
			wantStatusCode: http.StatusBadRequest,
			want: models.ArticleListResponse{
				HeaderResponse: models.HeaderResponse{
					Code:   http.StatusBadRequest,
					Status: "strconv.ParseInt: parsing \"a\": invalid syntax",
				},
			},
		},
		// {
		// 	name: "failed when encode body",
		// 	fields: fields{
		// 		articleUsecase: mockUsecase,
		// 		validate:       validator.New(),
		// 	},
		// 	args: args{
		// 		params: []httprouter.Param{{Key: "article_id", Value: "2"}},
		// 		request: models.ArticleUpdateRequest{
		// 			ID:       1,
		// 			Title:    "test",
		// 			Content:  "test",
		// 			UpdateAt: timeNow,
		// 		},
		// 	},
		// 	mock:           func(args args) {},
		// 	wantErr:        true,
		// 	wantStatusCode: http.StatusBadRequest,
		// 	want: models.ArticleListResponse{
		// 		HeaderResponse: models.HeaderResponse{
		// 			Code:   http.StatusBadRequest,
		// 			Status: "strconv.ParseInt: parsing \"a\": invalid syntax",
		// 		},
		// 	},
		// },
		{
			name: "failed when error validate",
			fields: fields{
				articleUsecase: mockUsecase,
				validate:       validator.New(),
			},
			args: args{
				params: []httprouter.Param{{Key: "article_id", Value: "2"}},
				request: models.ArticleUpdateRequest{
					ID:       1,
					Title:    "test",
					Content:  "test",
					UpdateAt: timeNow,
				},
			},
			mock:           func(args args) {},
			wantErr:        true,
			wantStatusCode: http.StatusBadRequest,
			want: models.ArticleListResponse{
				HeaderResponse: models.HeaderResponse{
					Code:   http.StatusBadRequest,
					Status: "Key: 'ArticleUpdateRequest.Title' Error:Field validation for 'Title' failed on the 'min' tag",
				},
			},
		},
		{
			name: "failed when update usecase was error",
			fields: fields{
				articleUsecase: mockUsecase,
				validate:       validator.New(),
			},
			args: args{
				params: []httprouter.Param{{Key: "article_id", Value: "2"}},
				request: models.ArticleUpdateRequest{
					ID:       2,
					Title:    "test data long",
					Content:  "test",
					UpdateAt: timeNow,
				},
			},
			mock: func(args args) {
				var res models.ArticleResponse
				mockUsecase.On("Update", mock.Anything, args.request).Return(res, errors.New("got error")).Once()
			},
			wantErr:        true,
			wantStatusCode: http.StatusInternalServerError,
			want: models.ArticleListResponse{
				HeaderResponse: models.HeaderResponse{
					Code:   http.StatusInternalServerError,
					Status: "got error",
				},
			},
		},
		{
			name: "success update article",
			fields: fields{
				articleUsecase: mockUsecase,
				validate:       validator.New(),
			},
			args: args{
				params: []httprouter.Param{{Key: "article_id", Value: "2"}},
				request: models.ArticleUpdateRequest{
					ID:       2,
					Title:    "test data long",
					Content:  "test",
					UpdateAt: timeNow,
				},
			},
			mock: func(args args) {
				var res models.ArticleResponse
				mockUsecase.On("Update", mock.Anything, args.request).Return(res, nil).Once()
			},
			wantErr:        false,
			wantStatusCode: http.StatusOK,
			want: models.ArticleListResponse{
				HeaderResponse: models.HeaderResponse{
					Code:   http.StatusOK,
					Status: "OK",
				},
				Data: []models.ArticleResponse{{}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var b = &bytes.Buffer{}
			if tt.args.request.Title != "" {
				err := json.NewEncoder(b).Encode(tt.args.request)
				if err != nil {
					t.Fatal(err)
				}
			}

			req := httptest.NewRequest("POST", "/test", b)
			w := httptest.NewRecorder()

			tt.mock(tt.args)

			d := &Delivery{
				articleUsecase: tt.fields.articleUsecase,
				validate:       tt.fields.validate,
			}

			d.Update(w, req, tt.args.params)
			resp := w.Result()
			body, _ := io.ReadAll(resp.Body)

			var result models.ArticleListResponse
			err := json.Unmarshal(body, &result)
			if err != nil {
				t.Errorf("error unmarshal %v", err)
			}

			assert.Equal(t, tt.wantStatusCode, resp.StatusCode)
			assert.Equal(t, tt.want, result)
		})
	}
}

func TestDelivery_Store(t *testing.T) {
	mockUsecase := new(usecase.MockArticleUsecase)
	timeNow := time.Now().UTC()
	type fields struct {
		articleUsecase usecase.ArticleUsecase
		validate       *validator.Validate
	}
	type args struct {
		params  httprouter.Params
		request models.ArticleCreateRequest
	}
	tests := []struct {
		name           string
		fields         fields
		args           args
		mock           func(args args)
		wantStatusCode int
		wantErr        bool
		want           models.ArticleListResponse
	}{
		// {
		// 	name: "failed when parse json body",
		// 		articleUsecase: mockUsecase,
		// 		validate:       validator.New(),
		// 	},
		// 	args: args{
		// 		params: []httprouter.Param{{Key: "article_id", Value: "2"}},
		// 		request: models.ArticleUpdateRequest{
		// 			ID:       1,
		// 			Title:    "test",
		// 			Content:  "test",
		// 			UpdateAt: timeNow,
		// 		},
		// 	},
		// 	mock:           func(args args) {},
		// 	wantErr:        true,
		// 	wantStatusCode: http.StatusBadRequest,
		// 	want: models.ArticleListResponse{
		// 		HeaderResponse: models.HeaderResponse{
		// 			Code:   http.StatusBadRequest,
		// 			Status: "strconv.ParseInt: parsing \"a\": invalid syntax",
		// 		},
		// 	},
		// },
		// },
		{
			name: "failed when error validate",
			fields: fields{
				articleUsecase: mockUsecase,
				validate:       validator.New(),
			},
			args: args{
				params: []httprouter.Param{{Key: "article_id", Value: "2"}},
				request: models.ArticleCreateRequest{
					Title:    "test",
					Content:  "test",
					CreateAt: timeNow,
				},
			},
			mock:           func(args args) {},
			wantErr:        true,
			wantStatusCode: http.StatusBadRequest,
			want: models.ArticleListResponse{
				HeaderResponse: models.HeaderResponse{
					Code:   http.StatusBadRequest,
					Status: "Key: 'ArticleCreateRequest.Title' Error:Field validation for 'Title' failed on the 'min' tag",
				},
			},
		},
		{
			name: "failed when update usecase was error",
			fields: fields{
				articleUsecase: mockUsecase,
				validate:       validator.New(),
			},
			args: args{
				params: []httprouter.Param{{Key: "article_id", Value: "2"}},
				request: models.ArticleCreateRequest{
					Title:    "test data long",
					Content:  "test",
					CreateAt: timeNow,
				},
			},
			mock: func(args args) {
				var res models.ArticleResponse
				mockUsecase.On("Store", mock.Anything, args.request).Return(res, errors.New("got error")).Once()
			},
			wantErr:        true,
			wantStatusCode: http.StatusInternalServerError,
			want: models.ArticleListResponse{
				HeaderResponse: models.HeaderResponse{
					Code:   http.StatusInternalServerError,
					Status: "got error",
				},
			},
		},
		{
			name: "success update article",
			fields: fields{
				articleUsecase: mockUsecase,
				validate:       validator.New(),
			},
			args: args{
				params: []httprouter.Param{{Key: "article_id", Value: "2"}},
				request: models.ArticleCreateRequest{
					Title:    "test data long",
					Content:  "test",
					CreateAt: timeNow,
				},
			},
			mock: func(args args) {
				var res models.ArticleResponse
				mockUsecase.On("Store", mock.Anything, args.request).Return(res, nil).Once()
			},
			wantErr:        false,
			wantStatusCode: http.StatusOK,
			want: models.ArticleListResponse{
				HeaderResponse: models.HeaderResponse{
					Code:   http.StatusOK,
					Status: "OK",
				},
				Data: []models.ArticleResponse{{}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var b = &bytes.Buffer{}
			if tt.args.request.Title != "" {
				err := json.NewEncoder(b).Encode(tt.args.request)
				if err != nil {
					t.Fatal(err)
				}
			}

			req := httptest.NewRequest("POST", "/test", b)
			w := httptest.NewRecorder()

			tt.mock(tt.args)

			d := &Delivery{
				articleUsecase: tt.fields.articleUsecase,
				validate:       tt.fields.validate,
			}

			d.Store(w, req, tt.args.params)
			resp := w.Result()
			body, _ := io.ReadAll(resp.Body)

			var result models.ArticleListResponse
			err := json.Unmarshal(body, &result)
			if err != nil {
				t.Errorf("error unmarshal %v", err)
			}
			assert.Equal(t, tt.wantStatusCode, resp.StatusCode)
			assert.Equal(t, tt.want, result)
		})
	}
}

func TestDelivery_Delete(t *testing.T) {
	mockUsecase := new(usecase.MockArticleUsecase)
	type fields struct {
		articleUsecase usecase.ArticleUsecase
		validate       *validator.Validate
	}
	type args struct {
		params  httprouter.Params
		request models.ArticleCreateRequest
	}
	tests := []struct {
		name           string
		fields         fields
		args           args
		mock           func(args args)
		wantStatusCode int
		wantErr        bool
		want           models.ArticleListResponse
	}{
		{
			name: "failed delete by id article",
			fields: fields{
				articleUsecase: mockUsecase,
			},
			args: args{
				params: []httprouter.Param{{Key: "article_id", Value: "1"}},
			},
			mock: func(args args) {
				id, _ := strconv.ParseInt(args.params.ByName("article_id"), 0, 6)
				mockUsecase.On("Delete", mock.Anything, id).Return(false, errors.New("got error")).Once()
			},
			wantStatusCode: http.StatusInternalServerError,
			want: models.ArticleListResponse{
				HeaderResponse: models.HeaderResponse{
					Code:   http.StatusInternalServerError,
					Status: "got error",
				},
			},
		},
		{
			name: "failed params not found",
			fields: fields{
				articleUsecase: mockUsecase,
			},
			args: args{
				params: []httprouter.Param{{Key: "article_id", Value: "a"}},
			},
			mock:           func(args args) {},
			wantStatusCode: http.StatusBadRequest,
			want: models.ArticleListResponse{
				HeaderResponse: models.HeaderResponse{
					Code:   http.StatusBadRequest,
					Status: "strconv.ParseInt: parsing \"a\": invalid syntax",
				},
			},
		},
		{
			name: "failed data not found",
			fields: fields{
				articleUsecase: mockUsecase,
			},
			args: args{
				params: []httprouter.Param{{Key: "article_id", Value: "0"}},
			},
			mock: func(args args) {
				id, _ := strconv.ParseInt(args.params.ByName("article_id"), 0, 64)
				mockUsecase.On("Delete", mock.Anything, id).Return(false, nil).Once()
			},
			wantStatusCode: http.StatusNotFound,
			want: models.ArticleListResponse{
				HeaderResponse: models.HeaderResponse{
					Code:   http.StatusNotFound,
					Status: "article_id was not zero",
				},
			},
		},
		{
			name: "success delete by id article",
			fields: fields{
				articleUsecase: mockUsecase,
			},
			args: args{
				params: []httprouter.Param{{Key: "article_id", Value: "1"}},
			},
			mock: func(args args) {
				id, _ := strconv.ParseInt(args.params.ByName("article_id"), 0, 64)
				mockUsecase.On("Delete", mock.Anything, id).Return(true, nil).Once()
			},
			wantStatusCode: http.StatusOK,
			want: models.ArticleListResponse{
				HeaderResponse: models.HeaderResponse{
					Code:   http.StatusOK,
					Status: "OK",
				},
			},
		},
		{
			name: "delete by article_id but data not found",
			fields: fields{
				articleUsecase: mockUsecase,
			},
			args: args{
				params: []httprouter.Param{{Key: "article_id", Value: "1"}},
			},
			mock: func(args args) {
				id, _ := strconv.ParseInt(args.params.ByName("article_id"), 0, 64)
				mockUsecase.On("Delete", mock.Anything, id).Return(false, nil).Once()
			},
			wantStatusCode: http.StatusInternalServerError,
			wantErr:        true,
			want: models.ArticleListResponse{
				HeaderResponse: models.HeaderResponse{
					Code:   http.StatusInternalServerError,
					Status: "unknown error",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var b = &bytes.Buffer{}
			if tt.args.request.Title != "" {
				err := json.NewEncoder(b).Encode(tt.args.request)
				if err != nil {
					t.Fatal(err)
				}
			}

			req := httptest.NewRequest("GET", "/test", b)
			w := httptest.NewRecorder()

			tt.mock(tt.args)

			d := &Delivery{
				articleUsecase: tt.fields.articleUsecase,
				validate:       tt.fields.validate,
			}
			d.Delete(w, req, tt.args.params)
			resp := w.Result()
			body, _ := io.ReadAll(resp.Body)

			var result models.ArticleListResponse
			err := json.Unmarshal(body, &result)
			if err != nil {
				t.Errorf("error unmarshal %v", err)
			}
			assert.Equal(t, tt.wantStatusCode, resp.StatusCode)
			assert.Equal(t, tt.want, result)
		})
	}
}
