package dto

import "time"

type MyCourseResponse struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	CourseID  string    `json:"course_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
