package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

// Our User Struct
type User struct {
	gorm.Model
	Name  string
	Email string
}

// our initial migration function
func initialMigration() {
	db, err := gorm.Open("mysql", "test.db")
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&User{})
}

func main() {
	fmt.Println("Go ORM Tutorial")

	// Add the call to our new initialMigration function
	initialMigration()

}
