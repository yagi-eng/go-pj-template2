package repository

import (
	"gorm.io/gorm"
)

func Tables() []any {
	return []any{
		&user{},
	}
}

type user struct {
	gorm.Model
}
