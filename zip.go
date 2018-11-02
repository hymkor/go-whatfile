package wfile

import (
	"strings"
)

func TryZip(fname string, bin []byte) string {
	var buffer strings.Builder

	buffer.WriteString("Zip Archive")
	if (bin[7] & 8) != 0 {
		buffer.WriteString(",utf8-flag-on")
	} else {
		buffer.WriteString(",utf8-flag-off")
	}
	if (bin[6] & 1) != 0 {
		buffer.WriteString(",with password")
	}
	return buffer.String()
}
