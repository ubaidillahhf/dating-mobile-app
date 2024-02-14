package domain

type PremiumPackage struct {
	Id             int64   `json:"id"`
	Name           string  `json:"name"`
	Description    string  `json:"description"`
	Price          float64 `json:"price"`
	DurationInDays int     `json:"durationInDays"`
	Repeat         int     `json:"repeat"`
}
