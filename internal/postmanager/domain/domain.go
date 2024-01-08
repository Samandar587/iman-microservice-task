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
	Page    int
}

func (p *Post) GetID() int {
	return p.id
}
func (p *Post) GetOriginalID() int {
	return p.original_post_id
}
func (p *Post) GetUserID() int {
	return p.user_id
}
func (p *Post) GetTitle() string {
	return p.title
}
func (p *Post) GetBody() string {
	return p.body
}
func (p *Post) GetPage() int {
	return p.page
}

type PostRepository interface {
	Save(newPost *NewPost) (int, error)
	FindByID(id int) (*Post, error)
	Update(id int, req *NewPost) (*Post, error)
	Delete(id int) (string, error)
}
