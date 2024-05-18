package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/mrakhaf/halo-suster/domain/user/interfaces"
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

	query := fmt.Sprintf("INSERT INTO users (id, nip, name, identityscanimage, created_at) VALUES ('%s', %d, '%s', '%s', '%s')", data.ID, data.NIP, data.Name, data.IdentityCardScanImg, data.CreatedAt)

	_, err = repo.databaseDB.Exec(query)

	if err != nil {

		if err.Error() == "pq: duplicate key value violates unique constraint \"users_nip_key\"" {
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

	queryUpdate := fmt.Sprintf("UPDATE users SET password = '%s' WHERE nip = %d", password, nurseIdInt)

	_, err = repo.databaseDB.Exec(queryUpdate)

	if err != nil {
		return
	}

	return
}

func (repo *repoHandler) GetNurseByNIP(nip int) (data entity.UserNurse, err error) {
	query := fmt.Sprintf("SELECT id, name, nip, password, identityscanimage, created_at FROM users WHERE nip = %d", nip)

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

func (repo *repoHandler) GetDataUsers(req request.GetUsers) (data []entity.Users, err error) {

	query := fmt.Sprintf("SELECT id, name, nip, created_at FROM users WHERE 1 = 1")

	if req.UserId != nil {
		query += fmt.Sprintf(" AND id = '%s'", *req.UserId)
	}

	if req.Nip != nil {
		query += fmt.Sprintf(" AND CAST(nip as text) like '%%%s%%'", *req.Nip)
	}

	if req.Name != nil {
		query += fmt.Sprintf(" AND name like '%s%%'", *req.Name)
	}

	if req.Role != nil {
		if *req.Role == "it" {
			//query first 3 digit NIP is 615
			query += fmt.Sprintf(" AND cast(nip as text) like '615%%'")
		} else if *req.Role == "nurse" {
			//query first 3 digit NIP is 303
			query += fmt.Sprintf(" AND cast(nip as text) like '303%%'")
		}
	}

	if req.CreatedAt != nil {
		if *req.CreatedAt == "asc" {
			query += fmt.Sprintf(" ORDER BY created_at ASC")
		} else if *req.CreatedAt == "desc" {
			query += fmt.Sprintf(" ORDER BY created_at DESC")
		}
	}

	if req.Limit != nil {
		query += fmt.Sprintf(" LIMIT %d", *req.Limit)
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
		fmt.Println(err.Error())
		err = errors.New("Get data failed")
		return
	}

	defer rows.Close()

	for rows.Next() {
		var dataRow entity.Users
		err = rows.Scan(&dataRow.ID, &dataRow.Name, &dataRow.NIP, &dataRow.CreatedAt)
		//convert iso 8601 to yyyy-mm-dd
		dataRow.CreatedAt = dataRow.CreatedAt[0:10]
		if err != nil {
			fmt.Println(err.Error())
			err = errors.New("Get data failed")
			return
		}
		data = append(data, dataRow)
	}

	return
}

func (repo *repoHandler) UpdateNurse(req request.EditNurse, nurseId string) (err error) {

	query := fmt.Sprintf("UPDATE users SET name = '%s', nip = %d WHERE id = '%s'", req.Name, req.NIP, nurseId)

	fmt.Println(query)

	sqlResult, err := repo.databaseDB.Exec(query)

	if err != nil {
		fmt.Println(err.Error())
		if err.Error() == "pq: duplicate key value violates unique constraint \"users_nip_key\"" {
			err = errors.New("NIP already exist")
			return
		}
		return
	}

	rows, err := sqlResult.RowsAffected()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if rows == 0 {
		err = errors.New("UserId not found")
		return
	}

	return
}

func (repo *repoHandler) DeleteNurse(nurseId string) (err error) {

	query := fmt.Sprintf("delete from users where id = '%s' and cast(nip as text) like '303%%'", nurseId)

	sqlResult, err := repo.databaseDB.Exec(query)

	if err != nil {
		return
	}

	rows, err := sqlResult.RowsAffected()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if rows == 0 {
		err = errors.New("not delete anything")
		return
	}

	return

}

func (repo *repoHandler) GetUserByID(id string) (data entity.Users, err error) {

	query := fmt.Sprintf("SELECT id, name, nip, created_at FROM users WHERE id = '%s'", id)

	fmt.Println(query)

	err = repo.databaseDB.QueryRow(query).Scan(&data.ID, &data.Name, &data.NIP, &data.CreatedAt)

	if err != nil {

		if err == sql.ErrNoRows {
			err = errors.New("UserId not found")
			return
		}

		fmt.Println(err.Error())
		err = errors.New("Get data failed")
		return
	}

	return
}
