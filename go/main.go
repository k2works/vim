package main

import "fmt"

func main() {
	message := Greeting()
	fmt.Println(message)
}

func Greeting() string {
	message := "Hello World"
	return message
}
