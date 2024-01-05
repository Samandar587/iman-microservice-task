package grpc

import (
	"context"
	"golang-project-template/internal/datafetcher/app"
	"golang-project-template/internal/datafetcher/ports/grpc/proto/pb"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct {
	postUsecase app.PostUsecase
	pb.UnimplementedSavePostsServiceServer
}

func NewDataFetcherServer(postUsecase app.PostUsecase) *server {
	return &server{
		postUsecase: postUsecase,
	}
}

func (s *server) CollectPosts(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	err := s.postUsecase.CollectPosts()
	if err != nil {
		return &pb.Response{}, status.Error(codes.Internal, "error collecting posts: "+err.Error())
	}

	res := &pb.Response{
		Message: "Successfully fetched data!",
	}

	return res, nil
}
