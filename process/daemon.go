package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

var pid = os.Getpid()

func main() {
	fmt.Printf("[%d] start\n", pid)
	fmt.Printf("[%d] PPID:%d\n", pid, os.Getppid())

	if len(os.Args) != 1 && os.Args[1] == "daemon" {
		runDaemon()
		return
	}

	if err := forkProcess(); err != nil {
		fmt.Println(err)
		return
	}

	if err := relaseResources(); err != nil {
		fmt.Println(err)
		return
	}
}

func runDaemon()  {
	for i := 1; i < 10; i++ {
		fmt.Println("daemon")
		time.Sleep(time.Second * 1)
	}
}

func forkProcess() error {
	cmd := exec.Command(os.Args[0], "daemon")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Dir = "/"
	return cmd.Start()
}

func relaseResources() error {
	p, err := os.FindProcess(pid)
	if err != nil {
		return err
	}
	return p.Release()
}