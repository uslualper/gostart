package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"

	"go-start/pkg/config"
)

type Swagger struct{}

func (s *Swagger) SetupRoutes(r fiber.Router) {
	r.Get("/swagger/*", swagger.New(swagger.Config{
		DeepLinking:  false,
		DocExpansion: "none",
		OAuth: &swagger.OAuthConfig{
			AppName: config.Config("APP_NAME"),
		},
	}))
}
