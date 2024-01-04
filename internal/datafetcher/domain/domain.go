package domain

type Post struct {
	id      int
	user_id int
	title   string
	body    string
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

// Setters
func (p *Post) SetID(id int) {
	p.id = id
}
func (p *Post) SetUserID(user_id int) {
	p.user_id = user_id
}

func (p *Post) SetTitle(title string) {
	p.title = title
}

func (p *Post) SetBody(body string) {
	p.body = body
}

type PostRepository interface {
	Save(post *Post) (int, error)
}

type PostProvider interface {
	GetPosts(page string) ([]Post, error)
}
