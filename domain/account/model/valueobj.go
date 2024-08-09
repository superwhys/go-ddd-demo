// File:		valueobj.go
// Created by:	Hoven
// Created on:	2024-08-08
//
// This file is part of the Example Project.
//
// (c) 2024 Example Corp. All rights reserved.

package model

import (
	"regexp"
)

type Gender uint

const (
	GenderUnknown Gender = iota
	GenderFemale
	GenderMale
)

func (g Gender) String() string {
	switch g {
	case GenderFemale:
		return "Female"
	case GenderMale:
		return "Male"
	default:
		return "Unknown"
	}
}

type Email struct {
	Address string `json:"address"`
}

func NewEmail(address string) *Email {
	return &Email{Address: address}
}

func (e *Email) Equal(other *Email) bool {
	if other == nil {
		return false
	}
	return e.Address == other.Address
}

func (e *Email) IsValidEmail() bool {
	if e.Address == "" {
		return true
	}

	regex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}$`)
	return regex.MatchString(e.Address)
}
