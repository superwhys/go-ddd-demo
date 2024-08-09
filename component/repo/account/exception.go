// File:		exception.go
// Created by:	Hoven
// Created on:	2024-08-09
//
// This file is part of the Example Project.
//
// (c) 2024 Example Corp. All rights reserved.

package account

import "github.com/pkg/errors"

var (
	ErrAccountNotFound = errors.New("account not exists")
)
