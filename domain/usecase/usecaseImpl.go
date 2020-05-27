package usecase

import (
	"belajar-gcp/domain/model"
	"belajar-gcp/domain/repository"
)

type implUseCase struct {
	repo repository.Repository
}

func NewUseCase(repo *repository.Repository) UseCase {
	return &implUseCase{
		repo: *repo,
	}
}

func (impl *implUseCase) FindAllUser() ([]model.User, error) {
	return impl.repo.FindAllUser()
}
