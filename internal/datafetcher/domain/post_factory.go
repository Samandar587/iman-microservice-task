package domain

type PostFactory struct {
}

func (p *PostFactory) ParseToDomain(original_post_id int, user_id int, title, body string, page int) *Post {
	return &Post{
		original_post_id: original_post_id,
		user_id:          user_id,
		title:            title,
		body:             body,
		page:             page,
	}
}
