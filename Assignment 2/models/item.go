package models

type Item struct {
	ItemID      uint   `gorm:"primaryKey"`
	ItemCode    string `gorm:"not null;type:varchar"`
	Description string `gorm:"not null;type:varchar"`
	Quantity    int    `gorm:"not null;type:int"`
	OrderID     uint
}
