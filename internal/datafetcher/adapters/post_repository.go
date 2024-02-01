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
	INSERT INTO posts (original_post_id, user_id, title, body, page)
	VALUES ($1, $2, $3, $4, $5)
	ON CONFLICT (original_post_id) DO UPDATE 
	SET original_post_id = EXCLUDED.original_post_id
	RETURNING id;
	`
	err := postRepo.db.QueryRow(sqlStatement, post.GetOriginalPostID(), post.GetUserID(), post.GetTitle(), post.GetBody(), post.GetPage()).Scan(&id)
	if err != nil {
		log.Printf("Error executing query: %v\nSQL: %s\nID: %d, Title: %s, Body: %s\n", err, sqlStatement, post.GetUserID(), post.GetTitle(), post.GetBody())
		return 0, fmt.Errorf("failed to save post: %w", err)
	}

	return id, nil
}
