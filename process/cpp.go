package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("ls", "-l")
	err := cmd.Start()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Command: ", cmd.Args[0])
	fmt.Println("Arguments: ", cmd.Args[1:])
	fmt.Println("Process Id: ", cmd.Process.Pid)

	cmd.Wait()
}
