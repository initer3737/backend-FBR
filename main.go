package main

import "github.com/gofiber/fiber/v2"

func main() {
    app := fiber.New()

    app.Get("/", func(c *fiber.Ctx) error {
        return c.JSON(fiber.Map{
			"name":"yotsusan_machi",
			"age":"secret",
			"hobby":"learning something news , watching some anime and analize game ui in every genre",
			"interest":"back end && front end",
		})
    })

    app.Listen(":3000")
}

