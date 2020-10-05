package main

import (
	"fmt"
	"reflect"
	"time"
)

func time1() {
	fmt.Println("=== time1 start ===")
	time.Sleep(1 * time.Second)
	fmt.Println("=== time1 end ===")
}

func time2(s int) {
	fmt.Println("=== time2 start ===")
	time.Sleep(time.Duration(s) * time.Second)
	fmt.Println("=== time2 end ===")
}

func makeTimeFunc(fn interface{}) interface{} {
	rf := reflect.TypeOf(fn)
	fmt.Println(rf)
	if rf.Kind() != reflect.Func {
		panic("not reflect.Func")
	}

	vf := reflect.ValueOf(fn)
	fmt.Println(vf)

	ft := func(in []reflect.Value) []reflect.Value {
		fmt.Println("=== ft start ===")
		out := vf.Call(in)
		fmt.Println("=== ft end ===")
		return out
	}

	mf := reflect.MakeFunc(rf, ft)

	return mf.Interface()
}

func main() {
	fmt.Println("ref")

	time1()
	time2(2)

	ft1 := makeTimeFunc(time1).(func())
	ft1()

	ft2 := makeTimeFunc(time2).(func(int))
	ft2(2)
}
