package dto

type SaveRefreshTokenInputBody struct {
	RefreshToken string `json:"refresh_token"`
	UserID       string `json:"user_id"`
}

type RefreshTokenInputBody struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
	Email        string `json:"email" binding:"required"`
}

type RefreshTokenResponse struct {
	AccessToken string `json:"access_token"`
}
