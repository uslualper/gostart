package main

import (
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/template/html"

	"go-start/config"
	db "go-start/db"
	"go-start/i18n"
	"go-start/router"
)

// @title API
// @version 1.0
// @termsOfService http://swagger.io/terms/

// @contact.name   Alper Uslu
// @contact.email  uslualper@outlook.com

// @BasePath /api
func main() {

	port := config.Config("APP_PORT")

	app := fiber.New(fiber.Config{
		Immutable: true,
		Views:     html.New("./views", ".html"),
	})

	allowOrigins := []string{}

	if config.Config("APP_DEBUG") == "true" {
		allowOrigins = append(allowOrigins, "*")
	} else {
		allowOrigins = append(allowOrigins, config.Config("APP_URL"))
	}

	app.Use(cors.New(cors.Config{
		AllowOrigins: strings.Join(allowOrigins, ","),
	}))

	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed,
	}))

	app.Use(limiter.New(limiter.Config{
		Max:                    100,
		Expiration:             1 * time.Minute,
		SkipSuccessfulRequests: true,
	}))

	db.SetupDB()

	router.SetupRoutes(app)

	i18n.SetupI18n()

	app.Listen(":" + port)
}
