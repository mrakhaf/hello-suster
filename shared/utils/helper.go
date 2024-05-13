package utils

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func GenerateUUID() string {
	id := uuid.Must(uuid.NewRandom())
	return id.String()
}

func HashPassword(password string) (result string, err error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return
	}

	result = string(bytes)
	return
}

func CheckPasswordHash(password, hash string) error {
	fmt.Println(password)
	fmt.Println(hash)
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	if err != nil {
		return err
	}

	return nil
}

func Contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func CheckImageType(url string) bool {
	url = url[strings.LastIndex(url, ".")+1:]
	if url == "png" || url == "jpg" || url == "jpeg" {
		return true
	}
	return false
}

func IsValidPhoneNumber(phoneNumber string) error {
	phoneNumberRegex := `^\+[0-9]{1,4}-?[0-9]{1,15}$`
	match, _ := regexp.MatchString(phoneNumberRegex, phoneNumber)
	if match {
		return nil
	}

	return errors.New("format Phone tidak valid")
}

func ValidateNIP(nip int, role string) error {

	nipString := strconv.Itoa(nip)

	roleCode := nipString[:3]
	if role == "IT" && roleCode != "615" {
		return errors.New("format NIP IT tidak valid")
	} else if role == "NURSE" && roleCode != "303" {
		return errors.New("format NIP Suster tidak valid")
	}

	gender := nipString[3:4]
	if !Contains([]string{"1", "2"}, gender) {
		return errors.New("format gender tidak valid")
	}

	year := nipString[4:8]
	yearInt, _ := strconv.Atoi(year)

	if !(yearInt >= 2000 && yearInt <= 2024) {
		return errors.New("format tahun tidak valid")
	}

	month := nipString[8:10]

	if !Contains([]string{"01", "02", "03", "04", "05", "06", "07", "08", "09", "10", "11", "12"}, month) {
		return errors.New("format bulan tidak valid")
	}

	return nil

}
