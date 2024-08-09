// File:		user_repo.go
// Created by:	Hoven
// Created on:	2024-08-08
//
// This file is part of the Example Project.
//
// (c) 2024 Example Corp. All rights reserved.

package account

import (
	"github.com/superwhys/go-ddd-demo/component/mysql"
	"github.com/superwhys/go-ddd-demo/domain/account/model"
	"gorm.io/gorm"
)

type UserRepoImpl struct {
	db mysql.BaseRepository[*User]
}

func NewUserRepo(db *gorm.DB) *UserRepoImpl {
	mysqlDb := mysql.NewMysqlRepository[*User](db)
	return &UserRepoImpl{db: mysqlDb}
}

func (u *UserRepoImpl) FindById(id int) (*model.User, error) {
	user, err := u.db.FindByID(id)
	if err != nil {
		return nil, err
	}

	return POToEntityUser(user), nil
}

func (u *UserRepoImpl) FindByUserName(name string) ([]*model.User, error) {
	users, err := u.db.FindAll("name = ?", name)
	if err != nil {
		return nil, err
	}

	var usersEntries []*model.User
	for _, user := range users {
		usersEntries = append(usersEntries, POToEntityUser(user))
	}

	return usersEntries, nil
}

func (u *UserRepoImpl) FindByIdCardAndUserName(userName string, idCard string) (*model.User, error) {
	user, err := u.db.First("id_card = ? and user_name = ?", idCard, userName)
	if err != nil {
		return nil, err
	}

	return POToEntityUser(user), nil
}

func (u *UserRepoImpl) SaveUser(user *model.User) (*model.User, error) {
	userPo := EntityToPOUser(user)
	err := u.db.Create(userPo)
	if err != nil {
		return nil, err
	}

	user.Id = userPo.ID
	return user, nil
}

func (u *UserRepoImpl) UpdateUser(user *model.User) error {
	userPo := EntityToPOUser(user)
	return u.db.Update(userPo)
}

func (u *UserRepoImpl) DeleteUser(user *model.User) error {
	userPo := EntityToPOUser(user)

	return u.db.Delete(userPo)
}
