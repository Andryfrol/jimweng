package main

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Our User Struct
type User struct {
	gorm.Model
	Name  string
	Email string
}

// our initial migration function
func init() {
	db, err := gorm.Open("mysql", "jim:password@tcp(localhost:3306)/demo_db?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.AutoMigrate(&User{})
}

func main() {
	fmt.Println("Go ORM Tutorial")

}
