package models

type CoffeeProduct struct {
	ID          uint    `gorm:"primaryKey"`
	Name        string  `gorm:"not null"`
	Price       float64 `gorm:"not null"`
	Description string  `gorm:"not null"`
	Stock       int     `gorm:"not null"` // 新增库存数量字段
}
