package main

import (
	"fmt"

	"github.com/sioncheng/gosystem/pkg-use/pa"
	//use of internal package github.com/sioncheng/gosystem/pkg-use/pa/internal not allowed
	//"github.com/sioncheng/gosystem/pkg-use/pa/internal"
)

func main() {
	a := pa.AClass{}
	//b := internal.BClass{}
	fmt.Println("package use", a)
}
