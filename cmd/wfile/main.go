package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	_ "github.com/mattn/getwild"

	"github.com/zetamatta/go-whatfile"
)

var flagMd5 = flag.Bool("md5", false, "Show md5sum")

func mains(args []string) error {
	if len(args) <= 0 {
		whatfile.Usage(os.Stdout)
		return nil
	}
	var f func(io.Reader) []string
	if *flagMd5 {
		f = whatfile.Md5
	}
	for i, fname := range args {
		result, err := whatfile.Report(fname, f)
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
	flag.Parse()
	if err := mains(flag.Args()); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
