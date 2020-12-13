# atexit

`atexit` register a function to be called at normal process termination .

## How to use

Define and registry your atexit function

```go
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

        // define atexit function
	hook := func() error {
		println("Bye ...")
		return nil
	}

        // registry function
	atexit.Exec(hook)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
```

```console
$ go run atexit_demo.go

Process running at PID: 24831
```


Open the second terminal, send `SIGINT(2)` or `SIGTERM(15)` to process.

```console
kill -2 24831
```

The atexit function will be called before process exit.

```console
$ go run atexit_demo.go

Process running at PID: 24831
PID 24831 receives signal: interrupt
Bye ...
```
