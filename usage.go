package whatfile

import (
	"fmt"
	"io"
	"os"
	"sort"
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

func Usage(w io.Writer) {
	fmt.Fprintf(w, "%s - what file is it ?\n\n", os.Args[0])
	fmt.Fprintf(w, "Usage: %s {filenames}\n\n", os.Args[0])
	fmt.Fprintln(w, "Signature database")

	suffixes := make([]string, 0, len(suffixTable))
	for suffix := range suffixTable {
		suffixes = append(suffixes, suffix)
	}
	sort.Strings(suffixes)

	for _, suffix := range suffixes {
		features := suffixTable[suffix]
		fmt.Fprintf(w, "for *.%s:\n", suffix)
		for _, f := range features {
			fmt.Print("  ")
			putBin(f.Magic, w)
			if f.Offset > 0 {
				fmt.Fprintf(w, " from %d", f.Offset)
			}
			fmt.Fprintf(w, " ... %s\n", f.Desc)
		}
	}
	fmt.Fprintf(w, "for *:\n")
	for _, f := range flatTable {
		fmt.Fprint(w, "  ")
		putBin(f.Magic, w)
		if f.Offset > 0 {
			fmt.Fprintf(w, " from %d", f.Offset)
		}
		fmt.Fprintf(w, " ... %s\n", f.Desc)
	}
}
