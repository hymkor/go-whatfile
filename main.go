package wfile

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func guess(fname string, bin []byte, features []*Feature) []string {
	for _, f := range features {
		if f.Offset+len(f.Magic) < len(bin) && bytes.Equal(bin[f.Offset:f.Offset+len(f.Magic)], f.Magic) {
			if f.Func != nil {
				return f.Func(fname, bin)
			} else {
				return []string{f.Desc}
			}
		}
	}
	return nil
}

func Report(fname string, f func(io.Reader) []string) ([]string, error) {
	fd, err := os.Open(fname)
	if err != nil {
		return nil, err
	}
	defer fd.Close()

	if stat, err := fd.Stat(); err != nil {
		return nil, err
	} else if stat.IsDir() {
		return nil, fmt.Errorf("%s: Directory", fname)
	}

	bin := make([]byte, 1024)
	n, err := fd.Read(bin)
	if err == io.EOF {
		return nil, fmt.Errorf("%s: zero byte file\n", fname)
	}
	if err != nil {
		return nil, err
	}
	output, err := report(fname, bin[:n])
	if err != nil {
		return nil, err
	}
	if f != nil {
		appendOutput := f(io.MultiReader(bytes.NewReader(bin[:n]), fd))
		if appendOutput != nil && len(appendOutput) > 0 {
			output = append(output, appendOutput...)
		}
	}
	return output, nil
}

func report(fname string, bin []byte) ([]string, error) {
	suffix := strings.TrimPrefix(strings.ToLower(filepath.Ext(fname)), ".")
	if features, ok := suffixTable[suffix]; ok {
		if result := guess(fname, bin, features); result != nil {
			return result, nil
		}
	}
	if result := guess(fname, bin, flatTable); result != nil {
		return result, nil
	}
	for _, features := range suffixTable {
		if result := guess(fname, bin, features); result != nil {
			return result, nil
		}
	}
	return TryText(fname, bin), nil
}
