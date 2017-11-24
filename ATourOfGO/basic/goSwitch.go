package main

import "fmt"

var numbers = [6]int{1, 2, 3, 4, 5}

func main() {
	for name, _ := range numbers {
		fmt.Println("name:", name)

		switch name {
		case 1:
			fmt.Println("name:", "John")
		case 2:
			fmt.Println("name:", "Jim")
		default:
			fmt.Println("name:", "Joker")
		}

	}
}
