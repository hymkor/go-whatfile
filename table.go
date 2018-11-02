package wfile

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

var suffixTable = map[string][]*Feature{
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

var flatTable = []*Feature{
	{Magic: B("#!"), Desc: "UNIX Executables"},
	{Magic: B("\xEF\xBB\xBF#!"), Desc: "Broken UNIX Executables(BOM)"},
	{Magic: B("\x7FELF"), Desc: "ELF - Executable and Linkable Format"},
}
