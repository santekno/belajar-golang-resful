package http

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-playground/validator"
	"github.com/julienschmidt/httprouter"
	"github.com/santekno/belajar-golang-restful/internal/models"
	"github.com/santekno/belajar-golang-restful/internal/usecase"
	"github.com/santekno/belajar-golang-restful/pkg/util"
)

type Delivery struct {
	articleUsecase usecase.ArticleUsecase
	validate       *validator.Validate
}

func New(articleUsecase usecase.ArticleUsecase) *Delivery {
	return &Delivery{
		articleUsecase: articleUsecase,
		validate:       validator.New(),
	}
}

func (d *Delivery) GetAll(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var response models.ArticleListResponse
	var statusCode int = http.StatusBadRequest

	defer func() {
		util.Response(w, response, statusCode)
	}()

	res, err := d.articleUsecase.GetAll(r.Context())
	if err != nil {
		statusCode = http.StatusInternalServerError
		response.Code = statusCode
		response.Status = err.Error()
		return
	}

	statusCode = http.StatusOK
	response.Data = append(response.Data, res...)
	response.Code = statusCode
	response.Status = "OK"
}

func (d *Delivery) GetByID(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var response models.ArticleListResponse
	var statusCode int = http.StatusBadRequest

	defer func() {
		util.Response(w, response, statusCode)
	}()

	articleIDString := params.ByName("article_id")
	articleID, err := strconv.ParseInt(articleIDString, 0, 64)
	if err != nil {
		statusCode = http.StatusBadRequest
		response.Code = statusCode
		response.Status = err.Error()
		return
	}

	res, err := d.articleUsecase.GetByID(r.Context(), articleID)
	if err != nil {
		statusCode = http.StatusInternalServerError
		response.Code = statusCode
		response.Status = err.Error()
		return
	}

	if res.ID == 0 {
		statusCode = http.StatusNotFound
		response.Code = statusCode
		response.Status = "data not found"
		return
	}

	statusCode = http.StatusOK
	response.Data = append(response.Data, res)
	response.Code = statusCode
	response.Status = "OK"
}

func (d *Delivery) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var response models.ArticleListResponse
	var request models.ArticleUpdateRequest
	var statusCode = http.StatusBadRequest

	defer func() {
		util.Response(w, response, statusCode)
	}()

	articleIDString := params.ByName("article_id")
	articleID, err := strconv.ParseInt(articleIDString, 0, 64)
	if err != nil {
		statusCode = http.StatusBadRequest
		response.Code = statusCode
		response.Status = err.Error()
		return
	}

	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&request)
	if err != nil {
		statusCode = http.StatusBadRequest
		response.Code = statusCode
		response.Status = err.Error()
		return
	}

	request.ID = articleID

	err = d.validate.Struct(request)
	if err != nil {
		statusCode = http.StatusBadRequest
		response.Code = statusCode
		response.Status = err.Error()
		return
	}

	res, err := d.articleUsecase.Update(r.Context(), request)
	if err != nil {
		statusCode = http.StatusInternalServerError
		response.Code = statusCode
		response.Status = err.Error()
		return
	}

	statusCode = http.StatusOK
	response.Data = append(response.Data, res)
	response.Code = statusCode
	response.Status = "OK"
}

func (d *Delivery) Store(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var response models.ArticleListResponse
	var request models.ArticleCreateRequest
	var statusCode int = http.StatusBadRequest

	defer func() {
		util.Response(w, response, statusCode)
	}()

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&request)
	if err != nil {
		statusCode = http.StatusBadRequest
		response.Code = statusCode
		response.Status = err.Error()
		return
	}

	err = d.validate.Struct(request)
	if err != nil {
		statusCode = http.StatusBadRequest
		response.Code = statusCode
		response.Status = err.Error()
		return
	}

	res, err := d.articleUsecase.Store(r.Context(), request)
	if err != nil {
		statusCode = http.StatusInternalServerError
		response.Code = statusCode
		response.Status = err.Error()
		return
	}

	response.Data = append(response.Data, res)
	statusCode = http.StatusOK
	response.Code = statusCode
	response.Status = "OK"
}

func (d *Delivery) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var response models.ArticleListResponse
	var statusCode int = http.StatusBadRequest

	defer func() {
		util.Response(w, response, statusCode)
	}()

	articleIDString := params.ByName("article_id")
	articleID, err := strconv.ParseInt(articleIDString, 0, 64)
	if err != nil {
		statusCode = http.StatusBadRequest
		response.Code = statusCode
		response.Status = err.Error()
		return
	}

	if articleID == 0 {
		statusCode = http.StatusNotFound
		response.Code = statusCode
		response.Status = "article_id was not zero"
		return
	}

	res, err := d.articleUsecase.Delete(r.Context(), articleID)
	if err != nil {
		statusCode = http.StatusInternalServerError
		response.Code = statusCode
		response.Status = err.Error()
		return
	}

	if !res {
		statusCode = http.StatusInternalServerError
		response.Code = statusCode
		response.Status = errors.New("unknown error").Error()
		return
	}

	statusCode = http.StatusOK
	response.Code = statusCode
	response.Status = "OK"
}
