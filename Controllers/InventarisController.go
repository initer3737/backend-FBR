package Controllers

import (
	"time"

	"github.com/go-playground/validator/v10"

	"github.com/gofiber/fiber/v2"
	"github.com/initer3737/go-fiber-api/Database"
	"github.com/initer3737/go-fiber-api/Entity/Inventaris"
	"github.com/initer3737/go-fiber-api/Requests"
	//"github.com/initer3737/go-fiber-api/Response"
)

type InventarisController struct{}

func NewInventarisController() InventarisController{
	return InventarisController{}
}
			//entities
	var InventarisEntity []Inventaris.InventarisEntity

func (InventarisController *InventarisController)Index(c *fiber.Ctx)error{
	data:=Database.DB.Find(&InventarisEntity)
		if data.Error != nil{
			return c.JSON(fiber.Map{
				"message":"data not found",
				"status":"404",
			})
		}
		//validate:=validator.new()
		// Response:= Response.InventarisResponse{
		// 	ID: 1,
		// 	Nama_Barang: "buah",
		// 	Stock: 12,
		// 	Kategory: "makanan",
		// 	Price: 12000,
		// 	Image: c.BaseURL()+"/images/buahbuahan.jpg",
		// 	CreatedAt: time.Now(),
		// 	UpdatedAt: time.Now(),
		// 	//DeletedAt: time.Now(),
		// }
	return c.JSON(fiber.Map{
		"data":&InventarisEntity,
	})
}

func (InventarisController *InventarisController)Show(c *fiber.Ctx)error{
			id:=c.Params("id")
		//validate:=validator.new()
			query:=Database.DB.First(&InventarisEntity,"id=?",id)
		if query.Error != nil{
			return c.JSON(fiber.Map{
				"message":"data not found",
			})
		}
	return c.JSON(fiber.Map{
		"data":&InventarisEntity,
	})
}


func (InventarisController *InventarisController)Create(c *fiber.Ctx)error{
	p:=new(Requests.InventarisRequest);
	if err:=c.BodyParser(p);err != nil {
		return err
	}
		// VALIDASI REQUEST
		validate := validator.New()
		errValidate := validate.Struct(p)
			if errValidate != nil {
				return c.Status(400).JSON(fiber.Map{
					"message": "failed",
					"error":   errValidate.Error(),
				})
		}
			createdInventaris :=Inventaris.InventarisEntity{
				Nama_Barang: p.Nama_Barang,
				Stock: p.Stock,
				Kategory: p.Kategory,
				Price:p.Price,
				Image: p.Image,
				CreatedAt: time.Now(),
			}


		created:=Database.DB.Create(&createdInventaris)
		
		if created.Error != nil {
			return c.JSON(fiber.Map{
					"message":"data fail to create",
				})
		}
	return c.JSON(fiber.Map{
		"message":"data has been created successfully",
	})
}

func (InventarisController *InventarisController)Update(c *fiber.Ctx)error{
		//id:=c.Params("id")
		request:=new(Requests.InventarisRequest)
		body:=c.BodyParser(request)
		if body != nil{
			return body
		}
		validate:=validator.New()
		if errvalidate:=validate.Struct(request); errvalidate!= nil{
			return c.Status(400).JSON(fiber.Map{
				"message": "failed",
				"error":   errvalidate.Error(),
			})
		}
		//updating
		// updating:=&InventarisEntity{
		// 	request.Nama_Barang:,
		// }

		message:="updated successfully"
		// updatingdata:=Database.DB.Model(&InventarisEntity).Where("id=?",id).Updates(updaupdating)
	return c.JSON(fiber.Map{
		"message":message,
	})
}

func (InventarisController *InventarisController)Destroy(c *fiber.Ctx)error{
		id:=c.Params("id")
		query:=Database.DB.Delete(&InventarisEntity,"id=?",id)

		if query.Error !=nil{
			c.JSON(fiber.Map{
				"message":"failed to delete!!",
			})
		}
	return c.JSON(fiber.Map{
		"data":"deleted successfully",
	})
}