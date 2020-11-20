package registry

import (
	"github.com/emerak/golang-bootcamp-2020/interface/controller"
	ip "github.com/emerak/golang-bootcamp-2020/interface/presenter"
	ir "github.com/emerak/golang-bootcamp-2020/interface/repository"
	"github.com/emerak/golang-bootcamp-2020/usecase/interactor"
	up "github.com/emerak/golang-bootcamp-2020/usecase/presenter"
	ur "github.com/emerak/golang-bootcamp-2020/usecase/repository"
)

func (r *registry) NewUserController() controller.UserController {
	return controller.NewUserController(r.NewUserInteractor())
}

func (r *registry) NewUserInteractor() interactor.UserInteractor {
	return interactor.NewUserInteractor(r.NewUserRepository(), r.NewUserPresenter())
}

func (r *registry) NewUserRepository() ur.UserRepository {
	return ir.NewUserRepository(r.db)
}

func (r *registry) NewUserPresenter() up.UserPresenter {
	return ip.NewUserPresenter()
}
