package mysql

import (
	"fmt"

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
		// initialize DB settings for connection max nums 10 and keep-alive
		if dbconfig := db.DB(); dbconfig != nil {
			dbconfig.SetMaxOpenConns(10)
			dbconfig.SetMaxIdleConns(0)
			dbconfig.SetConnMaxLifetime(-1)
		}
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
		// 抓取primary_key; 使用primary_key來判斷是否有建過該record，有的話更新。沒有則創建
		if err := db.Find(&utils.PKGContent{Name: pt.Name}).Update(&pt).Error; err != nil {
			if err := db.Create(&pt).Error; err != nil {
				fmt.Printf("Some error happened with msg : %v\n", err)
			}
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

// func main() {
// 	sc := &SQLConfig{
// 		DBName:   "pkg_lists",
// 		DBPort:   "3306",
// 		DBAddr:   "localhost",
// 		User:     "jim",
// 		Password: "password",
// 	}
// 	db, err := sc.openDB()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	err = insertData(db, &demo_pts)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }
