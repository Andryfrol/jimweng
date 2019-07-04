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
	Name  string `gorm:"primary_key"`
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
	create()
	deleteData()
}

// 實做CRUD
// Create
func create() {
	db, err := gorm.Open("mysql", "jim:password@tcp(localhost:3306)/demo_db?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	db.Create(&User{Name: "L1212", Email: "123"})
}

// Read
func read() {
	db, err := gorm.Open("mysql", "jim:password@tcp(localhost:3306)/demo_db?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	fmt.Printf("The value of db.value is %v\n", db.First(&User{Name: "L1212", Email: "123"}).Value)
}

// Update ... 更新相當於Read以後在把Read的資料改成新的資料；notes:在gorm裡面，更新以後也會更新updated_at的時間
func update() {
	db, err := gorm.Open("mysql", "jim:password@tcp(localhost:3306)/demo_db?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	fmt.Printf("The value of db.value is %v\n", db.First(&User{Name: "L1212", Email: "123"}).Update(&User{Name: "Jim", Email: "JimTest@example.com"}).Value)
}

// Delete ... 因為delete已經有預設方法，這邊改用deleteData來宣告該函數；notes:在gorm裡面刪除不是代表從db完全移除。而是去更改deleted_at的時間
func deleteData() {
	db, err := gorm.Open("mysql", "jim:password@tcp(localhost:3306)/demo_db?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	fmt.Printf("The value of db.value is %v\n", db.Delete(&User{Name: "L1212", Email: "123"}))
}
