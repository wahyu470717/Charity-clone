package request

type SignInRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RegisterRequest struct {
	Username    string `json:"username" binding:"required,min=3,max=50"`
	Email       string `json:"email" binding:"required,email"`
	Password    string `json:"password" binding:"required,min=6"`
	Fullname    string `json:"fullname" binding:"required"`
	PhoneNumber string `json:"phone_number,omitempty"`
	Address     string `json:"address,omitempty"`
	Role        int64  `json:"role,omitempty"`
}

type ChangePasswordRequest struct {
	Password string `json:"password" binding:"required,min=6"`
}

type ForgetPasswordRequest struct {
	Email string `json:"email" binding:"required,email"`
}
