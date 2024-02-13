package handler

import (
	"context"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/ubaidillahhf/dating-service/app/domain"
	"github.com/ubaidillahhf/dating-service/app/infra/exception"
	"github.com/ubaidillahhf/dating-service/app/infra/presenter"
	xvalidator "github.com/ubaidillahhf/dating-service/app/infra/validator"
	"github.com/ubaidillahhf/dating-service/app/usecases"
)

type IUserHandler interface {
	Register(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error
}

type userHandler struct {
	userUsecase usecases.IUserUsecase
}

func NewUserHandler(userUsecase *usecases.IUserUsecase) IUserHandler {
	return &userHandler{
		userUsecase: *userUsecase,
	}
}

func (co *userHandler) Register(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	request := new(domain.RegisterRequest)
	if err := c.BodyParser(&request); err != nil {
		return c.JSON(presenter.Error(err.Error(), nil, exception.BadRequestError))
	}

	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := validate.Struct(request); err != nil {
		return c.JSON(presenter.Error("error", xvalidator.GenerateHumanizeError(request, err), exception.BadRequestError))
	}

	newData := domain.User{
		Fullname: request.Fullname,
		Username: request.Username,
		Email:    request.Email,
		Password: request.Password,
	}
	res, resErr := co.userUsecase.Register(ctx, newData)
	if resErr != nil {
		return c.JSON(presenter.Error(resErr.Err.Error(), nil, resErr.Code))
	}

	newRes := domain.RegisterResponse{
		Id:    res.Id,
		Email: res.Email,
	}
	return c.JSON(presenter.Success("Success", newRes, nil))
}

func (co *userHandler) Login(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	request := new(domain.LoginRequest)
	if err := c.BodyParser(&request); err != nil {
		return c.JSON(presenter.Error(err.Error(), nil, exception.BadRequestError))
	}

	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := validate.Struct(request); err != nil {
		return c.JSON(presenter.Error("error", xvalidator.GenerateHumanizeError(request, err), exception.BadRequestError))
	}

	newData := domain.LoginRequest{
		Email:    request.Email,
		Password: request.Password,
	}
	res, resErr := co.userUsecase.Login(ctx, newData)
	if resErr != nil {
		return c.JSON(presenter.Error(resErr.Err.Error(), nil, resErr.Code))
	}

	newRes := domain.LoginResponse{
		Email: res.Email,
		Token: res.Token,
	}
	return c.JSON(presenter.Success("Success", newRes, nil))
}
