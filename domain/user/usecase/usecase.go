package usecase

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"log"
	"mime/multipart"
	"strconv"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
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

	dataUsers, err := u.repository.GetDataUsers(req)

	if err != nil {
		return
	}

	if len(dataUsers) == 0 {
		data = []interface{}{}
		return
	}

	data = dataUsers

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

func (u *usecase) UploadImage(file multipart.File, fileHeader *multipart.FileHeader) error {

	buf := bytes.NewBuffer(nil)
	if _, err := buf.ReadFrom(file); err != nil {
		return fmt.Errorf("failed to read file: %v", err)
	}

	fileType := fileHeader.Header.Get("Content-Type")

	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatalf("failed to load configuration, %v", err)
	}

	svc := s3.NewFromConfig(cfg)

	_, err = svc.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket:        aws.String("projectsprint-bucket-public-read"),
		Key:           aws.String(fileHeader.Filename),
		ACL:           "public-read",
		Body:          bytes.NewReader(buf.Bytes()),
		ContentLength: aws.Int64(fileHeader.Size),
		ContentType:   aws.String(fileType),
	})

	if err != nil {
		log.Fatalf("failed to put object, %v", err)
	}

	return nil
}
