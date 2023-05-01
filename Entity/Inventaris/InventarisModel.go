package Inventaris

import (
	"gorm.io/gorm"
)

type InventarisModel struct{
	gorm.Model
	Nama_Barang string `json:"nama_barang"`
	Stock 		string `json:"stock"`
	Kategory 	string `json:"kategory"`
	Price 		int `json:"price"`
	Image 		string `json:"image"`
}