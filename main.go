package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/initer3737/go-fiber-api/Routes"
	"github.com/initer3737/go-fiber-api/Routes/Api/Version1"
)

func main() {
    app := fiber.New()
		// version 1 : grouping
	Version1.Route__ApiV1__initialize(app)
		//default
	Routes.Route__initialize(app)
    app.Listen(":3000")
}

