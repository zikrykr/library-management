package payload

type CreateBookReq struct {
	Title          string `json:"title" binding:"required"`
	Description    string `json:"description" binding:"required"`
	ISBN           string `json:"isbn" binding:"required"`
	AuthorID       string `json:"author_id" binding:"required"`
	CategoryID     string `json:"category_id" binding:"required"`
	PublishedYear  int    `json:"published_year" binding:"required"`
	TotalStock     int    `json:"total_stock"`
	AvailableStock int    `json:"available_stock"`
}

type UpdateBookReq struct {
	Title          string `json:"title"`
	Description    string `json:"description"`
	ISBN           string `json:"isbn"`
	AuthorID       string `json:"author_id"`
	CategoryID     string `json:"category_id"`
	PublishedYear  int    `json:"published_year"`
	TotalStock     int    `json:"total_stock"`
	AvailableStock int    `json:"available_stock"`
}

type GetBooksReq struct {
	Title      string   `form:"title" json:"title"`
	ID         string   `form:"id" json:"id"`
	IDIN       []string `form:"id_in" json:"id_in"`
	AuthorID   string   `form:"author_id" json:"author_id"`
	CategoryID string   `form:"category_id" json:"category_id"`

	SortBy string `form:"sort" json:"sort"`
	Page   int    `json:"page" form:"page" default:"1"`
	Limit  int    `json:"limit" form:"limit" default:"15"`
}
