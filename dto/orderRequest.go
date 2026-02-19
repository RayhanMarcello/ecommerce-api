package dto

type CreateOrderRequest struct {
	CustomerID uint              `json:"customer_id" binding:"required"`
	Items      []CreateOrderItem `json:"items" binding:"required,min=1"`
}

type CreateOrderItem struct {
	ProductID uint `json:"product_id" binding:"required"`
	Qty       uint `json:"qty" binding:"required,min=1"`
}
