package atexit

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func Exec(f func() error) {
	var sChan = make(chan os.Signal)
	signal.Notify(sChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		for s := range sChan {
			fmt.Printf("PID %d receives signal: %s\n", os.Getpid(), s)

			if err := f(); err != nil {
				println(err.Error())
				os.Exit(1)
			}

			os.Exit(0)
		}
	}()
}
