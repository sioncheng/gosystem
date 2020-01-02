package main

import (
	"fmt"
)

func main() {
	a := [...]int{1,2,3}
	fmt.Println(a)
	//a = append(a, 4) , not compiled
	expect_array(a)

	b := []int{1,2,3}
	fmt.Println(b)
	b = append(b, 4)
	fmt.Println(b) 

	var vauto *[3]int = new([3]int)
	fmt.Println(vauto)

	as := a[0:1]
	fmt.Println(as)
	as[0] = 100
	fmt.Println(as)
	fmt.Println(a)
	
}

func expect_array(a [3]int) {
	fmt.Println(a)
}