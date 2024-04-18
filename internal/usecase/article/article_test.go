package article

import (
	"context"
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/santekno/learn-golang-restful/internal/models"
	repository "github.com/santekno/learn-golang-restful/internal/repository"
	"github.com/stretchr/testify/mock"
)

func TestNew(t *testing.T) {
	type args struct {
		repo repository.ArticleRepository
	}
	tests := []struct {
		name string
		args args
		want *Usecase
	}{
		{
			name: "success new",
			want: &Usecase{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.repo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUsecase_GetAll(t *testing.T) {
	var ctx = context.Background()
	var timeNow = time.Now()
	mockRepository := new(repository.MockArticleRepository)
	type fields struct {
		articleRepository repository.ArticleRepository
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []models.ArticleResponse
		wantErr bool
		mock    func(args args)
	}{
		{
			name: "success",
			args: args{
				ctx: ctx,
			},
			fields: fields{
				articleRepository: mockRepository,
			},
			mock: func(args args) {
				var res = []*models.Article{
					{
						ID:       1,
						Title:    "test",
						Content:  "test content",
						CreateAt: timeNow,
						UpdateAt: timeNow,
					},
				}
				mockRepository.On("GetAll", args.ctx).Return(res, nil).Once()
			},
			want: []models.ArticleResponse{
				{
					ID:       1,
					Title:    "test",
					Content:  "test content",
					CreateAt: timeNow,
					UpdateAt: timeNow,
				},
			},
			wantErr: false,
		},
		{
			name: "failed get all data",
			args: args{
				ctx: ctx,
			},
			fields: fields{
				articleRepository: mockRepository,
			},
			mock: func(args args) {
				var res = []*models.Article{}
				mockRepository.On("GetAll", args.ctx).Return(res, errors.New("got error")).Once()
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock(tt.args)
			u := &Usecase{
				articleRepository: tt.fields.articleRepository,
			}
			got, err := u.GetAll(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("Usecase.GetAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Usecase.GetAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUsecase_GetByID(t *testing.T) {
	var ctx = context.Background()
	var timeNow = time.Now()
	mockRepository := new(repository.MockArticleRepository)
	type fields struct {
		articleRepository repository.ArticleRepository
	}
	type args struct {
		ctx context.Context
		id  int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    models.ArticleResponse
		wantErr bool
		mock    func(args args)
	}{
		{
			name: "success",
			args: args{
				ctx: ctx,
				id:  1,
			},
			fields: fields{
				articleRepository: mockRepository,
			},
			mock: func(args args) {
				var res = &models.Article{
					ID:       1,
					Title:    "test",
					Content:  "test content",
					CreateAt: timeNow,
					UpdateAt: timeNow,
				}
				mockRepository.On("GetByID", args.ctx, args.id).Return(res, nil).Once()
			},
			want: models.ArticleResponse{
				ID:       1,
				Title:    "test",
				Content:  "test content",
				CreateAt: timeNow,
				UpdateAt: timeNow,
			},
			wantErr: false,
		},
		{
			name: "failed get by id",
			args: args{
				ctx: ctx,
			},
			fields: fields{
				articleRepository: mockRepository,
			},
			mock: func(args args) {
				var res = &models.Article{}
				mockRepository.On("GetByID", args.ctx, args.id).Return(res, errors.New("got error")).Once()
			},
			want:    models.ArticleResponse{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock(tt.args)
			u := &Usecase{
				articleRepository: tt.fields.articleRepository,
			}
			got, err := u.GetByID(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Usecase.GetByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Usecase.GetByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUsecase_Update(t *testing.T) {
	var ctx = context.Background()
	var timeNow = time.Now()
	mockRepository := new(repository.MockArticleRepository)
	type fields struct {
		articleRepository repository.ArticleRepository
	}
	type args struct {
		ctx     context.Context
		request models.ArticleUpdateRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    models.ArticleResponse
		wantErr bool
		mock    func(args args)
	}{
		{
			name: "failed when request id was zero",
			fields: fields{
				articleRepository: mockRepository,
			},
			args: args{
				ctx:     ctx,
				request: models.ArticleUpdateRequest{},
			},
			mock:    func(args args) {},
			want:    models.ArticleResponse{},
			wantErr: true,
		},
		{
			name: "failed when get by repo",
			fields: fields{
				articleRepository: mockRepository,
			},
			args: args{
				ctx: ctx,
				request: models.ArticleUpdateRequest{
					ID: 1,
				},
			},
			mock: func(args args) {
				var article *models.Article
				mockRepository.On("GetByID", args.ctx, args.request.ID).Return(article, errors.New("got error")).Once()
			},
			want:    models.ArticleResponse{},
			wantErr: true,
		},
		{
			name: "failed when get by id not found",
			fields: fields{
				articleRepository: mockRepository,
			},
			args: args{
				ctx: ctx,
				request: models.ArticleUpdateRequest{
					ID: 1,
				},
			},
			mock: func(args args) {
				var article = &models.Article{}
				mockRepository.On("GetByID", args.ctx, args.request.ID).Return(article, nil).Once()
			},
			want:    models.ArticleResponse{},
			wantErr: true,
		},
		{
			name: "failed when update repo",
			fields: fields{
				articleRepository: mockRepository,
			},
			args: args{
				ctx: ctx,
				request: models.ArticleUpdateRequest{
					ID:       1,
					Title:    "test update",
					Content:  "test update",
					UpdateAt: timeNow,
				},
			},
			mock: func(args args) {
				var article = &models.Article{
					ID:       1,
					Title:    "test",
					Content:  "test",
					CreateAt: timeNow,
					UpdateAt: timeNow,
				}
				mockRepository.On("GetByID", args.ctx, args.request.ID).Return(article, nil).Once()

				article.FromUpdateRequest(args.request)
				mockRepository.On("Update", args.ctx, article).Return(&models.Article{}, errors.New("got error"))
			},
			want:    models.ArticleResponse{},
			wantErr: true,
		},
		{
			name: "success update",
			fields: fields{
				articleRepository: mockRepository,
			},
			args: args{
				ctx: ctx,
				request: models.ArticleUpdateRequest{
					ID:       1,
					Title:    "test update",
					Content:  "test update",
					UpdateAt: timeNow,
				},
			},
			mock: func(args args) {
				var article = &models.Article{
					ID:       1,
					Title:    "test",
					Content:  "test",
					CreateAt: timeNow,
					UpdateAt: timeNow,
				}
				mockRepository.On("GetByID", args.ctx, args.request.ID).Return(article, nil).Once()

				article.FromUpdateRequest(args.request)
				mockRepository.On("Update", args.ctx, article).Return(article, nil).Once()
			},
			want: models.ArticleResponse{
				ID:       1,
				Title:    "test update",
				Content:  "test update",
				UpdateAt: timeNow,
				CreateAt: timeNow,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock(tt.args)
			u := &Usecase{
				articleRepository: tt.fields.articleRepository,
			}
			got, err := u.Update(tt.args.ctx, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("Usecase.Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.ID, tt.want.ID) {
				t.Errorf("Usecase.Update() = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got.Title, tt.want.Title) {
				t.Errorf("Usecase.Update() = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got.Content, tt.want.Content) {
				t.Errorf("Usecase.Update() = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got.CreateAt, tt.want.CreateAt) {
				t.Errorf("Usecase.Update() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUsecase_Store(t *testing.T) {
	var ctx = context.Background()
	var timeNow = time.Now()
	mockRepository := new(repository.MockArticleRepository)
	type fields struct {
		articleRepository repository.ArticleRepository
	}
	type args struct {
		ctx     context.Context
		request models.ArticleCreateRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    models.ArticleResponse
		wantErr bool
		mock    func(args args)
	}{
		{
			name: "failed store repository",
			fields: fields{
				articleRepository: mockRepository,
			},
			args: args{
				ctx: ctx,
				request: models.ArticleCreateRequest{
					Title:    "test",
					Content:  "test content",
					CreateAt: timeNow,
				},
			},
			mock: func(args args) {
				var result int64
				mockRepository.On("Store", args.ctx, mock.Anything).Return(result, errors.New("got error")).Once()
			},
			want:    models.ArticleResponse{},
			wantErr: true,
		},
		{
			name: "success store repository",
			fields: fields{
				articleRepository: mockRepository,
			},
			args: args{
				ctx: ctx,
				request: models.ArticleCreateRequest{
					Title:    "test",
					Content:  "test content",
					CreateAt: timeNow,
				},
			},
			mock: func(args args) {
				var result int64 = 1
				mockRepository.On("Store", args.ctx, mock.Anything).Return(result, nil).Once()
			},
			want: models.ArticleResponse{
				ID:       1,
				Title:    "test",
				Content:  "test content",
				CreateAt: timeNow,
				UpdateAt: timeNow,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock(tt.args)
			u := &Usecase{
				articleRepository: tt.fields.articleRepository,
			}
			got, err := u.Store(tt.args.ctx, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("Usecase.Store() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.ID, tt.want.ID) {
				t.Errorf("Usecase.Store() = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got.Title, tt.want.Title) {
				t.Errorf("Usecase.Store() = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got.Content, tt.want.Content) {
				t.Errorf("Usecase.Store() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUsecase_Delete(t *testing.T) {
	var ctx = context.Background()
	var timeNow = time.Now()
	mockRepository := new(repository.MockArticleRepository)
	type fields struct {
		articleRepository repository.ArticleRepository
	}
	type args struct {
		ctx context.Context
		id  int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
		mock    func(args args)
	}{
		{
			name: "failed when get by id was error",
			fields: fields{
				articleRepository: mockRepository,
			},
			args: args{
				ctx: ctx,
				id:  1,
			},
			mock: func(args args) {
				var article *models.Article
				mockRepository.On("GetByID", args.ctx, args.id).Return(article, errors.New("got error")).Once()
			},
			want:    false,
			wantErr: true,
		},
		{
			name: "failed when article not found",
			fields: fields{
				articleRepository: mockRepository,
			},
			args: args{
				ctx: ctx,
				id:  1,
			},
			mock: func(args args) {
				mockRepository.On("GetByID", args.ctx, args.id).Return(nil, nil).Once()

			},
			want:    false,
			wantErr: true,
		},
		{
			name: "failed when delete repo was error",
			fields: fields{
				articleRepository: mockRepository,
			},
			args: args{
				ctx: ctx,
				id:  1,
			},
			mock: func(args args) {
				var article = &models.Article{
					ID:       1,
					Title:    "test",
					Content:  "test content",
					CreateAt: timeNow,
					UpdateAt: timeNow,
				}
				mockRepository.On("GetByID", args.ctx, args.id).Return(article, nil).Once()

				mockRepository.On("Delete", args.ctx, args.id).Return(false, errors.New("got error")).Once()
			},
			want:    false,
			wantErr: true,
		},
		{
			name: "success delete article",
			fields: fields{
				articleRepository: mockRepository,
			},
			args: args{
				ctx: ctx,
				id:  1,
			},
			mock: func(args args) {
				var article = &models.Article{
					ID:       1,
					Title:    "test",
					Content:  "test content",
					CreateAt: timeNow,
					UpdateAt: timeNow,
				}
				mockRepository.On("GetByID", args.ctx, args.id).Return(article, nil).Once()
				mockRepository.On("Delete", args.ctx, args.id).Return(true, nil).Once()
			},
			want:    true,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock(tt.args)
			u := &Usecase{
				articleRepository: tt.fields.articleRepository,
			}
			got, err := u.Delete(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Usecase.Delete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Usecase.Delete() = %v, want %v", got, tt.want)
			}
		})
	}
}
