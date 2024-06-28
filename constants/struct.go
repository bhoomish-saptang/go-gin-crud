package constants

import "time"

type User struct {
	ID    string `json:"id" validate:"required,min=1,max=1000"`
	Name  string `json:"name" binding:"required"`
	Place string `json:"place" binding:"required"`
	Age   int64  `json:"age" binding:"required"`
}

type AuthInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type AuthUser struct {
	Username  string    `json:"username" validate:"required,min=2,max=100"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"createdAt"`
}
