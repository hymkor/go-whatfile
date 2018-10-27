package main

import (
	"bytes"
	"unicode/utf8"
)

func guessCode(bin []byte) string {
	if bytes.HasPrefix(bin, B("\xEF\xBB\xBF")) {
		return "UTF8 with BOM"
	}
	if bytes.HasPrefix(bin, B("\xFE\xFF")) {
		return "UTF16BE"
	}
	if bytes.HasPrefix(bin, B("\xFF\xFE")) {
		return "UTF16LE"
	}
	if pos := bytes.IndexByte(bin, 0); pos >= 0 {
		if pos%2 == 0 {
			return "UTF16BE"
		} else {
			return "UTF16LE"
		}
	}
	for bin != nil && len(bin) > 0 {
		pos := bytes.IndexByte(bin, '\n')
		var line []byte
		if pos >= 0 {
			line = bin[:pos]
			bin = bin[pos+1:]
		} else {
			line = bin
			bin = nil
		}
		if !utf8.Valid(line) {
			return "ANSI(MBCS)"
		}
		if utf8.RuneCount(line) != len(line) {
			return "UTF8"
		}
	}
	return "ANSI(SBCS)"
}

func TryText(fname string, bin []byte) string {
	code := guessCode(bin)
	if bytes.Contains(bin, B("\r\n")) || bytes.Contains(bin, B("\r\000\n")) {
		code = code + ",CRLF"
	} else if bytes.Contains(bin, B("\n")) {
		code = code + ",LF"
	}
	return code + " text data"
}
