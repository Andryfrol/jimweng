package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/goPractice/pkgserver/pkgserver"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

var db *OperationDatabase

// server is used to implement helloworld.GreeterServer.
type server struct{}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.Name)

	return &pb.HelloReply{
		Message: db.queryWithName(in.Name),
	}, nil
}

func main() {
	newDB := NewDBConfiguration("jim", "password", "mysql", "pkg_lists", "3306", "127.0.0.1")
	db, err := newDB.NewDBConnection()
	if err != nil {
		log.Fatal(err)
	}

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	defer db.Closed()
}

// Our DemoTable Struct
type DemoTable struct {
	// gorm.Model
	Name     string `gorm:"primary_key"`
	Parent   string `gorm:"primary_key"`
	Synopsis string
	Href     string
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
	queryWithName(name string) string
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

// Read
func (db *OperationDatabase) queryWithName(name string) string {
	// var dt = &utils.PKGContent{
	// Name: name,
	// }
	// if err := db.DB.Debug().Find(dt).Error; err != nil {
	// 	log.Fatalln("Can't find the parent with "+name, err)
	// }
	return "ok"
	// return dt.Parent
}
