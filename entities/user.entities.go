package entities

import (
	"time"
)

// User Entity
type User struct {
	ID        uint64     `gorm:"column:id; primary_key; AUTO_INCREMENT" json:"id"`
	Username  string     `gorm:"column:username; type:varchar(255); not null; unique" json:"username"`
	Password  string     `gorm:"column:password; type:varchar(255); not null" json:"-"`
	FullName  string     `gorm:"column:fullname; type:varchar(255); not null" json:"fullname"`
	Photo     string     `gorm:"column:photo; type:varchar(255)" json:"photo"`
	CreatedAt time.Time  `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time  `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at" json:"deleted_at"`
}
