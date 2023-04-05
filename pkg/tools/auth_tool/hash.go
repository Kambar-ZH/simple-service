package auth_tool

import "golang.org/x/crypto/bcrypt"

type Password string

func (p Password) Hash() (result string, err error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(p), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) (result bool) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
