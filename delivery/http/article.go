package http

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/santekno/belajar-golang-restful/models"
	usecase "github.com/santekno/belajar-golang-restful/usecease"
)

type Delivery struct {
	articleUsecase usecase.ArticleUsecase
}

func New(articleUsecase usecase.ArticleUsecase) *Delivery {
	return &Delivery{
		articleUsecase: articleUsecase,
	}
}

func (d *Delivery) GetAll(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var response models.ArticleListResponse
	res, err := d.articleUsecase.GetAll(r.Context())
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Status = err.Error()
	}

	response.Data = append(response.Data, res...)
	response.Code = http.StatusOK
	response.Status = "OK"
}

func (d *Delivery) GetByID(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var response models.ArticleListResponse

	articleIDString := params.ByName("article_id")
	articleID, err := strconv.ParseInt(articleIDString, 0, 64)
	if err != nil {
		response.Code = http.StatusBadRequest
		response.Status = err.Error()
	}

	res, err := d.articleUsecase.GetByID(r.Context(), articleID)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Status = err.Error()
	}

	response.Data = append(response.Data, res)
	response.Code = http.StatusOK
	response.Status = "OK"
}

func (d *Delivery) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var response models.ArticleListResponse
	var request models.ArticleUpdateRequest

	articleIDString := params.ByName("article_id")
	articleID, err := strconv.ParseInt(articleIDString, 0, 64)
	if err != nil {
		response.Code = http.StatusBadRequest
		response.Status = err.Error()
	}

	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&request)
	if err != nil {
		response.Code = http.StatusBadRequest
		response.Status = err.Error()
	}

	request.ID = articleID

	res, err := d.articleUsecase.Update(r.Context(), request)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Status = err.Error()
	}

	response.Data = append(response.Data, res)
	response.Code = http.StatusOK
	response.Status = "OK"
}

func (d *Delivery) Store(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var response models.ArticleListResponse
	var request models.ArticleCreateRequest

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&request)
	if err != nil {
		response.Code = http.StatusBadRequest
		response.Status = err.Error()
	}

	res, err := d.articleUsecase.Store(r.Context(), request)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Status = err.Error()
	}

	response.Data = append(response.Data, res)
	response.Code = http.StatusOK
	response.Status = "OK"
}

func (d *Delivery) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var response models.ArticleListResponse

	articleIDString := params.ByName("article_id")
	articleID, err := strconv.ParseInt(articleIDString, 0, 64)
	if err != nil {
		response.Code = http.StatusBadRequest
		response.Status = err.Error()
	}

	res, err := d.articleUsecase.Delete(r.Context(), articleID)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Status = err.Error()
	}

	if !res {
		response.Code = http.StatusInternalServerError
		response.Status = errors.New("unknown error").Error()
	}

	response.Code = http.StatusOK
	response.Status = "OK"
}
