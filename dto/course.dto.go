package dto

import (
	"time"
)

type CourseResponse struct {
	ID          string    `json:"id"`
	Certificate *bool     `json:"certificate"`
	Type        string    `json:"type"`
	Status      string    `json:"status"`
	Level       string    `json:"level"`
	MentorID    string    `json:"mentor_id"`
	Name        string    `json:"name"`
	CreatedAt   time.Time `json:"created_at"`
}

type CourseInputBody struct {
	Certificate *bool  `json:"certificate" binding:"required"`
	Type        string `json:"type" binding:"required"`
	Status      string `json:"status" binding:"required"`
	Level       string `json:"level" binding:"required"`
	MentorID    string `json:"mentor_id" binding:"required"`
	Name        string `json:"name" binding:"required"`
}

type UpdateCourseInputBody struct {
	Certificate *bool  `json:"certificate"`
	Type        string `json:"type"`
	Status      string `json:"status"`
	Level       string `json:"level"`
	MentorID    string `json:"mentor_id"`
	Name        string `json:"name"`
}
