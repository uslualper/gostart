package main

import (
	"strings"

	"gostart/pkg/cache"
	"gostart/pkg/config"
	"gostart/pkg/db"
	"gostart/pkg/i18n"
	"gostart/pkg/router"
	log "gostart/pkg/utils/logger"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/template/html"
)

// @title API
// @version 1.0
// @termsOfService http://swagger.io/terms/

// @contact.name   Alper Uslu
// @contact.email  uslualper@outlook.com

// @BasePath /api
func main() {

	config.Load(".env")

	port := config.GetString("APP_PORT")

	app := fiber.New(fiber.Config{
		Immutable: true,
		Views:     html.New("./pkg/views", ".html"),
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			log.Instance().Error(log.System, log.Context(err.Error(), "main"), 0)
			return err
		},
	})

	allowOrigins := []string{}

	if config.GetString("APP_DEBUG") == "true" {
		allowOrigins = append(allowOrigins, "*")
	} else {
		allowOrigins = append(allowOrigins, config.GetString("APP_URL"))
	}

	app.Use(cors.New(cors.Config{
		AllowOrigins: strings.Join(allowOrigins, ","),
	}))

	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed,
	}))

	db.SetupDB()

	cache.SetupCache()

	router.SetupRoutes(app)

	i18n.SetupI18n()

	log.InitLogger(config.GetBool("APP_DEBUG"), log.MongoLogHandler)

	app.Listen(":" + port)
}
