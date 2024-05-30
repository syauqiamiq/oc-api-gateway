package dto

type SaveRefreshTokenInputBody struct {
	RefreshToken string `json:"refresh_token"`
	UserID       string `json:"user_id"`
}
