package dto

import "time"

type MidtransPaymentResponse struct {
	Token       string `json:"token"`
	RedirectUrl string `json:"redirect_url"`
}

type CheckoutOrderResponse struct {
	AssignStatus    string                  `json:"assign_status"`
	MidtransPayment MidtransPaymentResponse `json:"midtrans_payment"`
}

type CheckoutOrderInputBody struct {
	UserID   string `json:"user_id" binding:"required"`
	CourseID string `json:"course_id" binding:"required"`
}

type MetaData struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Certificate bool      `json:"certificate"`
	Thumbnail   string    `json:"thumbnail"`
	Type        string    `json:"type"`
	Status      string    `json:"status"`
	Price       int       `json:"price"`
	Level       string    `json:"level"`
	Description string    `json:"description"`
	MentorID    string    `json:"mentor_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
type OrderResponse struct {
	ID              string    `json:"id"`
	Status          string    `json:"status"`
	UserID          string    `json:"user_id"`
	CourseID        string    `json:"course_id"`
	SnapURL         string    `json:"snap_url"`
	ExternalOrderID string    `json:"external_order_id"`
	MetaData        MetaData  `json:"meta_data"`
	CreatedAt       time.Time `json:"created_at"`
}
