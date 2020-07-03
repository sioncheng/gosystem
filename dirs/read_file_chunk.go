package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Please specify a file.")
		return
	}

	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Can't open file", err)
		return
	}

	defer f.Close()

	b := make([]byte, 16)

	for n := 0; err == nil; {
		n, err = f.Read(b)
		if err == nil {
			fmt.Println(string(b[:n]))
		} else if err == io.EOF {
			fmt.Println("\r\n EOF")
		} else {
			fmt.Println("\r\n Error", err)
		}
	}
}
