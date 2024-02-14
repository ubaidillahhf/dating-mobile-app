package presenter

import (
	"github.com/ubaidillahhf/dating-service/app/domain"
	"github.com/ubaidillahhf/dating-service/app/infra/utility/helper"
)

type FindMatchPresenter struct {
	Id       string `json:"id"`
	Fullname string `json:"fullname"`
	Image    string `json:"image"`
	Gender   string `json:"gender"`
	Age      string `json:"age"`
}

func FindMatchTransform(data []domain.User) (res []FindMatchPresenter) {

	for _, val := range data {
		res = append(res, FindMatchPresenter{
			Id:       val.Id,
			Fullname: val.Fullname,
			Image:    val.Image,
			Gender:   val.Gender,
			Age:      helper.GetAgeFromDob(val.Dob),
		})
	}

	return
}
