package main

import (
	"fmt"
	"reflect"
)

type t1 int
type t2 int

type SomeStructure struct {
	X int
	Y string
}

func main() {
	fmt.Println(1)

	x1 := t1(1)
	x2 := t2(1)

	fmt.Println(x1, x2)

	v1 := reflect.ValueOf(&x1)
	fmt.Println(v1, v1.Type())
	e1 := v1.Elem()
	fmt.Println(e1, e1.Type(),reflect.TypeOf(x1))
	if e1.CanSet() {
		fmt.Println("e1 can set")
		e1.SetInt(10)
		fmt.Println(x1)
	}

	ss := SomeStructure{100, "100s"}
	fmt.Println(ss)

	vss := reflect.ValueOf(&ss)
	fmt.Println(vss)

	tvss := vss.Elem()
	fmt.Println(tvss.Type())
	for i := 0 ; i < tvss.NumField(); i++ {
		fmt.Println(tvss.Field(i), tvss.Field(i).Type(), tvss.Field(i).Interface())
	}

	tss := reflect.TypeOf(ss)
	fmt.Println(tss, tss.Field(0), tss.Field(1))
}
