package domain

type Post struct {
	id               int
	original_post_id int
	user_id          int
	title            string
	body             string
	page             int
}

type NewPost struct {
	User_id int
	Title   string
	Body    string
}

type PostRepository interface {
	Create(newPost *NewPost) (int, error)
	FindByID(id int) (*Post, error)
	FindByPage(page int) (*[]Post, error)
	Update(id int) (*Post, error)
	Delete(id int) (string, error)
}
