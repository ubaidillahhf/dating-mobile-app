package repository

import (
	"context"

	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/ubaidillahhf/dating-service/app/domain"
	"github.com/ubaidillahhf/dating-service/app/infra/exception"
	"gorm.io/gorm"
)

type IUserRepository interface {
	Insert(ctx context.Context, newData domain.User) (domain.User, *exception.Error)
	FindByIdentifier(ctx context.Context, username, email string) (domain.User, *exception.Error)
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &userRepository{
		conn: db,
	}
}

type userRepository struct {
	conn *gorm.DB
}

func (repo *userRepository) Insert(ctx context.Context, newData domain.User) (res domain.User, err *exception.Error) {

	nd := domain.User{
		Id:       gonanoid.Must(),
		Email:    newData.Email,
		Password: newData.Password,
	}

	if err := repo.conn.WithContext(ctx).Table("users").Create(&nd).Error; err != nil {
		return res, &exception.Error{
			Code: exception.IntenalError,
			Err:  err,
		}
	}

	return nd, nil
}

func (repo *userRepository) FindByIdentifier(ctx context.Context, username, email string) (res domain.User, err *exception.Error) {

	if err := repo.conn.WithContext(ctx).Table("users").
		Where("username = ? OR email = ?", username, email).
		First(&res).Error; err != nil {
		return res, &exception.Error{
			Code: exception.IntenalError,
			Err:  err,
		}
	}

	return
}
