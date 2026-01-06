package order

type Status string

const (
	StatusPending   Status = "pending"
	StatusPaid      Status = "paid"
	StatusShipped   Status = "shipped"
	StatusCancelled Status = "cancelled"
)
