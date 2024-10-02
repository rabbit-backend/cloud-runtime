package http

import "github.com/dop251/goja"

type HttpModule struct {
	vm *goja.Runtime
}

func New(vm *goja.Runtime) *HttpModule {
	return &HttpModule{
		vm: vm,
	}
}
