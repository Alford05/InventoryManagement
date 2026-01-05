package product

type ListProductQuery struct {
	Category string  `form:"category"`
	MinPrice float64 `form:"min_price" validate:"omitempty,gte=0"`
	MaxPrice float64 `form:"max_price" validate:"omitempty,gte=0"`
	Search   string  `form:"search"`
	Page     int     `form:"page" validate:"omitempty,gte=1"`
	PageSize int     `form:"page_size" validate:"omitempty,gte=1,lte=100"`
}
