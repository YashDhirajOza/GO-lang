package main

import "fmt"

func main() {
	fmt.Println("Enter your name")
	var name string
	fmt.Scan(&name)
	fmt.Println("Your name is " + name)
}
