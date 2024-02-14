package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/ubaidillahhf/dating-service/app/domain"
	"github.com/ubaidillahhf/dating-service/app/infra/utility/helper"
	logx "github.com/ubaidillahhf/dating-service/app/infra/utility/logger"
	"gorm.io/gorm"
)

type IPremiumPackageRepository interface {
	Get(ctx context.Context, meta domain.Meta) ([]domain.PremiumPackage, int64, error)
	Find(ctx context.Context, id int64) (domain.PremiumPackage, error)
}

func NewPremiumPackageRepository(db *gorm.DB) IPremiumPackageRepository {
	return &premiumPackageRepository{
		conn: db,
	}
}

type premiumPackageRepository struct {
	conn *gorm.DB
}

func (repo *premiumPackageRepository) Get(ctx context.Context, meta domain.Meta) (res []domain.PremiumPackage, total int64, err error) {
	q := repo.conn.WithContext(ctx).Table("premium_packages")

	if meta.Order != "" && meta.OrderBy != "" {
		q.Order(fmt.Sprintf("%s %s", meta.OrderBy, meta.Order))
	}

	if err := q.
		Count(&total).
		Scopes(helper.GormPaginate(meta.Skip, meta.Limit)).
		Find(&res).Error; err != nil {
		logx.Create().Error().Msg(fmt.Sprintf("error: at repository when get premium package. Detail: %v", err))
		return res, total, errors.New("error: at repository when get premium package list")
	}

	return res, total, nil
}

func (repo *premiumPackageRepository) Find(ctx context.Context, id int64) (res domain.PremiumPackage, err error) {

	if err := repo.conn.WithContext(ctx).Table("premium_packages").
		Where("id = ?", id).
		First(&res).Error; err != nil {
		logx.Create().Error().Msg(fmt.Sprintf("error: at repository when get detail premium_packages. Detail: %v", err))
		return res, errors.New("error: when get detail premium_packages")
	}

	return
}
