package repository

import "gorm.io/gorm"

type BaseRepository struct {
	Manager *gorm.DB
}
