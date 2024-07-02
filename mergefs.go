package mergefs

import (
	"io/fs"
	"os"
)

type mfs struct{ fss []fs.FS }

func Merge(fss []fs.FS) fs.FS { return mfs{fss: fss} }

func (mfs mfs) Open(name string) (fs.File, error) {
	for _, fs := range mfs.fss {
		f, err := fs.Open(name)
		if err == nil {
			return f, nil
		}
	}
	return nil, os.ErrNotExist
}
