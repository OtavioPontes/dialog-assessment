package models

import (
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
	"github.com/google/uuid"
	"github.com/otaviopontes/api-go/src/security"
)

type User struct {
	Id        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Nick      string    `json:"nick"`
	Email     string    `json:"email"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

func (user *User) Prepare(isRegister bool) error {

	if err := user.validate(isRegister); err != nil {
		return err
	}
	user.format(isRegister)
	return nil
}

func (user *User) validate(isRegister bool) error {
	if user.Name == "" {
		return errors.New("the name is mandatory and cannot be left blank")
	}
	if user.Nick == "" {
		return errors.New("the nick is mandatory and cannot be left blank")
	}
	if user.Email == "" {
		return errors.New("the email is mandatory and cannot be left blank")
	}
	if err := checkmail.ValidateFormat(user.Email); err != nil {
		return errors.New("the email format is invalid")
	}

	if isRegister && user.Password == "" {
		return errors.New("the password is mandatory and cannot be left blank")
	}
	return nil
}

func (user *User) format(isRegister bool) error {
	user.Name = strings.TrimSpace(user.Name)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Email = strings.TrimSpace(user.Email)

	if isRegister {
		passwordHash, err := security.Hash(user.Password)
		if err != nil {
			return err
		}
		user.Password = string(passwordHash)
	}
	return nil
}
