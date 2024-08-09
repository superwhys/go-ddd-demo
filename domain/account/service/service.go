// File:		service.go
// Created by:	Hoven
// Created on:	2024-08-08
//
// This file is part of the Example Project.
//
// (c) 2024 Example Corp. All rights reserved.

package service

import (
	"github.com/go-puzzles/plog"
	"github.com/pkg/errors"
	"github.com/superwhys/go-ddd-demo/domain/account/model"
	"github.com/superwhys/go-ddd-demo/domain/account/repo"
	"github.com/superwhys/go-ddd-demo/pkg/password"
)

var (
	ErrPasswordNotCorrect        = errors.New("password is not correct")
	ErrAccountAlreadyExists      = errors.New("account already exists")
	ErrAccountNotExists          = errors.New("account not exists")
	ErrAccountIdAlreadyExists    = errors.New("account id already exists")
	ErrEmailInvalid              = errors.New("email address is not valid")
	ErrMissingAccount            = errors.New("missing account")
	ErrMissingAccountOrPasswords = errors.New("missing account or password")
	ErrAccountPasswordInValid    = errors.New("password is not valid")
	ErrAccountHasBindUser        = errors.New("account has bind user")
	ErrMissingUserNameOrCardId   = errors.New("missing user name or CardID")
)

type AccountService interface {
	Login(account, pwd string) (*model.Account, error)
	RegisterAccount(account *model.Account) error
	UpdateAccountId(oldAccId, newAccId string) error
	UpdateAccountPassword(accId, oldPwd, newPwd string) error
	UpdateAccountEmail(accId, email string) error
	DeleteAccount(account string) error

	BindUser(account, userName, cardId string) error
	UpdateUser(user *model.User) error
}

var _ AccountService = (*AccountServiceImpl)(nil)

type AccountServiceImpl struct {
	accountRepo repo.AccountRepository
	userRepo    repo.UserRepository
}

func NewAccountService(accountRepo repo.AccountRepository, userRepo repo.UserRepository) AccountService {
	return &AccountServiceImpl{
		accountRepo: accountRepo,
		userRepo:    userRepo,
	}
}

func (s *AccountServiceImpl) Login(account string, pwd string) (*model.Account, error) {
	if account == "" || pwd == "" {
		return nil, ErrMissingAccountOrPasswords
	}

	acc, err := s.accountRepo.FindByAccount(account)
	if err != nil {
		return nil, err
	}

	if !password.CheckPasswordHash(pwd, acc.Password) {
		return nil, ErrPasswordNotCorrect
	}

	return acc, nil
}

func (s *AccountServiceImpl) RegisterAccount(account *model.Account) (err error) {
	if err := account.Validate(); err != nil {
		return err
	}

	account.Password, err = password.HashPassword(account.Password)
	if err != nil {
		plog.Errorf("hash password error: %v", err)
		return ErrAccountPasswordInValid
	}

	if _, err := s.accountRepo.FindByAccount(account.Account); err == nil {
		return ErrAccountAlreadyExists
	}

	if _, err := s.accountRepo.SaveAccount(account); err != nil {
		return err
	}
	return nil
}

func (s *AccountServiceImpl) UpdateAccountId(oldAccId, newAccId string) error {
	if oldAccId == "" {
		return ErrMissingAccount
	}

	acc, err := s.accountRepo.FindByAccount(oldAccId)
	if err != nil {
		return err
	}

	newAcc, err := s.accountRepo.FindByAccount(newAccId)
	if err != nil && !errors.Is(err, repo.ErrAccountNotExists) {
		return err
	}

	if newAcc != nil {
		return ErrAccountIdAlreadyExists
	}

	acc.Account = newAccId
	return s.accountRepo.UpdateAccount(acc)
}

func (s *AccountServiceImpl) UpdateAccountPassword(accId, oldPwd, newPwd string) error {
	if accId == "" {
		return ErrMissingAccount
	}

	acc, err := s.accountRepo.FindByAccount(accId)
	if err != nil {
		return err
	}

	if !acc.CheckPassword(oldPwd) {
		return ErrPasswordNotCorrect
	}

	if err := acc.SetPassword(newPwd); err != nil {
		return err
	}

	return s.accountRepo.UpdateAccount(acc)
}

func (s *AccountServiceImpl) UpdateAccountEmail(accId, email string) error {
	if accId == "" {
		return ErrMissingAccount
	}

	acc, err := s.accountRepo.FindByAccount(accId)
	if err != nil {
		return err
	}

	e := model.NewEmail(email)
	if !e.IsValidEmail() {
		return ErrEmailInvalid
	}

	acc.MainEmail = e

	return s.accountRepo.UpdateAccount(acc)
}

func (s *AccountServiceImpl) UpdateUser(user *model.User) error {
	u, err := s.userRepo.FindById(user.Id)
	if err != nil {
		return nil
	}

	if u.Equal(user) {
		return nil
	}

	return s.userRepo.UpdateUser(user)
}

func (s *AccountServiceImpl) findUserByIdCardAndUserName(userName, idCard string) (*model.User, error) {
	user, err := s.userRepo.FindByIdCardAndUserName(userName, idCard)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *AccountServiceImpl) BindUser(account, userName, cardId string) error {
	if account == "" {
		return ErrMissingAccount
	}

	if cardId == "" || userName == "" {
		return ErrMissingUserNameOrCardId
	}

	acc, err := s.accountRepo.FindByAccount(account)
	if err != nil {
		return err
	}

	if acc.User != nil {
		return ErrAccountHasBindUser
	}

	// query whether a record of the user exists.
	user, err := s.findUserByIdCardAndUserName(userName, cardId)
	if err != nil && !errors.Is(err, repo.ErrUSerNotExists) {
		return err
	}

	if user == nil {
		// add the user record if it doesn't exist
		user, err = s.userRepo.SaveUser(&model.User{
			IdCard:   cardId,
			UserName: userName,
		})
		if err != nil {
			return err
		}
	}

	acc.User = user
	return s.accountRepo.UpdateAccount(acc)
}

func (s *AccountServiceImpl) DeleteAccount(account string) error {
	if account == "" {
		return ErrMissingAccount
	}

	acc, err := s.accountRepo.FindByAccount(account)
	if err != nil {
		return err
	}

	return s.accountRepo.DeleteAccount(acc)
}
