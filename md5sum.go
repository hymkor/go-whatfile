package whatfile

import (
	"crypto/md5"
	"fmt"
	"io"
)

func Md5(r io.Reader) []string {
	h := md5.New()
	io.Copy(h, r)
	return []string{fmt.Sprintf("md5sum: %x", h.Sum(nil))}
}
