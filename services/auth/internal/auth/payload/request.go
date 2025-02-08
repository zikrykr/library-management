package payload

type (
	SignUpReq struct {
		FullName string `json:"full_name" binding:"required"`
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
		Role     string `json:"role"`
	}

	LoginReq struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	GetUserFilter struct {
		Email string `json:"email"`

		Page   int    `json:"page"`
		Limit  int    `json:"limit"`
		SortBy string `json:"sort_by"`
	}
)
