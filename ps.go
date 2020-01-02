package main 

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	ps , err := exec.LookPath("ps")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(ps)

	command := []string{"ps", "-ax"}
	env := os.Environ()

	err = syscall.Exec(ps, command, env)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	os.Exit(0)
}