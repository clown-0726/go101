package main

import "fmt"

func main() {
	foo := "b"
	if foo == "a" {
		// do something
		fmt.Println("do something")
	} else if foo == "b" {
		// do something else
		fmt.Println("hello go")
	} else {
		// catch-all or default
		fmt.Println("hello go end")
	}
}
