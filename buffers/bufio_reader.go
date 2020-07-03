package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	sr := strings.NewReader("what are you doing now\ni am learning golang programming")
	br := bufio.NewReader(sr)
	line, more, err := br.ReadLine()
	fmt.Println(line, more, err)
	line, more, err = br.ReadLine()
	fmt.Println(line, more, err)
	line, more, err = br.ReadLine()
	fmt.Println(line, more, err)
}
