package email

import (
	"github.com/asaskevich/govalidator"
	"strings"
)

type Email string

func From(emailStr string) (Email, error) {
	trimmedEmailStr := strings.TrimSpace(emailStr)
	trimmedEmailStr = strings.ToLower(trimmedEmailStr)

	isValidEmail := govalidator.IsEmail(trimmedEmailStr)
	if !isValidEmail {
		return "", ErrWrongEmail
	}

	return Email(trimmedEmailStr), nil
}

func (e Email) String() string {
	return string(e)
}
