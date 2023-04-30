package Routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/initer3737/go-fiber-api/Controllers/HalloController"
	"github.com/initer3737/go-fiber-api/Controllers/AisatsuController"
)

func Route__initialize(R * fiber.App){
	R.Get("/",AisatsuController.Jikoshoukai)
	R.Get("/hello",HalloController.Hello)
}