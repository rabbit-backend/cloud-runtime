package require

import "github.com/dop251/goja"

type Module interface {
	Exports() map[string]interface{}
}

type Require struct {
	module_map map[string]interface{}
	VM         *goja.Runtime
}

func New(vm *goja.Runtime) *Require {
	return &Require{
		module_map: map[string]interface{}{},
		VM:         vm,
	}
}
