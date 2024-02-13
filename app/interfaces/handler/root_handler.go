package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ubaidillahhf/dating-service/app/infra/presenter"
)

type (
	Rootmessage struct {
		ServiceName string `json:"serviceName"`
		Version     string `json:"version"`
		Author      string `json:"author"`
		Year        string `json:"year"`
	}
)

func GetTopRoute(c *fiber.Ctx) error {
	var data Rootmessage

	data.ServiceName = "Dating Service"
	data.Version = "1.0"
	data.Author = "Ubaidillah Hakim Fadly"
	data.Year = "2024"

	return c.JSON(presenter.Success("Success", data, nil))
}
