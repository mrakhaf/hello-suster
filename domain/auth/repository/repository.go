package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/mrakhaf/halo-suster/domain/auth/interfaces"
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

func (repo *repoHandler) SaveUserIt(req request.Register) (data entity.User, err error) {

	password, err := utils.HashPassword(req.Password)

	if err != nil {
		return
	}

	data = entity.User{
		ID:        utils.GenerateUUID(),
		NIP:       req.NIP,
		Name:      req.Name,
		Password:  password,
		CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
	}

	query := fmt.Sprintf("INSERT INTO user_it (id, nip, name, password, created_at) VALUES ('%s', %d, '%s', '%s', '%s')", data.ID, data.NIP, data.Name, data.Password, data.CreatedAt)

	fmt.Println(query)

	_, err = repo.databaseDB.Exec(query)

	if err != nil {
		// if err duplicate nip
		if err.Error() == "pq: duplicate key value violates unique constraint \"user_it_nip_key\"" {
			err = errors.New("NIP already exist")
			return
		}

		err = errors.New("Register failed")
		return
	}

	return

}

func (repo *repoHandler) GetDataUserIt(nip int) (data entity.User, err error) {

	query := fmt.Sprintf("SELECT * FROM user_it WHERE nip = %d", nip)

	fmt.Println(query)

	err = repo.databaseDB.QueryRow(query).Scan(&data.ID, &data.Name, &data.NIP, &data.Password, &data.CreatedAt)

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
