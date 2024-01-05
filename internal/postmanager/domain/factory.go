package domain

type Factory struct {
}

func (f Factory) ParseToDomain(id, original_post_id, user_id int, title, body string, page int) *Post {
	return &Post{
		id:               id,
		original_post_id: original_post_id,
		user_id:          user_id,
		title:            title,
		body:             body,
		page:             page,
	}
}
