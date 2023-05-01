package Version1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/initer3737/go-fiber-api/Controllers"
	"github.com/initer3737/go-fiber-api/Controllers/HalloController"
)

func Route__ApiV1__initialize(R * fiber.App){
		v1:=R.Group("/api/v1")
		AisatsuController:=Controllers.NewAisatsuController()
	v1.Get("/",AisatsuController.Jikoshoukai)
	v1.Get("/hello",HalloController.Hello)
}