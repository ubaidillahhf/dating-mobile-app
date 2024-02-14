package repository

import (
	"context"
	"fmt"

	"github.com/ubaidillahhf/dating-service/app/domain"
	logx "github.com/ubaidillahhf/dating-service/app/infra/utility/logger"
	"gorm.io/gorm"
)

type ISwipeRepository interface {
	Insert(ctx context.Context, newData domain.Swipe) (domain.Swipe, error)
}

func NewSwipeRepository(db *gorm.DB) ISwipeRepository {
	return &swipeRepository{
		conn: db,
	}
}

type swipeRepository struct {
	conn *gorm.DB
}

func (repo *swipeRepository) Insert(ctx context.Context, newData domain.Swipe) (res domain.Swipe, err error) {

	if err := repo.conn.WithContext(ctx).Table("swipes").Create(&newData).Error; err != nil {

		logx.Create().Error().Msg(fmt.Sprintf("error: at repository when create user. Detail: %v", err))

		return res, fmt.Errorf("error: at repository when create user. Detail: %v", err)
	}

	return newData, nil
}
