package dto

import "time"

type MediaResponse struct {
	ID        string    `json:"id"`
	Images    string    `json:"images"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UploadMediaBody struct {
	Image string `json:"image" binding:"required"`
}
