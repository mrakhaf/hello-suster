package interfaces

import (
	"github.com/mrakhaf/halo-suster/models/entity"
	"github.com/mrakhaf/halo-suster/models/request"
)

type Repository interface {
	SaveMedicalRecord(req request.SaveMedicalRecord) (data entity.MedicalRecord, err error)
}

type Usecase interface {
	SaveMedicalRecord(req request.SaveMedicalRecord) (data interface{}, err error)
}
