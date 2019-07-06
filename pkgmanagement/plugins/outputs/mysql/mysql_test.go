package mysql

import (
	"log"
	"os"

	"github.com/goPractice/pkgmanagement/utils"
	_ "github.com/mattn/go-sqlite3"
)

var demo_pts = []*utils.PKGContent{
	&utils.PKGContent{Name: "name1", Parent: "parent1", Synopsis: "synopsis1", Href: "href1"},
	&utils.PKGContent{Name: "name2", Parent: "parent2", Synopsis: "synopsis2", Href: "href2"},
	&utils.PKGContent{Name: "name3", Parent: "parent3", Synopsis: "synopsis3", Href: "href3"},
	&utils.PKGContent{Name: "name1", Parent: "parentu1", Synopsis: "synopsisu1", Href: "hrefu1"},
}

func init() {
	clearTestEnv()
}

func clearTestEnv() {
	log.Println("Clear test env")
	if _, err := os.Stat("/tmp/gorm.db"); !os.IsNotExist(err) {
		log.Println("Remove Origin gorm.db Files")
		os.Remove("/tmp/gorm.db")
	}
}

// // Set mock DB env
// func mockTestEnv() (OPDB, error) {
// 	var dbc = DBConfig{}
// 	dbc.DBUri = "/tmp/gorm.db"
// 	dbc.DBType = "sqlite3"
// 	opdb, err := dbc.NewDBConnection()
// 	return opdb, err
// }
