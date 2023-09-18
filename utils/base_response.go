package utils

import "golang.org/x/crypto/bcrypt"

type BaseResponse struct {
	Status 	int 	   `json:"status"`
	Message string   `json:"message"`
	Data interface{} `json:"data"`
}

func EncryptPassword(param []byte) string {
	hash, err := bcrypt.GenerateFromPassword(param, bcrypt.DefaultCost)

	if err != nil {
		panic(err)
	}

	return string(hash)
}
