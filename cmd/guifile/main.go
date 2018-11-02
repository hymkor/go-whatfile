package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/mattn/msgbox"

	"github.com/zetamatta/wfile"
)

func message(msg, title string) {
	msgbox.Show(0, msg, title, msgbox.OK)
}

func main() {
	if len(os.Args) <= 1 {
		me := filepath.Base(os.Args[0])
		message(fmt.Sprintf("Usage:\n%s FILENAME", me), me)
		return
	}
	if result, err := wfile.Report(os.Args[1]); err != nil {
		message(err.Error(), os.Args[1])
	} else {
		message(result, os.Args[1])
	}
}
