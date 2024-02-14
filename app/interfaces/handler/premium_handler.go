package handler

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/ubaidillahhf/dating-service/app/domain"
	"github.com/ubaidillahhf/dating-service/app/infra/presenter"
	"github.com/ubaidillahhf/dating-service/app/infra/utility/helper"
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

// PaymentCallback implements IPremiumHandler.
func (*premiumHandler) PaymentCallback(c *fiber.Ctx) error {
	panic("unimplemented")
}

// OrderPackage implements IPremiumHandler.
func (*premiumHandler) OrderPackage(c *fiber.Ctx) error {
	panic("unimplemented")
}
