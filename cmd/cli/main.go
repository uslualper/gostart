package main

import (
	"fmt"
	"gostart/pkg/config"
	"log"
	"os"

	"github.com/urfave/cli"
)

var app = cli.NewApp()

func main() {

	config.Load(".env")

	info()
	commands()

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func info() {
	app.Name = config.GetString("APP_NAME")
	app.Usage = "cli app"
	app.Author = "uslualper"
	app.Version = "1.0.0"
}

func commands() {
	app.Commands = []cli.Command{
		{
			Name:    "test",
			Aliases: []string{"test"},
			Usage:   "Test",
			Action: func(c *cli.Context) {
				fmt.Println("Test Command")
			},
		},
	}
}
