package model

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	pagination "github.com/goPractice/memoServer/paginator"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type User struct {
	ID       int
	UserName string `gorm:"not full;size:100;unique"`
}

/* insert sample data */
func InsertSampleData() {
	db.Create(User{ID: 1, UserName: "biezhi"})
	db.Create(User{ID: 2, UserName: "rose"})
	db.Create(User{ID: 3, UserName: "jack"})
	db.Create(User{ID: 4, UserName: "lili"})
	db.Create(User{ID: 5, UserName: "bob"})
	db.Create(User{ID: 6, UserName: "tom"})
	db.Create(User{ID: 7, UserName: "anny"})
	db.Create(User{ID: 8, UserName: "wat"})
	log.Println("Insert OK!")
}

func InitDB() (*gorm.DB, error) {
	var err error
	if db, err = gorm.Open("sqlite3", "example.db"); err != nil {
		return nil, err
	}
	db.AutoMigrate(&User{})

	return db, err

}

func ReturnPageInfo(ctx *gin.Context) {
	var users []User
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "3"))

	paginator := pagination.Paging(&pagination.Param{
		DB:      db,
		Page:    page,
		Limit:   limit,
		OrderBy: []string{"id desc"},
		ShowSQL: true,
	}, &users)
	ctx.JSON(200, paginator)
}
