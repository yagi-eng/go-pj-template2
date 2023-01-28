package repository

import (
	"gorm.io/gorm"
)

func Tables() []any {
	return []any{
		&User{},
	}
}

type User struct {
	gorm.Model
}
