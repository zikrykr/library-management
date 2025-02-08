package payload

type CreateCategoryReq struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
}

type UpdateCategoryReq struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
}

type GetCategoriesReq struct {
	ID   string   `form:"id" json:"id"`
	IDIN []string `form:"id_in" json:"id_in"`

	Name string `form:"name" json:"name"`

	Search string `form:"keyword" json:"keyword"`
	SortBy string `form:"sort" json:"sort"`
	Page   int    `json:"page" form:"page" default:"1"`
	Limit  int    `json:"limit" form:"limit" default:"15"`
}
