package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"

	"gostart/pkg/config"
)

type Swagger struct{}

func (s *Swagger) SetupRoutes(r fiber.Router) {
	r.Get("/swagger/*", swagger.New(swagger.Config{
		DeepLinking:  false,
		DocExpansion: "none",
		OAuth: &swagger.OAuthConfig{
			AppName: config.GetString("APP_NAME"),
		},
	}))
}
