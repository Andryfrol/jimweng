package main

import (
	"fmt"

	"github.com/goPractice/memoServer/model"
	pagination "github.com/goPractice/memoServer/paginator"
	"github.com/goPractice/memoServer/router"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

// User 用户
type User struct {
	ID       int
	UserName string `gorm:"not null;size:100;unique"`
}

var db *gorm.DB

func init() {
	var err error
	db, err = model.InitDB()
	if err != nil {
		panic(err)
	}
}

func main() {

	var users []model.User

	pagination.Paging(&pagination.Param{
		DB:      db.Where("id > ?", 0),
		Page:    1,
		Limit:   3,
		OrderBy: []string{"id desc"},
		ShowSQL: true,
	}, &users)

	fmt.Println("users:", users)

	r := router.NewRouter()
	r.Run()
}
