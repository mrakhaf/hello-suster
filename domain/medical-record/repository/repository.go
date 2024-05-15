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

func (repo *repoHandler) GetMedicalRecord(req request.GetMedicalRecordParam) (data []entity.MedicalRecord, err error) {
	query := fmt.Sprintf("SELECT * FROM medical_record WHERE 1 = 1")

	if req.IdentityNumber != nil {
		query += fmt.Sprintf(" AND identitynumber = %d", *req.IdentityNumber)
	}

	if req.Name != nil {
		query += fmt.Sprintf(" AND LOWER(name) LIKE '%%%s%%'", *req.Name)
	}

	if req.PhoneNumber != nil {
		query += fmt.Sprintf(" AND phonenumber LIKE '%%%s%%'", *req.PhoneNumber)
	}

	if req.CreatedAt != nil {
		if *req.CreatedAt == "asc" {
			query += " ORDER BY created_at ASC"
		} else if *req.CreatedAt == "desc" {
			query += " ORDER BY created_at DESC"
		}
	}

	if req.Limit != nil {
		if *req.Limit > 5 {
			query += fmt.Sprintf(" LIMIT %d", *req.Limit)
		} else {
			query += fmt.Sprintf(" LIMIT 5")
		}
	} else {
		query += fmt.Sprintf(" LIMIT 5")
	}

	if req.Offset != nil {
		query += fmt.Sprintf(" OFFSET %d", *req.Offset)
	} else {
		query += fmt.Sprintf(" OFFSET 0")
	}

	fmt.Println(query)
	rows, err := repo.databaseDB.Query(query)

	if err != nil {
		err = errors.New("Get medical record failed")
		return
	}

	defer rows.Close()

	medical_record := entity.MedicalRecord{}

	for rows.Next() {
		err = rows.Scan(&medical_record.ID, &medical_record.IdentityNumber, &medical_record.Name, &medical_record.Gender, &medical_record.BirthDate, &medical_record.PhoneNumber, &medical_record.IdentityCardScanImg, &medical_record.CreatedAt)
		data = append(data, medical_record)
	}

	return
}
