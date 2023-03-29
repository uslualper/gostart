package handlers

import (
	"github.com/gofiber/fiber/v2"

	"go-start/pkg/cache"
	payloadScheme "go-start/pkg/schema/payload/v1/test"
	responseScheme "go-start/pkg/schema/response/v1/test"
	"go-start/pkg/utils/validate"
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

	cache.Cache().Set("test", testValid.Message)
	cacheValue, cacheErr := cache.Cache().Get("test")

	if cacheErr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(cacheErr)
	}

	testResponse := responseScheme.Test{
		Message: testValid.Message,
		Cache:   cacheValue,
	}

	return c.JSON(testResponse)
}
