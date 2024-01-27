package models

import "time"

type ArticleCreateRequest struct {
	Title    string
	Content  string
	CreateAt time.Time
}

type ArticleUpdateRequest struct {
	ID       int64
	Title    string
	Content  string
	UpdateAt time.Time
}

type ArticleListResponse struct {
	HeaderResponse
	Data []ArticleResponse
}

type ArticleResponse struct {
	ID       int64
	Title    string
	Content  string
	CreateAt time.Time
	UpdateAt time.Time
}

type HeaderResponse struct {
	Code   int
	Status string
}
