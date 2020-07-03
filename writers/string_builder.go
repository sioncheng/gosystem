package main

import "strings"

import "fmt"

func main() {
	b := strings.Builder{}
	b.WriteString("Hello")
	fmt.Println(b)
	b.WriteString("Wolrd")
	fmt.Println(b)

	c := b
	c.WriteString("!")
}
