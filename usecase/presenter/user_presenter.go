package presenter

import "github.com/emerak/golang-bootcamp-2020/domain/model"

type UserPresenter interface {
	ResponseUsers(u []*model.User) []*model.User
}
