package test

import (
	"gostart/pkg/utils/validate"

	"github.com/gofiber/fiber/v2"
)

type Test struct {
	Message string `json:"message" validate:"required"`
}

func (t *Test) Init(c *fiber.Ctx) (*Test, error) {
	if err := validate.InitStruct(c, t); err != nil {
		return nil, err
	}
	return t, nil
}
