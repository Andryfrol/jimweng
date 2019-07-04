package main

import (
	"fmt"
	"log"

	"github.com/goPractice/pkgmanagement/plugins/outputs"
	"github.com/goPractice/pkgmanagement/utils"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type MySQLClient struct {
}

type SQLConfig struct {
	DBName   string
	DBPort   string
	DBAddr   string
	User     string
	Password string
}

func (s *SQLConfig) openDB() (*gorm.DB, error) {
	connectionUrl := s.User + ":" + s.Password + "@tcp(" + s.DBAddr + ":" + s.DBPort + ")/" + s.DBName + "?charset=utf8&parseTime=True&loc=Local"
	if db, err := gorm.Open("mysql", connectionUrl); err != nil {
		return nil, err
	} else {
		db.AutoMigrate(&utils.PKGContent{})
		return db, nil
	}
}

func (s *SQLConfig) closeDB(db *gorm.DB) error {
	if err := db.Close(); err != nil {
		return err
	}
	return nil
}

func insertData(db *gorm.DB, points *[]*utils.PKGContent) error {
	for _, pt := range *points {
		if exists := db.NewRecord(pt); exists {
			fmt.Printf("%v\n", exists)
			fmt.Printf("%v\n", db.Find(&utils.PKGContent{Name: "123"}).Value)
			db.Find(&utils.PKGContent{Name: pt.Name}).Update(pt)
		} else {
			db.Create(pt)
		}
	}
	return nil
}

func (s *SQLConfig) Write(points *[]*utils.PKGContent) error {
	s.openDB()
	return nil
}

func init() {
	outputs.Add("mysql", func() utils.Output {
		return &SQLConfig{}
	})
}

var demo_pts = []*utils.PKGContent{
	&utils.PKGContent{Name: "jim", Parent: "jsasdfjsdf12adfasf3im", Synopsis: "jjimm", Href: "jimm"},
}

func main() {
	sc := &SQLConfig{
		DBName:   "pkg_lists",
		DBPort:   "3306",
		DBAddr:   "localhost",
		User:     "jim",
		Password: "password",
	}
	db, err := sc.openDB()
	if err != nil {
		log.Fatal(err)
	}
	err = insertData(db, &demo_pts)
	if err != nil {
		log.Fatal(err)
	}
	// db.Create(&utils.PKGContent{Name: "jim", Parent: "jjim", Synopsis: "jjimm", Href: "jimm"})
}
