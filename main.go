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

	// Setup Service
	useCase := usecases.NewAppUseCase(
		userRepository,
	)

	router.Init(useCase, configuration)
}
