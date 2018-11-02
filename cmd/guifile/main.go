package main

import (
	"fmt"
	"strings"

	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"

	"github.com/zetamatta/wfile"
)

func main() {
	var textEdit *walk.TextEdit

	MainWindow{
		Title:   "What file is it?",
		MinSize: Size{400, 300},
		Layout:  VBox{},
		OnDropFiles: func(files []string) {
			var buffer strings.Builder
			for i, fname := range files {
				if i > 0 {
					buffer.WriteString("\r\n\r\n")
				}
				if result, err := wfile.Report(fname); err != nil {
					fmt.Fprintf(&buffer, "%s:\r\n  %s", fname, err)
				} else {
					fmt.Fprintf(&buffer, "%s:\r\n  %s", fname, result)
				}
			}
			textEdit.SetText(buffer.String())
		},
		Children: []Widget{
			TextEdit{
				AssignTo: &textEdit,
				ReadOnly: true,
				Text:     "Drop files here, from windows explorer...",
			},
		},
	}.Run()
}
