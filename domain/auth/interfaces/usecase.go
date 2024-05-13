package interfaces

import "github.com/mrakhaf/halo-suster/models/request"

type Usecase interface {
	Register(req request.Register) (data interface{}, err error)
	Login(req request.Login) (data interface{}, err error)
}
