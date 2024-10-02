package fs

import (
	"bytes"
	"os"
)

func (fs *FsModule) Read(path string) string {
	file, err := os.Open(path)
	if err != nil {
		panic(fs.vm.ToValue(err))
	}

	buf := new(bytes.Buffer)
	if _, err = file.WriteTo(buf); err != nil {
		panic(fs.vm.ToValue(err))
	}

	return buf.String()
}
