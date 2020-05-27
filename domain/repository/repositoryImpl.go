package repository

import (
	"belajar-gcp/domain/model"

	"github.com/jinzhu/gorm"
)

type implRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &implRepository{
		db: db,
	}
}

func (impl *implRepository) FindAllUser() ([]model.User, error) {
	users := make([]model.User, 0)
	// q := `select * from "user";`
	// if err := impl.db.Raw(q).Scan(&users).Error; err != nil {
	// 	return users, err
	// }
	if err := impl.db.Table("user").Scan(&users).Error; err != nil {
		return users, err
	}
	return users, nil
}
