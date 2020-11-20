package repository

import "github.com/emerak/golang-bootcamp-2020/domain/model"

type UserRepository interface {
	FindAll(u []*model.User) ([]*model.User, error)
}
