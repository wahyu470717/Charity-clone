package response

type SignInResponse struct {
	UserID   int64  `json:"user_id"`
	UserName string `json:"username"`
	Role     string `json:"role"`
	Token    string `json:"token"`
}

type RegisterResponse struct {
	UserID   int64  `json:"user_id"`
	UserName string `json:"username"`
	Email    string `json:"email"`
	Fullname string `json:"fullname"`
	RoleID   int64  `json:"role"`
}

