package main

import (
	"fmt"
)

func main() {
	fmt.Println("switch use")

	a := 0
	switch a {
	case 0:
		fmt.Println("zero")
	case 1:
		fmt.Println("one")
	default:
		fmt.Println("more")
	}
}
