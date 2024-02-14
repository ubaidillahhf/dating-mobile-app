package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/ubaidillahhf/dating-service/app/domain"
	logx "github.com/ubaidillahhf/dating-service/app/infra/utility/logger"
	"gorm.io/gorm"
)

type ISubscriptionRepository interface {
	Create(ctx context.Context, newData domain.Subscription) (domain.Subscription, error)
	CreateTx(ctx context.Context, tx *gorm.DB, newData domain.Subscription) (domain.Subscription, error)
	UpdateTx(ctx context.Context, tx *gorm.DB, data domain.Subscription) (bool, error)
	Find(ctx context.Context, id int64) (domain.Subscription, error)
}

func NewSubscriptionRepository(db *gorm.DB) ISubscriptionRepository {
	return &subscriptionRepository{
		conn: db,
	}
}

type subscriptionRepository struct {
	conn *gorm.DB
}

func (repo *subscriptionRepository) Create(ctx context.Context, newData domain.Subscription) (res domain.Subscription, err error) {

	if err := repo.conn.WithContext(ctx).Table("subscriptions").Create(&newData).Error; err != nil {
		logx.Create().Error().Msg(fmt.Sprintf("error: at repository when create subscriptions. Detail: %v", err))
		return res, errors.New("error: when create subscriptions")
	}

	return newData, nil
}

func (repo *subscriptionRepository) CreateTx(ctx context.Context, tx *gorm.DB, newData domain.Subscription) (res domain.Subscription, err error) {

	if err := tx.WithContext(ctx).Table("subscriptions").Create(&newData).Error; err != nil {
		logx.Create().Error().Msg(fmt.Sprintf("error: at repository when create subscriptions. Detail: %v", err))
		return res, errors.New("error: when create subscriptions")
	}

	return newData, nil
}

func (repo *subscriptionRepository) UpdateTx(ctx context.Context, tx *gorm.DB, newData domain.Subscription) (res bool, err error) {
	newUpdate := make(map[string]interface{})

	if newData.Id == 0 {
		return res, errors.New("error: at repo subscription when update. Empty id update")
	}

	if newData.Status != "" {
		newUpdate["status"] = newData.Status
	}

	if err := tx.WithContext(ctx).Table("subscriptions").
		Where("id = ?", newData.Id).
		Updates(newUpdate).Error; err != nil {
		logx.Create().Error().Msg(fmt.Sprintf("error: at repository when update subscriptions. Detail: %v", err))
		return false, errors.New("error: when update subscriptions")
	}

	return true, nil
}

func (repo *subscriptionRepository) Find(ctx context.Context, id int64) (res domain.Subscription, err error) {

	if err := repo.conn.WithContext(ctx).Table("subscriptions").
		Where("id = ?", id).
		First(&res).Error; err != nil {
		logx.Create().Error().Msg(fmt.Sprintf("error: at repository when get detail subsriptions. Detail: %v", err))
		return res, errors.New("error: when get detail subsriptions")
	}

	return
}
