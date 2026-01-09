package order

type CreateOrderRequest struct {
	CustomerID uint                     `json:"customer_id" validate:"required,gt=0"`
	Items      []CreateOrderItemRequest `json:"items" validate:"required,min=1,dive"`
}

type CreateOrderItemRequest struct {
	ProductID uint `json:"product_id" validate:"required,gt=0"`
	Quantity  int  `json:"quantity" validate:"required,gt=0"`
}
