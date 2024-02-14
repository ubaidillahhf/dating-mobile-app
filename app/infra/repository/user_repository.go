package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/jinzhu/now"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/ubaidillahhf/dating-service/app/domain"
	"github.com/ubaidillahhf/dating-service/app/infra/utility/helper"
	logx "github.com/ubaidillahhf/dating-service/app/infra/utility/logger"
	"gorm.io/gorm"
)

type IUserRepository interface {
	Insert(ctx context.Context, newData domain.User) (domain.User, error)
	FindByIdentifier(ctx context.Context, username, email string) (domain.User, error)
	Update(ctx context.Context, newData domain.User) (bool, error)
	Get(ctx context.Context, meta domain.Meta, myId string, excludeSeenDaily bool) ([]domain.User, int64, error)
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &userRepository{
		conn: db,
	}
}

type userRepository struct {
	conn *gorm.DB
}

func (repo *userRepository) Insert(ctx context.Context, newData domain.User) (res domain.User, err error) {

	nd := domain.User{
		Id:       gonanoid.Must(),
		Email:    newData.Email,
		Password: newData.Password,
	}

	if err := repo.conn.WithContext(ctx).Table("users").Create(&nd).Error; err != nil {
		logx.Create().Error().Msg(fmt.Sprintf("error: at repository when create user. Detail: %v", err))
		return res, errors.New("error: when create user")
	}

	return nd, nil
}

func (repo *userRepository) FindByIdentifier(ctx context.Context, username, email string) (res domain.User, err error) {

	if err := repo.conn.WithContext(ctx).Table("users").
		Where("username = ? OR email = ?", username, email).
		First(&res).Error; err != nil {
		logx.Create().Error().Msg(fmt.Sprintf("error: at repository when get detail user. Detail: %v", err))
		return res, errors.New("error: when get detail user")
	}

	return
}

func (repo *userRepository) Update(ctx context.Context, newData domain.User) (res bool, err error) {
	newUpdate := make(map[string]interface{})

	if newData.Fullname != "" {
		newUpdate["fullname"] = newData.Fullname
	}

	if err := repo.conn.WithContext(ctx).Table("users").
		Where("id = ?", newData.Id).
		Updates(newUpdate).Error; err != nil {
		logx.Create().Error().Msg(fmt.Sprintf("error: at repository when update user. Detail: %v", err))
		return false, errors.New("error: when update user")
	}

	return
}

func (repo *userRepository) Get(ctx context.Context, meta domain.Meta, myId string, excludeSeenDaily bool) (res []domain.User, total int64, err error) {

	q := repo.conn.WithContext(ctx).Table("users")

	if excludeSeenDaily {
		startAt := now.BeginningOfDay().UTC()
		endAt := now.EndOfDay().UTC()

		subQ := repo.conn.
			Select("sender_id").
			Where("sender_id = ?", myId).
			Where("created_at >= ? AND created_at <= ?", startAt, endAt).
			Table("swipes")

		q = q.Where("id NOT IN (?)", subQ)
	}

	if err := q.
		Count(&total).
		Order("RAND()").
		Scopes(helper.Paginate(meta.Page, meta.PerPage)).
		Find(&res).Error; err != nil {
		logx.Create().Error().Msg(fmt.Sprintf("error: at repository when get user. Detail: %v", err))
		return res, total, errors.New("error: at repository when get media")
	}

	return res, total, nil
}
