package Models

import (
	"time"
)

type User struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique"`
	Password []byte `json:"-"`
}

type List struct {
	UserID    uint      `json:"user_id"`
	ID        uint      `json:"id"`
	Task      string    `json:"task"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
}
