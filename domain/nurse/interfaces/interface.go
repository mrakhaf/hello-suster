package interfaces

import (
	"github.com/mrakhaf/halo-suster/models/entity"
	"github.com/mrakhaf/halo-suster/models/request"
)

type Repository interface {
	SaveUserNurse(req request.RegisterNurse) (data entity.UserNurse, err error)
	AddAccessNurse(password string, nurseIdInt int) (err error)
	GetNurseByNIP(nip int) (data entity.UserNurse, err error)
}

type Usecase interface {
	Register(req request.RegisterNurse) (data interface{}, err error)
	GiveAccessNurse(password string, nurseIdInt int) (err error)
	LoginNurse(req request.Login) (data interface{}, err error)
}
