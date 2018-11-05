package main

import (
	"fmt"
	"os"

	"github.com/mattn/go-zglob"

	. "github.com/zetamatta/wfile"
)

func main1(fname string) {
	if result, err := Report(fname); err != nil {
		fmt.Fprintf(os.Stderr, "%s: %s\n", fname, err)
	} else {
		fmt.Printf("%s: %s\n", fname, result)
	}
}

func main() {
	if len(os.Args) <= 1 {
		Usage(os.Stdout)
		return
	}
	for _, arg1 := range os.Args[1:] {
		matches, err := zglob.Glob(arg1)
		if err != nil || matches == nil || len(matches) <= 0 {
			main1(arg1)
		} else {
			for _, fname := range matches {
				main1(fname)
			}
		}
	}
}
