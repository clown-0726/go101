package main

import "fmt"
import "go101/grammar/init/a"
import "go101/grammar/init/b"

func init() {
	fmt.Println("main init")
}

func main() {
	a.SayHello()
	b.SayHello()
	fmt.Println("test")
}
