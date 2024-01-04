package domain

type PostFactory struct {
}

func (p *PostFactory) ParseToDomain(user_id int, title, body string) *Post {
	return &Post{
		user_id: user_id,
		title:   title,
		body:    body,
	}
}
