package grpc

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"golang-project-template/internal/postmanager/app"
	"golang-project-template/internal/postmanager/domain"
	"golang-project-template/internal/postmanager/ports/grpc/proto/pb"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct {
	pb.UnimplementedManagePostsServiceServer
	postUsecase app.PostUsecaseInterface
}

func NewPostManagerGrpcServer(postUsecase app.PostUsecaseInterface) *server {
	return &server{
		postUsecase: postUsecase,
	}
}

func (s *server) Create(ctx context.Context, req *pb.NewPostRequest) (*pb.CreateResponse, error) {
	var newPost = domain.NewPost{
		User_id: int(req.UserId),
		Title:   req.Title,
		Body:    req.Body,
		Page:    int(req.Page),
	}

	id, err := s.postUsecase.Create(newPost)
	if err != nil {
		return &pb.CreateResponse{}, status.Error(codes.Internal, "error creating a new post: "+err.Error())
	}
	return &pb.CreateResponse{Id: int64(id)}, nil
}

func (s *server) GetByID(ctx context.Context, req *pb.IdRequest) (*pb.PostResponse, error) {
	fmt.Println("coming to getbyid. .....")
	post, err := s.postUsecase.GetByID(int(req.Id))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &pb.PostResponse{}, status.Error(codes.NotFound, "post with the provided id is not found")
		}
		return &pb.PostResponse{}, status.Error(codes.Internal, "error getting post with id: "+err.Error())
	}
	var response = &pb.PostResponse{
		Id:             int64(post.GetID()),
		OriginalPostId: int64(post.GetOriginalID()),
		UserId:         int64(post.GetUserID()),
		Title:          post.GetTitle(),
		Body:           post.GetBody(),
		Page:           int64(post.GetPage()),
	}
	return response, nil
}

func (s *server) Update(ctx context.Context, req *pb.UpdateRequest) (*pb.PostResponse, error) {
	var request = domain.NewPost{
		User_id: int(req.UserId),
		Title:   req.Title,
		Body:    req.Body,
		Page:    int(req.Page),
	}
	var id = int(req.Id)

	post, err := s.postUsecase.Update(id, &request)
	if err != nil {
		return &pb.PostResponse{}, status.Error(codes.Internal, "error updating the post: "+err.Error())
	}

	var response = &pb.PostResponse{
		Id:             int64(post.GetID()),
		OriginalPostId: int64(post.GetID()),
		UserId:         int64(post.GetUserID()),
		Title:          post.GetTitle(),
		Body:           post.GetTitle(),
		Page:           int64(post.GetPage()),
	}

	return response, nil
}

func (s *server) Delete(ctx context.Context, req *pb.IdRequest) (*pb.DeleteResponse, error) {
	var id = int64(req.Id)
	msg, err := s.postUsecase.Delete(int(id))
	if err != nil {
		return &pb.DeleteResponse{}, status.Error(codes.Internal, "error deleting the post: "+err.Error())
	}

	var response = &pb.DeleteResponse{
		Msg: msg,
	}

	return response, nil
}
