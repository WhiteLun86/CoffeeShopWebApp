package models

type Order struct {
	ID         uint           `json:"id" gorm:"primaryKey"`
	UserID     uint           `json:"user_id"`
	Products   []OrderProduct `json:"products" gorm:"-"` // 使用自定义结构体
	TotalPrice float64        `json:"total_price"`
	Status     string         `json:"status"`
}

type OrderProduct struct {
	OrderID   uint          `gorm:"column:order_id"`
	ProductID uint          `gorm:"column:product_id"`
	Quantity  int           `json:"quantity"`
	Price     float64       `json:"price"`   // 添加价格字段
	Product   CoffeeProduct `json:"product"` // 包含产品的详细信息
}
