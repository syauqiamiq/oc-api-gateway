package dto

import (
	"time"
)

type ChapterResponse struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CourseID  string    `json:"course_id"`
	CreatedAt time.Time `json:"created_at"`
}

type ChapterInputBody struct {
	Name     string `json:"name" binding:"required"`
	CourseID string `json:"course_id"  binding:"required"`
}

type UpdateChapterInputBody struct {
	Name     string `json:"name" `
	CourseID string `json:"course_id"  `
}
