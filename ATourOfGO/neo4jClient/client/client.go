package client

// package main

import (
	"io"
)

type Client interface {
	Query(command string) error
	WriteStream(b io.Reader) error
	Close() error
}

type WriteParams struct {
	Database        string
	RetentionPolicy string
	Precision       string
	Consistency     string
}

// func main(){
// 	test := WriteParams{
// 		Database :"test",
// 	}
// 	fmt.Println(test)
// }
