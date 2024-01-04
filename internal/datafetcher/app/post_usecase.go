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

	var allPosts []domain.Post

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

			allPosts = append(allPosts, posts...)
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

	for _, post := range allPosts {
		exists, err := p.postRepository.UserIdExists(post.GetUserID())

		if err != nil {
			log.Printf("Error checking page existence: %v", err)
			return errors.New("Unable to check page existence at the moment. Please try again later!")
		}

		if !exists {
			_, err = p.postRepository.Save(&post)
			if err != nil {
				log.Printf("Error: " + err.Error())
				return errors.New("Unable to save the data into database. Please, try again later!")
			}
		} else {
			log.Println("warning: post with user id already exists")
		}

	}

	return nil
}
