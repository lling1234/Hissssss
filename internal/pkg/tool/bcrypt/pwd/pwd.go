package pwd

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
)

func Encrypt(Password string) (string, error) {
	bcryptBytes, err := bcrypt.GenerateFromPassword([]byte(Password), bcrypt.DefaultCost)
	if err != nil {
		return "", errors.New("密码加密错误")
	}
	return string(bcryptBytes), nil
}

func Verify(Password string, bcryptPwd string) bool {
	return bcrypt.CompareHashAndPassword([]byte(bcryptPwd), []byte(Password)) == nil
}
