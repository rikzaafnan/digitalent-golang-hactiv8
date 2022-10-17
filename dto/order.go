package dto

type OrderRequest struct {
	CustomerName string             `json:"customerName" valid:"required~customer name cannot be empty" `
	OrderedAt    string             `json:"orderedAt" valid:"required~ordered at cannot be empty" `
	OrderItems   []OrderItemRequest `json:"items"`
}

type OrderItemRequest struct {
	ItemCode    string `json:"itemCode" valid:"required~itemCode cannot be empty" `
	Description string `json:"description" valid:"required~description cannot be empty" `
	Quantity    int    `json:"quantity" valid:"required~quantity cannot be empty" `
}

type OrderResponse struct {
	ID           int64               `json:"id"`
	CustomerName string              `json:"customerName" valid:"required~customer name cannot be empty" `
	OrderedAt    string              `json:"orderedAt" valid:"required~ordered at cannot be empty" `
	OrderItems   []OrderItemResponse `json:"items"`
}

type OrderItemResponse struct {
	ID int64 `json:"id"`
	OrderItemRequest
}

type OrderRequestUpdate struct {
	CustomerName string              `json:"customerName" valid:"required~customer name cannot be empty" `
	OrderedAt    string              `json:"orderedAt" valid:"required~ordered at cannot be empty" `
	OrderItems   []OrderItemResponse `json:"items"`
}

type OrderItemUpdate struct {
	OrderItemResponse
}
