package atexit

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func Exec(f func() error) {
	var c = make(chan os.Signal)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		for sig := range c {
			fmt.Printf("PID %d receives signal: %s\n", os.Getpid(), sig)
			if err := f(); err != nil {
				println(err.Error())
				os.Exit(1)
			}
			os.Exit(0)
		}
	}()
}
