package user

type RegisterUserDTO struct {
	Username     string `json:"username"`
	Password     string `json:"password"`
	Email        string `json:"email"`
	Type         string `json:"type"`
	ProfilePhoto string `json:"profile_photo"`
}

type LoginUserDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
