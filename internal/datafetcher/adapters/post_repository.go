package adapters

import (
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
