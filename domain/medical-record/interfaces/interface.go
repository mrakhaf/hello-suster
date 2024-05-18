package interfaces

import (
	"github.com/mrakhaf/halo-suster/models/dto"
	"github.com/mrakhaf/halo-suster/models/entity"
	"github.com/mrakhaf/halo-suster/models/request"
)

type Repository interface {
	SavePatient(req request.SavePatient) (data entity.Patient, err error)
	GetPatientByIdentity(identitynumber int) (data entity.Patient, err error)
	SaveMedicalRecord(req request.SaveMedicalRecord, nip int) (data entity.MedicalRecord, err error)
	GetPatients(req request.GetPatientsParam) (data []entity.Patient, err error)
	GetMedicalRecords(req request.GetMedicalRecordsParam) (data []dto.MedicalRecordResponse, err error)
}

type Usecase interface {
	SavePatient(req request.SavePatient) (data interface{}, err error)
	SaveMedicalRecord(req request.SaveMedicalRecord, nip int) (data interface{}, err error)
	GetPatients(req request.GetPatientsParam) (data interface{}, err error)
	GetMedicalRecords(req request.GetMedicalRecordsParam) (data interface{}, err error)
}
