package app

import (
	"golang-project-template/internal/postmanager/adapters"
	"golang-project-template/internal/postmanager/domain"
)

type PostUsecase struct {
	postRepo *adapters.PostRepo
}

type PostUsecaseInterface interface {
	Create(req domain.NewPost) (int, error)
	GetByID(id int) (*domain.Post, error)
	Update(id int, req *domain.NewPost) (*domain.Post, error)
	Delete(id int) (string, error)
}

func NewPostUsecase(postRepo *adapters.PostRepo) PostUsecaseInterface {
	return &PostUsecase{
		postRepo: postRepo,
	}
}

func (p *PostUsecase) Create(req domain.NewPost) (int, error) {

	id, err := p.postRepo.Save(&req)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (p *PostUsecase) Delete(id int) (string, error) {
	str, err := p.postRepo.Delete(id)
	if err != nil {
		return "", err
	}
	return str, nil
}

func (p *PostUsecase) GetByID(id int) (*domain.Post, error) {
	post, err := p.postRepo.FindByID(id)
	if err != nil {
		return nil, err
	}
	return post, nil
}

func (p *PostUsecase) Update(id int, req *domain.NewPost) (*domain.Post, error) {
	post, err := p.postRepo.Update(id, req)
	if err != nil {
		return nil, err
	}
	return post, nil
}
