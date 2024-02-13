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
	Register(ctx context.Context, request domain.User) (domain.User, *exception.Error)
	Login(ctx context.Context, request domain.LoginRequest) (domain.LoginResponse, *exception.Error)
}

func NewUserUsecase(repo *repository.IUserRepository) IUserUsecase {
	return &userUsecase{
		repo: *repo,
	}
}

type userUsecase struct {
	repo repository.IUserRepository
}

func (uc *userUsecase) Register(ctx context.Context, request domain.User) (res domain.User, err *exception.Error) {

	data, _ := uc.repo.FindByIdentifier(ctx, request.Username, request.Email)
	if data != (domain.User{}) {
		return res, &exception.Error{
			Code: 400,
			Err:  errors.New("username or email already registered"),
		}
	}

	if request.Username == "" {
		request.Username = helper.RandomUsername(request.Fullname)
	}

	hashPwd, _ := helper.HashPassword(request.Password)
	newData := domain.User{
		Email:    request.Email,
		Password: hashPwd,
	}

	p, pErr := uc.repo.Insert(ctx, newData)
	if pErr != nil {
		return res, pErr
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
