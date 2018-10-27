package main

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type Feature struct {
	Offset int
	Magic  []byte
	Func   func(fname string, bin []byte) string
	Desc   string
}

type B = []byte

var FeatureTable = map[string][]*Feature{
	"exe": {{Magic: B("MZ"), Func: tryExe}},
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
}

var FeatureTable2 = []*Feature{
	{Magic: B("#!"), Desc: "UNIX Executables"},
	{Magic: B("\xEF\xBB\xBF#!"), Desc: "Broken UNIX Executables(BOM)"},
}

func testFeatures(fname string, bin []byte, features []*Feature) bool {
	for _, f := range features {
		if f.Offset+len(f.Magic) < len(bin) && bytes.Equal(bin[f.Offset:f.Offset+len(f.Magic)], f.Magic) {
			if f.Func != nil {
				fmt.Printf("%s: %s\n", fname, f.Func(fname, bin))
			}
			if f.Desc != "" {
				fmt.Printf("%s: %s\n", fname, f.Desc)
			}
			return true
		}
	}
	return false
}

func eachFile(fname string) error {
	fd, err := os.Open(fname)
	if err != nil {
		return err
	}
	bin := make([]byte, 1024)
	n, err := fd.Read(bin)
	if err != nil {
		println(err)
	}
	defer fd.Close()
	bin = bin[:n]

	suffix := strings.TrimPrefix(strings.ToLower(filepath.Ext(fname)), ".")
	if features, ok := FeatureTable[suffix]; ok {
		if testFeatures(fname, bin, features) {
			return nil
		}
	}
	if testFeatures(fname, bin, FeatureTable2) {
		return nil
	}
	for _, features := range FeatureTable {
		if testFeatures(fname, bin, features) {
			return nil
		}
	}
	return nil
}

func main() {
	for _, fname := range os.Args[1:] {
		if err := eachFile(fname); err != nil {
			fmt.Fprintf(os.Stderr, "%s: %s\n", fname, err.Error())
		}
	}
}
