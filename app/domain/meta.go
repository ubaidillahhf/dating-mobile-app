package domain

type Meta struct {
	Page    int64  `json:"page" `
	PerPage int64  `json:"perPage"`
	Order   string `json:"order" `
	OrderBy string `json:"orderBy"`
	Skip    int
	Limit   int
}
