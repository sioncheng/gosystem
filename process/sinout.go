package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	buf := bytes.NewBuffer(nil)
	cmd := exec.Command("cat")
	cmd.Stdin = buf
	cmd.Stdout = os.Stdout

	fmt.Fprintf(buf,"Hello World! I am using this memory address: %p\n", buf)

	if err := cmd.Start(); err != nil {
		fmt.Println(err)
		return
	}

	cmd.Wait()
}
