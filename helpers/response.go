package helpers

import (
	"github.com/gofiber/fiber/v3"
)

type Response struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

type ResponseData struct {
	Code    int         `json:"code"`
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    any         `json:"data"`
	Meta    *Pagination `json:"meta,omitempty"`
}

func (r *Response) SendResponse(c fiber.Ctx) error {
	c.Status(r.Code)
	return c.JSON(r)
}

func (r *ResponseData) SendResponse(c fiber.Ctx) error {
	c.Status(r.Code)
	return c.JSON(r)
}
