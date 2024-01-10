package app

import (
	"context"
	fetcher "golang-project-template/internal/datafetcher/ports/proto/pb"
	"log"
	"time"

	manager "golang-project-template/internal/postmanager/ports/grpc/proto/pb"
)

type Service interface {
	DataFetcherService
	PostManagerService
}

type DataFetcherService interface {
	CollectPosts() (map[string]string, error)
}

type PostManagerService interface {
	Create(*NewPost) (int, error)
	GetById(id int) (*Post, error)
	Update(field *NewPost) (*Post, error)
	Delete(id int) (map[string]string, error)
}

type service struct {
	fClient fetcher.SavePostsServiceClient
	mClient manager.ManagePostsServiceClient
}

// init service
func NewUsecase(
	fClient fetcher.SavePostsServiceClient,
	mClient manager.ManagePostsServiceClient,
) Service {
	return &service{
		fClient: fClient,
		mClient: mClient,
	}
}

func (s *service) CollectPosts() (map[string]string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	req := fetcher.Request{}

	res, err := s.fClient.CollectPosts(ctx, &req)
	if err != nil {
		log.Println("error calling grpx server: ", err)
		return nil, err
	}

	var response = map[string]string{
		"message": res.Message,
	}

	return response, nil
}

func (s *service) Create(req *NewPost) (int, error) {
	ctx := context.Background()
	reqBody := manager.NewPostRequest{
		UserId: int64(req.User_id),
		Title:  req.Title,
		Body:   req.Body,
		Page:   int64(req.Page),
	}
	response, err := s.mClient.Create(ctx, &reqBody)
	if err != nil {
		return 0, err
	}
	var id = int(response.Id)

	return id, nil
}

func (s *service) Delete(id int) (map[string]string, error) {
	ctx := context.Background()
	req := manager.IdRequest{
		Id: int64(id),
	}
	response, err := s.mClient.Delete(ctx, &req)
	if err != nil {
		return nil, err
	}
	msg := response.Msg
	return msg, nil
}

func (s *service) GetById(id int) (*Post, error) {
	ctx := context.Background()
	req := manager.IdRequest{
		Id: int64(id),
	}

	res, err := s.mClient.GetByID(ctx, &req)
	if err != nil {
		return nil, err
	}

	post := parseToModel(
		int(res.Id),
		int(res.OriginalPostId),
		int(res.UserId),
		res.Title,
		res.Body,
		int(res.Page),
	)

	return post, nil
}

func (s *service) Update(field *NewPost) (*Post, error) {
	ctx := context.Background()
	req := manager.UpdateRequest{
		Id:     int64(field.ID),
		UserId: int64(field.User_id),
		Title:  field.Title,
		Body:   field.Body,
		Page:   int64(field.Page),
	}

	res, err := s.mClient.Update(ctx, &req)
	if err != nil {
		return nil, err
	}
	post := parseToModel(
		field.ID,
		int(res.OriginalPostId),
		int(res.UserId),
		res.Title,
		res.Body,
		int(res.Page),
	)

	return post, nil
}

func parseToModel(id, original_post_id, user_id int, title, body string, page int) *Post {
	return &Post{
		ID:               id,
		Original_post_id: original_post_id,
		User_id:          user_id,
		Title:            title,
		Body:             body,
		Page:             page,
	}
}
