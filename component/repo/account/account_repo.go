// File:		account_repo.go
// Created by:	Hoven
// Created on:	2024-08-08
//
// This file is part of the Example Project.
//
// (c) 2024 Example Corp. All rights reserved.

package account

import (
	"errors"

	"github.com/superwhys/go-ddd-demo/component/mysql"
	"github.com/superwhys/go-ddd-demo/domain/account/model"
	"gorm.io/gorm"
)

type AccountRepoImpl struct {
	db mysql.BaseRepository[*Account]
}

func NewAccountRepo(db *gorm.DB) *AccountRepoImpl {
	mysqlDb := mysql.NewMysqlRepository[*Account](db)
	return &AccountRepoImpl{db: mysqlDb}
}

func (a *AccountRepoImpl) FindByAccount(account string) (*model.Account, error) {
	acc, err := a.db.First("account = ?", account)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	} else if err != nil {
		return nil, ErrAccountNotFound
	}

	return POToEntityAccount(acc), nil
}

func (a *AccountRepoImpl) SaveAccount(account *model.Account) (*model.Account, error) {
	accPo := EntityToPOAccount(account)

	err := a.db.Create(accPo)
	if err != nil {
		return nil, err
	}

	return account, nil

}

func (a *AccountRepoImpl) DeleteAccount(account *model.Account) error {
	accPo := EntityToPOAccount(account)

	return a.db.Delete(accPo)
}

func (a *AccountRepoImpl) UpdateAccount(account *model.Account) error {
	accPo := EntityToPOAccount(account)

	return a.db.Update(accPo)
}
