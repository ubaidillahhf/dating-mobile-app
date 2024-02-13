package usecases

import "github.com/ubaidillahhf/dating-service/app/infra/repository"

type AppUseCase struct {
	UserUsecase IUserUsecase
}

func NewAppUseCase(
	UserRepo repository.IUserRepository,
) AppUseCase {
	return AppUseCase{
		UserUsecase: NewUserUsecase(&UserRepo),
	}
}
