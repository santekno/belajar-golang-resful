package models

import "time"

type Article struct {
	ID       int64
	Title    string
	Content  string
	CreateAt time.Time
	UpdateAt time.Time
}

func (m *Article) FromCreateRequest(request ArticleCreateRequest) {
	m.Title = request.Title
	m.Content = request.Content
	m.CreateAt = time.Now()
	m.UpdateAt = time.Now()
}

func (m *Article) FromUpdateRequest(request ArticleUpdateRequest) {
	m.ID = request.ID
	m.Title = request.Title
	m.Content = request.Content
	m.UpdateAt = time.Now()
}

func (m *Article) ToArticleResponse() ArticleResponse {
	return ArticleResponse{
		ID:       m.ID,
		Title:    m.Title,
		Content:  m.Content,
		CreateAt: m.CreateAt,
		UpdateAt: m.UpdateAt,
	}
}
