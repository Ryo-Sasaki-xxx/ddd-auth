package domains

import (
	"errors"
	"unicode/utf8"
)

type UserName struct {
	name string
}

func NewUserName(name string) (*UserName, error) {
	if len := utf8.RuneCountInString(name); 6 > len || len > 20 {
		return nil, errors.New("validation error")
	}

	userName := &UserName{}
	userName.name = name

	return userName, nil
}
