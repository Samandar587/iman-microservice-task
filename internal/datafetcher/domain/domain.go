package domain

type Post struct {
	id      int
	user_id int
	title   string
	body    string
	page    int
}

// Getters
func (p *Post) GetID() int {
	return p.id
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
	Save(post *Post) (int, error)
	UserIdExists(page int) (bool, error)
}

type PostProvider interface {
	GetPosts(page string) ([]Post, error)
}
