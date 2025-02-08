package payload

type CreateAuthorReq struct {
	Name string `json:"name" binding:"required"`
	Bio  string `json:"bio" binding:"required"`
}

type UpdateAuthorReq struct {
	Name string `json:"name" binding:"required"`
	Bio  string `json:"bio" binding:"required"`
}

type GetAuthorsReq struct {
	ID   string   `form:"id" json:"id"`
	IDIN []string `form:"id_in" json:"id_in"`

	Name string `form:"name" json:"name"`

	SortBy string `form:"sort" json:"sort"`
	Page   int    `json:"page" form:"page" default:"1"`
	Limit  int    `json:"limit" form:"limit" default:"15"`
}
