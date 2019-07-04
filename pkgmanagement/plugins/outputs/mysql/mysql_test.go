package mysql

import (
	"testing"

	"github.com/goPractice/pkgmanagement/utils"
	"github.com/stretchr/testify/assert"
)

var demo_pts = []*utils.PKGContent{
	&utils.PKGContent{Name: "jim", Parent: "jjim", Synopsis: "jjimm", Href: "jimm"},
}

func TestSomething(t *testing.T) {
	sc := &SQLConfig{
		DBName:   "pkg_lists",
		DBPort:   "3306",
		DBAddr:   "127.0.0.1",
		User:     "jim",
		Password: "password",
	}
	db, err := sc.openDB()
	assert.Nil(t, err)
	err = sc.closeDB(db)
	assert.Nil(t, err)
}

func TestInsertData(t *testing.T) {
	sc := &SQLConfig{
		DBName:   "pkg_lists",
		DBPort:   "3306",
		DBAddr:   "127.0.0.1",
		User:     "jim",
		Password: "password",
	}
	db, err := sc.openDB()
	assert.Nil(t, err)
	// assert.Equal(t, "", fmt.Sprintf("%v\n", db))
	err = insertData(db, &demo_pts)
	assert.Nil(t, err)

	err = sc.closeDB(db)
	assert.Nil(t, err)
}
