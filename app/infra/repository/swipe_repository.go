package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/jinzhu/now"
	"github.com/ubaidillahhf/dating-service/app/domain"
	logx "github.com/ubaidillahhf/dating-service/app/infra/utility/logger"
	"gorm.io/gorm"
)

type ISwipeRepository interface {
	Create(ctx context.Context, newData domain.Swipe) (domain.Swipe, error)
	CountBySenderId(ctx context.Context, senderId string, onlyToday bool) (int64, error)
}

func NewSwipeRepository(db *gorm.DB) ISwipeRepository {
	return &swipeRepository{
		conn: db,
	}
}

type swipeRepository struct {
	conn *gorm.DB
}

func (repo *swipeRepository) Create(ctx context.Context, newData domain.Swipe) (res domain.Swipe, err error) {

	if err := repo.conn.WithContext(ctx).Table("swipes").Create(&newData).Error; err != nil {

		logx.Create().Error().Msg(fmt.Sprintf("error: at repository when create user. Detail: %v", err))

		return res, fmt.Errorf("error: at repository when create user. Detail: %v", err)
	}

	return newData, nil
}

func (repo *swipeRepository) CountBySenderId(ctx context.Context, senderId string, onlyToday bool) (total int64, err error) {
	q := repo.conn.WithContext(ctx).Table("swipes")

	if onlyToday {
		startAt := now.BeginningOfDay().UTC()
		endAt := now.EndOfDay().UTC()

		q.Where("created_at >= ? AND created_at <= ?", startAt, endAt)
	}

	if err := q.Where("sender_id = ?", senderId).
		Count(&total).Error; err != nil {
		logx.Create().Error().Msg(fmt.Sprintf("error: at repository when count today total swipe. Detail: %v", err))
		return total, errors.New("error: at repository when count today total swipe")
	}

	return
}
