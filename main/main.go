package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/project104/Database"
	"github.com/project104/Routes"
)

func main() {

	Database.Connect()
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))
	Routes.Setup(app)

	app.Listen(":3000")

}
