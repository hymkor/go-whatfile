package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/mattn/msgbox"

	"github.com/zetamatta/wfile"
)

func message(msg string) {
	msgbox.Show(0, msg, os.Args[0], msgbox.OK)
}

func main() {
	if len(os.Args) <= 1 {
		message(fmt.Sprintf("Usage:\n%s FILENAME", os.Args[0]))
	}
	var out strings.Builder
	var errmsg strings.Builder

	if err := wfile.Report(os.Args[1], &out, &errmsg); err != nil {
		message(err.Error())
	} else if errmsg.Len() > 0 {
		message(errmsg.String())
	} else {
		message(out.String())
	}
}
