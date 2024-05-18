package usecase

import (
	"errors"
	"strconv"

	"github.com/mrakhaf/halo-suster/domain/user/interfaces"
	"github.com/mrakhaf/halo-suster/models/request"
	"github.com/mrakhaf/halo-suster/shared/common/jwt"
	"github.com/mrakhaf/halo-suster/shared/utils"
)

type usecase struct {
	repository interfaces.Repository
	jwtAccess  *jwt.JWT
}

func NewUsecase(repository interfaces.Repository, jwtAccess *jwt.JWT) interfaces.Usecase {
	return &usecase{
		repository: repository,
		jwtAccess:  jwtAccess,
	}
}

func (u *usecase) Register(req request.RegisterNurse) (data interface{}, err error) {

	dataNurse, err := u.repository.SaveUserNurse(req)

	if err != nil {
		return
	}

	data = map[string]interface{}{
		"userId": dataNurse.ID,
		"nip":    dataNurse.NIP,
		"name":   dataNurse.Name,
	}

	return

}

func (u *usecase) GiveAccessNurse(password string, nurseId string) (err error) {

	// _, err = u.repository.GetUserByID(nurseId)

	// if err != nil {
	// 	return
	// }

	err = u.repository.AddAccessNurse(password, nurseId)

	if err != nil {
		return
	}

	return

}

func (u *usecase) LoginNurse(req request.Login) (data interface{}, err error) {

	dataNurse, err := u.repository.GetNurseByNIP(req.NIP)

	if err != nil {
		return
	}

	if dataNurse.Password == nil {
		err = errors.New("NIP not have access")
		return
	}

	err = utils.CheckPasswordHash(req.Password, *dataNurse.Password)

	if err != nil {
		err = errors.New("Wrong password")
		return
	}

	NipString := strconv.Itoa(dataNurse.NIP)

	accessToken, err := u.jwtAccess.GenerateToken(NipString)

	if err != nil {
		err = errors.New("Generate access token failed")
		return
	}

	data = map[string]interface{}{
		"userId":      dataNurse.ID,
		"nip":         dataNurse.NIP,
		"name":        dataNurse.Name,
		"accessToken": accessToken,
	}

	return

}

func (u *usecase) GetUsers(req request.GetUsers) (data interface{}, err error) {

	data, err = u.repository.GetDataUsers(req)

	if err != nil {
		return
	}

	return

}

func (u *usecase) UpdateNurse(req request.EditNurse, nurseId string) (err error) {

	err = u.repository.UpdateNurse(req, nurseId)

	return

}

func (u *usecase) DeleteNurse(nurseId string) (err error) {

	err = u.repository.DeleteNurse(nurseId)

	return

}
