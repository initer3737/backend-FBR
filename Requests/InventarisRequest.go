package requests


type InventarisRequest struct{
	Nama_Barang string 	`validate:"required"`
	Stock 		string 	`validate:"required"`
	Kategory 	string 	`validate:"required"`
	Price 		int 	`validate:"required"`
	Image 		string 	`validate:"required"`
}