package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/initer3737/go-fiber-api/Routes"
)

func main() {
    app := fiber.New()

    Routes.Route__initialize(app)

    app.Listen(":3000")
}

