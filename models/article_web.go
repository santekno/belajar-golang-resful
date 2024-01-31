package models

import "time"

type ArticleCreateRequest struct {
	Title    string `validate:"required,max=160,min=0"`
	Content  string `validate:"required"`
	CreateAt time.Time
}

type ArticleUpdateRequest struct {
	ID       int64  `validate:"required"`
	Title    string `validate:"required,max=160,min=0"`
	Content  string `validate:"required"`
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
