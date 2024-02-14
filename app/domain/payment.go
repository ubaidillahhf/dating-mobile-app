package domain

type Payment struct {
	Id         int64   `json:"id"`
	UserId     string  `json:"userId"`
	RefContext string  `json:"refContext"`
	RefId      string  `json:"refId"`
	Amount     float64 `json:"amount"`
	ExternalId string  `json:"externalId"`
	Method     string  `json:"method"`
	Status     string  `json:"status"`
}

type PaymentCallbackRequest struct {
	Id         int64  `json:"paymentId" validate:"required"`
	UserId     string `json:"userId" validate:"required"`
	RefContext string `json:"refContext" validate:"required,oneof=subscription"`
	RefId      string `json:"refId" validate:"required"`
	Status     string `json:"status" validate:"required,oneof=success failed"`
}

const (
	PaymentWaiting = "waiting"
	PaymentSuccess = "success"
	PaymentFailed  = "failed"
)

const (
	RefContextSubs = "subscription"
)
