package Version1
import (
	"github.com/gofiber/fiber/v2"
	"github.com/initer3737/go-fiber-api/Controllers/HalloController"
	"github.com/initer3737/go-fiber-api/Controllers/AisatsuController"
)

func Route__ApiV1__initialize(R * fiber.App){
		v1:=R.Group("/api/v1")
	v1.Get("/",AisatsuController.Jikoshoukai)
	v1.Get("/hello",HalloController.Hello)
}