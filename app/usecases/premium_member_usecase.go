package usecases

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/ubaidillahhf/dating-service/app/domain"
	"github.com/ubaidillahhf/dating-service/app/infra/exception"
	"github.com/ubaidillahhf/dating-service/app/infra/repository"
	logx "github.com/ubaidillahhf/dating-service/app/infra/utility/logger"
)

type IPremiumMemberUsecase interface {
	GetPackagePremium(ctx context.Context, meta domain.Meta) ([]domain.PremiumPackage, int64, *exception.Error)
	SubscribePackage(ctx context.Context, myId string, packageId int64) (domain.Subscription, *exception.Error)
	PaymentCallback(ctx context.Context, data domain.PaymentCallbackRequest) (domain.Payment, *exception.Error)
}

func NewPremiumMemberUsecase(
	repo repository.IPremiumPackageRepository,
	userRepo repository.IUserRepository,
	subsRepo repository.ISubscriptionRepository,
	paymentRepo repository.IPaymentRepository,
	gormTx repository.IGormTx,
) IPremiumMemberUsecase {
	return &premiumMemberUsecase{
		repo:        repo,
		userRepo:    userRepo,
		subsRepo:    subsRepo,
		paymentRepo: paymentRepo,
		gormTx:      gormTx,
	}
}

type premiumMemberUsecase struct {
	repo        repository.IPremiumPackageRepository
	userRepo    repository.IUserRepository
	subsRepo    repository.ISubscriptionRepository
	paymentRepo repository.IPaymentRepository
	gormTx      repository.IGormTx
}

func (uc *premiumMemberUsecase) GetPackagePremium(ctx context.Context, meta domain.Meta) (res []domain.PremiumPackage, total int64, err *exception.Error) {
	data, total, dErr := uc.repo.Get(ctx, meta)
	if dErr != nil {
		return res, total, &exception.Error{
			Code: exception.IntenalError,
			Err:  dErr,
		}
	}

	return data, total, nil
}

func (uc *premiumMemberUsecase) PaymentCallback(ctx context.Context, data domain.PaymentCallbackRequest) (res domain.Payment, err *exception.Error) {

	paymentStatus := domain.PaymentWaiting
	subsStatus := domain.SubsPending

	userId := data.UserId
	paymentId := data.Id
	refIdAsInt, riaiErr := strconv.ParseInt(data.RefId, 10, 64)
	subsId := refIdAsInt

	if ok := uc.paymentRepo.ValidateCallback(ctx, paymentId, userId, data.RefId); !ok {
		return res, &exception.Error{
			Code: exception.IntenalError,
			Err:  errors.New("ops something error, try again later"),
		}

		// can insert to queue for manual retry
	}

	tx, txErr := uc.gormTx.Begin()
	if txErr != nil {
		return res, &exception.Error{
			Code: exception.IntenalError,
			Err:  txErr,
		}
	}

	switch data.Status {
	case domain.PaymentSuccess:

		paymentStatus = domain.PaymentSuccess
		subsStatus = domain.SubsActive

		if _, err := uc.userRepo.UpdateTx(ctx, tx, domain.User{
			Id:        userId,
			IsPremium: 1,
		}); err != nil {
			uc.gormTx.Rollback(tx)

			return res, &exception.Error{
				Code: exception.IntenalError,
				Err:  err,
			}
		}

	case domain.PaymentFailed:

		paymentStatus = domain.PaymentFailed
		subsStatus = domain.SubsCancel

	}

	if _, err := uc.paymentRepo.UpdateTx(ctx, tx, domain.Payment{
		Id:     paymentId,
		Status: paymentStatus,
	}); err != nil {
		uc.gormTx.Rollback(tx)

		return res, &exception.Error{
			Code: exception.IntenalError,
			Err:  err,
		}
	}

	if riaiErr != nil {
		uc.gormTx.Rollback(tx)

		return res, &exception.Error{
			Code: exception.IntenalError,
			Err:  riaiErr,
		}
	}
	if _, err := uc.subsRepo.UpdateTx(ctx, tx, domain.Subscription{
		Id:     subsId,
		Status: subsStatus,
	}); err != nil {
		uc.gormTx.Rollback(tx)

		return res, &exception.Error{
			Code: exception.IntenalError,
			Err:  err,
		}
	}

	uc.gormTx.Commit(tx)

	return
}

func (uc *premiumMemberUsecase) SubscribePackage(ctx context.Context, myId string, packageId int64) (res domain.Subscription, err *exception.Error) {

	detailPremiumPkg, dpmErr := uc.repo.Find(ctx, packageId)
	if dpmErr != nil {
		return res, &exception.Error{
			Code: exception.IntenalError,
			Err:  dpmErr,
		}
	}

	tx, txErr := uc.gormTx.Begin()
	if txErr != nil {
		return res, &exception.Error{
			Code: exception.IntenalError,
			Err:  txErr,
		}
	}

	newData := domain.Subscription{
		UserId:            myId,
		PremiumPackagesId: packageId,
		Status:            domain.SubsPending,
	}
	p, pErr := uc.subsRepo.CreateTx(ctx, tx, newData)
	if pErr != nil {
		uc.gormTx.Rollback(tx)
		return res, &exception.Error{
			Code: exception.IntenalError,
			Err:  pErr,
		}
	}

	// create waiting payment
	payment, paErr := uc.paymentRepo.CreateTx(ctx, tx, domain.Payment{
		UserId:     myId,
		RefContext: domain.RefContextSubs,
		RefId:      fmt.Sprintf("%d", p.Id),
		Amount:     detailPremiumPkg.Price,
		Status:     domain.PaymentWaiting,
	})
	if paErr != nil {
		uc.gormTx.Rollback(tx)
		return p, &exception.Error{
			Code: exception.IntenalError,
			Err:  paErr,
		}
	}

	uc.gormTx.Commit(tx)

	go uc.requestPayment(payment.Id)

	return p, nil
}

func (uc *premiumMemberUsecase) requestPayment(paymentId int64) error {

	logx.Create().Info().Msg(fmt.Sprintf("initiate payment create with id: %d. call payment gateway...", paymentId))

	// do requset to payment gateway (dummy process)
	// do with payload:
	/**
	 * - status (req),
	 * meta: {
	 * - paymentId (req),
	 * - userId (req),
	 * - refContext (req),
	 * - subsId (req)
	 * }
	 */

	return nil

}
