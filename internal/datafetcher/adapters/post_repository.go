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
	var id int
	sqlStatement := `
	INSERT into POSTS (user_id, title, body)
	VALUES ($1, $2, $3)
	RETURNING id
	`

	err := postRepo.db.QueryRow(sqlStatement, post.GetUserID(), post.GetTitle(), post.GetBody()).Scan(&id)
	if err != nil {
		log.Println("internal error: " + err.Error())
		return 0, err
	}

	return id, nil
}

func (postRepo *postRepository) PageExists(page int) (bool, error) {
	if postRepo.db == nil {
		return false, errors.New("database connection is nil")
	}

	if page < 1 {
		return false, fmt.Errorf("invalid page number: %d", page)
	}
	var exists bool
	sqlStatement := `
		SELECT 1
		FROM posts
		WHERE page = $1
	`

	err := postRepo.db.QueryRow(sqlStatement, page).Scan(&exists)
	if err != nil {
		log.Printf("Error executing query: %v\nSQL: %s\nPage: %d\n", err, sqlStatement, page)
		return false, err
	}
	return exists, nil
}
