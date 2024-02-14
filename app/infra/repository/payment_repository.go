package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/ubaidillahhf/dating-service/app/domain"
	logx "github.com/ubaidillahhf/dating-service/app/infra/utility/logger"
	"gorm.io/gorm"
)

type IPaymentRepository interface {
	Create(ctx context.Context, newData domain.Payment) (domain.Payment, error)
	Update(ctx context.Context, newData domain.Payment) (bool, error)
	UpdateTx(ctx context.Context, tx *gorm.DB, newData domain.Payment) (bool, error)
	ValidateCallback(ctx context.Context, id int64, userId, refId string) bool
	CreateTx(ctx context.Context, tx *gorm.DB, newData domain.Payment) (domain.Payment, error)
}

func NewPaymentRepository(db *gorm.DB) IPaymentRepository {
	return &paymentRepository{
		conn: db,
	}
}

type paymentRepository struct {
	conn *gorm.DB
}

func (repo *paymentRepository) Create(ctx context.Context, newData domain.Payment) (res domain.Payment, err error) {

	if err := repo.conn.WithContext(ctx).Table("payments").Create(&newData).Error; err != nil {
		logx.Create().Error().Msg(fmt.Sprintf("error: at repository when create payments. Detail: %v", err))
		return res, errors.New("error: when create payments")
	}

	return newData, nil
}

func (repo *paymentRepository) CreateTx(ctx context.Context, tx *gorm.DB, newData domain.Payment) (res domain.Payment, err error) {

	if err := tx.WithContext(ctx).Table("payments").Create(&newData).Error; err != nil {
		logx.Create().Error().Msg(fmt.Sprintf("error: at repository when create payments. Detail: %v", err))
		return res, errors.New("error: when create payments")
	}

	return newData, nil
}

func (repo *paymentRepository) Update(ctx context.Context, newData domain.Payment) (res bool, err error) {
	newUpdate := make(map[string]interface{})

	if newData.Id == 0 {
		return res, errors.New("error: at repo payment when update. Empty id update")
	}

	if newData.Status != "" {
		newUpdate["status"] = newData.Status
	}

	if err := repo.conn.WithContext(ctx).Table("payments").
		Where("id = ?", newData.Id).
		Updates(newUpdate).Error; err != nil {
		logx.Create().Error().Msg(fmt.Sprintf("error: at repository when update payments. Detail: %v", err))
		return false, errors.New("error: when update payments")
	}

	return true, nil
}

func (repo *paymentRepository) UpdateTx(ctx context.Context, tx *gorm.DB, newData domain.Payment) (res bool, err error) {
	newUpdate := make(map[string]interface{})

	if newData.Id == 0 {
		return res, errors.New("error: at repo payment when update. Empty id update")
	}

	if newData.Status != "" {
		newUpdate["status"] = newData.Status
	}

	if err := tx.WithContext(ctx).Table("payments").
		Where("id = ?", newData.Id).
		Updates(newUpdate).Error; err != nil {
		logx.Create().Error().Msg(fmt.Sprintf("error: at repository when update payments. Detail: %v", err))
		return false, errors.New("error: when update payments")
	}

	return true, nil
}

func (repo *paymentRepository) ValidateCallback(ctx context.Context, id int64, userId, refId string) bool {
	res := domain.Payment{}

	if err := repo.conn.WithContext(ctx).Table("payments").
		Where("id = ? AND user_id = ? AND ref_id = ?", id, userId, refId).
		First(&res).Error; err != nil {
		logx.Create().Error().Msg(fmt.Sprintf("error: at repository when ValidateCallback. Detail: %v", err))
		return false
	}

	return res.Id != 0
}
