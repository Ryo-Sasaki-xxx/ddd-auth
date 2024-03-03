package domains

import (
	"errors"
	"net/mail"
)

type UserEmail struct {
	email string
}

func NewUserEmail(email string) (*UserEmail, error) {
	if _, err := mail.ParseAddress(email); err != nil {
		return nil, errors.New("email validation error")
	}

	userEmail := &UserEmail{}
	userEmail.email = email

	return userEmail, nil
}
