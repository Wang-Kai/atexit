package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Wang-Kai/atexit"
)

func main() {
	fmt.Printf("Process running at PID: %d\n", os.Getpid())

	hook := func() error {
		println("Bye ...")
		return nil
	}

	atexit.Exec(hook)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
