package Controllers

import (
	"github.com/gofiber/fiber/v2"
)

//ibaratkan seperti class jika di oop
type AnimeController struct{

}
	func NewAnimeController() AnimeController{
		return AnimeController{} //mengembalikan struct AnimeController
	}

		//method Jikoshoukai milik dari struct NewAisatsuController
func (Anime *AnimeController)Jikoshoukai(c *fiber.Ctx) error{
	return c.JSON(fiber.Map{
		"name":"yotsusan_machi",
		"age":"secret",
		"hobby":"learning something news , watching some anime and analize game ui in every genre",
		"interest":"back end && front end",
	})
}

func (Anime *AnimeController)AnimeChars(c *fiber.Ctx)error{
		//type data slice dan string itu sama yang membedakkan adalah slice tidak perlu di tentukan panjang dari indexnya []string{} tapi array perlu [3]string{}
	return c.JSON(fiber.Map{
		"animes":[3]string{
			"kimetsu no yaiba",
			"spongebob squarpanties",
			"penjahit bulan",
		},
	})
}

//inisialisasi object
var AnimeControllers=NewAnimeController()