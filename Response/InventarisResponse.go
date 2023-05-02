package Response

import (
	"time"

	"gorm.io/gorm"
)

type InventarisResponse struct {
	ID           uint 	`json:"id" gorm:"primaryKey"`
	Nama_Barang string 	`json:"nama_barang"`
	Stock 		int 	`json:"stock"`
	Kategory 	string 	`json:"kategory"`
	Price 		int 	`json:"price"`
	Image 		string 	`json:"image"`
	CreatedAt time.Time	`json:"created_at"`
	UpdatedAt time.Time	`json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}