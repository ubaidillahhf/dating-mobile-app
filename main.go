package main

import (
	"github.com/ubaidillahhf/dating-service/app/infra/config"
	"github.com/ubaidillahhf/dating-service/app/infra/repository"
	"github.com/ubaidillahhf/dating-service/app/infra/router"
	"github.com/ubaidillahhf/dating-service/app/usecases"
)

func main() {
	// load config
	configuration := config.New(".env")

	// error monitoring
	config.SentryInit(configuration)

	// conn mongo
	database := config.NewGormPostgres(configuration)

	// Setup Repository
	userRepository := repository.NewUserRepository(database)
	swipeRepository := repository.NewSwipeRepository(database)
	premiumRepository := repository.NewPremiumPackageRepository(database)
	subsRepository := repository.NewSubscriptionRepository(database)
	paymentRepository := repository.NewPaymentRepository(database)
	gormTx := repository.NewGormTx(database)

	// Setup Service
	useCase := usecases.NewAppUseCase(
		userRepository,
		swipeRepository,
		premiumRepository,
		subsRepository,
		paymentRepository,
		gormTx,
	)

	router.Init(useCase, configuration)
}
