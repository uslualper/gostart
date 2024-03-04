package handlers

import (
	"gostart/pkg/i18n"

	"github.com/gofiber/fiber/v2"

	payloadScheme "gostart/pkg/schema/payload/v1/test"
	responseScheme "gostart/pkg/schema/response/v1/test"
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
	testValid, err := (&payloadScheme.Test{}).Init(c)
	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(err)
	}

	testResponse := responseScheme.Test{
		Message: testValid.Message,
		Translate: map[string]string{
			"en": i18n.Translate("en", "welcome"),
			"tr": i18n.Translate("tr", "welcome"),
		},
	}

	return c.JSON(testResponse)
}
