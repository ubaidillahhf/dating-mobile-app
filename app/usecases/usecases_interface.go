package usecases

import "github.com/ubaidillahhf/dating-service/app/infra/repository"

type AppUseCase struct {
	UserUsecase  IUserUsecase
	SwipeUsecase ISwipeUsecase
}

func NewAppUseCase(
	UserRepo repository.IUserRepository,
	SwipeRepo repository.ISwipeRepository,
) AppUseCase {
	return AppUseCase{
		UserUsecase:  NewUserUsecase(UserRepo),
		SwipeUsecase: NewSwipeUsecase(SwipeRepo, UserRepo),
	}
}
