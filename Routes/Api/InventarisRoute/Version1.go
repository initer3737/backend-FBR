package InventarisRoute

import (
	"github.com/gofiber/fiber/v2"
	"github.com/initer3737/go-fiber-api/Controllers"
)


func Inventaris__ApiV1__Initialize(R *fiber.App){
	api:=R.Group("/api")
	V1:=api.Group("/v1")
	   Controller:=Controllers.NewInventarisController()
		V1.Get("/",Controller.Index)
		V1.Get("/:id",Controller.Show)
		V1.Post("/",Controller.Create)
		V1.Get("/inventaris/delete/:id",Controller.Destroy)
}