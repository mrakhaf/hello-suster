package usecase

import (
	"github.com/mrakhaf/halo-suster/domain/medical-record/interfaces"
	"github.com/mrakhaf/halo-suster/models/request"
	"github.com/mrakhaf/halo-suster/shared/common/jwt"
)

type usecase struct {
	repository interfaces.Repository
	jwtAccess  *jwt.JWT
}

func NewUsecase(repository interfaces.Repository, jwtAccess *jwt.JWT) interfaces.Usecase {
	return &usecase{
		repository: repository,
		jwtAccess:  jwtAccess,
	}
}

func (u *usecase) SaveMedicalRecord(req request.SaveMedicalRecord) (data interface{}, err error) {
	data, err = u.repository.SaveMedicalRecord(req)
	if err != nil {
		return
	}

	return
}
