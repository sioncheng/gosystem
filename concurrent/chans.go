package main

import "fmt"

func main() {
	intchan := make(chan int)
	//close(intchan)
	intchan = nil
	go func() {
		select {
		case intchan <- 1:
		default:
			fmt.Println("do not panic")
		}

	}()
	if v, ok := <- intchan; ok {
		fmt.Println(v)
	}
}
