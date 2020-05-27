package usecase

import "belajar-gcp/domain/model"

type UseCase interface {
	FindAllUser() ([]model.User, error)
}
