package Controllers

import (
	"github.com/gofiber/fiber/v2"
)

//ibaratkan seperti class jika di oop
type AisatsuController struct{

}
	func NewAisatsuController() AisatsuController{
		return AisatsuController{} //mengembalikan struct aisatsucontroller
	}

		//method Jikoshoukai milik dari struct NewAisatsuController
func (Aisatsu *AisatsuController)Jikoshoukai(c *fiber.Ctx) error{
	return c.JSON(fiber.Map{
		"name":"yotsusan_machi",
		"age":"secret",
		"hobby":"learning something news , watching some anime and analize game ui in every genre",
		"interest":"back end && front end",
	})
}