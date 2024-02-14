package usecases

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/jinzhu/copier"
	"github.com/ubaidillahhf/dating-service/app/domain"
	"github.com/ubaidillahhf/dating-service/app/infra/config"
	"github.com/ubaidillahhf/dating-service/app/infra/exception"
	"github.com/ubaidillahhf/dating-service/app/infra/repository"
)

type ISwipeUsecase interface {
	Swipe(ctx context.Context, request domain.Swipe) (domain.SwipeResponse, *exception.Error)
}

func NewSwipeUsecase(repo repository.ISwipeRepository, userRepo repository.IUserRepository) ISwipeUsecase {
	return &swipeUsecase{
		repo:     repo,
		userRepo: userRepo,
	}
}

type swipeUsecase struct {
	repo     repository.ISwipeRepository
	userRepo repository.IUserRepository
}

func (uc *swipeUsecase) Swipe(ctx context.Context, request domain.Swipe) (res domain.SwipeResponse, err *exception.Error) {

	if request.SenderId == request.ReceiverId {
		return res, &exception.Error{
			Code: exception.BadRequestError,
			Err:  errors.New("error: can't swipe yourself"),
		}
	}

	valid, vErr := uc.userRepo.SenderReceiverValidation(ctx, request.SenderId, request.ReceiverId)
	if vErr != nil {
		return res, &exception.Error{
			Code: exception.IntenalError,
			Err:  vErr,
		}
	}
	if !valid {
		return res, &exception.Error{
			Code: exception.BadRequestError,
			Err:  errors.New("error: invalid receiverId"),
		}
	}

	if err := uc.swipeQuotaValidation(ctx, request.SenderId); err != nil {
		return res, &exception.Error{
			Code: exception.BadRequestError,
			Err:  err,
		}
	}

	p, pErr := uc.repo.Create(ctx, request)
	if pErr != nil {

		// when duplicate, no insert db but return success
		if strings.Contains(pErr.Error(), "23505") {
			copier.Copy(&res, &request)
			res.Notes = "sensational! you meet same user again, swipe success but unchange last direction"

			return res, nil
		}

		return res, &exception.Error{
			Code: exception.IntenalError,
			Err:  pErr,
		}
	}

	copier.Copy(&res, &p)
	res.Notes = "gotcha! new swipe success keep it up"

	return res, nil
}

func (uc *swipeUsecase) swipeQuotaValidation(ctx context.Context, userId string) error {

	basicUserQuota, _ := strconv.ParseInt(config.GetEnv("BASIC_USER_QUOTA"), 10, 64)
	if basicUserQuota == 0 {
		basicUserQuota = 10
	}

	user, err := uc.userRepo.Find(ctx, userId)
	if err != nil {
		return err
	}

	if user.IsPremium == 0 {
		totalSwipe, tsErr := uc.repo.CountBySenderId(ctx, userId, true)
		if tsErr != nil {
			return tsErr
		}

		if totalSwipe >= basicUserQuota {
			return fmt.Errorf("ops, basic user only %d swipe a day. Upgrade to premium for unlimited swipe", basicUserQuota)
		}

	}

	return nil
}
