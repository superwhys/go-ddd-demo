// File:		entity.go
// Created by:	Hoven
// Created on:	2024-08-08
//
// This file is part of the Example Project.
//
// (c) 2024 Example Corp. All rights reserved.

package model

import (
	"errors"

	"github.com/superwhys/go-ddd-demo/pkg/password"
)

type Account struct {
	Account   string `json:"account"`
	Password  string `json:"password"`
	User      *User  `json:"user"`
	MainEmail *Email `json:"main_email"`
}

func (a *Account) GetUser() *User {
	return a.User
}

func (a *Account) Validate() error {
	if a.Account == "" {
		return errors.New("missing account")
	}

	if a.Password == "" {
		return errors.New("missing password")
	}

	if !a.MainEmail.IsValidEmail() {
		return errors.New("invalid email")
	}

	return nil
}

func (a *Account) CheckPassword(pwd string) bool {
	return password.CheckPasswordHash(pwd, a.Password)
}

func (a *Account) SetPassword(pwd string) error {
	p, err := password.HashPassword(pwd)
	if err != nil {
		return err
	}

	a.Password = p
	return nil
}

func (a *Account) Equal(newAcc *Account) bool {
	if a.Account != newAcc.Account && !a.MainEmail.Equal(newAcc.MainEmail) {
		return false
	}

	if newAcc.Password != "" {
		return password.CheckPasswordHash(newAcc.Password, a.Password)
	}
	return true
}

type User struct {
	Id       int    `json:"user_id"`
	IdCard   string `json:"id_card"`
	UserName string `json:"user_name"`
	Phone    string `json:"phone"`
	Gender   Gender `json:"gender"`

	account string
}

func (a *User) GetId() int {
	if a == nil {
		return 0
	}
	return a.Id
}

func (a *User) GetIdCard() string {
	if a == nil {
		return ""
	}

	return a.IdCard
}

func (a *User) GetUserName() string {
	if a == nil {
		return ""
	}

	return a.UserName
}

func (a *User) GetPhone() string {
	if a == nil {
		return ""
	}

	return a.Phone
}

func (a *User) GetGender() Gender {
	if a == nil {
		return GenderUnknown
	}

	return a.Gender
}

func (a *User) GetAccount() string {
	return a.account
}

func (a *User) SetAccount(account string) {
	a.account = account
}

func (a *User) Validate() error {
	if a.UserName == "" {
		return errors.New("missing user name")
	}

	if a.Phone == "" {
		return errors.New("missing phone")
	}

	if a.Gender == 0 {
		return errors.New("missing gender")
	}

	return nil
}

func (a *User) Equal(b *User) bool {
	if a.Phone != b.Phone || a.Gender != b.Gender {
		return false
	}

	return true
}
