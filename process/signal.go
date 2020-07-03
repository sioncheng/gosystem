package main

import (
	"log"
	"os"
	"os/signal"
)

func main() {
	log.Println("start application...")
	c := make(chan os.Signal)
	signal.Notify(c)

	s := <-c
	log.Println("exit with signal:", s)
}
