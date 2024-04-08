package models

import (
	"reflect"
	"testing"
	"time"
)

func TestArticle_FromCreateRequest(t *testing.T) {
	timeNow := time.Now()
	type fields struct {
		ID       int64
		Title    string
		Content  string
		CreateAt time.Time
		UpdateAt time.Time
	}
	type args struct {
		request ArticleCreateRequest
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "success",
			fields: fields{
				ID:       1,
				Title:    "test",
				Content:  "content",
				CreateAt: timeNow,
				UpdateAt: timeNow,
			},
			args: args{
				request: ArticleCreateRequest{
					Title:    "test",
					Content:  "content",
					CreateAt: timeNow,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Article{
				ID:       tt.fields.ID,
				Title:    tt.fields.Title,
				Content:  tt.fields.Content,
				CreateAt: tt.fields.CreateAt,
				UpdateAt: tt.fields.UpdateAt,
			}
			m.FromCreateRequest(tt.args.request)
		})
	}
}

func TestArticle_FromUpdateRequest(t *testing.T) {
	timeNow := time.Now()
	type fields struct {
		ID       int64
		Title    string
		Content  string
		CreateAt time.Time
		UpdateAt time.Time
	}
	type args struct {
		request ArticleUpdateRequest
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "success",
			fields: fields{
				ID:       1,
				Title:    "test",
				Content:  "content",
				CreateAt: timeNow,
				UpdateAt: timeNow,
			},
			args: args{
				request: ArticleUpdateRequest{
					ID:       1,
					Title:    "test",
					Content:  "content",
					UpdateAt: timeNow,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Article{
				ID:       tt.fields.ID,
				Title:    tt.fields.Title,
				Content:  tt.fields.Content,
				CreateAt: tt.fields.CreateAt,
				UpdateAt: tt.fields.UpdateAt,
			}
			m.FromUpdateRequest(tt.args.request)
		})
	}
}

func TestArticle_ToArticleResponse(t *testing.T) {
	timeNow := time.Now()
	type fields struct {
		ID       int64
		Title    string
		Content  string
		CreateAt time.Time
		UpdateAt time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   ArticleResponse
	}{
		{
			name: "success",
			fields: fields{
				ID:       1,
				Title:    "test",
				Content:  "content",
				CreateAt: timeNow,
				UpdateAt: timeNow,
			},
			want: ArticleResponse{
				ID:       1,
				Title:    "test",
				Content:  "content",
				CreateAt: timeNow,
				UpdateAt: timeNow,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Article{
				ID:       tt.fields.ID,
				Title:    tt.fields.Title,
				Content:  tt.fields.Content,
				CreateAt: tt.fields.CreateAt,
				UpdateAt: tt.fields.UpdateAt,
			}
			if got := m.ToArticleResponse(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Article.ToArticleResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}
