package fs

import (
	"bytes"
	"os"

	"github.com/dop251/goja"
)

func (fs *FsModule) Read(path string) goja.ArrayBuffer {
	file, err := os.Open(path)
	if err != nil {
		panic(fs.vm.ToValue(err))
	}

	buf := new(bytes.Buffer)
	if _, err = file.WriteTo(buf); err != nil {
		panic(fs.vm.ToValue(err))
	}

	return fs.vm.NewArrayBuffer(buf.Bytes())
}
