package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

type Feature struct {
	Offset int
	Magic  []byte
	Func   func(fname string, bin []byte) string
	Desc   string
}

type B = []byte

var jpeg = []*Feature{
	{Magic: B("\xFF\xD8\xFF\xDB"), Desc: "JPEG Image"},
	{Magic: B("\xFF\xD8\xFF\xE0\x00\x10\x4A\x46\x49\x46\x00\x01"), Desc: "JPEG Image"},
	{Magic: B("\xFF\xD8\xFF\xEE"), Desc: "JPEG Image"},
	{Magic: B("\xFF\xD8\xFF\xE1"), Desc: "JPEG Image"},
}

var exe = []*Feature{
	{Magic: B("MZ"), Func: tryExe, Desc: "Portable Executable"},
}

var FeatureTable = map[string][]*Feature{
	"exe": exe,
	"dll": exe,
	"arx": exe,
	"brx": exe,
	"dwg": {
		{Magic: B("AC1003"), Desc: "AutoCAD EX-II"},
		{Magic: B("AC1006"), Desc: "AutoCAD GX-III"},
		{Magic: B("AC1009"), Desc: "AutoCAD 12,12,GX-5"},
		{Magic: B("AC1012"), Desc: "AutoCAD 13"},
		{Magic: B("AC1014"), Desc: "AutoCAD 14"},
		{Magic: B("AC1015"), Desc: "AutoCAD 2000,2000i,2002"},
		{Magic: B("AC1018"), Desc: "AutoCAD 2004,2005,2006"},
		{Magic: B("AC1021"), Desc: "AutoCAD 2007,2008,2009"},
		{Magic: B("AC1024"), Desc: "AutoCAD 2010,2011,2012"},
		{Magic: B("AC1027"), Desc: "AutoCAD 2013,2014,2015,2016,2017"},
		{Magic: B("AC1032"), Desc: "AutoCAD 2018"},
	},
	"gz":    {{Magic: B("\x1F\x8B"), Desc: "gzip compressed"}},
	"class": {{Magic: B("\xCA\xFE\xBA\xBE"), Desc: "Java class file"}},
	"lzh":   {{Offset: 2, Magic: B("-lh"), Desc: "LHA Archive"}},
	"pdf":   {{Magic: B("%PDF-"), Desc: "PDF"}},
	"zip":   {{Magic: B("PK\003\004"), Func: TryZip, Desc: "ZIP Archive"}},
	"png":   {{Magic: B("\x89\x50\x4E\x47\x0D\x0A\x1A\x0A"), Desc: "Portable Network Graphics"}},
	"jpg":   jpeg,
	"jpeg":  jpeg,
}

var FeatureTable2 = []*Feature{
	{Magic: B("#!"), Desc: "UNIX Executables"},
	{Magic: B("\xEF\xBB\xBF#!"), Desc: "Broken UNIX Executables(BOM)"},
	{Magic: B("\x7FELF"), Desc: "ELF - Executable and Linkable Format"},
}

func testFeatures(fname string, bin []byte, features []*Feature, w io.Writer) bool {
	for _, f := range features {
		if f.Offset+len(f.Magic) < len(bin) && bytes.Equal(bin[f.Offset:f.Offset+len(f.Magic)], f.Magic) {
			if f.Func != nil {
				fmt.Fprintf(w, "%s: %s\n", fname, f.Func(fname, bin))
			} else if f.Desc != "" {
				fmt.Fprintf(w, "%s: %s\n", fname, f.Desc)
			}
			return true
		}
	}
	return false
}

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

func eachFile(fname string, w io.Writer, errOut io.Writer) error {
	fd, err := os.Open(fname)
	if err != nil {
		return err
	}
	defer fd.Close()

	if stat, err := fd.Stat(); err != nil {
		return err
	} else if stat.IsDir() {
		fmt.Fprintf(w, "%s: Directory\n", fname)
		return nil
	}

	bin := make([]byte, 1024)
	n, err := fd.Read(bin)
	if err == io.EOF {
		fmt.Fprintf(w, "%s: zero byte file\n", fname)
		return nil
	}
	if err != nil {
		fmt.Fprintf(errOut, "%s: %s\n", err)
	}
	bin = bin[:n]

	suffix := strings.TrimPrefix(strings.ToLower(filepath.Ext(fname)), ".")
	if features, ok := FeatureTable[suffix]; ok {
		if testFeatures(fname, bin, features, w) {
			return nil
		}
	}
	if testFeatures(fname, bin, FeatureTable2, w) {
		return nil
	}
	for _, features := range FeatureTable {
		if testFeatures(fname, bin, features, w) {
			return nil
		}
	}
	fmt.Fprintf(w, "%s: %s\n", fname, TryText(fname, bin))
	return nil
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
		if err := eachFile(fname, os.Stdout, os.Stderr); err != nil {
			fmt.Fprintf(os.Stderr, "%s: %s\n", fname, err.Error())
		}
	}
}
