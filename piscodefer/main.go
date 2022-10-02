package main

import "fmt"

func main() {
	defer fmt.Println("hello")
	defer fmt.Println("hello1")
	defer fmt.Println("hello12")

	fmt.Println("defer")
}
