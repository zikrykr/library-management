package payload

type (
	LoginResp struct {
		AccessToken string `json:"access_token"`
	}

	GetProfileResp struct {
		ID       string `json:"id"`
		FullName string `json:"full_name"`
		Email    string `json:"email"`
		Role     string `json:"role"`
	}
)
