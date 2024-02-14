package domain

type Swipe struct {
	Id         int64  `json:"id"`
	SenderId   string `json:"senderId"`
	ReceiverId string `json:"receiverId" validate:"required"`
	Direction  string `json:"direction" validate:"required,oneof=left right"`
}
