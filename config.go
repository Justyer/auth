package auth

import (
	"gorm.io/gorm"
)

type Config struct {
	DB  *gorm.DB
	Err *Err
}
