package usecase

import "github.com/mrakhaf/halo-suster/domain/nurse/interfaces"

type usecase struct {
	repository interfaces.Repository
}

func NewUsecase(repository interfaces.Repository) interfaces.Usecase {
	return &usecase{
		repository: repository,
	}
}
