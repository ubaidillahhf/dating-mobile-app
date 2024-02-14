package usecases

import (
	"context"
	"errors"
	"strings"

	"github.com/ubaidillahhf/dating-service/app/domain"
	"github.com/ubaidillahhf/dating-service/app/infra/exception"
	"github.com/ubaidillahhf/dating-service/app/infra/repository"
)

type ISwipeUsecase interface {
	Swipe(ctx context.Context, request domain.Swipe) (domain.Swipe, *exception.Error)
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

func (uc *swipeUsecase) Swipe(ctx context.Context, request domain.Swipe) (res domain.Swipe, err *exception.Error) {

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

	p, pErr := uc.repo.Insert(ctx, request)
	if pErr != nil {

		if strings.Contains(pErr.Error(), "23505") {
			return res, &exception.Error{
				Code: exception.BadRequestError,
				Err:  errors.New("error: duplicate action"),
			}
		}

		return res, &exception.Error{
			Code: exception.IntenalError,
			Err:  pErr,
		}
	}

	return p, nil
}
