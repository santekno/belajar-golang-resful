package mysql

import (
	"context"
	"database/sql"
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/santekno/belajar-golang-restful/internal/models"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	type args struct {
		conn *sql.DB
	}
	tests := []struct {
		name string
		args args
		want *ArticleStore
	}{
		{
			name: "success",
			args: args{},
			want: &ArticleStore{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.conn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArticleStore_GetAll(t *testing.T) {
	ctx := context.Background()
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf("error")
	}

	// expect query
	const expetcqueryGetAll = `SELECT id, title, content, create_at, update_at FROM articles`
	var column = []string{"id", "title", "content", "create_at", "update_at"}
	var timeNow = time.Now()
	var expectResult []*models.Article

	type fields struct {
		db *sql.DB
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*models.Article
		wantErr bool
		mock    func()
	}{
		{
			name: "success get article",
			fields: fields{
				db: mockDB,
			},
			mock: func() {
				mock.ExpectQuery(expetcqueryGetAll).WillReturnRows(
					sqlmock.NewRows(column).AddRow(
						1, "test", "test content", timeNow, timeNow,
					),
				)
			},
			want: []*models.Article{
				{
					ID:       1,
					Title:    "test",
					Content:  "test content",
					CreateAt: timeNow,
					UpdateAt: timeNow,
				},
			},
			args: args{
				ctx: ctx,
			},
			wantErr: false,
		},
		{
			name: "failed when query error",
			fields: fields{
				db: mockDB,
			},
			mock: func() {
				mock.ExpectQuery(expetcqueryGetAll).WillReturnError(sql.ErrNoRows)
			},
			want: expectResult,
			args: args{
				ctx: ctx,
			},
			wantErr: true,
		},
		{
			name: "failed rows next",
			fields: fields{
				db: mockDB,
			},
			mock: func() {
				mock.ExpectQuery(expetcqueryGetAll).WillReturnRows(
					sqlmock.NewRows(column).AddRow(
						1, "test", "test content", "test", "test",
					),
				)
			},
			want: expectResult,
			args: args{
				ctx: ctx,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &ArticleStore{
				db: tt.fields.db,
			}
			tt.mock()
			got, err := r.GetAll(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("ArticleStore.GetAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.EqualValues(t, got, tt.want)
		})
	}
}

func TestArticleStore_GetByID(t *testing.T) {
	ctx := context.Background()
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf("error")
	}

	// expect query
	const expectQueryGetById = `SELECT id, title, content, create_at, update_at FROM articles WHERE id=(.*)`
	var column = []string{"id", "title", "content", "create_at", "update_at"}
	var timeNow = time.Now()

	type fields struct {
		db *sql.DB
	}
	type args struct {
		ctx context.Context
		id  int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.Article
		wantErr bool
		mock    func()
	}{
		{
			name: "success get article",
			fields: fields{
				db: mockDB,
			},
			mock: func() {
				mock.ExpectQuery(expectQueryGetById).WillReturnRows(
					sqlmock.NewRows(column).AddRow(
						1, "test", "test content", timeNow, timeNow,
					),
				)
			},
			want: &models.Article{
				ID:       1,
				Title:    "test",
				Content:  "test content",
				CreateAt: timeNow,
				UpdateAt: timeNow,
			},
			args: args{
				ctx: ctx,
			},
			wantErr: false,
		},
		{
			name: "failed when query error",
			fields: fields{
				db: mockDB,
			},
			mock: func() {
				mock.ExpectQuery(expectQueryGetById).WillReturnError(sql.ErrNoRows)
			},
			want: &models.Article{},
			args: args{
				ctx: ctx,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &ArticleStore{
				db: tt.fields.db,
			}
			tt.mock()
			got, err := r.GetByID(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("ArticleStore.GetByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.EqualExportedValues(t, got, tt.want)
		})
	}
}

func TestArticleStore_Update(t *testing.T) {
	ctx := context.Background()
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf("error")
	}

	// expect query
	const expectQueryUpdate = `UPDATE articles SET title=(.*), content=(.*), update_at=(.*) WHERE id=(.*)`
	var timeNow = time.Now()

	type fields struct {
		db *sql.DB
	}
	type args struct {
		ctx     context.Context
		article *models.Article
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.Article
		wantErr bool
		mock    func()
	}{
		{
			name: "success update article",
			fields: fields{
				db: mockDB,
			},
			mock: func() {
				mock.ExpectExec(expectQueryUpdate).WillReturnResult(
					sqlmock.NewResult(1, 1),
				)
			},
			want: &models.Article{
				ID:       1,
				Title:    "test",
				Content:  "test content",
				CreateAt: timeNow,
				UpdateAt: timeNow,
			},
			args: args{
				ctx: ctx,
				article: &models.Article{
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
			name: "failed when query error",
			fields: fields{
				db: mockDB,
			},
			mock: func() {
				mock.ExpectExec(expectQueryUpdate).WillReturnError(errors.New("got error"))
			},
			want: nil,
			args: args{
				ctx: ctx,
				article: &models.Article{
					ID:       1,
					Title:    "test",
					Content:  "test content",
					CreateAt: timeNow,
					UpdateAt: timeNow,
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &ArticleStore{
				db: tt.fields.db,
			}
			tt.mock()
			got, err := r.Update(tt.args.ctx, tt.args.article)
			if (err != nil) != tt.wantErr {
				t.Errorf("ArticleStore.Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.EqualExportedValues(t, got, tt.want)
		})
	}
}

func TestArticleStore_Store(t *testing.T) {
	ctx := context.Background()
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf("error")
	}

	// expect query
	const expectQueryInsert = `INSERT INTO articles\(title, content, create_at, update_at\) VALUES\((.*),(.*),(.*),(.*)\)`
	var timeNow = time.Now()

	type fields struct {
		db *sql.DB
	}
	type args struct {
		ctx     context.Context
		article *models.Article
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
		mock    func()
	}{
		{
			name: "success create article",
			fields: fields{
				db: mockDB,
			},
			mock: func() {
				mock.ExpectExec(expectQueryInsert).WillReturnResult(
					sqlmock.NewResult(1, 1),
				)
			},
			want: 1,
			args: args{
				ctx: ctx,
				article: &models.Article{
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
			name: "failed when query error",
			fields: fields{
				db: mockDB,
			},
			mock: func() {
				mock.ExpectExec(expectQueryInsert).WillReturnError(errors.New("got error"))
			},
			want: 0,
			args: args{
				ctx: ctx,
				article: &models.Article{
					ID:       1,
					Title:    "test",
					Content:  "test content",
					CreateAt: timeNow,
					UpdateAt: timeNow,
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &ArticleStore{
				db: tt.fields.db,
			}
			tt.mock()
			got, err := r.Store(tt.args.ctx, tt.args.article)
			if (err != nil) != tt.wantErr {
				t.Errorf("ArticleStore.Store() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, got, tt.want)
		})
	}
}

func TestArticleStore_Delete(t *testing.T) {
	ctx := context.Background()
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf("error")
	}

	// expect query
	const expectQueryDelete = `DELETE FROM articles WHERE id=(.*)`
	type fields struct {
		db *sql.DB
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
		mock    func()
	}{
		{
			name: "success delete article",
			fields: fields{
				db: mockDB,
			},
			mock: func() {
				mock.ExpectExec(expectQueryDelete).WillReturnResult(
					sqlmock.NewResult(1, 1),
				)
			},
			want: true,
			args: args{
				ctx: ctx,
				id:  1,
			},
			wantErr: false,
		},
		{
			name: "failed when query error",
			fields: fields{
				db: mockDB,
			},
			mock: func() {
				mock.ExpectExec(expectQueryDelete).WillReturnError(errors.New("got error"))
			},
			want: false,
			args: args{
				ctx: ctx,
				id:  1,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &ArticleStore{
				db: tt.fields.db,
			}
			tt.mock()
			got, err := r.Delete(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("ArticleStore.Delete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, got, tt.want)
		})
	}
}
