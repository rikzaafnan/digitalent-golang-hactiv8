package dto

import "gopkg.in/guregu/null.v4"

type ItemRequest struct {
	ItemCode    string `json:"itemCode" valid:"required~itemCode cannot be empty" `
	Description string `json:"description" valid:"required~description cannot be empty" `
	Quantity    int    `json:"quantity" valid:"required~quantity cannot be empty" `
	OrderId     int    `json:"orderId" valid:"required~orderId cannot be empty" `
}

type ItemResponse struct {
	ID int64 `json:"id"`
	ItemRequest
	CreatedAt null.String `json:"createdAt"`
	UpdatedAt null.String `json:"updatedAt"`
}
