package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Please specify a file path.")
		return
	}

	b, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Can't open file.", err)
	}

	fmt.Println(string(b))
}
