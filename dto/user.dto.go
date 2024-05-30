package dto

type RegisterInputBody struct {
	Name       string `json:"name" binding:"required"`
	Email      string `json:"email" binding:"required"`
	Password   string `json:"password" binding:"required"`
	Profession string `json:"profession" binding:"required"`
	Avatar     string `json:"avatar"`
}

type UpdateUserInputBody struct {
	Name       string `json:"name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Profession string `json:"profession"`
	Avatar     string `json:"avatar"`
}

type LogoutBody struct {
	UserID string `json:"user_id" binding:"required"`
}

type LoginInputBody struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type User struct {
	ID    string
	Name  string
	Email string
}
