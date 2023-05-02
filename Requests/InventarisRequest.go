package Requests


type InventarisRequest struct{
	Nama_Barang string 	`validate:"required" json:"nama_barang"`
	Stock 		int 	`validate:"required" json:"stock"`
	Kategory 	string 	`validate:"required" json:"kategory"`
	Price 		int 	`validate:"required" json:"price"`
	Image 		string 	`validate:"required" json:"image"`
}