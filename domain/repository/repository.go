package repository

import "belajar-gcp/domain/model"

type Repository interface {
	FindAllUser() ([]model.User, error)
}
