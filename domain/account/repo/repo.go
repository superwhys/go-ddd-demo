// File:		user.go
// Created by:	Hoven
// Created on:	2024-08-08
//
// This file is part of the Example Project.
//
// (c) 2024 Example Corp. All rights reserved.

package repo

import (
	"github.com/pkg/errors"
	"github.com/superwhys/go-ddd-demo/domain/account/model"
)

var (
	ErrAccountNotExists = errors.New("account does not exist")
	ErrUSerNotExists    = errors.New("user does not exist")
)

type AccountRepository interface {
	FindByAccount(account string) (*model.Account, error)
	SaveAccount(account *model.Account) (*model.Account, error)
	DeleteAccount(account *model.Account) error
	UpdateAccount(account *model.Account) error
}

type UserRepository interface {
	FindById(id int) (*model.User, error)
	FindByUserName(name string) ([]*model.User, error)
	FindByIdCardAndUserName(userName, idCard string) (*model.User, error)

	SaveUser(user *model.User) (*model.User, error)
	UpdateUser(user *model.User) error
	DeleteUser(user *model.User) error
}
