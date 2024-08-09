// File:		account.go
// Created by:	Hoven
// Created on:	2024-08-08
//
// This file is part of the Example Project.
//
// (c) 2024 Example Corp. All rights reserved.

package account

import (
	"github.com/superwhys/go-ddd-demo/application/account/dto"
	"github.com/superwhys/go-ddd-demo/domain/account/model"
	"github.com/superwhys/go-ddd-demo/domain/account/service"
)

type AccountApp struct {
	accountService service.AccountService
}

func NewAccountApp(accSrv service.AccountService) *AccountApp {
	return &AccountApp{
		accountService: accSrv,
	}
}

func (a *AccountApp) Login(info *dto.LoginRequest) (*dto.Account, error) {
	account, err := a.accountService.Login(info.Account, info.Password)
	if err != nil {
		return nil, err
	}

	return dto.AccountEntityToDto(account), nil
}

func (a *AccountApp) Register(info *dto.RegisterRequest) (*dto.Account, error) {
	acc := &model.Account{
		Account:   info.Account,
		Password:  info.Password,
		MainEmail: model.NewEmail(info.Email),
	}

	err := a.accountService.RegisterAccount(acc)
	if err != nil {
		return nil, err
	}

	return dto.AccountEntityToDto(acc), nil
}
