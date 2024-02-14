package router

import (
	"github.com/gofiber/contrib/fibersentry"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/ubaidillahhf/dating-service/app/infra/config"
	"github.com/ubaidillahhf/dating-service/app/interfaces/handler"
	"github.com/ubaidillahhf/dating-service/app/interfaces/middleware"
	"github.com/ubaidillahhf/dating-service/app/usecases"
)

func Init(useCase usecases.AppUseCase, conf config.IConfig) {
	router := fiber.New()

	// middleware
	allowCors := cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowHeaders:     "Authorization, Origin, Content-Length, Content-Type, User-Agent, Referrer, Host, Token, CSRF-Token",
		AllowMethods:     "GET, POST, PATCH, OPTIONS, PUT, DELETE",
		ExposeHeaders:    "Content-Length",
		AllowCredentials: true,
	})
	logging := logger.New(logger.Config{
		Format: "${pid} ${locals:requestid} ${status} - ${method} ${path}\n",
	})
	router.Use(allowCors, logging, recover.New())
	router.Use(fibersentry.New(fibersentry.Config{
		Repanic: true,
	}))

	// hadler
	userHandler := handler.NewUserHandler(&useCase.UserUsecase)
	swipeHandler := handler.NewSwipeHandler(&useCase.SwipeUsecase)
	premiumHandler := handler.NewPremiumHandler(&useCase.PremiumUsecase)

	// service route
	router.Get("/", handler.GetTopRoute)

	api := router.Group("/api")
	v1 := api.Group("/v1")

	user := v1.Group("/users")
	user.Post("/register", userHandler.Register)
	user.Post("/login", userHandler.Login)
	user.Patch("/", middleware.ValidateToken, userHandler.Update)
	user.Get("/find-match", middleware.ValidateToken, userHandler.GetRandomProfiles)
	user.Get("/my-profile", middleware.ValidateToken, userHandler.MyProfile)

	swipe := v1.Group("/swipes")
	swipe.Post("/", middleware.ValidateToken, swipeHandler.Swipe)

	premium := v1.Group("/premiums")
	premium.Get("/", middleware.ValidateToken, premiumHandler.GetPackagePremium)
	premium.Post("/order", middleware.ValidateToken, premiumHandler.OrderPackage)

	callback := v1.Group("/callbacks")
	callback.Post("/payment", premiumHandler.PaymentCallback)

	router.Listen(":" + conf.Get("PORT"))
}
