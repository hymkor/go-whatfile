package whatfile

import (
	"github.com/zetamatta/vf1s/peinfo"
)

func GetVersionInfo(fname string) (*Version, error) {
	vi, err := peinfo.GetVersionInfo(fname)
	if err != nil {
		return nil, err
	}
	file, product, err := vi.Number()
	if err != nil {
		return nil, err
	}
	v := &Version{}
	copy(v.File[:], file)
	copy(v.Product[:], product)
	return v, nil
}
