package models

import (
	"time"
)

type Order struct {
	OrderID      uint   `gorm:"primaryKey;column:id"`
	CustomerName string `gorm:"not null;type:varchar"`
	Items        []Item
	OrderedAt    time.Time
}
