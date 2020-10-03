package main

import (
	"fmt"
	"os"
	"strings"

	_ "github.com/mattn/getwild"

	. "github.com/zetamatta/wfile"
)

func mains(args []string) error {
	if len(args) <= 0 {
		Usage(os.Stdout)
		return nil
	}
	for i, fname := range args {
		result, err := Report(fname)
		if err != nil {
			return fmt.Errorf("%s: %w", fname, err)
		}
		if i > 0 {
			fmt.Println()
		}
		fmt.Printf("%s\n    %s\n", fname, strings.Join(result, "\n    "))
	}
	return nil
}

func main() {
	if err := mains(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
