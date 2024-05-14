package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/mrakhaf/halo-suster/domain/nurse/interfaces"
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

func (repo *repoHandler) SaveUserNurse(req request.RegisterNurse) (data entity.UserNurse, err error) {

	data = entity.UserNurse{
		ID:                  utils.GenerateUUID(),
		NIP:                 req.NIP,
		Name:                req.Name,
		IdentityCardScanImg: req.IdentityCardScanImage,
		CreatedAt:           time.Now().Format("2006-01-02 15:04:05"),
	}

	query := fmt.Sprintf("INSERT INTO user_nurse (id, nip, name, identityscanimage, created_at) VALUES ('%s', %d, '%s', '%s', '%s')", data.ID, data.NIP, data.Name, data.IdentityCardScanImg, data.CreatedAt)

	_, err = repo.databaseDB.Exec(query)

	if err != nil {

		if err.Error() == "pq: duplicate key value violates unique constraint \"user_nurse_nip_key\"" {
			err = errors.New("NIP already exist")
			return
		}
		fmt.Println(err.Error())
		err = errors.New("Register failed")
		return
	}

	return
}

func (repo *repoHandler) AddAccessNurse(password string, nurseIdInt int) (err error) {

	password, err = utils.HashPassword(password)

	if err != nil {
		return
	}

	queryUpdate := fmt.Sprintf("UPDATE user_nurse SET password = '%s' WHERE nip = %d", password, nurseIdInt)

	_, err = repo.databaseDB.Exec(queryUpdate)

	if err != nil {
		return
	}

	return
}

func (repo *repoHandler) GetNurseByNIP(nip int) (data entity.UserNurse, err error) {
	query := fmt.Sprintf("SELECT * FROM user_nurse WHERE nip = %d", nip)

	fmt.Println(query)

	err = repo.databaseDB.QueryRow(query).Scan(&data.ID, &data.Name, &data.NIP, &data.Password, &data.IdentityCardScanImg, &data.CreatedAt)

	if err != nil {
		fmt.Println(err.Error())
		if err == sql.ErrNoRows {
			err = errors.New("NIP not found")
			return
		}
		err = errors.New("Get data failed")
		return
	}

	return
}
