package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/initer3737/go-fiber-api/Database"
	"github.com/initer3737/go-fiber-api/Migration"
	"github.com/initer3737/go-fiber-api/Routes/Api/InventarisRoute"
	//"github.com/initer3737/go-fiber-api/Routes"
	//"github.com/initer3737/go-fiber-api/Routes/Api/Version1"
)

func main() {
		//debeh koneksi
		Database.Databse__initialize()
		Migration.RunMigrations()
    app := fiber.New()
		// version 1 : grouping
	//Version1.Route__ApiV1__initialize(app)
		//default
	//Routes.Route__initialize(app)

		//inventaris route
	InventarisRoute.Inventaris__ApiV1__Initialize(app)

    app.Listen(":3000")
}

