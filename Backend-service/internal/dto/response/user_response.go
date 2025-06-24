package response

type UserProfileResponse struct {
	UserID         int64  `json:"id"`
	Username       string `json:"username"`
	Fullname       string `json:"fullname"`
	Email          string `json:"email"`
	RoleID         int64  `json:"role_id"`
	ProfilePicture string `json:"profile_picture"`
	PhoneNumber    string `json:"phone_number"`
	Address        string `json:"address"`
}