package interfaces

import (
	"github.com/mrakhaf/halo-suster/models/entity"
	"github.com/mrakhaf/halo-suster/models/request"
)

type Repository interface {
	SaveUserIt(req request.Register) (data entity.User, err error)
	GetDataUserIt(nip int) (data entity.User, err error)
}
