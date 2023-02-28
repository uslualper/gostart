package handlers

import (
	"github.com/gofiber/fiber/v2"

	payloadScheme "go-start/schema/payload/v1/test"
	responseScheme "go-start/schema/response/v1/test"
	"go-start/utils/validate"
)

type Test struct{}

// Test
// @Summary Test
// @Description get test data
// @Tags Test
// @Accept json
// @Produce json
// @Param payload body payloadScheme.Test true "Test Payload"
// @Response 200 {object} responseScheme.Test
// @Router /v1/test [post]
func (t *Test) Test(c *fiber.Ctx) error {
	testValid := new(payloadScheme.Test)
	c.BodyParser(testValid)
	if err := validate.ValidateStruct(testValid); err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(err)
	}

	testResponse := responseScheme.Test{
		Message: testValid.Message + " (Test)",
	}

	return c.JSON(testResponse)
}
