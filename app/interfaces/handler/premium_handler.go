package handler

import (
	"context"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/ubaidillahhf/dating-service/app/domain"
	"github.com/ubaidillahhf/dating-service/app/infra/exception"
	"github.com/ubaidillahhf/dating-service/app/infra/presenter"
	"github.com/ubaidillahhf/dating-service/app/infra/utility/helper"
	xvalidator "github.com/ubaidillahhf/dating-service/app/infra/validator"
	"github.com/ubaidillahhf/dating-service/app/usecases"
)

type IPremiumHandler interface {
	GetPackagePremium(c *fiber.Ctx) error
	OrderPackage(c *fiber.Ctx) error
	PaymentCallback(c *fiber.Ctx) error
}

type premiumHandler struct {
	uc usecases.IPremiumUsecase
}

func NewPremiumHandler(uc *usecases.IPremiumUsecase) IPremiumHandler {
	return &premiumHandler{
		uc: *uc,
	}
}

func (co *premiumHandler) GetPackagePremium(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	queryParam := domain.Meta{}
	if err := c.QueryParser(&queryParam); err != nil {
		return c.JSON(presenter.Error(err.Error(), nil, fiber.StatusBadRequest))
	}

	skip, limit, defPage, defPerPage := helper.NormalizeAndGetDefaultPagination(queryParam.Page, queryParam.PerPage)

	responses, total, err := co.uc.GetPackagePremium(ctx, domain.Meta{
		Skip:  skip,
		Limit: limit,
	})
	if err != nil {
		return c.JSON(presenter.Error(err.Err.Error(), nil, err.Code))
	}

	return c.JSON(presenter.Success("Success", responses, presenter.Meta(presenter.MetaProps{
		Page:    defPage,
		PerPage: defPerPage,
		Total:   total,
	})))
}

func (co *premiumHandler) PaymentCallback(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	request := domain.PaymentCallbackRequest{}
	if err := c.BodyParser(&request); err != nil {
		return c.JSON(presenter.Error(err.Error(), nil, exception.BadRequestError))
	}

	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := validate.Struct(request); err != nil {
		return c.JSON(presenter.Error("error", xvalidator.GenerateHumanizeError(request, err), exception.BadRequestError))
	}

	res, resErr := co.uc.PaymentCallback(ctx, request)
	if resErr != nil {
		return c.JSON(presenter.Error(resErr.Err.Error(), nil, resErr.Code))
	}

	return c.JSON(presenter.Success("Success", res, nil))
}

func (co *premiumHandler) OrderPackage(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	myId := c.Locals("myId").(string)

	request := domain.Subscription{}
	if err := c.BodyParser(&request); err != nil {
		return c.JSON(presenter.Error(err.Error(), nil, exception.BadRequestError))
	}

	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := validate.Struct(request); err != nil {
		return c.JSON(presenter.Error("error", xvalidator.GenerateHumanizeError(request, err), exception.BadRequestError))
	}

	res, resErr := co.uc.OrderPackage(ctx, myId, request.PremiumPackagesId)
	if resErr != nil {
		return c.JSON(presenter.Error(resErr.Err.Error(), nil, resErr.Code))
	}

	return c.JSON(presenter.Success("Success", res, nil))
}
