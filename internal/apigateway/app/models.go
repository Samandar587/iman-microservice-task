package app

type NewPost struct {
	ID      int    `json:"id"`
	User_id int    `json:"user_id"`
	Title   string `json:"title"`
	Body    string `json:"body"`
	Page    int    `json:"page"`
}

type Post struct {
	ID               int    `json:"id"`
	Original_post_id int    `json:"original_post_id"`
	User_id          int    `json:"user_id"`
	Title            string `json:"title"`
	Body             string `json:"body"`
	Page             int    `json:"page"`
}
