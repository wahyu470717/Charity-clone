package request

type UpdateProfileRequest struct {
	Fullname       string `json:"fullname"`
	PhoneNumber    string `json:"phone_number"`
	Address        string `json:"address"`
	ProfilePicture string `json:"profile_picture"`
}
