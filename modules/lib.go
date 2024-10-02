package modules

import (
	"github.com/dop251/goja"
	"github.com/rabbit-backend/cloud-runtime/modules/env"
	"github.com/rabbit-backend/cloud-runtime/modules/fs"
	"github.com/rabbit-backend/cloud-runtime/modules/http"
)

type Module interface {
	Exports() map[string]interface{}
}

type Factory = func(vm *goja.Runtime) Module

func CreateFactory(vm *goja.Runtime) map[string]Module {
	factory := map[string]Module{
		"rabbit:http": http.New(vm),
		"rabbit:env":  env.New(vm),
		"rabbit:fs":   fs.New(vm),
	}

	return factory
}
