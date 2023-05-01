package Database

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
	 var DB *gorm.DB 
func Databse__initialize(){
	var err error
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	const Mysql = "root:@tcp(127.0.0.1:3306)/todo_list?charset=utf8mb4&parseTime=True&loc=Local"

		dsn:=Mysql
			db,err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil{
			//panic("koneksi gagal ")
			log.Fatal(err)
		}
		fmt.Println("koneksi berhasil!! waifumu bangga mas")
		DB=db
}