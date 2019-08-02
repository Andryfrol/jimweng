package model

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	pagination "github.com/goPractice/memoServer/paginator"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

// meomo column: Title; Description; Start Date; End Date; Category[Pesonal, Business, Others]
type MemoList struct {
	gorm.Model
	Title       string `gorm:"not full;size:50"`   // 50 chars
	Description string `gorm:"not full;size:1000"` // 1000 chars
	Category    int    // 0[Others], 1[Personal], 2[Business]
}

/* insert sample data */
func InsertSampleData() {
	db.AutoMigrate(&MemoList{})

	db.Create(&MemoList{Title: "TestTitle1", Description: "1001 words included", Category: 0})
	db.Create(&MemoList{Title: "TestTitle2", Description: "1002 words included", Category: 0})
	db.Create(&MemoList{Title: "TestTitle3", Description: "1003 words included", Category: 0})
	db.Create(&MemoList{Title: "TestTitle4", Description: "1004 words included", Category: 0})
	db.Create(&MemoList{Title: "TestTitle5", Description: "1005 words included", Category: 0})
	db.Create(&MemoList{Title: "TestTitle6", Description: "1006 words included", Category: 0})
	db.Create(&MemoList{Title: "TestTitle7", Description: "1007 words included", Category: 0})
	db.Create(&MemoList{Title: "TestTitle8", Description: "1008 words included", Category: 0})
	db.Create(&MemoList{Title: "TestTitle9", Description: "1009 words included", Category: 0})
	log.Println("Insert OK!")
}

func InitDB() (*gorm.DB, error) {
	var err error
	if db, err = gorm.Open("sqlite3", "example.db"); err != nil {
		return nil, err
	}
	db.AutoMigrate(&MemoList{})

	return db, err

}

func ReturnPageInfo(ctx *gin.Context) {
	var memoss []MemoList
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "3"))

	paginator := pagination.Paging(&pagination.Param{
		DB:      db,
		Page:    page,
		Limit:   limit,
		OrderBy: []string{"id"},
		ShowSQL: true,
	}, &memoss)
	ctx.JSON(200, paginator)
}
