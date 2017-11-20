package main

import (
	"fmt"
	"sort"
)

func main(){
	// claim a array
	cases := []struct {
		in, want string
	}{
		{"Hello, world", "dlrow ,olleH"},
		{"Hello, 世界", "界世 ,olleH"},
		{"", ""},
	}

	

	fmt.Println(cases)
}

