package main

import (
	"fmt"
	"os"

	. "github.com/zetamatta/wfile"
)

func main() {
	if len(os.Args) <= 1 {
		Usage(os.Stdout)
		return
	}
	for _, fname := range os.Args[1:] {
		if result, err := Report(fname); err != nil {
			fmt.Fprintf(os.Stderr, "%s: %s\n", fname, err)
		} else {
			fmt.Printf("%s: %s\n", fname, result)
		}
	}
}
