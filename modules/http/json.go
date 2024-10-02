package http

import (
	"bytes"
	"encoding/json"

	"github.com/dop251/goja"
	"github.com/go-zoox/fetch"
)

type JSON struct {
	vm *goja.Runtime
}

func (j *JSON) Json(response *fetch.Response) interface{} {
	var result interface{}

	buf := new(bytes.Buffer)
	buf.Write(response.Body)

	if err := json.NewDecoder(buf).Decode(&result); err != nil {
		panic(j.vm.ToValue(err))
	}

	return result
}
