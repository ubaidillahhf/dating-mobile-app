package domain

type Subscription struct {
	Id                int64  `json:"id"`
	UserId            string `json:"userId"`
	PremiumPackagesId int64  `json:"premiumPackagesId" validate:"required"`
	Status            string `json:"status"`
}

const (
	SubsActive      = "active"
	SubsGracePeriod = "grace_period"
	SubsEnd         = "end"
	SubsPending     = "pending"
	SubsCancel      = "cancel"
)
