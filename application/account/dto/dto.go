// File:		dto.go
// Created by:	Hoven
// Created on:	2024-08-08
//
// This file is part of the Example Project.
//
// (c) 2024 Example Corp. All rights reserved.

package dto

import "github.com/superwhys/go-ddd-demo/domain/account/model"

type Account struct {
	Account string       `json:"account"`
	Email   string       `json:"email"`
	Phone   string       `json:"phone"`
	Gender  model.Gender `json:"gender"`
}

func AccountEntityToDto(acc *model.Account) *Account {
	return &Account{
		Account: acc.Account,
		Email:   acc.MainEmail.Address,
		Phone:   acc.GetUser().GetPhone(),
		Gender:  acc.GetUser().GetGender(),
	}
}

// AccountDtoToEntity
func AccountDtoToEntity(a *Account) *model.Account {
	return &model.Account{
		Account:   a.Account,
		MainEmail: model.NewEmail(a.Email),
	}
}

type LoginRequest struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

type RegisterRequest struct {
	Account  string `json:"account"`
	Password string `json:"password"`
	Email    string `json:"email"`
}
