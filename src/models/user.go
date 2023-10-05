package models

import (
	"api/src/config"
	"errors"
	"github.com/badoux/checkmail"
	"strings"
	"time"
)

type User struct {
	ID        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Nick      string    `json:"nick,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
}

func (user *User) Prepare(stepUpdate string) error {
	if err := user.validate(stepUpdate); err != nil {
		return err
	}
	if err := user.format(stepUpdate); err != nil {
		return err
	}
	return nil
}

func (user *User) validate(stepUpdate string) error {
	if user.Name == "" {
		return errors.New("The field Name is required and can't be blank")
	}

	if user.Nick == "" {
		return errors.New("The field Nick is required and can't be blank")
	}

	if user.Email == "" {
		return errors.New("The field Email is required and can't be blank")
	}

	if err := checkmail.ValidateFormat(user.Email); err != nil {
		return errors.New("Email not valid")
	}

	if stepUpdate == "register" && user.Password == "" {
		return errors.New("The field Password is required and can't be blank")
	}
	return nil
}

func (user *User) format(stepUpdate string) error {
	user.Name = strings.TrimSpace(user.Name)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Email = strings.TrimSpace(user.Email)

	if stepUpdate == "register" {
		passHash, err := config.Hash(user.Password)
		if err != nil {
			return err
		}
		user.Password = string(passHash)
	}
	return nil
}
