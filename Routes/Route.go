package Routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/initer3737/go-fiber-api/Controllers"
	"github.com/initer3737/go-fiber-api/Controllers/HalloController"
)

func Route__initialize(R * fiber.App){
			//keuntungan menggunakkan methode dibawah adalah tidak boros import export cukup panggil package controllers tapi harus membuat variable yang menampung data dari controller yang akan dipanggil
	R.Get("/animes",Controllers.AnimeControllers.AnimeChars)
		AisatsuController:=Controllers.NewAisatsuController()

	R.Get("/",AisatsuController.Jikoshoukai)
				//keuntungan menggunakkan methode dibawah adalah tidak perlu inisialisasi variable yang menampung sebuah method yang mengembalikan struct (secara sederhana struct adalah sebuah data yang menghimpun method-methods yang telah didaftarkan menjadi milik dari struct atau class jika anda mengetahui konsep oop {object oriented programming})
	R.Get("/hello",HalloController.Hello)

		//baseurl
		R.Get("/baseurl",func(c *fiber.Ctx) error {
			 return c.JSON(fiber.Map{
				"url":c.BaseURL(),
			 })
		});
}