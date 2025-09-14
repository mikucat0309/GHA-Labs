package main

import "fmt"

func main() {
	println(hello("World"))
}

func hello(s string) string {
	return fmt.Sprintf("Hello, %s!", s)
}
