package app

import (
	"errors"
	"golang-project-template/internal/datafetcher/domain"
	"log"
	"strconv"
	"sync"
)

type postUsecase struct {
	postRepository domain.PostRepository
	postProvider   domain.PostProvider
}

type PostUsecase interface {
	CollectPosts() error
}

func NewPostUsecase(postRepo domain.PostRepository, postProvider domain.PostProvider) PostUsecase {
	return &postUsecase{
		postRepository: postRepo,
		postProvider:   postProvider,
	}
}
func (p *postUsecase) CollectPosts() error {
	var wg sync.WaitGroup
	errCh := make(chan error, 1)
	var mu sync.Mutex // Mutex for protecting concurrent writes to postList
	var postList []domain.Post

	for i := 1; i <= 50; i++ {
		wg.Add(1)

		go func(pageNum int) {
			defer wg.Done()

			pageStr := strconv.Itoa(pageNum)
			posts, err := p.postProvider.GetPosts(pageStr)
			if err != nil {
				errCh <- err
				return
			}

			mu.Lock()
			postList = append(postList, posts...)
			mu.Unlock()
		}(i)
	}

	go func() {
		wg.Wait()
		close(errCh)
	}()

	for err := range errCh {
		if err != nil {
			return err
		}
	}

	for _, post := range postList {
		_, err := p.postRepository.Save(&post)
		if err != nil {
			log.Printf("Error: " + err.Error())
			return errors.New("Unable to save the data into the database. Please try again later!")
		}
	}

	return nil
}
