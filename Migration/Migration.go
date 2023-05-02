package Migration

import (
	"fmt"
	"log"

	"github.com/initer3737/go-fiber-api/Database"
	"github.com/initer3737/go-fiber-api/Entity/Inventaris"
)

func RunMigrations(){
	err:=Database.DB.AutoMigrate(&Inventaris.InventarisEntity{})
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println("migrasi suksess dilakukan waifumu pasti bangga!")
}