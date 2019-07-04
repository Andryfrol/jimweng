package mysql

import (
	"testing"

	"github.com/goPractice/pkgmanagement/utils"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

var demo_pts = []*utils.PKGContent{
	&utils.PKGContent{Name: "name1", Parent: "parent1", Synopsis: "synopsis1", Href: "href1"},
	&utils.PKGContent{Name: "name2", Parent: "parent2", Synopsis: "synopsis2", Href: "href2"},
	&utils.PKGContent{Name: "name3", Parent: "parent3", Synopsis: "synopsis3", Href: "href3"},
	&utils.PKGContent{Name: "name1", Parent: "parentu1", Synopsis: "synopsisu1", Href: "hrefu1"},
}

func TestSomething(t *testing.T) {
	db, mock, _ := sqlmock.New()
	models.Db, _ = gorm.Open("mysql", db)
}

// func TestSomething(t *testing.T) {
// 	sc := &SQLConfig{
// 		DBName:   "pkg_lists",
// 		DBPort:   "3306",
// 		DBAddr:   "127.0.0.1",
// 		User:     "jim",
// 		Password: "password",
// 	}
// 	db, err := sc.openDB()
// 	assert.Nil(t, err)
// 	err = sc.closeDB(db)
// 	assert.Nil(t, err)
// }

// func TestInsertData(t *testing.T) {
// 	sc := &SQLConfig{
// 		DBName:   "pkg_lists",
// 		DBPort:   "3306",
// 		DBAddr:   "127.0.0.1",
// 		User:     "jim",
// 		Password: "password",
// 	}
// 	db, err := sc.openDB()
// 	assert.Nil(t, err)
// 	// assert.Equal(t, "", fmt.Sprintf("%v\n", db))
// 	err = insertData(db, &demo_pts)
// 	assert.Nil(t, err)

// 	err = sc.closeDB(db)
// 	assert.Nil(t, err)
// }
