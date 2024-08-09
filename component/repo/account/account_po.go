// File:		account_po.go
// Created by:	Hoven
// Created on:	2024-08-08
//
// This file is part of the Example Project.
//
// (c) 2024 Example Corp. All rights reserved.

package account

import (
	"database/sql"
	"time"

	"github.com/superwhys/go-ddd-demo/domain/account/model"
)

type User struct {
	ID       int    `json:"user_id" gorm:"primarykey"`
	IdCard   string `json:"id_card" gorm:"type:varchar(20);uniqueIndex:idx_card_name"`
	UserName string `json:"user_name" gorm:"type:varchar(64);uniqueIndex:idx_card_name"`
	Phone    string `json:"phone"`
	Gender   uint   `json:"gender"`

	Accounts []*Account `json:"accounts"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime `gorm:"index"`
}

type Account struct {
	Account   string `json:"account" gorm:"primarykey"`
	Password  string `json:"password"`
	MainEmail string `json:"main_email"`
	UserID    *int   `json:"user_id"`
	User      *User

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime `gorm:"index"`
}

// Convert PO to Entity
func POToEntityUser(poUser *User) *model.User {
	if poUser == nil {
		return nil
	}

	accounts := make([]*model.Account, len(poUser.Accounts))
	for i, poAccount := range poUser.Accounts {
		accounts[i] = POToEntityAccount(poAccount)
	}

	return &model.User{
		Id:       int(poUser.ID),
		UserName: poUser.UserName,
		Phone:    poUser.Phone,
		Gender:   model.Gender(poUser.Gender),
	}
}

func POToEntityAccount(poAccount *Account) *model.Account {
	if poAccount == nil {
		return nil
	}

	return &model.Account{
		Account:   poAccount.Account,
		Password:  poAccount.Password,
		MainEmail: model.NewEmail(poAccount.MainEmail),
		User:      POToEntityUser(poAccount.User),
	}
}

// Convert Entity to PO
func EntityToPOUser(user *model.User) *User {
	if user == nil {
		return nil
	}

	return &User{
		ID:       user.Id,
		UserName: user.UserName,
		Phone:    user.Phone,
		Gender:   uint(user.Gender),
	}
}

func EntityToPOAccount(account *model.Account) *Account {
	if account == nil {
		return nil
	}

	return &Account{
		Account:   account.Account,
		Password:  account.Password,
		MainEmail: account.MainEmail.Address,
		User:      EntityToPOUser(account.User),
		UserID: func() *int {
			if account.User == nil {
				return nil
			}
			return &account.User.Id
		}(),
	}
}
