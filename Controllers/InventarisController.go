package Controllers

import (
	//"github.com/go-playground/validator/v10"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/initer3737/go-fiber-api/Response"
)

type InventarisController struct{}

func NewInventarisController() InventarisController{
	return InventarisController{}
}

func (Inventaris *InventarisController)Index(c *fiber.Ctx)error{
		//validate:=validator.new()
		Response:= Response.InventarisResponse{
			ID: 1,
			Nama_Barang: "buah",
			Stock: 12,
			Kategory: "makanan",
			Price: 12000,
			Image: c.BaseURL()+"/images/buahbuahan.jpg",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			//DeletedAt: time.Now(),
		}
	return c.JSON(fiber.Map{
		"data":Response,
	})
}

func (Inventaris *InventarisController)Show(c *fiber.Ctx)error{
			id:=c.Params("id")
		//validate:=validator.new()
	return c.JSON(fiber.Map{
		"data":"id is "+id,
	})
}


func (Inventaris *InventarisController)Create(c *fiber.Ctx)error{
	//validate:=validator.new()
	return c.JSON(fiber.Map{
		"message":"data has been created successfully",
	})
}

func (Inventaris *InventarisController)Update(c *fiber.Ctx)error{
		//id:=c.Params("id")
		//validate:=validator.new()
		message:="updated successfully"
	return c.JSON(fiber.Map{
		"message":message,
	})
}

func (Inventaris *InventarisController)Destroy(c *fiber.Ctx)error{
		//id:=c.Params("id")

	return c.JSON(fiber.Map{
		"data":"deleted successfully",
	})
}