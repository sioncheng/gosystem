package main

import (
	"fmt"
)

func main() {
	fmt.Println(f1())
}

func f1() (r int) {
	r = 100
	defer func(){
		r = r + 1
	}()
	return r
}