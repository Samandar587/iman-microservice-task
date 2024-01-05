package adapters

import (
	"context"
	"fmt"
	"golang-project-template/internal/postmanager/domain"
	"log"
	"time"

	"github.com/jackc/pgx"
)

type PostRepo struct {
	db *pgx.Conn
	f  domain.Factory
}

func NewpostRepsitory(db *pgx.Conn) domain.PostRepository {
	return &PostRepo{
		db: db,
	}
}

func (p *PostRepo) Save(newPost *domain.NewPost) (int, error) {
	var id int
	sqlStatement := `
		INSERT INTO posts (user_id, title, body, page)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := p.db.QueryRowEx(ctx, sqlStatement, nil, newPost.User_id, newPost.Title, newPost.Body, newPost.Page).Scan(&id)
	if err != nil {
		log.Printf("Error while creating a new post: SQL:%v\n, VALUES:%v\n, ERR:%v\n", sqlStatement, newPost, err)
		return 0, fmt.Errorf("failed to save post: %w", err)
	}
	return id, nil
}

func (p *PostRepo) FindByID(id int) (*domain.Post, error) {
	var original_post_id, user_id, page int
	var title, body string
	sqlStatement := `
		SELECT id, original_post_id, user_id, title, body, page
		FROM posts
		WHERE id = $1
	`

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := p.db.QueryRowEx(ctx, sqlStatement, nil, id).Scan(
		&id,
		&original_post_id,
		&user_id,
		&title,
		&body,
		&page,
	)
	if err != nil {
		log.Printf("Error while finding a post: SQL:%v\n, ID:%v\n, ERR:%v\n", sqlStatement, id, err)
		return nil, fmt.Errorf("failed to get post: %w", err)
	}
	post := p.f.ParseToDomain(id, original_post_id, user_id, title, body, page)

	return post, nil
}

func (p *PostRepo) Update(id int, req *domain.NewPost) (*domain.Post, error) {
	var original_post_id, user_id int

	sqlStatement := `
		UPDATE posts
		SET title = $1, body = $2
		WHERE id = $3
	`
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := p.db.ExecEx(ctx, sqlStatement, nil, req.Title, req.Body, id)
	if err != nil {
		log.Printf("Error while updating post: SQL:%v\n, TITLE:%v\n, BODY:%v\n, ID:%v\n, ERR:%v\n", sqlStatement, req.Title, req.Body, id, err)
		return nil, fmt.Errorf("failed to get posts by page: %w", err)
	}
	sql := `
		SELECT *
		FROM posts
		WHERE id = $1
	`

	err = p.db.QueryRowEx(ctx, sql, nil, id).Scan(
		&id,
		&original_post_id,
		&user_id,
		&req.Title,
		&req.Body,
		&req.Page,
	)
	if err != nil {
		log.Printf("Error while finding a post: SQL:%v\n, ID:%v\n, ERR:%v\n", sqlStatement, id, err)
		return nil, fmt.Errorf("failed to get post: %w", err)
	}
	post := p.f.ParseToDomain(id, original_post_id, user_id, req.Title, req.Body, req.Page)
	return post, nil
}

func (p *PostRepo) Delete(id int) (string, error) {
	if id < 0 {
		return "", fmt.Errorf("invalid post ID: %d", id)
	}
	var msg = "Successfully deleted!"
	sql := `
		DELETE
		FROM posts
		WHERE id = $1
	`
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := p.db.ExecEx(ctx, sql, nil, id)
	if err != nil {
		log.Printf("Error deleting post with ID %d: SQL: %s, Error: %v", id, sql, err)
		return "", fmt.Errorf("failed to delete post with ID %d: %w", id, err)

	}
	return msg, nil
}
