package dto

type MentorResponse struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Profile    string `json:"email"`
	Profession string `json:"profession"`
}

type MentorInputBody struct {
	Name       string `json:"name" binding:"required"`
	Profile    string `json:"email" binding:"required"`
	Profession string `json:"profession" binding:"required"`
}

type UpdateMentorInputBody struct {
	Name       string `json:"name" `
	Profile    string `json:"email" `
	Profession string `json:"profession" `
}
