package dto

import (
	"time"
)

type LessonResponse struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Video     string    `json:"video"`
	ChapterID string    `json:"chapter_id"`
	CreatedAt time.Time `json:"created_at"`
}

type LessonInputBody struct {
	Name      string `json:"name" binding:"required"`
	Video     string `json:"video" binding:"required"`
	ChapterID string `json:"chapter_id"  binding:"required"`
}

type UpdateLessonInputBody struct {
	Name      string `json:"name"`
	Video     string `json:"video"`
	ChapterID string `json:"chapter_id"`
}
