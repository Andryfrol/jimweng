package main

import (
	"flag"

	"github.com/goPractice/memoServer/model"
	"github.com/goPractice/memoServer/router"
	_ "github.com/mattn/go-sqlite3"
)

// User 用户
type User struct {
	ID       int
	UserName string `gorm:"not null;size:100;unique"`
}

func init() {
	_, err := model.InitDB()
	if err != nil {
		panic(err)
	}
}

var fSample = flag.Bool("sample", false, "show debug mode help")

func main() {

	flag.Parse()
	if *fSample {
		model.InsertSampleData()
	}

	r := router.NewRouter()
	r.Run()
}
