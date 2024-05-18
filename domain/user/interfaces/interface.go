package interfaces

import (
	"github.com/mrakhaf/halo-suster/models/entity"
	"github.com/mrakhaf/halo-suster/models/request"
)

type Repository interface {
	SaveUserNurse(req request.RegisterNurse) (data entity.UserNurse, err error)
	AddAccessNurse(password string, nurseId string) (err error)
	GetNurseByNIP(nip int) (data entity.UserNurse, err error)
	GetDataUsers(req request.GetUsers) (data []entity.Users, err error)
	UpdateNurse(req request.EditNurse, nurseId string) (err error)
	DeleteNurse(nurseId string) (err error)
	GetUserByID(id string) (data entity.Users, err error)
}

type Usecase interface {
	Register(req request.RegisterNurse) (data interface{}, err error)
	GiveAccessNurse(password string, nurseId string) (err error)
	LoginNurse(req request.Login) (data interface{}, err error)
	GetUsers(req request.GetUsers) (data interface{}, err error)
	UpdateNurse(req request.EditNurse, nurseId string) (err error)
	DeleteNurse(nip string) (err error)
}
