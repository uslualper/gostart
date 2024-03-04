package validate

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var (
	validate = validator.New()
)

func ValidateStruct(s interface{}) error {
	return validate.Struct(s)
}

func InitStruct(c *fiber.Ctx, s interface{}) error {
	if err := c.BodyParser(s); err != nil {
		return err
	}
	if isValid := ValidateStruct(s); isValid != nil {
		return isValid
	}
	return nil
}
