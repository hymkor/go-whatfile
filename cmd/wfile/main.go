package main

import (
	"fmt"
	"io"
	"os"
	"sort"

	. "github.com/zetamatta/wfile"
)

func putBin(bin []byte, w io.Writer) {
	for _, c := range bin {
		switch c {
		case '\r':
			fmt.Fprint(w, "\\r")
		case '\n':
			fmt.Fprint(w, "\\n")
		case '\t':
			fmt.Fprint(w, "\\t")
		case '\a':
			fmt.Fprint(w, "\\a")
		case '\b':
			fmt.Fprint(w, "\\b")
		default:
			if 0x20 <= c && c < 0x7F {
				fmt.Fprintf(w, "%c", c)
			} else {
				fmt.Fprintf(w, "\\x%02X", c)
			}
		}
	}
}

func main() {
	if len(os.Args) <= 1 {
		fmt.Printf("%s - what file is it ?\n\n", os.Args[0])
		fmt.Printf("Usage: %s {filenames}\n\n", os.Args[0])
		fmt.Println("Signature database")

		suffixes := make([]string, 0, len(FeatureTable))
		for suffix := range FeatureTable {
			suffixes = append(suffixes, suffix)
		}
		sort.Strings(suffixes)

		for _, suffix := range suffixes {
			features := FeatureTable[suffix]
			fmt.Printf("for *.%s:\n", suffix)
			for _, f := range features {
				fmt.Print("  ")
				putBin(f.Magic, os.Stdout)
				if f.Offset > 0 {
					fmt.Printf(" from %d", f.Offset)
				}
				fmt.Printf(" ... %s\n", f.Desc)
			}
		}
		fmt.Printf("for *:\n")
		for _, f := range FeatureTable2 {
			fmt.Print("  ")
			putBin(f.Magic, os.Stdout)
			if f.Offset > 0 {
				fmt.Printf(" from %d", f.Offset)
			}
			fmt.Printf(" ... %s\n", f.Desc)
		}
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
