package main

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Our DemoTable Struct
type DemoTable struct {
	// gorm.Model
	Name  string `gorm:"primary_key"`
	Email string
}

type DBConfig struct {
	User      string
	Password  string
	DBType    string
	DBName    string
	DBAddress string
	DBPort    string
}

type opdb struct {
	DB *gorm.DB
}

type OPDB interface {
	create(name string, email string) error
	queryWithName(name string) (string, error)
	update_email(name string, email string) error
	deleteData(name string, email string) error
	Closed()
	debug()
}

func (dbc *DBConfig) NewDBConnection() (OPDB, error) {
	connection := dbc.User + ":" + dbc.Password + "@tcp(" + dbc.DBAddress + ":" + dbc.DBPort + ")/" + dbc.DBName + "?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(dbc.DBType, connection)
	if err != nil {
		return nil, err
	}
	db = db.AutoMigrate(&DemoTable{})
	return &opdb{DB: db}, err
}

func NewDBConfiguration(user string, password string, dbtype string, dbname string, dbport string, dbaddress string) *DBConfig {
	return &DBConfig{
		User:      user,
		Password:  password,
		DBType:    dbtype,
		DBName:    dbname,
		DBPort:    dbport,
		DBAddress: dbaddress,
	}
}

func main() {
	fmt.Println("Go ORM Tutorial")
	newDB := NewDBConfiguration("jim", "password", "mysql", "demo_db", "3306", "127.0.0.1")
	db, err := newDB.NewDBConnection()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Closed()
	db.debug()
}

func (db *opdb) Closed() {
	log.Printf("Going to close DB")
	if err := db.DB.Close(); err != nil {
		log.Fatal(err)
	}
}

func (db *opdb) debug() {
	db.DB.Debug().Where("name =?", "jim").First(&DemoTable{})
}

// 實做CRUD
// Create
func (db *opdb) create(name string, email string) error {
	log.Printf("The %s's Email has been created with %s", name, db.DB.Create(&DemoTable{Name: name, Email: email}).Value)
	return nil
}

// Read
func (db *opdb) queryWithName(name string) (string, error) {
	// log.Printf("The %s's Email has been found with %s", name, db.DB.Find(&DemoTable{Name: name}).Value)
	// return fmt.Sprintf("%v", db.DB.Select("email").Where("name = ?", name).Value), nil
	// return fmt.Sprintf("%v", db.DB.Select("email").Find(&DemoTable{Name: name}).Where("name = ?", name).Value), nil
	return fmt.Sprintf("%v", db.DB.Select("email").Find(&DemoTable{Name: name}).Value), nil
}

// Update ... 更新相當於Read以後在把Read的資料改成新的資料；notes:在gorm裡面，更新以後也會更新updated_at的時間
func (db *opdb) update_email(name string, email string) error {
	log.Printf("The %s's Email has been update to %s", name, db.DB.First(&DemoTable{Name: name}).Update(&DemoTable{Name: name, Email: email}).Value)
	return nil
}

// Delete ... 因為delete已經有預設方法，這邊改用deleteData來宣告該函數；notes:在gorm裡面刪除不是代表從db完全移除。而是去更改deleted_at的時間
func (db *opdb) deleteData(name string, email string) error {
	log.Printf("The %s's Email has been delete (%s)", name, db.DB.Delete(&DemoTable{Name: name, Email: email}).Value)
	return nil
}
