package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Please specify a file")
		return
	}

	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Can't open file", err)
		return
	}

	defer f.Close()

	r := bufio.NewReader(f)
	var rowCount int
	for nil == err {
		var b string
		b, err = r.ReadString('\n')
		fmt.Println(b)
		rowCount++
	}
	if err == io.EOF {
		fmt.Println("\r\nEOF")
	}

	fmt.Println("Row count:", rowCount)
}
