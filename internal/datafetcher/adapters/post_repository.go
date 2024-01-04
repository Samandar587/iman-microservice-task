package adapters

import (
	"errors"
	"fmt"
	"golang-project-template/internal/datafetcher/domain"
	"log"

	"github.com/jackc/pgx"
)

type postRepository struct {
	db *pgx.Conn
}

func NewPostRepository(db *pgx.Conn) *postRepository {
	return &postRepository{
		db: db,
	}
}

func (postRepo *postRepository) Save(post *domain.Post) (int, error) {
	if postRepo.db == nil {
		return 0, errors.New("database connection is nil")
	}

	var id int
	sqlStatement := `
	INSERT INTO posts (user_id, title, body, page)
	VALUES ($1, $2, $3, $4)
	RETURNING id
	`
	err := postRepo.db.QueryRow(sqlStatement, post.GetUserID(), post.GetTitle(), post.GetBody(), post.GetPage()).Scan(&id)
	if err != nil {
		log.Printf("Error executing query: %v\nSQL: %s\nUserID: %d, Title: %s, Body: %s\n", err, sqlStatement, post.GetUserID(), post.GetTitle(), post.GetBody())
		return 0, fmt.Errorf("failed to save post: %w", err)
	}

	return id, nil
}

func (postRepo *postRepository) UserIdExists(userID int) (bool, error) {
	if postRepo.db == nil {
		return false, errors.New("database connection is nil")
	}

	if userID < 1 {
		return false, fmt.Errorf("invalid user ID: %d", userID)
	}

	var exists bool
	sqlStatement := `
		SELECT EXISTS (
			SELECT 1
			FROM posts
			WHERE user_id = $1
		)
	`
	err := postRepo.db.QueryRow(sqlStatement, userID).Scan(&exists)
	if err != nil {
		log.Printf("Error executing query: %v\nSQL: %s\nPage: %d\n", err, sqlStatement, userID)
		return false, err
	}
	return exists, nil
}
