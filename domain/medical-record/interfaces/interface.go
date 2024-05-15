package interfaces

import (
	"github.com/mrakhaf/halo-suster/models/entity"
	"github.com/mrakhaf/halo-suster/models/request"
)

type Repository interface {
	SavePatient(req request.SavePatient) (data entity.Patient, err error)
	GetPatients(req request.GetPatientsParam) (data []entity.Patient, err error)
}

type Usecase interface {
	SavePatient(req request.SavePatient) (data interface{}, err error)
	GetPatients(req request.GetPatientsParam) (data interface{}, err error)
}
