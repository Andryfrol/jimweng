// https://blog.golang.org/error-handling-and-go
package main

import (
	"log"
	"os"
)

func Open(name string) (file *File, err error)

func main() {
	f, err := os.Open("filename.ext")
	if err != nil {
		log.Fatal(err)
	}

}
