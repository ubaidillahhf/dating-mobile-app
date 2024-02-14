package usecases

import (
	"context"
	"errors"
	"strconv"

	"github.com/ubaidillahhf/dating-service/app/domain"
	"github.com/ubaidillahhf/dating-service/app/infra/config"
	"github.com/ubaidillahhf/dating-service/app/infra/exception"
	"github.com/ubaidillahhf/dating-service/app/infra/repository"
	"github.com/ubaidillahhf/dating-service/app/infra/utility/helper"
	"github.com/ubaidillahhf/dating-service/app/interfaces/middleware"
	"golang.org/x/crypto/bcrypt"
)

type IUserUsecase interface {
	Register(ctx context.Context, request domain.RegisterRequest) (domain.User, *exception.Error)
	Login(ctx context.Context, request domain.LoginRequest) (domain.LoginResponse, *exception.Error)
	UpdateProfile(ctx context.Context, request domain.User) (bool, *exception.Error)
	GetRandomProfiles(ctx context.Context, meta domain.Meta, myId string) ([]domain.User, int64, *exception.Error)
}

func NewUserUsecase(repo repository.IUserRepository) IUserUsecase {
	return &userUsecase{
		repo: repo,
	}
}

type userUsecase struct {
	repo repository.IUserRepository
}

func (uc *userUsecase) Register(ctx context.Context, request domain.RegisterRequest) (res domain.User, err *exception.Error) {

	data, _ := uc.repo.FindByIdentifier(ctx, request.Username, request.Email)
	if data != (domain.User{}) {
		return res, &exception.Error{
			Code: exception.BadRequestError,
			Err:  errors.New("username or email already registered"),
		}
	}

	if request.Username == "" {
		request.Username = helper.RandomUsername(request.Fullname)
	}

	hashPwd, _ := helper.HashPassword(request.Password)
	newData := domain.User{
		Username: request.Username,
		Fullname: request.Fullname,
		Email:    request.Email,
		Password: hashPwd,
		Gender:   domain.Undisclosed,
	}

	p, pErr := uc.repo.Insert(ctx, newData)
	if pErr != nil {
		return res, &exception.Error{
			Code: exception.IntenalError,
			Err:  pErr,
		}
	}

	return p, nil
}

func (uc *userUsecase) Login(ctx context.Context, req domain.LoginRequest) (res domain.LoginResponse, err *exception.Error) {

	secret := config.GetEnv("ACCESS_TOKEN_SECRET")
	exp := config.GetEnv("ACCESS_TOKEN_EXPIRY_HOUR")
	expAsInt, _ := strconv.Atoi(exp)

	data, _ := uc.repo.FindByIdentifier(ctx, "", req.Email)
	match := bcrypt.CompareHashAndPassword([]byte(data.Password), []byte(req.Password))
	if match != nil {
		return res, &exception.Error{
			Code: 400,
			Err:  errors.New("wrong password"),
		}
	}

	newToken, _ := middleware.CreateAccessToken(&data, secret, int(expAsInt))

	res = domain.LoginResponse{
		Email: data.Email,
		Token: newToken,
	}

	return res, nil
}

func (uc *userUsecase) UpdateProfile(ctx context.Context, request domain.User) (res bool, err *exception.Error) {

	if request.Id == "" {
		return res, &exception.Error{
			Code: exception.BadRequestError,
			Err:  errors.New("error: unknown userId"),
		}
	}

	if request.Username != "" || request.Email != "" || request.Password != "" {
		return res, &exception.Error{
			Code: exception.BadRequestError,
			Err:  errors.New("username, email, and password can't be change from this ep"),
		}
	}

	p, pErr := uc.repo.Update(ctx, request)
	if pErr != nil {
		return res, &exception.Error{
			Code: exception.IntenalError,
			Err:  pErr,
		}
	}

	return p, nil
}

func (uc *userUsecase) GetRandomProfiles(ctx context.Context, meta domain.Meta, myId string) (res []domain.User, t int64, err *exception.Error) {

	data, total, dErr := uc.repo.Get(ctx, meta, myId, true)
	if dErr != nil {
		return res, t, &exception.Error{
			Code: exception.IntenalError,
			Err:  dErr,
		}
	}

	return data, total, nil
}
