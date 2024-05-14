package usecase

import (
	"errors"
	"strconv"

	"github.com/mrakhaf/halo-suster/domain/auth/interfaces"
	"github.com/mrakhaf/halo-suster/models/request"
	"github.com/mrakhaf/halo-suster/shared/common/jwt"
	"github.com/mrakhaf/halo-suster/shared/utils"
)

type (
	usecase struct {
		repository interfaces.Repository
		JwtAccess  *jwt.JWT
	}
)

func NewUsecase(repository interfaces.Repository, JwtAccess *jwt.JWT) interfaces.Usecase {
	return &usecase{
		repository: repository,
		JwtAccess:  JwtAccess,
	}
}

func (u *usecase) Register(req request.Register) (data interface{}, err error) {

	userIt, err := u.repository.SaveUserIt(req)

	if err != nil {
		return
	}

	accessToken, err := u.JwtAccess.GenerateToken(userIt.ID)

	if err != nil {
		err = errors.New("Generate access token failed")
		return
	}

	data = map[string]interface{}{
		"userId":      userIt.ID,
		"nip":         userIt.NIP,
		"name":        userIt.Name,
		"accessToken": accessToken,
	}

	return
}

func (u *usecase) Login(req request.Login) (data interface{}, err error) {
	//get data user it
	userIt, err := u.repository.GetDataUserIt(req.NIP)

	if err != nil {
		return
	}

	//check password
	err = utils.CheckPasswordHash(req.Password, userIt.Password)

	if err != nil {
		err = errors.New("Wrong password")
		return
	}

	NipString := strconv.Itoa(userIt.NIP)

	//generate access token
	accessToken, err := u.JwtAccess.GenerateToken(NipString)

	if err != nil {
		err = errors.New("Generate access token failed")
		return
	}

	data = map[string]interface{}{
		"userId":      userIt.ID,
		"nip":         userIt.NIP,
		"name":        userIt.Name,
		"accessToken": accessToken,
	}

	return
}
