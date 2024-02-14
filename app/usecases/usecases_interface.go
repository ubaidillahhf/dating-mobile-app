package usecases

import "github.com/ubaidillahhf/dating-service/app/infra/repository"

type AppUseCase struct {
	UserUsecase    IUserUsecase
	SwipeUsecase   ISwipeUsecase
	PremiumUsecase IPremiumUsecase
}

func NewAppUseCase(
	UserRepo repository.IUserRepository,
	SwipeRepo repository.ISwipeRepository,
	PremiumRepo repository.IPremiumPackageRepository,
	SubsRepo repository.ISubscriptionRepository,
	PaymentRepo repository.IPaymentRepository,
	GormTx repository.IGormTx,
) AppUseCase {
	return AppUseCase{
		UserUsecase:    NewUserUsecase(UserRepo),
		SwipeUsecase:   NewSwipeUsecase(SwipeRepo, UserRepo),
		PremiumUsecase: NewPremiumUsecase(PremiumRepo, UserRepo, SubsRepo, PaymentRepo, GormTx),
	}
}
