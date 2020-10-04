package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"

	"github.com/zetamatta/wfile"
)

func report(files []string) string {
	var buffer strings.Builder
	for i, fname := range files {
		if i > 0 {
			buffer.WriteString("\r\n\r\n")
		}
		if result, err := wfile.Report(fname, nil); err != nil {
			fmt.Fprintf(&buffer, "%s:\r\n  %s", fname, err)
		} else {
			fmt.Fprintf(&buffer, "%s:\r\n  %s", fname, strings.Join(result, "\r\n  "))
		}
	}
	return buffer.String()
}

func main() {
	var textEdit *walk.TextEdit
	defaultText := "Drop files here, from windows explorer..."
	empty := true

	if len(os.Args) >= 2 {
		defaultText = report(os.Args[1:])
		empty = false
	}

	MainWindow{
		Title:   "What file is it?",
		MinSize: Size{400, 300},
		Layout:  VBox{},
		OnDropFiles: func(files []string) {
			if empty {
				textEdit.SetText(report(files))
				empty = false
			} else {
				textEdit.AppendText("\r\n\r\n")
				textEdit.AppendText(report(files))
			}
		},
		Children: []Widget{
			TextEdit{
				AssignTo: &textEdit,
				ReadOnly: true,
				Text:     defaultText,
				Font:     Font{PointSize: 12},
			},
		},
	}.Run()
}
