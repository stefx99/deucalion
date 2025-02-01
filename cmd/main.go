package main

import (
	"os"
	"os/signal"
)

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go func() {
		<-c
		// TODO: graceful shutdown logic

		os.Exit(1)
	}()

}
