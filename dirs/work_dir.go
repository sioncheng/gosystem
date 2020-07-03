package main

import (
	"fmt"
	"os"
)

func main() {
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("starting dir:", wd)

	if err := os.Chdir("/Users"); err != nil {
		fmt.Println(err)
		return
	}

	wd, err = os.Getwd()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("final dir:", wd)
}
