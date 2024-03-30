package article

import (
	"context"
	"errors"

	"github.com/santekno/belajar-golang-restful/models"
	repository "github.com/santekno/belajar-golang-restful/repository"
)

type Usecase struct {
	articleRepository repository.ArticleRepository
}

func New(repo repository.ArticleRepository) *Usecase {
	return &Usecase{articleRepository: repo}
}

func (u *Usecase) GetAll(ctx context.Context) ([]models.ArticleResponse, error) {
	var response []models.ArticleResponse

	res, err := u.articleRepository.GetAll(ctx)
	if err != nil {
		return response, err
	}

	for _, v := range res {
		response = append(response, v.ToArticleResponse())
	}

	return response, nil
}

func (u *Usecase) GetByID(ctx context.Context, id int64) (models.ArticleResponse, error) {
	var response models.ArticleResponse

	res, err := u.articleRepository.GetByID(ctx, id)
	if err != nil {
		return response, err
	}

	response = res.ToArticleResponse()
	return response, nil
}

func (u *Usecase) Update(ctx context.Context, request models.ArticleUpdateRequest) (models.ArticleResponse, error) {
	var article = new(models.Article)
	var response models.ArticleResponse

	if request.ID == 0 {
		return response, errors.New("request article_id do not zero or empty")
	}

	// validate data from database was found
	article, err := u.articleRepository.GetByID(ctx, request.ID)
	if err != nil {
		return response, err
	}

	if article.ID == 0 {
		return response, errors.New("data not found")
	}

	// from request to article struct
	article.FromUpdateRequest(request)

	// execute update
	article, err = u.articleRepository.Update(ctx, article)
	if err != nil {
		return response, err
	}

	// append data update to response
	response = article.ToArticleResponse()
	return response, nil
}

func (u *Usecase) Store(ctx context.Context, request models.ArticleCreateRequest) (models.ArticleResponse, error) {
	var article = new(models.Article)
	var response models.ArticleResponse

	// convert request to article struct
	article.FromCreateRequest(request)

	// executed store
	articleID, err := u.articleRepository.Store(ctx, article)
	if err != nil {
		return response, err
	}

	article.ID = articleID
	response = article.ToArticleResponse()
	return response, nil
}

func (u *Usecase) Delete(ctx context.Context, id int64) (bool, error) {
	var isSuccss bool

	// validate data from database was found
	article, err := u.articleRepository.GetByID(ctx, id)
	if err != nil {
		return isSuccss, err
	}

	if article == nil {
		return isSuccss, errors.New("article not found")
	}

	isSuccss, err = u.articleRepository.Delete(ctx, id)
	if err != nil {
		return isSuccss, err
	}

	return isSuccss, nil
}
