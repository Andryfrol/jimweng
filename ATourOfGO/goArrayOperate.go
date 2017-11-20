package main

import (
	"fmt"
)

func main(){
	cases := []struct {
		in, want string
	}{
		{"Hello, world", "dlrow ,olleH"},
		{"Hello, 世界", "界世 ,olleH"},
		{"", ""},
	}

	fmt.Println(cases)
}

