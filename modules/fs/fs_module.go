package fs

import "github.com/dop251/goja"

type FsModule struct {
	vm *goja.Runtime
}

func New(vm *goja.Runtime) *FsModule {
	return &FsModule{
		vm: vm,
	}
}
