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
	DBUri     string
}

type OperationDatabase struct {
	DB *gorm.DB
}

type OPDB interface {
	create(name string, email string) error
	queryWithName(name string) (string, error)
	updateEmail(name string, email string) error
	deleteData(name string, email string) error
	Closed() error
	debug()
}

func (dbc *DBConfig) NewDBConnection() (OPDB, error) {
	// connection :=
	db, err := gorm.Open(dbc.DBType, dbc.DBUri)
	if err != nil {
		return nil, err
	}
	db = db.AutoMigrate(&DemoTable{})
	return &OperationDatabase{DB: db}, err
}

func NewDBConfiguration(user string, password string, dbtype string, dbname string, dbport string, dbaddress string) *DBConfig {
	return &DBConfig{
		User:      user,
		Password:  password,
		DBType:    dbtype,
		DBName:    dbname,
		DBPort:    dbport,
		DBAddress: dbaddress,
		DBUri:     user + ":" + password + "@tcp(" + dbaddress + ":" + dbport + ")/" + dbname + "?charset=utf8&parseTime=True&loc=Local",
	}
}

func main() {
	fmt.Println("Go ORM Tutorial")
	newDB := NewDBConfiguration("jim", "password", "mysql", "demo_db", "3306", "127.0.0.1")
	db, err := newDB.NewDBConnection()
	if err != nil {
		log.Fatal(err)
	}

	if queryrest, err := db.queryWithName("jim"); err == nil {
		fmt.Printf("The query result is %v\n", queryrest)
	} else {
		fmt.Printf("Error happend while querying %s\n", err)
	}

	defer db.Closed()
}

func (db *OperationDatabase) Closed() error {
	if err := db.DB.Close(); err != nil {
		return fmt.Errorf("Error happended while closing db : %v\n", err)
	}
	log.Fatalln("Going to close DB")
	return nil
}

// 透過使用Debug()可以轉譯語言為SQL語法
func (db *OperationDatabase) debug() {
	db.DB = db.DB.Debug()
}

// 實做CRUD
// Create
func (db *OperationDatabase) create(name string, email string) error {
	var dt = &DemoTable{
		Name:  name,
		Email: email,
	}
	if err := db.DB.Create(dt).Error; err != nil {
		return err
	}
	return nil
}

// Read
func (db *OperationDatabase) queryWithName(name string) (string, error) {
	// log.Printf("The %s's Email has been found with %s", name, db.DB.Find(&DemoTable{Name: name}).Value)
	// return fmt.Sprintf("%v", db.DB.Select("email").Where("name = ?", name).Value), nil
	// return fmt.Sprintf("%v", db.DB.Select("email").Find(&DemoTable{Name: name}).Where("name = ?", name).Value), nil
	// return fmt.Sprintf("%v", db.DB.Select("email").Find(&DemoTable{Name: name}).Value), nil
	// var dt DemoTable
	var dt = &DemoTable{
		Name: name,
	}
	if err := db.DB.Select("email").Find(dt).Error; err != nil {
		return "Can't find the email with " + name, err
	}
	return dt.Email, nil
}

// Update ... 更新相當於Read以後在把Read的資料改成新的資料；notes:在gorm裡面，更新以後也會更新updated_at的時間
func (db *OperationDatabase) updateEmail(name string, email string) error {
	// log.Printf("The %s's Email has been update to %s", name, db.DB.First(&DemoTable{Name: name}).Update(&DemoTable{Name: name, Email: email}).Value)
	if err := db.DB.First(&DemoTable{Name: name}).Update(&DemoTable{Name: name, Email: email}).Error; err != nil {
		return err
	}
	return nil
}

// Delete ... 因為delete已經有預設方法，這邊改用deleteData來宣告該函數；notes:在gorm裡面刪除不是代表從db完全移除。而是去更改deleted_at的時間
func (db *OperationDatabase) deleteData(name string, email string) error {
	// log.Printf("The %s's Email has been delete (%s)", name, db.DB.Delete(&DemoTable{Name: name, Email: email}).Value)
	if err := db.DB.Where("email = ?", email).Delete(&DemoTable{}).Error; err != nil {
		log.Fatal("Encount Error with no data to delete")
		return err
	}
	return nil
}
