// File:		mysql.go
// Created by:	Hoven
// Created on:	2024-08-08
//
// This file is part of the Example Project.
//
// (c) 2024 Example Corp. All rights reserved.

package mysql

import (
	"gorm.io/gorm"
)

type BaseRepository[T any] interface {
	DB() *gorm.DB
	Create(entity T) error
	Update(entity T) error
	Delete(entity T) error
	FindByID(id int) (T, error)
	All() ([]T, error)
	FindAll(query any, args ...any) ([]T, error)
	First(query any, args ...any) (T, error)
}

var _ BaseRepository[MockT] = (*MysqlRepository[MockT])(nil)

type MockT struct{}

type MysqlRepository[T any] struct {
	db *gorm.DB
}

func NewMysqlRepository[T any](db *gorm.DB) *MysqlRepository[T] {
	return &MysqlRepository[T]{db: db}
}

func (r *MysqlRepository[T]) DB() *gorm.DB {
	return r.db
}

func (r *MysqlRepository[T]) Create(entity T) error {
	return r.db.Create(entity).Error
}

func (r *MysqlRepository[T]) Update(entity T) error {
	return r.db.Save(entity).Error
}

func (r *MysqlRepository[T]) Delete(entity T) error {
	return r.db.Delete(entity).Error
}

func (r *MysqlRepository[T]) FindByID(id int) (T, error) {
	var entity T
	if err := r.db.First(&entity, id).Error; err != nil {
		return entity, err
	}
	return entity, nil
}

func (r *MysqlRepository[T]) All() ([]T, error) {
	var entities []T
	if err := r.db.Find(&entities).Error; err != nil {
		return nil, err
	}
	return entities, nil
}

func (r *MysqlRepository[T]) FindAll(query any, args ...any) ([]T, error) {
	var entities []T
	if err := r.db.Where(query, args...).Find(&entities).Error; err != nil {
		return nil, err
	}
	return entities, nil
}

func (r *MysqlRepository[T]) First(query any, args ...any) (T, error) {
	var entity T
	err := r.db.Where(query, args...).First(&entity).Error
	return entity, err
}
