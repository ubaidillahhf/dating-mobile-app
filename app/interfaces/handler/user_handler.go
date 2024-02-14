package handler

import (
	"context"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/copier"
	"github.com/ubaidillahhf/dating-service/app/domain"
	"github.com/ubaidillahhf/dating-service/app/infra/exception"
	"github.com/ubaidillahhf/dating-service/app/infra/presenter"
	"github.com/ubaidillahhf/dating-service/app/infra/utility/helper"
	xvalidator "github.com/ubaidillahhf/dating-service/app/infra/validator"
	"github.com/ubaidillahhf/dating-service/app/usecases"
)

type IUserHandler interface {
	Register(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	GetRandomProfiles(*fiber.Ctx) error
	MyProfile(*fiber.Ctx) error
}

type userHandler struct {
	uc usecases.IUserUsecase
}

func NewUserHandler(uc *usecases.IUserUsecase) IUserHandler {
	return &userHandler{
		uc: *uc,
	}
}

func (co *userHandler) Register(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	request := domain.RegisterRequest{}
	if err := c.BodyParser(&request); err != nil {
		return c.JSON(presenter.Error(err.Error(), nil, exception.BadRequestError))
	}

	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := validate.Struct(request); err != nil {
		return c.JSON(presenter.Error("error", xvalidator.GenerateHumanizeError(request, err), exception.BadRequestError))
	}

	res, resErr := co.uc.Register(ctx, request)
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
	res, resErr := co.uc.Login(ctx, newData)
	if resErr != nil {
		return c.JSON(presenter.Error(resErr.Err.Error(), nil, resErr.Code))
	}

	newRes := domain.LoginResponse{
		Email: res.Email,
		Token: res.Token,
	}
	return c.JSON(presenter.Success("Success", newRes, nil))
}

func (co *userHandler) GetRandomProfiles(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	myId := c.Locals("myId").(string)

	queryParam := domain.Meta{}
	if err := c.QueryParser(&queryParam); err != nil {
		return c.JSON(presenter.Error(err.Error(), nil, fiber.StatusBadRequest))
	}

	skip, limit, defPage, defPerPage := helper.NormalizeAndGetDefaultPagination(queryParam.Page, queryParam.PerPage)

	responses, total, err := co.uc.GetRandomProfiles(ctx, domain.Meta{Skip: skip, Limit: limit}, myId)
	if err != nil {
		return c.JSON(presenter.Error(err.Err.Error(), nil, err.Code))
	}

	return c.JSON(presenter.Success("Success", presenter.FindMatchTransform(responses), presenter.Meta(presenter.MetaProps{
		Page:    defPage,
		PerPage: defPerPage,
		Total:   total,
	})))
}

func (co *userHandler) Update(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	request := domain.User{}
	if err := c.BodyParser(&request); err != nil {
		return c.JSON(presenter.Error(err.Error(), nil, exception.BadRequestError))
	}

	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := validate.Struct(request); err != nil {
		return c.JSON(presenter.Error("error", xvalidator.GenerateHumanizeError(request, err), exception.BadRequestError))
	}

	myId := c.Locals("myId").(string)
	request.Id = myId

	res, resErr := co.uc.UpdateProfile(ctx, request)
	if resErr != nil {
		return c.JSON(presenter.Error(resErr.Err.Error(), nil, resErr.Code))
	}

	return c.JSON(presenter.Success("Success", res, nil))
}

func (co *userHandler) MyProfile(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	myId := c.Locals("myId").(string)

	res, resErr := co.uc.GetProfile(ctx, myId)
	if resErr != nil {
		return c.JSON(presenter.Error(resErr.Err.Error(), nil, resErr.Code))
	}

	profileResponse := domain.ProfileResponse{}
	copier.Copy(&profileResponse, &res)

	return c.JSON(presenter.Success("Success", profileResponse, nil))
}
