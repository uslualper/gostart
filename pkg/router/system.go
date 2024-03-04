package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"github.com/gofiber/fiber/v2/middleware/monitor"

	"gostart/pkg/config"
)

type System struct{}

func (s *System) SetupRoutes(r fiber.Router) {

	systemMiddleware := r.Group("/system").Use(basicauth.New(basicauth.Config{
		Users: map[string]string{
			config.GetString("SYSTEM_USERNAME"): config.GetString("SYSTEM_PASSWORD"),
		},
	}))

	systemMiddleware.Get("/metrics", monitor.New(monitor.Config{Title: "MyService Metrics Page"}))
}
