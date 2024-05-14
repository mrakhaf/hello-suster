package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/mrakhaf/halo-suster/domain/medical-record/interfaces"
	"github.com/mrakhaf/halo-suster/models/entity"
	"github.com/mrakhaf/halo-suster/models/request"
	"github.com/mrakhaf/halo-suster/shared/utils"
)

type repoHandler struct {
	databaseDB *sql.DB
}

func NewRepository(databaseDB *sql.DB) interfaces.Repository {
	return &repoHandler{
		databaseDB: databaseDB,
	}
}

func (repo *repoHandler) SaveMedicalRecord(req request.SaveMedicalRecord) (data entity.MedicalRecord, err error) {

	data = entity.MedicalRecord{
		ID:                  utils.GenerateUUID(),
		IdentityNumber:      req.IdentityNumber,
		PhoneNumber:         req.PhoneNumber,
		Name:                req.Name,
		BirthDate:           req.BirthDate,
		Gender:              req.Gender,
		IdentityCardScanImg: req.IdentityCardScanImage,
		CreatedAt:           time.Now().Format("2006-01-02 15:04:05"),
	}

	query := fmt.Sprintf("INSERT INTO medical_record (id, identitynumber, name, birthdate, phonenumber, gender, identityscanimage, created_at) VALUES ('%s', %d, '%s', '%s', '%s', '%s', '%s', '%s')", data.ID, data.IdentityNumber, data.Name, data.BirthDate, data.PhoneNumber, data.Gender, data.IdentityCardScanImg, data.CreatedAt)

	_, err = repo.databaseDB.Exec(query)

	if err != nil {
		if err.Error() == "pq: duplicate key value violates unique constraint \"medical_record_identitynumber_key\"" {
			err = errors.New("Identity number already exist")
			return
		}

		err = errors.New("Save medical record failed")
		return
	}

	return
}
