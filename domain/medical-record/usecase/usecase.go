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

func (u *usecase) SavePatient(req request.SavePatient) (data interface{}, err error) {
	data, err = u.repository.SavePatient(req)
	if err != nil {
		return
	}

	return
}

func (u *usecase) GetPatients(req request.GetPatientsParam) (data interface{}, err error) {
	patient, err := u.repository.GetPatients(req)

	if err != nil {
		err = fmt.Errorf("error get patients record: %v", err)
		return
	}

	response := []dto.PatientResponse{}

	for _, data := range patient {
		response = append(response, dto.PatientResponse{
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
