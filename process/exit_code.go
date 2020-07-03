package main

import (
	"fmt"
	"os/exec"
	"syscall"
)

func main() {
	cmd := exec.Command("ls", "-l")
	if err := cmd.Run(); err != nil {
		fmt.Print(err)
	}

	status, ok := cmd.ProcessState.Sys().(syscall.WaitStatus)
	if !ok {
		fmt.Println("!ok")
	}
	fmt.Println(status.ExitStatus())

}
