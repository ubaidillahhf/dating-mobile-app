package domain

import "time"

type Subscription struct {
	Id                int64     `json:"id"`
	UserId            string    `json:"userId"`
	PremiumPackagesId int64     `json:"premiumPackagesId" validate:"required"`
	Status            string    `json:"status"`
	EndAt             time.Time `json:"endAt"`
}

type SubscriptionResponse struct {
	Id                int64     `json:"id"`
	UserId            string    `json:"userId"`
	PremiumPackagesId int64     `json:"premiumPackagesId"`
	Status            string    `json:"status"`
	EndAt             time.Time `json:"endAt"`
	PaymentId         int64     `json:"paymentId"`
}

const (
	SubsActive      = "active"
	SubsGracePeriod = "grace_period"
	SubsEnd         = "end"
	SubsPending     = "pending"
	SubsCancel      = "cancel"
)
