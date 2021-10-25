package params

type CreatePostRequest struct {
	Title      string `json:"title" bindding:"required"`
	Content    string `json:"content" bindding:"required"`
	AuthorId   int64  `json:"author_id" bindding:"required"`
	Status     string `json:"status" bindding:"required"`
	CategoryId string `json:"category_id" bindding:"required"`
}
