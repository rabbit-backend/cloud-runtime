package env

import "github.com/dop251/goja"

type EnvModule struct {
	vm *goja.Runtime
}

func New(vm *goja.Runtime) *EnvModule {
	return &EnvModule{
		vm: vm,
	}
}
