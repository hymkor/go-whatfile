package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/floyernick/fleep-go"
)

var extensions = map[string]func(fname string, bin []byte) string{
	"exe": tryExe,
}

func eachFile(fname string) error {
	file, err := ioutil.ReadFile(fname)
	if err != nil {
		return err
	}
	info, err := fleep.GetInfo(file)
	if err != nil {
		return err
	}
	fmt.Printf("%s\n", fname)
	functions := make([]func(string, []byte) string, 0, 10)
	for i := range info.Type {
		fmt.Printf("  %-13s: %-4s: %s\n", info.Type[i], info.Extension[i], info.Mime[i])
		if extension1, ok := extensions[info.Extension[i]]; ok {
			functions = append(functions, extension1)
		}
	}
	for _, f := range functions {
		if detail := f(fname, file); detail != "" {
			for _, line := range strings.Split(detail, "\n") {
				fmt.Printf("  %s\n", line)
			}
		}
	}
	return nil
}

func main() {
	for _, fname := range os.Args[1:] {
		if err := eachFile(fname); err != nil {
			fmt.Fprintf(os.Stderr, "%s: %s\n", fname, err.Error())
		}
		fmt.Println()
	}
}
