package models

import "time"

type ArticleCreateRequest struct {
	Title    string `json:"title" validate:"required,max=160,min=10"`
	Content  string `json:"content" validate:"required"`
	CreateAt time.Time
}

type ArticleUpdateRequest struct {
	ID       int64  `json:"id" validate:"required"`
	Title    string `json:"title" validate:"required,max=160,min=10"`
	Content  string `json:"content" validate:"required"`
	UpdateAt time.Time
}

type ArticleListResponse struct {
	HeaderResponse
	Data []ArticleResponse
}

type ArticleResponse struct {
	ID       int64     `json:"id"`
	Title    string    `json:"title"`
	Content  string    `json:"content"`
	CreateAt time.Time `json:"create_at"`
	UpdateAt time.Time `json:"update_at"`
}

type HeaderResponse struct {
	Code   int
	Status string
}
