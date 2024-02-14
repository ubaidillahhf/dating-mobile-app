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

type ISwipeHandler interface {
	Swipe(c *fiber.Ctx) error
}

type swipeHandler struct {
	uc usecases.ISwipeUsecase
}

func NewSwipeHandler(uc *usecases.ISwipeUsecase) ISwipeHandler {
	return &swipeHandler{
		uc: *uc,
	}
}

func (co *swipeHandler) Swipe(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	request := domain.Swipe{}
	if err := c.BodyParser(&request); err != nil {
		return c.JSON(presenter.Error(err.Error(), nil, exception.BadRequestError))
	}

	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := validate.Struct(request); err != nil {
		return c.JSON(presenter.Error("error", xvalidator.GenerateHumanizeError(request, err), exception.BadRequestError))
	}

	myId := c.Locals("myId").(string)
	request.SenderId = myId

	res, resErr := co.uc.Swipe(ctx, request)
	if resErr != nil {
		return c.JSON(presenter.Error(resErr.Err.Error(), nil, resErr.Code))
	}

	return c.JSON(presenter.Success("Success", res, nil))
}
