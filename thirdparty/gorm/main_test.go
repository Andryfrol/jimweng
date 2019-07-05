package main

import (
	"testing"

	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestNewConnection(t *testing.T) {
	var (
		user     = "jim"
		password = "pw"
		dbtype   = "dbtype"
		dbname   = "demo_db"
		dbport   = "3306"
		dbaddr   = "127.0.0.1"
	)
	NewDBConfig := NewDBConfiguration(user, password, dbtype, dbname, dbport, dbaddr)
	assert.Equal(t, user, NewDBConfig.User)
	assert.Equal(t, password, NewDBConfig.Password)
	assert.Equal(t, dbtype, NewDBConfig.DBType)
	assert.Equal(t, dbname, NewDBConfig.DBName)
	assert.Equal(t, dbport, NewDBConfig.DBPort)
	assert.Equal(t, dbaddr, NewDBConfig.DBAddress)
}

func TestNewDBConnection(t *testing.T) {
	opDB := opdb{}
	db, mock, _ := sqlmock.New()
	mock.ExpectBegin()

	// columns := []string{"name", "email"}
	opDB.DB, _ = gorm.Open("mysql", db)
	err := opDB.create("jim", "123")
	assert.Nil(t, err)
	responseString, err := opDB.queryWithName("jim")
	assert.Nil(t, err)
	assert.Equal(t, "jim", responseString)

	// sqlRows := sqlmock.NewRows(columns).AddRow("jim", "123")
	// mock.ExpectQuery("SELECT * FROM demo_tables").WithArgs(1).WillReturnRows(sqlRows)
	// // mock.ExpectExec("INSERT INTO demo_tables (name, email) VALUES (\"jim\",\"123\"").WillReturnResult(sqlmock.NewResult(1, 1))
	// assert.Equal(t, "", mock.ExpectationsWereMet())

}
