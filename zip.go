package wfile

func TryZip(fname string, bin []byte) []string {
	result := make([]string, 0, 3)
	result = append(result, "Zip Archive")

	if (bin[7] & 8) != 0 {
		result = append(result, "utf8-flag-on")
	} else {
		result = append(result, "utf8-flag-off")
	}
	if (bin[6] & 1) != 0 {
		result = append(result, "has password")
	}
	return result
}
