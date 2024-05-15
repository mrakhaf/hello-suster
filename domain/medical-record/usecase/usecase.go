package usecase

import (
	"fmt"

	"github.com/mrakhaf/halo-suster/domain/medical-record/interfaces"
	"github.com/mrakhaf/halo-suster/models/dto"
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

func (u *usecase) GetMedicalRecord(req request.GetMedicalRecordParam) (data interface{}, err error) {
	medicalRecord, err := u.repository.GetMedicalRecord(req)

	if err != nil {
		err = fmt.Errorf("error get medical record: %v", err)
		return
	}

	response := []dto.MedicalRecordResponse{}

	for _, data := range medicalRecord {
		response = append(response, dto.MedicalRecordResponse{
			IdentityNumber: data.IdentityNumber,
			PhoneNumber:    data.PhoneNumber,
			Name:           data.Name,
			BirthDate:      data.BirthDate,
			Gender:         data.Gender,
			CreatedAt:      data.CreatedAt,
		})
	}

	data = response

	return
}
